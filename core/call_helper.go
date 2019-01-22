package core

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/WinchainOS/Winchain/common"
	"github.com/WinchainOS/Winchain/consensus/ethash"
	"github.com/WinchainOS/Winchain/core/state"
	"github.com/WinchainOS/Winchain/core/types"
	"github.com/WinchainOS/Winchain/core/vm"
	"github.com/WinchainOS/Winchain/crypto"
	"github.com/WinchainOS/Winchain/ethdb"
	"github.com/WinchainOS/Winchain/params"
)

//TODO【Winchain】================the part had been changed==================
// callHelper makes it easier to do proper calls and use the state transition object.
// It also manages the nonces of the caller and keeps private and public state, which
// can be freely modified outside of the helper.
type callHelper struct {
	db ethdb.Database

	nonces map[common.Address]uint64
	header types.Header
	gp     *GasPool

	PrivateState, PublicState *state.StateDB
}

// TxNonce returns the pending nonce
func (cg *callHelper) TxNonce(addr common.Address) uint64 {
	return cg.nonces[addr]
}

// MakeCall makes does a call to the recipient using the given input. It can switch between private and public
// by setting the private boolean flag. It returns an error if the call failed.
func (cg *callHelper) MakeCall(private bool, key *ecdsa.PrivateKey, to common.Address, input []byte) error {
	var (
		from = crypto.PubkeyToAddress(key.PublicKey)
		err  error
	)

	// TODO(joel): these are just stubbed to the same values as in dual_state_test.go
	cg.header.Number = new(big.Int)
	cg.header.Time = new(big.Int).SetUint64(43)
	cg.header.Difficulty = new(big.Int).SetUint64(1000488)
	cg.header.GasLimit = uint64(4700000)

	//TODO【Winchain】================the part had been changed==================
	signer := types.MakeSigner(params.WinchainTestChainConfig, cg.header.Number)
	tx, err := types.SignTx(types.NewTransaction(cg.TxNonce(from), to, new(big.Int), uint64(1000000), new(big.Int), input), signer, key)
	if err != nil {
		return err
	}
	defer func() { cg.nonces[from]++ }()
	msg, err := tx.AsMessage(signer)
	if err != nil {
		return err
	}

	publicState, _ := cg.PublicState, cg.PrivateState
	//privateState = publicState
	//TODO【Winchain】================the part had been changed==================
	// 不分私有还是公有
	if !private {
	} else {
		//tx.SetPrivate()
	}

	// TODO(joel): can we just pass nil instead of bc?
	//TODO【Winchain】================the part had been changed==================
	bc, _ := NewBlockChain(cg.db, nil, params.WinchainTestChainConfig, ethash.NewFaker(), vm.Config{})
	context := NewEVMContext(msg, &cg.header, bc, &from)
	vmenv := vm.NewEVM(context, publicState, params.WinchainTestChainConfig, vm.Config{})
	_, _, _, err = ApplyMessage(vmenv, msg, cg.gp)
	if err != nil {
		return err
	}
	return nil
}

// MakeCallHelper returns a new callHelper
func MakeCallHelper() *callHelper {
	memdb := ethdb.NewMemDatabase()
	db := state.NewDatabase(memdb)

	publicState, err := state.New(common.Hash{}, db)
	if err != nil {
		panic(err)
	}
	privateState, err := state.New(common.Hash{}, db)
	if err != nil {
		panic(err)
	}
	cg := &callHelper{
		db:           memdb,
		nonces:       make(map[common.Address]uint64),
		gp:           new(GasPool).AddGas(uint64(5000000)),
		PublicState:  publicState,
		PrivateState: privateState,
	}
	return cg
}
