// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package blockchain

import mock "github.com/stretchr/testify/mock"

// MockEvents is an autogenerated mock type for the Events type
type MockEvents struct {
	mock.Mock
}

// SequencedBroadcastBatch provides a mock function with given fields: batch, additionalInfo
func (_m *MockEvents) SequencedBroadcastBatch(batch BroadcastBatch, additionalInfo map[string]interface{}) {
	_m.Called(batch, additionalInfo)
}

// TransactionUpdate provides a mock function with given fields: txTrackingID, txState, errorMessage, additionalInfo
func (_m *MockEvents) TransactionUpdate(txTrackingID string, txState TransactionState, errorMessage string, additionalInfo map[string]interface{}) {
	_m.Called(txTrackingID, txState, errorMessage, additionalInfo)
}