package mocks

import "github.com/taubyte/go-interfaces/vm"

func NewPlugin(memories []*vm.HostModuleMemoryDefinition, globals []*vm.HostModuleGlobalDefinition, failInstance bool) MockedPlugin {
	return &mockPlugin{Memories: memories, Globals: globals, InstanceFail: failInstance}
}

func NewPluginInstance(memories []*vm.HostModuleMemoryDefinition, globals []*vm.HostModuleGlobalDefinition) MockedPluginInstance {
	return &mockPluginInstance{Memories: memories, Globals: globals}
}

func NewModuleInstance() MockedModuleInstance {
	return &mockModuleInstance{}
}

func NewFunctionInstance() MockedFunctionInstance {
	return &mockFunctionInstance{}
}

func NewReturn() MockedReturn {
	return &mockReturn{}
}
