// Copyright © 2020, 2023 Attestant Limited.
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

package http

import (
	"bytes"
	"context"

	"github.com/attestantio/go-eth2-client/api"
	apiv1 "github.com/attestantio/go-eth2-client/api/v1"
)

// NodeSyncing provides the syncing information for the node.
func (s *Service) NodeSyncing(ctx context.Context) (*api.Response[*apiv1.SyncState], error) {
	httpResponse, err := s.get2(ctx, "/eth/v1/node/syncing")
	if err != nil {
		return nil, err
	}

	data, metadata, err := decodeJSONResponse(bytes.NewReader(httpResponse.body), &apiv1.SyncState{})
	if err != nil {
		return nil, err
	}

	return &api.Response[*apiv1.SyncState]{
		Data:     data,
		Metadata: metadata,
	}, nil
}
