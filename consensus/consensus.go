// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package consensus implements different Ethereum consensus engines.
// 共识引擎描述
// 在CPU挖矿部分，CpuAgent的mine函数，执行挖矿操作的时候调用了self.engine.Seal函数。
// 这里的engine是就是共识引擎。Seal为其中很重要的一个接口。
// 它实现了nonce值的寻找和hash的计算。并且该函数是保证共识并且不能伪造的一个重要的函数。
// 再PoW共识算法中，Seal函数实现了工作证明。该部分源码在consensus/ethhash下。
package consensus

import (
	"math/big"
	"github.com/WinchainOS/Winchain/common"
	"github.com/WinchainOS/Winchain/core/state"
	"github.com/WinchainOS/Winchain/core/types"
	"github.com/WinchainOS/Winchain/params"
	"github.com/WinchainOS/Winchain/rpc"
	"github.com/WinchainOS/Winchain/p2p"
)

// ChainReader defines a small collection of methods needed to access the local
// blockchain during header and/or uncle verification.
type ChainReader interface {
	// Config retrieves the blockchain's chain configuration.
	Config() *params.ChainConfig

	// CurrentHeader retrieves the current header from the local chain.
	CurrentHeader() *types.Header

	// GetHeader retrieves a block header from the database by hash and number.
	GetHeader(hash common.Hash, number uint64) *types.Header

	// GetHeaderByNumber retrieves a block header from the database by number.
	GetHeaderByNumber(number uint64) *types.Header

	// GetHeaderByHash retrieves a block header from the database by its hash.
	GetHeaderByHash(hash common.Hash) *types.Header

	// GetBlock retrieves a block from the database by hash and number.
	GetBlock(hash common.Hash, number uint64) *types.Block
}

// Engine is an algorithm agnostic consensus engine.
// 在CPU挖矿部分，CpuAgent的mine函数，执行挖矿操作的时候调用了self.engine.Seal函数。
// 这里的engine是就是共识引擎。Seal为其中很重要的一个接口。
// 它实现了nonce值的寻找和hash的计算。并且该函数是保证共识并且不能伪造的一个重要的函数。
// 再PoW共识算法中，Seal函数实现了工作证明。该部分源码在consensus/ethhash下。
type Engine interface {
	// Author retrieves the Ethereum address of the account that minted the given
	// block, which may be different from the header's coinbase if a consensus
	// engine is based on signatures.
	// 获取区块挖掘者, 即coinbase
	Author(header *types.Header) (common.Address, error)

	// VerifyHeader checks whether a header conforms to the consensus rules of a
	// given engine. Verifying the seal may be done optionally here, or explicitly
	// via the VerifySeal method.
	// VerifyHeader 用于校验区块头，通过共识规则来校验，验证区块可以在这里进行也科通通过VerifySeal方法
	VerifyHeader(chain ChainReader, header *types.Header, seal bool) error

	// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers
	// concurrently. The method returns a quit channel to abort the operations and
	// a results channel to retrieve the async verifications (the order is that of
	// the input slice).
	// VerifyHeaders与VerifyHeader相似,同时这个用于批量操作校验头。这个方法返回一个退出信号
	// 用于终止操作，用于异步校验。
	VerifyHeaders(chain ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error)

	// VerifyUncles verifies that the given block's uncles conform to the consensus
	// rules of a given engine.
	// VerifyUncles 用于校验叔块以符合共识引擎的规则
	VerifyUncles(chain ChainReader, block *types.Block) error

	// VerifySeal checks whether the crypto seal on a header is valid according to
	// the consensus rules of the given engine.
	// VerifySeal根据共识算法的规则校验区块头
	VerifySeal(chain ChainReader, header *types.Header) error

	// Prepare initializes the consensus fields of a block header according to the
	// rules of a particular engine. The changes are executed inline.
	// Prepare 用于初始化区块头的共识字段根据共识引擎。这些改变都是内联执行的。
	Prepare(chain ChainReader, header *types.Header) error

	// Finalize runs any post-transaction state modifications (e.g. block rewards)
	// and assembles the final block.
	// Note: The block header and state database might be updated to reflect any
	// consensus rules that happen at finalization (e.g. block rewards).
	// Finalize 完成所有的状态修改，并最终组装成块。
	// 区块头和状态数据库在最终确认的时候可以被更新使之符合共识规则。
	Finalize(chain ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction,
		uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error)

	// Seal generates a new sealing request for the given input block and pushes
	// the result into the given channel.
	//
	// Note, the method returns immediately and will send the result async. More
	// than one result may also be returned depending on the consensus algorithm.
	// Seal 根据输入区块打包生产一个新的区块
	Seal(chain ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error

	// SealHash returns the hash of a block prior to it being sealed.
	SealHash(header *types.Header) common.Hash

	// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
	// that a new block should have.
	// CalcDifficulty 是难度调整算法，它返回新的区块的难度值。
	CalcDifficulty(chain ChainReader, time uint64, parent *types.Header) *big.Int

	// APIs returns the RPC APIs this consensus engine provides.
	// APIs 返回由共识引擎提供的RPC APIs
	APIs(chain ChainReader) []rpc.API

	// Close terminates any background threads maintained by the consensus engine.
	Close() error

	//TODO【Winchain】================the part had been changed==================
	// Protocol returns the protocol for this consensus
	Protocol() Protocol
}

//TODO【Winchain】================the part had been changed==================
// Handler should be implemented is the consensus needs to handle and send peer's message
type Handler interface {
	// NewChainHead handles a new head block comes
	NewChainHead() error

	// HandleMsg handles a message from peer
	HandleMsg(address common.Address, data p2p.Msg) (bool, error)

	// SetBroadcaster sets the broadcaster to send message to peers
	SetBroadcaster(Broadcaster)
}

// PoW is a consensus engine based on proof-of-work.
type PoW interface {
	Engine

	// Hashrate returns the current mining hashrate of a PoW consensus engine.
	Hashrate() float64
}

//TODO【Winchain】================the part had been changed==================
// Istanbul is a consensus engine to avoid byzantine failure
type Istanbul interface {
	Engine

	// Start starts the engine
	Start(chain ChainReader, currentBlock func() *types.Block, hasBadBlock func(hash common.Hash) bool) error

	// Stop stops the engine
	Stop() error
}

//TODO【Winchain】================the part had been changed==================
type Wpoa interface {
	Engine

	ValidateTransaction(chain ChainReader, block *types.Block, tx *types.Transaction) error
}