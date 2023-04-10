package mocks

import "github.com/taubyte/go-interfaces/vm"

type MockedPlugin interface {
	vm.Plugin
}

type mockPlugin struct {
	Memories     []*vm.HostModuleMemoryDefinition
	Globals      []*vm.HostModuleGlobalDefinition
	InstanceFail bool
}

type MockedPluginInstance interface {
	vm.PluginInstance
}

type mockPluginInstance struct {
	Memories []*vm.HostModuleMemoryDefinition
	Globals  []*vm.HostModuleGlobalDefinition
}

type MockedModuleInstance interface {
	vm.ModuleInstance
}

type mockModuleInstance struct{}

type MockedFunctionInstance interface {
	vm.FunctionInstance
}

type mockFunctionInstance struct{}

type MockedReturn interface {
	vm.Return
}

type mockReturn struct{}
