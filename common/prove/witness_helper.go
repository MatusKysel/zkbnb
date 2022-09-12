/*
 * Copyright © 2021 ZkBNB Protocol
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package prove

import (
	"fmt"
	bsmt "github.com/bnb-chain/zkbnb-smt"
	"github.com/bnb-chain/zkbnb/tree"
	"github.com/bnb-chain/zkbnb/types"
)

type WitnessHelper struct {
	treeCtx *tree.Context

	accountModel AccountModel

	// Trees
	accountTree   bsmt.SparseMerkleTree
	assetTrees    *[]bsmt.SparseMerkleTree
	liquidityTree bsmt.SparseMerkleTree
	nftTree       bsmt.SparseMerkleTree
}

func NewWitnessHelper(treeCtx *tree.Context, accountTree, liquidityTree, nftTree bsmt.SparseMerkleTree,
	assetTrees *[]bsmt.SparseMerkleTree, accountModel AccountModel) *WitnessHelper {
	return &WitnessHelper{
		treeCtx:       treeCtx,
		accountModel:  accountModel,
		accountTree:   accountTree,
		assetTrees:    assetTrees,
		liquidityTree: liquidityTree,
		nftTree:       nftTree,
	}
}

func (w *WitnessHelper) ConstructTxWitness(oTx *Tx, finalityBlockNr uint64,
) (cryptoTx *TxWitness, err error) {
	switch oTx.TxType {
	case types.TxTypeEmpty:
		return nil, fmt.Errorf("there should be no empty tx")
	default:
		cryptoTx, err = w.FillInTxWitness(oTx, finalityBlockNr)
		if err != nil {
			return nil, err
		}
	}
	return cryptoTx, nil
}

func (w *WitnessHelper) FillInTxWitness(oTx *Tx, finalityBlockNr uint64) (witness *TxWitness, err error) {
	witness.TxType = uint8(oTx.TxType)
	witness.Nonce = oTx.Nonce
	switch oTx.TxType {
	case types.TxTypeRegisterZns:
		return w.constructRegisterZnsTxWitness(witness, oTx)
	case types.TxTypeCreatePair:
		return w.constructCreatePairTxWitness(witness, oTx)
	case types.TxTypeUpdatePairRate:
		return w.constructUpdatePairRateTxWitness(witness, oTx)
	case types.TxTypeDeposit:
		return w.constructDepositTxWitness(witness, oTx)
	case types.TxTypeDepositNft:
		return w.constructDepositNftTxWitness(witness, oTx)
	case types.TxTypeTransfer:
		return w.constructTransferTxWitness(witness, oTx)
	case types.TxTypeSwap:
		return w.constructSwapTxWitness(witness, oTx)
	case types.TxTypeAddLiquidity:
		return w.constructAddLiquidityTxWitness(witness, oTx)
	case types.TxTypeRemoveLiquidity:
		return w.constructRemoveLiquidityTxWitness(witness, oTx)
	case types.TxTypeWithdraw:
		return w.constructWithdrawTxWitness(witness, oTx)
	case types.TxTypeCreateCollection:
		return w.constructCreateCollectionTxWitness(witness, oTx)
	case types.TxTypeMintNft:
		return w.constructMintNftTxWitness(witness, oTx)
	case types.TxTypeTransferNft:
		return w.constructTransferNftTxWitness(witness, oTx)
	case types.TxTypeAtomicMatch:
		return w.constructAtomicMatchTxWitness(witness, oTx)
	case types.TxTypeCancelOffer:
		return w.constructCancelOfferTxWitness(witness, oTx)
	case types.TxTypeWithdrawNft:
		return w.constructWithdrawNftTxWitness(witness, oTx)
	case types.TxTypeFullExit:
		return w.constructFullExitTxWitness(witness, oTx)
	case types.TxTypeFullExitNft:
		return w.constructFullExitNftTxWitness(witness, oTx)
	default:
		return nil, fmt.Errorf("tx type error")
	}
}

func SetFixedAccountArray(proof [][]byte) (res [AccountMerkleLevels][]byte, err error) {
	if len(proof) != AccountMerkleLevels {
		return res, fmt.Errorf("invalid size")
	}
	copy(res[:], proof[:])
	return res, nil
}

func SetFixedAccountAssetArray(proof [][]byte) (res [AssetMerkleLevels][]byte, err error) {
	if len(proof) != AssetMerkleLevels {
		return res, fmt.Errorf("invalid size")
	}
	copy(res[:], proof[:])
	return res, nil
}

func SetFixedLiquidityArray(proof [][]byte) (res [LiquidityMerkleLevels][]byte, err error) {
	if len(proof) != LiquidityMerkleLevels {
		return res, fmt.Errorf("invalid size")
	}
	copy(res[:], proof[:])
	return res, nil
}

func SetFixedNftArray(proof [][]byte) (res [NftMerkleLevels][]byte, err error) {
	if len(proof) != NftMerkleLevels {
		return res, fmt.Errorf("invalid size")
	}
	copy(res[:], proof[:])
	return res, nil
}
