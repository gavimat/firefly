// Copyright © 2021 Kaleido, Inc.
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

package apiserver

import (
	"bytes"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/firefly/mocks/datamocks"
	"github.com/hyperledger/firefly/pkg/fftypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDataBlob(t *testing.T) {
	o, r := newTestAPIServer()
	mdm := &datamocks.Manager{}
	o.On("Data").Return(mdm)
	req := httptest.NewRequest("GET", "/api/v1/namespaces/mynamespace/data/abcd1234/blob", nil)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res := httptest.NewRecorder()

	blobHash := fftypes.NewRandB32()
	mdm.On("DownloadBLOB", mock.Anything, "mynamespace", "abcd1234").
		Return(&fftypes.Blob{
			Hash: blobHash,
			Size: 12345,
		}, ioutil.NopCloser(bytes.NewReader([]byte("hello"))), nil)
	r.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Result().StatusCode)
	b, err := ioutil.ReadAll(res.Body)
	assert.NoError(t, err)
	assert.Equal(t, "hello", string(b))
	assert.Equal(t, "12345", res.Result().Header.Get(fftypes.HTTPHeadersBlobSize))
	assert.Equal(t, blobHash.String(), res.Result().Header.Get(fftypes.HTTPHeadersBlobHashSHA256))
}
