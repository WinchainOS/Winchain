// Copyright 2017 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"time"
	"unicode"
	"github.com/WinchainOS/Winchain/consensus/raft"
	"github.com/WinchainOS/Winchain/p2p/discover"

	"gopkg.in/urfave/cli.v1"

	"github.com/naoina/toml"
	"github.com/WinchainOS/Winchain/cmd/utils"
	"github.com/WinchainOS/Winchain/dashboard"
	"github.com/WinchainOS/Winchain/eth"
	"github.com/WinchainOS/Winchain/node"
	"github.com/WinchainOS/Winchain/params"
	whisper "github.com/WinchainOS/Winchain/whisper/whisperv6"
)

var (
	dumpConfigCommand = cli.Command{
		Action:      utils.MigrateFlags(dumpConfig),
		Name:        "dumpconfig",
		Usage:       "Show configuration values",
		ArgsUsage:   "",
		Flags:       append(append(nodeFlags, rpcFlags...), whisperFlags...),
		Category:    "MISCELLANEOUS COMMANDS",
		Description: `The dumpconfig command shows configuration values.`,
	}

	configFileFlag = cli.StringFlag{
		Name:  "config",
		Usage: "TOML configuration file",
	}
)

// These settings ensure that TOML keys use the same names as Go struct fields.
var tomlSettings = toml.Config{
	NormFieldName: func(rt reflect.Type, key string) string {
		return key
	},
	FieldToKey: func(rt reflect.Type, field string) string {
		return field
	},
	MissingField: func(rt reflect.Type, field string) error {
		link := ""
		if unicode.IsUpper(rune(rt.Name()[0])) && rt.PkgPath() != "main" {
			link = fmt.Sprintf(", see https://godoc.org/%s#%s for available fields", rt.PkgPath(), rt.Name())
		}
		return fmt.Errorf("field '%s' is not defined in %s%s", field, rt.String(), link)
	},
}

type ethstatsConfig struct {
	URL string `toml:",omitempty"`
}

type gethConfig struct {
	Eth       eth.Config
	Shh       whisper.Config
	Node      node.Config
	Ethstats  ethstatsConfig
	Dashboard dashboard.Config
}

func loadConfig(file string, cfg *gethConfig) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	err = tomlSettings.NewDecoder(bufio.NewReader(f)).Decode(cfg)
	// Add file name to errors that have a line number.
	if _, ok := err.(*toml.LineError); ok {
		err = errors.New(file + ", " + err.Error())
	}
	return err
}
// 默认的Node配置
func defaultNodeConfig() node.Config {
	cfg := node.DefaultConfig
	cfg.Name = clientIdentifier
	cfg.Version = params.VersionWithCommit(gitCommit)
	//TODO【Winchain】================the part had been changed==================
	cfg.ChainName=winchainClientIdentifier
	cfg.ChainVersion=params.WinchainVersionWithCommit(gitCommit)
	cfg.HTTPModules = append(cfg.HTTPModules, "eth", "shh")
	cfg.WSModules = append(cfg.WSModules, "eth", "shh")
	cfg.IPCPath = "geth.ipc"
	return cfg
}

func makeConfigNode(ctx *cli.Context) (*node.Node, gethConfig) {
	// Load defaults.
	cfg := gethConfig{
		Eth:       eth.DefaultConfig,
		Shh:       whisper.DefaultConfig,
		Node:      defaultNodeConfig(),//默认node配置
		Dashboard: dashboard.DefaultConfig,
	}

	// Load config file.
	if file := ctx.GlobalString(configFileFlag.Name); file != "" {
		if err := loadConfig(file, &cfg); err != nil {
			utils.Fatalf("%v", err)
		}
	}

	// Apply flags.
	utils.SetNodeConfig(ctx, &cfg.Node)
	stack, err := node.New(&cfg.Node)
	if err != nil {
		utils.Fatalf("Failed to create the protocol stack: %v", err)
	}
	utils.SetEthConfig(ctx, stack, &cfg.Eth)
	if ctx.GlobalIsSet(utils.EthStatsURLFlag.Name) {
		cfg.Ethstats.URL = ctx.GlobalString(utils.EthStatsURLFlag.Name)
	}

	utils.SetShhConfig(ctx, stack, &cfg.Shh)
	utils.SetDashboardConfig(ctx, &cfg.Dashboard)

	//TODO【Winchain】================the part had been changed==================
	cfg.Eth.RaftMode = ctx.GlobalBool(utils.RaftModeFlag.Name)

	return stack, cfg
}

// enableWhisper returns true in case one of the whisper flags is set.
func enableWhisper(ctx *cli.Context) bool {
	for _, flag := range whisperFlags {
		if ctx.GlobalIsSet(flag.GetName()) {
			return true
		}
	}
	return false
}
// 创建全节点
func makeFullNode(ctx *cli.Context) *node.Node {
	// 根据命令行参数和一些特殊的配置来创建一个node
	stack, cfg := makeConfigNode(ctx)

	// 把eth的服务注册到这个节点上面。 eth服务是以太坊的主要的服务。 是以太坊功能的提供者。
	ethChan := utils.RegisterEthService(stack, &cfg.Eth)

	//TODO【Winchain】================the part had been changed==================
	if ctx.GlobalBool(utils.RaftModeFlag.Name) {
		RegisterRaftService(stack, ctx, cfg, ethChan)
	} else if ctx.GlobalBool(utils.DashboardEnabledFlag.Name) {
		utils.RegisterDashboardService(stack, &cfg.Dashboard, gitCommit)
	}
	// Whisper must be explicitly enabled by specifying at least 1 whisper flag or in dev mode
	// Whisper是一个新的模块，用来进行加密通讯的功能。 需要显式的提供参数来启用，或者是处于开发模式。
	shhEnabled := enableWhisper(ctx)
	shhAutoEnabled := !ctx.GlobalIsSet(utils.WhisperEnabledFlag.Name) && ctx.GlobalIsSet(utils.DeveloperFlag.Name)
	if shhEnabled || shhAutoEnabled {
		if ctx.GlobalIsSet(utils.WhisperMaxMessageSizeFlag.Name) {
			cfg.Shh.MaxMessageSize = uint32(ctx.Int(utils.WhisperMaxMessageSizeFlag.Name))
		}
		if ctx.GlobalIsSet(utils.WhisperMinPOWFlag.Name) {
			cfg.Shh.MinimumAcceptedPOW = ctx.Float64(utils.WhisperMinPOWFlag.Name)
		}
		// 注册Shh服务
		utils.RegisterShhService(stack, &cfg.Shh)
	}

	// Add the Ethereum Stats daemon if requested.
	if cfg.Ethstats.URL != "" {
		// 注册 以太坊的状态服务。 默认情况下是没有启动的。
		utils.RegisterEthStatsService(stack, cfg.Ethstats.URL)
	}
	return stack
}

// dumpConfig is the dumpconfig command.
func dumpConfig(ctx *cli.Context) error {
	_, cfg := makeConfigNode(ctx)
	comment := ""

	if cfg.Eth.Genesis != nil {
		cfg.Eth.Genesis = nil
		comment += "# Note: this config doesn't contain the genesis block.\n\n"
	}

	out, err := tomlSettings.Marshal(&cfg)
	if err != nil {
		return err
	}
	io.WriteString(os.Stdout, comment)
	os.Stdout.Write(out)
	return nil
}

//TODO【Winchain】================the part had been changed==================
func RegisterRaftService(stack *node.Node, ctx *cli.Context, cfg gethConfig, ethChan <-chan *eth.Ethereum) {
	blockTimeMillis := ctx.GlobalInt(utils.RaftBlockTimeFlag.Name)
	datadir := ctx.GlobalString(utils.DataDirFlag.Name)
	joinExistingId := ctx.GlobalInt(utils.RaftJoinExistingFlag.Name)

	raftPort := uint16(ctx.GlobalInt(utils.RaftPortFlag.Name))

	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		privkey := cfg.Node.NodeKey()
		strId := discover.PubkeyID(&privkey.PublicKey).String()
		blockTimeNanos := time.Duration(blockTimeMillis) * time.Millisecond
		peers := cfg.Node.StaticNodes()

		var myId uint16
		var joinExisting bool

		if joinExistingId > 0 {
			myId = uint16(joinExistingId)
			joinExisting = true
		} else if len(peers) == 0 {
			utils.Fatalf("Raft-based consensus requires either (1) an initial peers list (in static-nodes.json) including this enode hash (%v), or (2) the flag --raftjoinexisting RAFT_ID, where RAFT_ID has been issued by an existing cluster member calling `raft.addPeer(ENODE_ID)` with an enode ID containing this node's enode hash.", strId)
		} else {
			peerIds := make([]string, len(peers))

			for peerIdx, peer := range peers {
				if !peer.HasRaftPort() {
					utils.Fatalf("raftport querystring parameter not specified in static-node enode ID: %v. please check your static-nodes.json file.", peer.String())
				}

				peerId := peer.ID.String()
				peerIds[peerIdx] = peerId
				if peerId == strId {
					myId = uint16(peerIdx) + 1
				}
			}

			if myId == 0 {
				utils.Fatalf("failed to find local enode ID (%v) amongst peer IDs: %v", strId, peerIds)
			}
		}

		ethereum := <-ethChan

		return raft.New(ctx, ethereum.ChainConfig(), myId, raftPort, joinExisting, blockTimeNanos, ethereum, peers, datadir)
	}); err != nil {
		utils.Fatalf("Failed to register the Raft service: %v", err)
	}
}
