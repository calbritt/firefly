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

package definitions

import (
	"context"

	"github.com/hyperledger/firefly-common/pkg/fftypes"
	"github.com/hyperledger/firefly-common/pkg/log"
	"github.com/hyperledger/firefly/pkg/core"
	"github.com/hyperledger/firefly/pkg/database"
)

func (dh *definitionHandlers) persistFFI(ctx context.Context, ffi *core.FFI) (valid bool, err error) {
	if err := dh.contracts.ValidateFFIAndSetPathnames(ctx, ffi); err != nil {
		log.L(ctx).Warnf("Unable to process FFI %s - validate failed: %s", ffi.ID, err)
		return false, nil
	}

	err = dh.database.UpsertFFI(ctx, ffi)
	if err != nil {
		return false, err
	}

	for _, method := range ffi.Methods {
		err := dh.database.UpsertFFIMethod(ctx, method)
		if err != nil {
			return false, err
		}
	}

	for _, event := range ffi.Events {
		err := dh.database.UpsertFFIEvent(ctx, event)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (dh *definitionHandlers) persistContractAPI(ctx context.Context, api *core.ContractAPI) (valid bool, err error) {
	existing, err := dh.database.GetContractAPIByName(ctx, api.Namespace, api.Name)
	if err != nil {
		return false, err // retryable
	}
	if existing != nil {
		if !api.LocationAndLedgerEquals(existing) {
			return false, nil // not retryable
		}
	}
	err = dh.database.UpsertContractAPI(ctx, api)
	if err != nil {
		if err == database.IDMismatch {
			log.L(ctx).Errorf("Invalid contract API '%s'. ID mismatch with existing record", api.ID)
			return false, nil // not retryable
		}
		log.L(ctx).Errorf("Failed to insert contract API '%s': %s", api.ID, err)
		return false, err // retryable
	}
	return true, nil
}

func (dh *definitionHandlers) handleFFIBroadcast(ctx context.Context, state DefinitionBatchState, msg *core.Message, data core.DataArray, tx *fftypes.UUID) (HandlerResult, error) {
	l := log.L(ctx)
	var broadcast core.FFI
	valid := dh.getSystemBroadcastPayload(ctx, msg, data, &broadcast)
	if valid {
		if validationErr := broadcast.Validate(ctx, true); validationErr != nil {
			l.Warnf("Unable to process contract definition broadcast %s - validate failed: %s", msg.Header.ID, validationErr)
			valid = false
		} else {
			var err error
			broadcast.Message = msg.Header.ID
			valid, err = dh.persistFFI(ctx, &broadcast)
			if err != nil {
				return HandlerResult{Action: ActionRetry}, err
			}
		}
	}

	if !valid {
		l.Warnf("Contract interface rejected id=%s author=%s", broadcast.ID, msg.Header.Author)
		return HandlerResult{Action: ActionReject}, nil
	}

	l.Infof("Contract interface created id=%s author=%s", broadcast.ID, msg.Header.Author)
	state.AddFinalize(func(ctx context.Context) error {
		event := core.NewEvent(core.EventTypeContractInterfaceConfirmed, broadcast.Namespace, broadcast.ID, tx, broadcast.Topic())
		return dh.database.InsertEvent(ctx, event)
	})
	return HandlerResult{Action: ActionConfirm}, nil
}

func (dh *definitionHandlers) handleContractAPIBroadcast(ctx context.Context, state DefinitionBatchState, msg *core.Message, data core.DataArray, tx *fftypes.UUID) (HandlerResult, error) {
	l := log.L(ctx)
	var broadcast core.ContractAPI
	valid := dh.getSystemBroadcastPayload(ctx, msg, data, &broadcast)
	if valid {
		if validateErr := broadcast.Validate(ctx, true); validateErr != nil {
			l.Warnf("Unable to process contract API broadcast %s - validate failed: %s", msg.Header.ID, validateErr)
			valid = false
		} else {
			var err error
			broadcast.Message = msg.Header.ID
			valid, err = dh.persistContractAPI(ctx, &broadcast)
			if err != nil {
				return HandlerResult{Action: ActionRetry}, err
			}
		}
	}

	if !valid {
		l.Warnf("Contract API rejected id=%s author=%s", broadcast.ID, msg.Header.Author)
		return HandlerResult{Action: ActionReject}, nil
	}

	l.Infof("Contract API created id=%s author=%s", broadcast.ID, msg.Header.Author)
	state.AddFinalize(func(ctx context.Context) error {
		event := core.NewEvent(core.EventTypeContractAPIConfirmed, broadcast.Namespace, broadcast.ID, tx, core.SystemTopicDefinitions)
		return dh.database.InsertEvent(ctx, event)
	})
	return HandlerResult{Action: ActionConfirm}, nil
}
