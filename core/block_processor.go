package core

import (
	"fmt"

	"github.com/bnb-chain/zkbnb/core/executor"
	"github.com/bnb-chain/zkbnb/dao/tx"
)

type Processor interface {
	Process(tx *tx.Tx) error
}

type CommitProcessor struct {
	bc *BlockChain
}

func NewCommitProcessor(bc *BlockChain) Processor {
	return &CommitProcessor{
		bc: bc,
	}
}

func (p *CommitProcessor) Process(tx *tx.Tx) error {
	p.bc.setCurrentBlockTimeStamp()
	defer p.bc.resetCurrentBlockTimeStamp()

	executor, err := executor.NewTxExecutor(p.bc, tx)
	if err != nil {
		return fmt.Errorf("new tx executor failed")
	}

	err = executor.Prepare()
	if err != nil {
		return err
	}
	err = executor.VerifyInputs()
	if err != nil {
		return err
	}
	err = executor.ApplyTransaction()
	if err != nil {
		panic(err)
	}
	err = executor.GeneratePubData()
	if err != nil {
		panic(err)
	}
	tx, err = executor.GetExecutedTx()
	if err != nil {
		panic(err)
	}

	p.bc.Statedb.Txs = append(p.bc.Statedb.Txs, tx)

	return nil
}

type WitnessProcessor struct {
	bc *BlockChain
}

func NewWitnessProcessor(bc *BlockChain) Processor {
	return &WitnessProcessor{
		bc: bc,
	}
}

func (p *WitnessProcessor) Process(tx *tx.Tx) error {
	executor, err := executor.NewTxExecutor(p.bc, tx)
	if err != nil {
		return fmt.Errorf("new tx executor failed")
	}
	err = executor.Prepare()
	if err != nil {
		return err
	}
	// TODO
	witness, err := executor.GenerateWitness(0)
	if err != nil {
		return err
	}
	err = executor.ApplyTransaction()
	if err != nil {
		panic(err)
	}
	p.bc.Statedb.Witnesses = append(p.bc.Statedb.Witnesses, witness)
	return nil
}
