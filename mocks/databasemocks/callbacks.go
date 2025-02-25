// Code generated by mockery v1.0.0. DO NOT EDIT.

package databasemocks

import (
	database "github.com/hyperledger/firefly/pkg/database"
	fftypes "github.com/hyperledger/firefly/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// Callbacks is an autogenerated mock type for the Callbacks type
type Callbacks struct {
	mock.Mock
}

// HashCollectionNSEvent provides a mock function with given fields: resType, eventType, ns, hash
func (_m *Callbacks) HashCollectionNSEvent(resType database.HashCollectionNS, eventType fftypes.ChangeEventType, ns string, hash *fftypes.Bytes32) {
	_m.Called(resType, eventType, ns, hash)
}

// OrderedCollectionEvent provides a mock function with given fields: resType, eventType, sequence
func (_m *Callbacks) OrderedCollectionEvent(resType database.OrderedCollection, eventType fftypes.ChangeEventType, sequence int64) {
	_m.Called(resType, eventType, sequence)
}

// OrderedUUIDCollectionNSEvent provides a mock function with given fields: resType, eventType, ns, id, sequence
func (_m *Callbacks) OrderedUUIDCollectionNSEvent(resType database.OrderedUUIDCollectionNS, eventType fftypes.ChangeEventType, ns string, id *fftypes.UUID, sequence int64) {
	_m.Called(resType, eventType, ns, id, sequence)
}

// UUIDCollectionEvent provides a mock function with given fields: resType, eventType, id
func (_m *Callbacks) UUIDCollectionEvent(resType database.UUIDCollection, eventType fftypes.ChangeEventType, id *fftypes.UUID) {
	_m.Called(resType, eventType, id)
}

// UUIDCollectionNSEvent provides a mock function with given fields: resType, eventType, ns, id
func (_m *Callbacks) UUIDCollectionNSEvent(resType database.UUIDCollectionNS, eventType fftypes.ChangeEventType, ns string, id *fftypes.UUID) {
	_m.Called(resType, eventType, ns, id)
}
