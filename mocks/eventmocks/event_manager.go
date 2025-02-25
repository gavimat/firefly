// Code generated by mockery v1.0.0. DO NOT EDIT.

package eventmocks

import (
	context "context"

	blockchain "github.com/hyperledger/firefly/pkg/blockchain"

	dataexchange "github.com/hyperledger/firefly/pkg/dataexchange"

	fftypes "github.com/hyperledger/firefly/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"

	system "github.com/hyperledger/firefly/internal/events/system"

	tokens "github.com/hyperledger/firefly/pkg/tokens"
)

// EventManager is an autogenerated mock type for the EventManager type
type EventManager struct {
	mock.Mock
}

// AddSystemEventListener provides a mock function with given fields: ns, el
func (_m *EventManager) AddSystemEventListener(ns string, el system.EventListener) error {
	ret := _m.Called(ns, el)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, system.EventListener) error); ok {
		r0 = rf(ns, el)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BLOBReceived provides a mock function with given fields: dx, peerID, hash, size, payloadRef
func (_m *EventManager) BLOBReceived(dx dataexchange.Plugin, peerID string, hash fftypes.Bytes32, size int64, payloadRef string) error {
	ret := _m.Called(dx, peerID, hash, size, payloadRef)

	var r0 error
	if rf, ok := ret.Get(0).(func(dataexchange.Plugin, string, fftypes.Bytes32, int64, string) error); ok {
		r0 = rf(dx, peerID, hash, size, payloadRef)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BatchPinComplete provides a mock function with given fields: bi, batch, signingIdentity
func (_m *EventManager) BatchPinComplete(bi blockchain.Plugin, batch *blockchain.BatchPin, signingIdentity string) error {
	ret := _m.Called(bi, batch, signingIdentity)

	var r0 error
	if rf, ok := ret.Get(0).(func(blockchain.Plugin, *blockchain.BatchPin, string) error); ok {
		r0 = rf(bi, batch, signingIdentity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ChangeEvents provides a mock function with given fields:
func (_m *EventManager) ChangeEvents() chan<- *fftypes.ChangeEvent {
	ret := _m.Called()

	var r0 chan<- *fftypes.ChangeEvent
	if rf, ok := ret.Get(0).(func() chan<- *fftypes.ChangeEvent); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan<- *fftypes.ChangeEvent)
		}
	}

	return r0
}

// ContractEvent provides a mock function with given fields: event
func (_m *EventManager) ContractEvent(event *blockchain.ContractEvent) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(*blockchain.ContractEvent) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateUpdateDurableSubscription provides a mock function with given fields: ctx, subDef, mustNew
func (_m *EventManager) CreateUpdateDurableSubscription(ctx context.Context, subDef *fftypes.Subscription, mustNew bool) error {
	ret := _m.Called(ctx, subDef, mustNew)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Subscription, bool) error); ok {
		r0 = rf(ctx, subDef, mustNew)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteDurableSubscription provides a mock function with given fields: ctx, subDef
func (_m *EventManager) DeleteDurableSubscription(ctx context.Context, subDef *fftypes.Subscription) error {
	ret := _m.Called(ctx, subDef)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.Subscription) error); ok {
		r0 = rf(ctx, subDef)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeletedSubscriptions provides a mock function with given fields:
func (_m *EventManager) DeletedSubscriptions() chan<- *fftypes.UUID {
	ret := _m.Called()

	var r0 chan<- *fftypes.UUID
	if rf, ok := ret.Get(0).(func() chan<- *fftypes.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan<- *fftypes.UUID)
		}
	}

	return r0
}

// MessageReceived provides a mock function with given fields: dx, peerID, data
func (_m *EventManager) MessageReceived(dx dataexchange.Plugin, peerID string, data []byte) (string, error) {
	ret := _m.Called(dx, peerID, data)

	var r0 string
	if rf, ok := ret.Get(0).(func(dataexchange.Plugin, string, []byte) string); ok {
		r0 = rf(dx, peerID, data)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dataexchange.Plugin, string, []byte) error); ok {
		r1 = rf(dx, peerID, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewEvents provides a mock function with given fields:
func (_m *EventManager) NewEvents() chan<- int64 {
	ret := _m.Called()

	var r0 chan<- int64
	if rf, ok := ret.Get(0).(func() chan<- int64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan<- int64)
		}
	}

	return r0
}

// NewPins provides a mock function with given fields:
func (_m *EventManager) NewPins() chan<- int64 {
	ret := _m.Called()

	var r0 chan<- int64
	if rf, ok := ret.Get(0).(func() chan<- int64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan<- int64)
		}
	}

	return r0
}

// NewSubscriptions provides a mock function with given fields:
func (_m *EventManager) NewSubscriptions() chan<- *fftypes.UUID {
	ret := _m.Called()

	var r0 chan<- *fftypes.UUID
	if rf, ok := ret.Get(0).(func() chan<- *fftypes.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan<- *fftypes.UUID)
		}
	}

	return r0
}

// OperationUpdate provides a mock function with given fields: plugin, operationID, txState, errorMessage, opOutput
func (_m *EventManager) OperationUpdate(plugin fftypes.Named, operationID *fftypes.UUID, txState fftypes.OpStatus, errorMessage string, opOutput fftypes.JSONObject) error {
	ret := _m.Called(plugin, operationID, txState, errorMessage, opOutput)

	var r0 error
	if rf, ok := ret.Get(0).(func(fftypes.Named, *fftypes.UUID, fftypes.OpStatus, string, fftypes.JSONObject) error); ok {
		r0 = rf(plugin, operationID, txState, errorMessage, opOutput)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Start provides a mock function with given fields:
func (_m *EventManager) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscriptionUpdates provides a mock function with given fields:
func (_m *EventManager) SubscriptionUpdates() chan<- *fftypes.UUID {
	ret := _m.Called()

	var r0 chan<- *fftypes.UUID
	if rf, ok := ret.Get(0).(func() chan<- *fftypes.UUID); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan<- *fftypes.UUID)
		}
	}

	return r0
}

// TokenPoolCreated provides a mock function with given fields: ti, pool
func (_m *EventManager) TokenPoolCreated(ti tokens.Plugin, pool *tokens.TokenPool) error {
	ret := _m.Called(ti, pool)

	var r0 error
	if rf, ok := ret.Get(0).(func(tokens.Plugin, *tokens.TokenPool) error); ok {
		r0 = rf(ti, pool)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TokensTransferred provides a mock function with given fields: ti, transfer
func (_m *EventManager) TokensTransferred(ti tokens.Plugin, transfer *tokens.TokenTransfer) error {
	ret := _m.Called(ti, transfer)

	var r0 error
	if rf, ok := ret.Get(0).(func(tokens.Plugin, *tokens.TokenTransfer) error); ok {
		r0 = rf(ti, transfer)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TransferResult provides a mock function with given fields: dx, trackingID, status, update
func (_m *EventManager) TransferResult(dx dataexchange.Plugin, trackingID string, status fftypes.OpStatus, update fftypes.TransportStatusUpdate) error {
	ret := _m.Called(dx, trackingID, status, update)

	var r0 error
	if rf, ok := ret.Get(0).(func(dataexchange.Plugin, string, fftypes.OpStatus, fftypes.TransportStatusUpdate) error); ok {
		r0 = rf(dx, trackingID, status, update)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitStop provides a mock function with given fields:
func (_m *EventManager) WaitStop() {
	_m.Called()
}
