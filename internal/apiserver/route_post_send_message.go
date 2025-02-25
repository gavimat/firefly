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

package apiserver

import (
	"context"
	"net/http"

	"github.com/hyperledger/firefly/internal/config"
	"github.com/hyperledger/firefly/internal/i18n"
	"github.com/hyperledger/firefly/internal/oapispec"
	"github.com/hyperledger/firefly/pkg/fftypes"
)

var postSendMessage = &oapispec.Route{
	Name:   "postSendMessage",
	Path:   "namespaces/{ns}/send/message",
	Method: http.MethodPost,
	PathParams: []*oapispec.PathParam{
		{Name: "ns", ExampleFromConf: config.NamespacesDefault, Description: i18n.MsgTBD},
	},
	QueryParams:     nil,
	FilterFactory:   nil,
	Description:     i18n.MsgTBD,
	JSONInputValue:  func() interface{} { return &fftypes.MessageInOut{} },
	JSONInputSchema: func(ctx context.Context) string { return privateSendSchema },
	JSONOutputValue: func() interface{} { return &fftypes.Message{} },
	JSONOutputCodes: []int{http.StatusAccepted}, // Async operation
	JSONHandler: func(r *oapispec.APIRequest) (output interface{}, err error) {
		// This (old) route is always async, and returns the message
		output, err = getOr(r.Ctx).PrivateMessaging().SendMessage(r.Ctx, r.PP["ns"], r.Input.(*fftypes.MessageInOut), false)
		return output, err
	},
	Deprecated: true, // moving to more intutitive route/return structure
}
