// Copyright © 2022 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fftypes

import (
	"context"
	"crypto/sha256"
	"database/sql/driver"
	"encoding/json"

	"github.com/hyperledger/firefly/internal/i18n"
)

type Batch struct {
	ID        *UUID       `json:"id"`
	Namespace string      `json:"namespace"`
	Type      MessageType `json:"type"`
	Node      *UUID       `json:"node,omitempty"`
	Identity
	Group      *Bytes32     `jdon:"group,omitempty"`
	Hash       *Bytes32     `json:"hash"`
	Created    *FFTime      `json:"created"`
	Confirmed  *FFTime      `json:"confirmed"`
	Payload    BatchPayload `json:"payload"`
	PayloadRef string       `json:"payloadRef,omitempty"`
	Blobs      []*Bytes32   `json:"blobs,omitempty"` // only used in-flight
}

type BatchPayload struct {
	TX       TransactionRef `json:"tx"`
	Messages []*Message     `json:"messages"`
	Data     []*Data        `json:"data"`
}

// Value implements sql.Valuer
func (ma BatchPayload) Value() (driver.Value, error) {
	return json.Marshal(&ma)
}

func (ma *BatchPayload) Hash() *Bytes32 {
	b, _ := json.Marshal(&ma)
	var b32 Bytes32 = sha256.Sum256(b)
	return &b32
}

// Scan implements sql.Scanner
func (ma *BatchPayload) Scan(src interface{}) error {
	switch src := src.(type) {
	case nil:
		return nil

	case []byte:
		return json.Unmarshal(src, &ma)

	case string:
		if src == "" {
			return nil
		}
		return json.Unmarshal([]byte(src), &ma)

	default:
		return i18n.NewError(context.Background(), i18n.MsgScanFailed, src, ma)
	}

}

func (b *Batch) Manifest() *Manifest {
	if b == nil {
		return nil
	}
	tm := &Manifest{
		Messages: make([]MessageRef, len(b.Payload.Messages)),
		Data:     make([]DataRef, len(b.Payload.Data)),
	}
	for i, m := range b.Payload.Messages {
		tm.Messages[i].ID = m.Header.ID
		tm.Messages[i].Hash = m.Hash
	}
	for i, d := range b.Payload.Data {
		tm.Data[i].ID = d.ID
		tm.Data[i].Hash = d.Hash
	}
	return tm
}
