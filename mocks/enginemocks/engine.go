// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package enginemocks

import (
	context "context"

	fftypes "github.com/kaleido-io/firefly/internal/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// Engine is an autogenerated mock type for the Engine type
type Engine struct {
	mock.Mock
}

// BroadcastSchemaDefinition provides a mock function with given fields: ctx, author, s
func (_m *Engine) BroadcastSchemaDefinition(ctx context.Context, author string, s *fftypes.Schema) (*fftypes.MessageExpanded, error) {
	ret := _m.Called(ctx, author, s)

	var r0 *fftypes.MessageExpanded
	if rf, ok := ret.Get(0).(func(context.Context, string, *fftypes.Schema) *fftypes.MessageExpanded); ok {
		r0 = rf(ctx, author, s)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.MessageExpanded)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *fftypes.Schema) error); ok {
		r1 = rf(ctx, author, s)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *Engine) Close() {
	_m.Called()
}

// Init provides a mock function with given fields: ctx
func (_m *Engine) Init(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
