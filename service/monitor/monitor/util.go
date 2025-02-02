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
 */

package monitor

import (
	"encoding/hex"
	"strconv"

	"github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"
	"github.com/ethereum/go-ethereum/common"
)

func ComputeL1TxTxHash(requestId int64, txHash string) string {
	hFunc := mimc.NewMiMC()
	hFunc.Write([]byte(strconv.FormatInt(requestId, 10)))
	hFunc.Write(common.FromHex(txHash))
	return hex.EncodeToString(hFunc.Sum(nil))
}
