package logic

import (
	"encoding/json"
	"github.com/zecrey-labs/zecrey-legend/common/tree"
	"github.com/zecrey-labs/zecrey-legend/common/util"
	"github.com/zeromicro/go-zero/core/logx"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func DefaultBlockHeader() StorageStoredBlockInfo {
	var (
		pendingOnchainOperationsHash [32]byte
		stateRoot                    [32]byte
		commitment                   [32]byte
	)
	copy(pendingOnchainOperationsHash[:], common.FromHex(util.EmptyStringKeccak)[:])
	copy(stateRoot[:], tree.NilStateRoot[:])
	copy(commitment[:], common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000000")[:])
	return StorageStoredBlockInfo{
		BlockSize:                    0,
		BlockNumber:                  0,
		PriorityOperations:           0,
		PendingOnchainOperationsHash: pendingOnchainOperationsHash,
		Timestamp:                    big.NewInt(0),
		StateRoot:                    stateRoot,
		Commitment:                   commitment,
	}
}

/*
	ConvertBlocksForCommitToCommitBlockInfos: helper function to convert blocks to commit block infos
*/
func ConvertBlocksForCommitToCommitBlockInfos(oBlocks []*BlockForCommit) (commitBlocks []ZecreyLegendCommitBlockInfo, err error) {
	for _, oBlock := range oBlocks {
		var newStateRoot [32]byte
		var pubDataOffsets []uint32
		copy(newStateRoot[:], common.FromHex(oBlock.StateRoot)[:])
		err = json.Unmarshal([]byte(oBlock.PublicDataOffsets), &pubDataOffsets)
		if err != nil {
			logx.Errorf("[ConvertBlocksForCommitToCommitBlockInfos] unable to unmarshal: %s", err.Error())
			return nil, err
		}
		commitBlock := ZecreyLegendCommitBlockInfo{
			NewStateRoot:      newStateRoot,
			PublicData:        common.FromHex(oBlock.PublicData),
			Timestamp:         big.NewInt(oBlock.Timestamp),
			PublicDataOffsets: pubDataOffsets,
			BlockNumber:       uint32(oBlock.BlockHeight),
			BlockSize:         oBlock.BlockSize,
		}
		commitBlocks = append(commitBlocks, commitBlock)
	}
	return commitBlocks, nil
}

func ConvertBlocksToVerifyAndExecuteBlockInfos(oBlocks []*Block) (verifyAndExecuteBlocks []ZecreyLegendVerifyBlockInfo, err error) {
	for _, oBlock := range oBlocks {
		var pendingOnChainOpsPubData [][]byte
		if oBlock.PendingOnChainOperationsPubData != "" {
			err = json.Unmarshal([]byte(oBlock.PendingOnChainOperationsPubData), &pendingOnChainOpsPubData)
			if err != nil {
				logx.Errorf("[ConvertBlocksToVerifyAndExecuteBlockInfos] unable to unmarshal pending pub data: %s", err.Error())
				return nil, err
			}
		}
		verifyAndExecuteBlock := ZecreyLegendVerifyBlockInfo{
			BlockHeader:              util.ConstructStoredBlockInfo(oBlock),
			PendingOnchainOpsPubData: pendingOnChainOpsPubData,
		}
		verifyAndExecuteBlocks = append(verifyAndExecuteBlocks, verifyAndExecuteBlock)
	}
	return verifyAndExecuteBlocks, nil
}
