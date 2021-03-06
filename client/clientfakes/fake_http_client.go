// This file was generated by counterfeiter
package clientfakes

import (
	"net/http"
	"sync"

	"github.com/cloudfoundry-incubator/credhub-cli/client"
)

type FakeHttpClient struct {
	DoStub        func(req *http.Request) (resp *http.Response, err error)
	doMutex       sync.RWMutex
	doArgsForCall []struct {
		req *http.Request
	}
	doReturns struct {
		result1 *http.Response
		result2 error
	}
	invocations map[string][][]interface{}
}

func (fake *FakeHttpClient) Do(req *http.Request) (resp *http.Response, err error) {
	fake.doMutex.Lock()
	fake.doArgsForCall = append(fake.doArgsForCall, struct {
		req *http.Request
	}{req})
	fake.guard("Do")
	fake.invocations["Do"] = append(fake.invocations["Do"], []interface{}{req})
	fake.doMutex.Unlock()
	if fake.DoStub != nil {
		return fake.DoStub(req)
	} else {
		return fake.doReturns.result1, fake.doReturns.result2
	}
}

func (fake *FakeHttpClient) DoCallCount() int {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return len(fake.doArgsForCall)
}

func (fake *FakeHttpClient) DoArgsForCall(i int) *http.Request {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return fake.doArgsForCall[i].req
}

func (fake *FakeHttpClient) DoReturns(result1 *http.Response, result2 error) {
	fake.DoStub = nil
	fake.doReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeHttpClient) Invocations() map[string][][]interface{} {
	return fake.invocations
}

func (fake *FakeHttpClient) guard(key string) {
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
}

var _ client.HttpClient = new(FakeHttpClient)
