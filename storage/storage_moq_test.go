// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package storage

import (
	"context"
	"sync"
)

var (
	lockFirestoreMockGet sync.RWMutex
	lockFirestoreMockSet sync.RWMutex
)

// Ensure, that FirestoreMock does implement Firestore.
// If this is not the case, regenerate this file with moq.
var _ Firestore = &FirestoreMock{}

// FirestoreMock is a mock implementation of Firestore.
//
//     func TestSomethingThatUsesFirestore(t *testing.T) {
//
//         // make and configure a mocked Firestore
//         mockedFirestore := &FirestoreMock{
//             GetFunc: func(ctx context.Context, dataTo interface{}, documentID string) error {
// 	               panic("mock out the Get method")
//             },
//             SetFunc: func(ctx context.Context, documentID string, data interface{}) error {
// 	               panic("mock out the Set method")
//             },
//         }
//
//         // use mockedFirestore in code that requires Firestore
//         // and then make assertions.
//
//     }
type FirestoreMock struct {
	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, dataTo interface{}, documentID string) error

	// SetFunc mocks the Set method.
	SetFunc func(ctx context.Context, documentID string, data interface{}) error

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DataTo is the dataTo argument value.
			DataTo interface{}
			// DocumentID is the documentID argument value.
			DocumentID string
		}
		// Set holds details about calls to the Set method.
		Set []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DocumentID is the documentID argument value.
			DocumentID string
			// Data is the data argument value.
			Data interface{}
		}
	}
}

// Get calls GetFunc.
func (mock *FirestoreMock) Get(ctx context.Context, dataTo interface{}, documentID string) error {
	if mock.GetFunc == nil {
		panic("FirestoreMock.GetFunc: method is nil but Firestore.Get was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		DataTo     interface{}
		DocumentID string
	}{
		Ctx:        ctx,
		DataTo:     dataTo,
		DocumentID: documentID,
	}
	lockFirestoreMockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	lockFirestoreMockGet.Unlock()
	return mock.GetFunc(ctx, dataTo, documentID)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedFirestore.GetCalls())
func (mock *FirestoreMock) GetCalls() []struct {
	Ctx        context.Context
	DataTo     interface{}
	DocumentID string
} {
	var calls []struct {
		Ctx        context.Context
		DataTo     interface{}
		DocumentID string
	}
	lockFirestoreMockGet.RLock()
	calls = mock.calls.Get
	lockFirestoreMockGet.RUnlock()
	return calls
}

// Set calls SetFunc.
func (mock *FirestoreMock) Set(ctx context.Context, documentID string, data interface{}) error {
	if mock.SetFunc == nil {
		panic("FirestoreMock.SetFunc: method is nil but Firestore.Set was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		DocumentID string
		Data       interface{}
	}{
		Ctx:        ctx,
		DocumentID: documentID,
		Data:       data,
	}
	lockFirestoreMockSet.Lock()
	mock.calls.Set = append(mock.calls.Set, callInfo)
	lockFirestoreMockSet.Unlock()
	return mock.SetFunc(ctx, documentID, data)
}

// SetCalls gets all the calls that were made to Set.
// Check the length with:
//     len(mockedFirestore.SetCalls())
func (mock *FirestoreMock) SetCalls() []struct {
	Ctx        context.Context
	DocumentID string
	Data       interface{}
} {
	var calls []struct {
		Ctx        context.Context
		DocumentID string
		Data       interface{}
	}
	lockFirestoreMockSet.RLock()
	calls = mock.calls.Set
	lockFirestoreMockSet.RUnlock()
	return calls
}
