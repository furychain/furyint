// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	client "github.com/cosmos/cosmos-sdk/client"

	coretypes "github.com/tendermint/tendermint/rpc/core/types"

	cosmosaccount "github.com/ignite/cli/ignite/pkg/cosmosaccount"

	cosmosclient "github.com/furychain/cosmosclient/cosmosclient"

	mock "github.com/stretchr/testify/mock"

	rollapptypes "github.com/furychain/furya/x/rollapp/types"

	sequencertypes "github.com/furychain/furya/x/sequencer/types"

	types "github.com/cosmos/cosmos-sdk/types"
)

// CosmosClient is an autogenerated mock type for the CosmosClient type
type CosmosClient struct {
	mock.Mock
}

// BroadcastTx provides a mock function with given fields: accountName, msgs
func (_m *CosmosClient) BroadcastTx(accountName string, msgs ...types.Msg) (cosmosclient.Response, error) {
	_va := make([]interface{}, len(msgs))
	for _i := range msgs {
		_va[_i] = msgs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, accountName)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 cosmosclient.Response
	if rf, ok := ret.Get(0).(func(string, ...types.Msg) cosmosclient.Response); ok {
		r0 = rf(accountName, msgs...)
	} else {
		r0 = ret.Get(0).(cosmosclient.Response)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...types.Msg) error); ok {
		r1 = rf(accountName, msgs...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Context provides a mock function with given fields:
func (_m *CosmosClient) Context() client.Context {
	ret := _m.Called()

	var r0 client.Context
	if rf, ok := ret.Get(0).(func() client.Context); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(client.Context)
	}

	return r0
}

// EventListenerQuit provides a mock function with given fields:
func (_m *CosmosClient) EventListenerQuit() <-chan struct{} {
	ret := _m.Called()

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func() <-chan struct{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	return r0
}

// GetAccount provides a mock function with given fields: accountName
func (_m *CosmosClient) GetAccount(accountName string) (cosmosaccount.Account, error) {
	ret := _m.Called(accountName)

	var r0 cosmosaccount.Account
	if rf, ok := ret.Get(0).(func(string) cosmosaccount.Account); ok {
		r0 = rf(accountName)
	} else {
		r0 = ret.Get(0).(cosmosaccount.Account)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(accountName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRollappClient provides a mock function with given fields:
func (_m *CosmosClient) GetRollappClient() rollapptypes.QueryClient {
	ret := _m.Called()

	var r0 rollapptypes.QueryClient
	if rf, ok := ret.Get(0).(func() rollapptypes.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(rollapptypes.QueryClient)
		}
	}

	return r0
}

// GetSequencerClient provides a mock function with given fields:
func (_m *CosmosClient) GetSequencerClient() sequencertypes.QueryClient {
	ret := _m.Called()

	var r0 sequencertypes.QueryClient
	if rf, ok := ret.Get(0).(func() sequencertypes.QueryClient); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sequencertypes.QueryClient)
		}
	}

	return r0
}

// StartEventListener provides a mock function with given fields:
func (_m *CosmosClient) StartEventListener() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopEventListener provides a mock function with given fields:
func (_m *CosmosClient) StopEventListener() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SubscribeToEvents provides a mock function with given fields: ctx, subscriber, query, outCapacity
func (_m *CosmosClient) SubscribeToEvents(ctx context.Context, subscriber string, query string, outCapacity ...int) (<-chan coretypes.ResultEvent, error) {
	_va := make([]interface{}, len(outCapacity))
	for _i := range outCapacity {
		_va[_i] = outCapacity[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, subscriber, query)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 <-chan coretypes.ResultEvent
	if rf, ok := ret.Get(0).(func(context.Context, string, string, ...int) <-chan coretypes.ResultEvent); ok {
		r0 = rf(ctx, subscriber, query, outCapacity...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan coretypes.ResultEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, ...int) error); ok {
		r1 = rf(ctx, subscriber, query, outCapacity...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewCosmosClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewCosmosClient creates a new instance of CosmosClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCosmosClient(t mockConstructorTestingTNewCosmosClient) *CosmosClient {
	mock := &CosmosClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
