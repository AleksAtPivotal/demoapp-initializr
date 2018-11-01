// Code generated by counterfeiter. DO NOT EDIT.
package resourcefakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
)

type FakeResourceInstance struct {
	SourceStub        func() atc.Source
	sourceMutex       sync.RWMutex
	sourceArgsForCall []struct{}
	sourceReturns     struct {
		result1 atc.Source
	}
	sourceReturnsOnCall map[int]struct {
		result1 atc.Source
	}
	ParamsStub        func() atc.Params
	paramsMutex       sync.RWMutex
	paramsArgsForCall []struct{}
	paramsReturns     struct {
		result1 atc.Params
	}
	paramsReturnsOnCall map[int]struct {
		result1 atc.Params
	}
	VersionStub        func() atc.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct{}
	versionReturns     struct {
		result1 atc.Version
	}
	versionReturnsOnCall map[int]struct {
		result1 atc.Version
	}
	ResourceTypeStub        func() resource.ResourceType
	resourceTypeMutex       sync.RWMutex
	resourceTypeArgsForCall []struct{}
	resourceTypeReturns     struct {
		result1 resource.ResourceType
	}
	resourceTypeReturnsOnCall map[int]struct {
		result1 resource.ResourceType
	}
	ResourceCacheStub        func() db.UsedResourceCache
	resourceCacheMutex       sync.RWMutex
	resourceCacheArgsForCall []struct{}
	resourceCacheReturns     struct {
		result1 db.UsedResourceCache
	}
	resourceCacheReturnsOnCall map[int]struct {
		result1 db.UsedResourceCache
	}
	ContainerOwnerStub        func() db.ContainerOwner
	containerOwnerMutex       sync.RWMutex
	containerOwnerArgsForCall []struct{}
	containerOwnerReturns     struct {
		result1 db.ContainerOwner
	}
	containerOwnerReturnsOnCall map[int]struct {
		result1 db.ContainerOwner
	}
	LockNameStub        func(string) (string, error)
	lockNameMutex       sync.RWMutex
	lockNameArgsForCall []struct {
		arg1 string
	}
	lockNameReturns struct {
		result1 string
		result2 error
	}
	lockNameReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	FindOnStub        func(lager.Logger, worker.Worker) (worker.Volume, bool, error)
	findOnMutex       sync.RWMutex
	findOnArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.Worker
	}
	findOnReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	findOnReturnsOnCall map[int]struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceInstance) Source() atc.Source {
	fake.sourceMutex.Lock()
	ret, specificReturn := fake.sourceReturnsOnCall[len(fake.sourceArgsForCall)]
	fake.sourceArgsForCall = append(fake.sourceArgsForCall, struct{}{})
	fake.recordInvocation("Source", []interface{}{})
	fake.sourceMutex.Unlock()
	if fake.SourceStub != nil {
		return fake.SourceStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.sourceReturns.result1
}

func (fake *FakeResourceInstance) SourceCallCount() int {
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	return len(fake.sourceArgsForCall)
}

func (fake *FakeResourceInstance) SourceReturns(result1 atc.Source) {
	fake.SourceStub = nil
	fake.sourceReturns = struct {
		result1 atc.Source
	}{result1}
}

func (fake *FakeResourceInstance) SourceReturnsOnCall(i int, result1 atc.Source) {
	fake.SourceStub = nil
	if fake.sourceReturnsOnCall == nil {
		fake.sourceReturnsOnCall = make(map[int]struct {
			result1 atc.Source
		})
	}
	fake.sourceReturnsOnCall[i] = struct {
		result1 atc.Source
	}{result1}
}

func (fake *FakeResourceInstance) Params() atc.Params {
	fake.paramsMutex.Lock()
	ret, specificReturn := fake.paramsReturnsOnCall[len(fake.paramsArgsForCall)]
	fake.paramsArgsForCall = append(fake.paramsArgsForCall, struct{}{})
	fake.recordInvocation("Params", []interface{}{})
	fake.paramsMutex.Unlock()
	if fake.ParamsStub != nil {
		return fake.ParamsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.paramsReturns.result1
}

func (fake *FakeResourceInstance) ParamsCallCount() int {
	fake.paramsMutex.RLock()
	defer fake.paramsMutex.RUnlock()
	return len(fake.paramsArgsForCall)
}

func (fake *FakeResourceInstance) ParamsReturns(result1 atc.Params) {
	fake.ParamsStub = nil
	fake.paramsReturns = struct {
		result1 atc.Params
	}{result1}
}

func (fake *FakeResourceInstance) ParamsReturnsOnCall(i int, result1 atc.Params) {
	fake.ParamsStub = nil
	if fake.paramsReturnsOnCall == nil {
		fake.paramsReturnsOnCall = make(map[int]struct {
			result1 atc.Params
		})
	}
	fake.paramsReturnsOnCall[i] = struct {
		result1 atc.Params
	}{result1}
}

func (fake *FakeResourceInstance) Version() atc.Version {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct{}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.versionReturns.result1
}

func (fake *FakeResourceInstance) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeResourceInstance) VersionReturns(result1 atc.Version) {
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeResourceInstance) VersionReturnsOnCall(i int, result1 atc.Version) {
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 atc.Version
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 atc.Version
	}{result1}
}

func (fake *FakeResourceInstance) ResourceType() resource.ResourceType {
	fake.resourceTypeMutex.Lock()
	ret, specificReturn := fake.resourceTypeReturnsOnCall[len(fake.resourceTypeArgsForCall)]
	fake.resourceTypeArgsForCall = append(fake.resourceTypeArgsForCall, struct{}{})
	fake.recordInvocation("ResourceType", []interface{}{})
	fake.resourceTypeMutex.Unlock()
	if fake.ResourceTypeStub != nil {
		return fake.ResourceTypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.resourceTypeReturns.result1
}

func (fake *FakeResourceInstance) ResourceTypeCallCount() int {
	fake.resourceTypeMutex.RLock()
	defer fake.resourceTypeMutex.RUnlock()
	return len(fake.resourceTypeArgsForCall)
}

func (fake *FakeResourceInstance) ResourceTypeReturns(result1 resource.ResourceType) {
	fake.ResourceTypeStub = nil
	fake.resourceTypeReturns = struct {
		result1 resource.ResourceType
	}{result1}
}

func (fake *FakeResourceInstance) ResourceTypeReturnsOnCall(i int, result1 resource.ResourceType) {
	fake.ResourceTypeStub = nil
	if fake.resourceTypeReturnsOnCall == nil {
		fake.resourceTypeReturnsOnCall = make(map[int]struct {
			result1 resource.ResourceType
		})
	}
	fake.resourceTypeReturnsOnCall[i] = struct {
		result1 resource.ResourceType
	}{result1}
}

func (fake *FakeResourceInstance) ResourceCache() db.UsedResourceCache {
	fake.resourceCacheMutex.Lock()
	ret, specificReturn := fake.resourceCacheReturnsOnCall[len(fake.resourceCacheArgsForCall)]
	fake.resourceCacheArgsForCall = append(fake.resourceCacheArgsForCall, struct{}{})
	fake.recordInvocation("ResourceCache", []interface{}{})
	fake.resourceCacheMutex.Unlock()
	if fake.ResourceCacheStub != nil {
		return fake.ResourceCacheStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.resourceCacheReturns.result1
}

func (fake *FakeResourceInstance) ResourceCacheCallCount() int {
	fake.resourceCacheMutex.RLock()
	defer fake.resourceCacheMutex.RUnlock()
	return len(fake.resourceCacheArgsForCall)
}

func (fake *FakeResourceInstance) ResourceCacheReturns(result1 db.UsedResourceCache) {
	fake.ResourceCacheStub = nil
	fake.resourceCacheReturns = struct {
		result1 db.UsedResourceCache
	}{result1}
}

func (fake *FakeResourceInstance) ResourceCacheReturnsOnCall(i int, result1 db.UsedResourceCache) {
	fake.ResourceCacheStub = nil
	if fake.resourceCacheReturnsOnCall == nil {
		fake.resourceCacheReturnsOnCall = make(map[int]struct {
			result1 db.UsedResourceCache
		})
	}
	fake.resourceCacheReturnsOnCall[i] = struct {
		result1 db.UsedResourceCache
	}{result1}
}

func (fake *FakeResourceInstance) ContainerOwner() db.ContainerOwner {
	fake.containerOwnerMutex.Lock()
	ret, specificReturn := fake.containerOwnerReturnsOnCall[len(fake.containerOwnerArgsForCall)]
	fake.containerOwnerArgsForCall = append(fake.containerOwnerArgsForCall, struct{}{})
	fake.recordInvocation("ContainerOwner", []interface{}{})
	fake.containerOwnerMutex.Unlock()
	if fake.ContainerOwnerStub != nil {
		return fake.ContainerOwnerStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.containerOwnerReturns.result1
}

func (fake *FakeResourceInstance) ContainerOwnerCallCount() int {
	fake.containerOwnerMutex.RLock()
	defer fake.containerOwnerMutex.RUnlock()
	return len(fake.containerOwnerArgsForCall)
}

func (fake *FakeResourceInstance) ContainerOwnerReturns(result1 db.ContainerOwner) {
	fake.ContainerOwnerStub = nil
	fake.containerOwnerReturns = struct {
		result1 db.ContainerOwner
	}{result1}
}

func (fake *FakeResourceInstance) ContainerOwnerReturnsOnCall(i int, result1 db.ContainerOwner) {
	fake.ContainerOwnerStub = nil
	if fake.containerOwnerReturnsOnCall == nil {
		fake.containerOwnerReturnsOnCall = make(map[int]struct {
			result1 db.ContainerOwner
		})
	}
	fake.containerOwnerReturnsOnCall[i] = struct {
		result1 db.ContainerOwner
	}{result1}
}

func (fake *FakeResourceInstance) LockName(arg1 string) (string, error) {
	fake.lockNameMutex.Lock()
	ret, specificReturn := fake.lockNameReturnsOnCall[len(fake.lockNameArgsForCall)]
	fake.lockNameArgsForCall = append(fake.lockNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("LockName", []interface{}{arg1})
	fake.lockNameMutex.Unlock()
	if fake.LockNameStub != nil {
		return fake.LockNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lockNameReturns.result1, fake.lockNameReturns.result2
}

func (fake *FakeResourceInstance) LockNameCallCount() int {
	fake.lockNameMutex.RLock()
	defer fake.lockNameMutex.RUnlock()
	return len(fake.lockNameArgsForCall)
}

func (fake *FakeResourceInstance) LockNameArgsForCall(i int) string {
	fake.lockNameMutex.RLock()
	defer fake.lockNameMutex.RUnlock()
	return fake.lockNameArgsForCall[i].arg1
}

func (fake *FakeResourceInstance) LockNameReturns(result1 string, result2 error) {
	fake.LockNameStub = nil
	fake.lockNameReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceInstance) LockNameReturnsOnCall(i int, result1 string, result2 error) {
	fake.LockNameStub = nil
	if fake.lockNameReturnsOnCall == nil {
		fake.lockNameReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.lockNameReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceInstance) FindOn(arg1 lager.Logger, arg2 worker.Worker) (worker.Volume, bool, error) {
	fake.findOnMutex.Lock()
	ret, specificReturn := fake.findOnReturnsOnCall[len(fake.findOnArgsForCall)]
	fake.findOnArgsForCall = append(fake.findOnArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.Worker
	}{arg1, arg2})
	fake.recordInvocation("FindOn", []interface{}{arg1, arg2})
	fake.findOnMutex.Unlock()
	if fake.FindOnStub != nil {
		return fake.FindOnStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findOnReturns.result1, fake.findOnReturns.result2, fake.findOnReturns.result3
}

func (fake *FakeResourceInstance) FindOnCallCount() int {
	fake.findOnMutex.RLock()
	defer fake.findOnMutex.RUnlock()
	return len(fake.findOnArgsForCall)
}

func (fake *FakeResourceInstance) FindOnArgsForCall(i int) (lager.Logger, worker.Worker) {
	fake.findOnMutex.RLock()
	defer fake.findOnMutex.RUnlock()
	return fake.findOnArgsForCall[i].arg1, fake.findOnArgsForCall[i].arg2
}

func (fake *FakeResourceInstance) FindOnReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.FindOnStub = nil
	fake.findOnReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceInstance) FindOnReturnsOnCall(i int, result1 worker.Volume, result2 bool, result3 error) {
	fake.FindOnStub = nil
	if fake.findOnReturnsOnCall == nil {
		fake.findOnReturnsOnCall = make(map[int]struct {
			result1 worker.Volume
			result2 bool
			result3 error
		})
	}
	fake.findOnReturnsOnCall[i] = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceInstance) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	fake.paramsMutex.RLock()
	defer fake.paramsMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	fake.resourceTypeMutex.RLock()
	defer fake.resourceTypeMutex.RUnlock()
	fake.resourceCacheMutex.RLock()
	defer fake.resourceCacheMutex.RUnlock()
	fake.containerOwnerMutex.RLock()
	defer fake.containerOwnerMutex.RUnlock()
	fake.lockNameMutex.RLock()
	defer fake.lockNameMutex.RUnlock()
	fake.findOnMutex.RLock()
	defer fake.findOnMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceInstance) recordInvocation(key string, args []interface{}) {
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

var _ resource.ResourceInstance = new(FakeResourceInstance)