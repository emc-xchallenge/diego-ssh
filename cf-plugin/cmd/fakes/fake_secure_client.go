// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/diego-ssh/cf-plugin/cmd"
)

type FakeSecureClient struct {
	NewSessionStub        func() (cmd.SecureSession, error)
	newSessionMutex       sync.RWMutex
	newSessionArgsForCall []struct{}
	newSessionReturns struct {
		result1 cmd.SecureSession
		result2 error
	}
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	closeReturns struct {
		result1 error
	}
}

func (fake *FakeSecureClient) NewSession() (cmd.SecureSession, error) {
	fake.newSessionMutex.Lock()
	fake.newSessionArgsForCall = append(fake.newSessionArgsForCall, struct{}{})
	fake.newSessionMutex.Unlock()
	if fake.NewSessionStub != nil {
		return fake.NewSessionStub()
	} else {
		return fake.newSessionReturns.result1, fake.newSessionReturns.result2
	}
}

func (fake *FakeSecureClient) NewSessionCallCount() int {
	fake.newSessionMutex.RLock()
	defer fake.newSessionMutex.RUnlock()
	return len(fake.newSessionArgsForCall)
}

func (fake *FakeSecureClient) NewSessionReturns(result1 cmd.SecureSession, result2 error) {
	fake.NewSessionStub = nil
	fake.newSessionReturns = struct {
		result1 cmd.SecureSession
		result2 error
	}{result1, result2}
}

func (fake *FakeSecureClient) Close() error {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		return fake.CloseStub()
	} else {
		return fake.closeReturns.result1
	}
}

func (fake *FakeSecureClient) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeSecureClient) CloseReturns(result1 error) {
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

var _ cmd.SecureClient = new(FakeSecureClient)