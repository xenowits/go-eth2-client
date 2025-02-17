// Copyright © 2021 Attestant Limited.
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
	"context"
	"encoding/json"

	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

type voluntaryExitPoolJSON struct {
	Data []*phase0.SignedVoluntaryExit `json:"data"`
}

// VoluntaryExitPool obtains the voluntary exit pool.
func (s *Service) VoluntaryExitPool(ctx context.Context) ([]*phase0.SignedVoluntaryExit, error) {
	respBodyReader, err := s.get(ctx, "/eth/v1/beacon/pool/voluntary_exits")
	if err != nil {
		return nil, errors.Wrap(err, "failed to request voluntary exit pool")
	}
	if respBodyReader == nil {
		return nil, errors.New("failed to obtain voluntary exit pool")
	}

	var voluntaryExitPoolJSON voluntaryExitPoolJSON
	if err := json.NewDecoder(respBodyReader).Decode(&voluntaryExitPoolJSON); err != nil {
		return nil, errors.Wrap(err, "failed to parse voluntary exit pool")
	}

	// Ensure the data returned to us is as expected given our input.
	if voluntaryExitPoolJSON.Data == nil {
		return nil, errors.New("voluntary exit pool not returned")
	}

	return voluntaryExitPoolJSON.Data, nil
}
