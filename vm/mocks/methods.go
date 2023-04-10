package mocks

import (
	"errors"
	"time"

	"github.com/taubyte/go-interfaces/vm"
)

func (m *mockPlugin) New(instance vm.Instance) (vm.PluginInstance, error) {
	if instance == nil {
		return nil, errors.New("instance is nil")
	}

	if m.InstanceFail {
		return nil, errors.New("mock failure")
	}

	return &mockPluginInstance{Memories: m.Memories, Globals: m.Globals}, nil
}

func (m *mockPlugin) Name() string {
	return "mock"
}

func (m *mockPluginInstance) Load(hostModule vm.HostModule) (vm.ModuleInstance, error) {
	if hostModule == nil {
		return nil, errors.New("host module is nil")
	}

	if len(m.Memories) > 0 {
		for _, memory := range m.Memories {
			if err := hostModule.Memory(memory); err != nil {
				return nil, err
			}
		}
	}

	if len(m.Globals) > 0 {
		if err := hostModule.Globals(m.Globals); err != nil {
			return nil, err
		}
	}

	return &mockModuleInstance{}, nil
}

func (m *mockPluginInstance) Close() error {
	return nil
}

func (m *mockPluginInstance) LoadFactory(factory vm.Factory, hm vm.HostModule) error {
	if factory == nil || hm == nil {
		return errors.New("params are nil")
	}

	return nil
}

func (m *mockModuleInstance) Function(name string) (vm.FunctionInstance, error) {
	if len(name) == 0 {
		return nil, errors.New("name is empty")
	}

	return &mockFunctionInstance{}, nil
}

func (m *mockFunctionInstance) Timeout(time.Duration) vm.FunctionInstance {
	return m
}

func (m *mockFunctionInstance) Cancel() error {
	return nil
}

func (m *mockFunctionInstance) Call(...interface{}) vm.Return {
	return &mockReturn{}
}

func (m *mockReturn) Error() error {
	return nil
}

func (m *mockReturn) Reflect(...interface{}) error {
	return nil
}
