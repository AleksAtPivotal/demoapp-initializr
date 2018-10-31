// Code generated by counterfeiter. DO NOT EDIT.
package migrationfakes

import (
	"sync"

	"github.com/concourse/atc/db/migration"
)

type FakeBindata struct {
	AssetNamesStub        func() []string
	assetNamesMutex       sync.RWMutex
	assetNamesArgsForCall []struct{}
	assetNamesReturns     struct {
		result1 []string
	}
	assetNamesReturnsOnCall map[int]struct {
		result1 []string
	}
	AssetStub        func(name string) ([]byte, error)
	assetMutex       sync.RWMutex
	assetArgsForCall []struct {
		name string
	}
	assetReturns struct {
		result1 []byte
		result2 error
	}
	assetReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBindata) AssetNames() []string {
	fake.assetNamesMutex.Lock()
	ret, specificReturn := fake.assetNamesReturnsOnCall[len(fake.assetNamesArgsForCall)]
	fake.assetNamesArgsForCall = append(fake.assetNamesArgsForCall, struct{}{})
	fake.recordInvocation("AssetNames", []interface{}{})
	fake.assetNamesMutex.Unlock()
	if fake.AssetNamesStub != nil {
		return fake.AssetNamesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.assetNamesReturns.result1
}

func (fake *FakeBindata) AssetNamesCallCount() int {
	fake.assetNamesMutex.RLock()
	defer fake.assetNamesMutex.RUnlock()
	return len(fake.assetNamesArgsForCall)
}

func (fake *FakeBindata) AssetNamesReturns(result1 []string) {
	fake.AssetNamesStub = nil
	fake.assetNamesReturns = struct {
		result1 []string
	}{result1}
}

func (fake *FakeBindata) AssetNamesReturnsOnCall(i int, result1 []string) {
	fake.AssetNamesStub = nil
	if fake.assetNamesReturnsOnCall == nil {
		fake.assetNamesReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.assetNamesReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *FakeBindata) Asset(name string) ([]byte, error) {
	fake.assetMutex.Lock()
	ret, specificReturn := fake.assetReturnsOnCall[len(fake.assetArgsForCall)]
	fake.assetArgsForCall = append(fake.assetArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("Asset", []interface{}{name})
	fake.assetMutex.Unlock()
	if fake.AssetStub != nil {
		return fake.AssetStub(name)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.assetReturns.result1, fake.assetReturns.result2
}

func (fake *FakeBindata) AssetCallCount() int {
	fake.assetMutex.RLock()
	defer fake.assetMutex.RUnlock()
	return len(fake.assetArgsForCall)
}

func (fake *FakeBindata) AssetArgsForCall(i int) string {
	fake.assetMutex.RLock()
	defer fake.assetMutex.RUnlock()
	return fake.assetArgsForCall[i].name
}

func (fake *FakeBindata) AssetReturns(result1 []byte, result2 error) {
	fake.AssetStub = nil
	fake.assetReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeBindata) AssetReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.AssetStub = nil
	if fake.assetReturnsOnCall == nil {
		fake.assetReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.assetReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeBindata) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.assetNamesMutex.RLock()
	defer fake.assetNamesMutex.RUnlock()
	fake.assetMutex.RLock()
	defer fake.assetMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBindata) recordInvocation(key string, args []interface{}) {
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

var _ migration.Bindata = new(FakeBindata)
