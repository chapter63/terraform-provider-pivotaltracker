// Code generated by counterfeiter. DO NOT EDIT.
package ptfakes

import (
	"net/http"
	"sync"

	"github.com/xchapter7x/terraform-provider-pivotaltracker/pkg/pt"
)

type FakeRequestDoer struct {
	DoStub        func(*http.Request, interface{}) (*http.Response, error)
	doMutex       sync.RWMutex
	doArgsForCall []struct {
		arg1 *http.Request
		arg2 interface{}
	}
	doReturns struct {
		result1 *http.Response
		result2 error
	}
	doReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	NewRequestStub        func(string, string, interface{}) (*http.Request, error)
	newRequestMutex       sync.RWMutex
	newRequestArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 interface{}
	}
	newRequestReturns struct {
		result1 *http.Request
		result2 error
	}
	newRequestReturnsOnCall map[int]struct {
		result1 *http.Request
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRequestDoer) Do(arg1 *http.Request, arg2 interface{}) (*http.Response, error) {
	fake.doMutex.Lock()
	ret, specificReturn := fake.doReturnsOnCall[len(fake.doArgsForCall)]
	fake.doArgsForCall = append(fake.doArgsForCall, struct {
		arg1 *http.Request
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("Do", []interface{}{arg1, arg2})
	fake.doMutex.Unlock()
	if fake.DoStub != nil {
		return fake.DoStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.doReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRequestDoer) DoCallCount() int {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return len(fake.doArgsForCall)
}

func (fake *FakeRequestDoer) DoCalls(stub func(*http.Request, interface{}) (*http.Response, error)) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = stub
}

func (fake *FakeRequestDoer) DoArgsForCall(i int) (*http.Request, interface{}) {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	argsForCall := fake.doArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRequestDoer) DoReturns(result1 *http.Response, result2 error) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = nil
	fake.doReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeRequestDoer) DoReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.doMutex.Lock()
	defer fake.doMutex.Unlock()
	fake.DoStub = nil
	if fake.doReturnsOnCall == nil {
		fake.doReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.doReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeRequestDoer) NewRequest(arg1 string, arg2 string, arg3 interface{}) (*http.Request, error) {
	fake.newRequestMutex.Lock()
	ret, specificReturn := fake.newRequestReturnsOnCall[len(fake.newRequestArgsForCall)]
	fake.newRequestArgsForCall = append(fake.newRequestArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 interface{}
	}{arg1, arg2, arg3})
	fake.recordInvocation("NewRequest", []interface{}{arg1, arg2, arg3})
	fake.newRequestMutex.Unlock()
	if fake.NewRequestStub != nil {
		return fake.NewRequestStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newRequestReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRequestDoer) NewRequestCallCount() int {
	fake.newRequestMutex.RLock()
	defer fake.newRequestMutex.RUnlock()
	return len(fake.newRequestArgsForCall)
}

func (fake *FakeRequestDoer) NewRequestCalls(stub func(string, string, interface{}) (*http.Request, error)) {
	fake.newRequestMutex.Lock()
	defer fake.newRequestMutex.Unlock()
	fake.NewRequestStub = stub
}

func (fake *FakeRequestDoer) NewRequestArgsForCall(i int) (string, string, interface{}) {
	fake.newRequestMutex.RLock()
	defer fake.newRequestMutex.RUnlock()
	argsForCall := fake.newRequestArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRequestDoer) NewRequestReturns(result1 *http.Request, result2 error) {
	fake.newRequestMutex.Lock()
	defer fake.newRequestMutex.Unlock()
	fake.NewRequestStub = nil
	fake.newRequestReturns = struct {
		result1 *http.Request
		result2 error
	}{result1, result2}
}

func (fake *FakeRequestDoer) NewRequestReturnsOnCall(i int, result1 *http.Request, result2 error) {
	fake.newRequestMutex.Lock()
	defer fake.newRequestMutex.Unlock()
	fake.NewRequestStub = nil
	if fake.newRequestReturnsOnCall == nil {
		fake.newRequestReturnsOnCall = make(map[int]struct {
			result1 *http.Request
			result2 error
		})
	}
	fake.newRequestReturnsOnCall[i] = struct {
		result1 *http.Request
		result2 error
	}{result1, result2}
}

func (fake *FakeRequestDoer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	fake.newRequestMutex.RLock()
	defer fake.newRequestMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRequestDoer) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ pt.RequestDoer = new(FakeRequestDoer)
