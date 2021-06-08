// Code generated by mockery v1.0.0. DO NOT EDIT.

package blockchainmocks

import (
	config "github.com/kaleido-io/firefly/internal/config"
	blockchain "github.com/kaleido-io/firefly/pkg/blockchain"

	context "context"

	fftypes "github.com/kaleido-io/firefly/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// Plugin is an autogenerated mock type for the Plugin type
type Plugin struct {
	mock.Mock
}

// Capabilities provides a mock function with given fields:
func (_m *Plugin) Capabilities() *blockchain.Capabilities {
	ret := _m.Called()

	var r0 *blockchain.Capabilities
	if rf, ok := ret.Get(0).(func() *blockchain.Capabilities); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*blockchain.Capabilities)
		}
	}

	return r0
}

// Init provides a mock function with given fields: ctx, prefix, callbacks
func (_m *Plugin) Init(ctx context.Context, prefix config.Prefix, callbacks blockchain.Callbacks) error {
	ret := _m.Called(ctx, prefix, callbacks)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, config.Prefix, blockchain.Callbacks) error); ok {
		r0 = rf(ctx, prefix, callbacks)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitPrefix provides a mock function with given fields: prefix
func (_m *Plugin) InitPrefix(prefix config.Prefix) {
	_m.Called(prefix)
}

// Name provides a mock function with given fields:
func (_m *Plugin) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *Plugin) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubmitBatchPin provides a mock function with given fields: ctx, ledgerID, identity, batch
func (_m *Plugin) SubmitBatchPin(ctx context.Context, ledgerID *fftypes.UUID, identity *fftypes.Identity, batch *blockchain.BatchPin) (string, error) {
	ret := _m.Called(ctx, ledgerID, identity, batch)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, *fftypes.Identity, *blockchain.BatchPin) string); ok {
		r0 = rf(ctx, ledgerID, identity, batch)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, *fftypes.Identity, *blockchain.BatchPin) error); ok {
		r1 = rf(ctx, ledgerID, identity, batch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyIdentitySyntax provides a mock function with given fields: ctx, identity
func (_m *Plugin) VerifyIdentitySyntax(ctx context.Context, identity *fftypes.Identity) error {
	ret := _m.Called(ctx, identity)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Identity) error); ok {
		r0 = rf(ctx, identity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
