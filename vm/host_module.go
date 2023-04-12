package vm

import (
	"io"

	"github.com/spf13/afero"
	"golang.org/x/net/context"
)

// HostFunction is the function handler of a HostModuleFunctionDefinition
type HostFunction interface{}

// HostModuleFunctionDefinition is the definition of a Function within a HostModule
type HostModuleFunctionDefinition struct {
	Name    string
	Handler HostFunction
}

// HostModuleGlobalDefinition is Global Value stored within the HostModule
type HostModuleGlobalDefinition struct {
	Name  string
	Value interface{}
}

// HostModuleMemoryDefinition is the memory definition of the Host Module.
type HostModuleMemoryDefinition struct {
	Name  string
	Pages struct {
		Min   uint64
		Max   uint64
		Maxed bool
	}
}

type HostModule interface {
	// Functions adds the function definitions to the HostModule
	Functions(...*HostModuleFunctionDefinition) error

	// Memory adds the memory definitions to the HostModule
	Memories(...*HostModuleMemoryDefinition) error

	// Globals adds the global definitions to the HostModule
	Globals(...*HostModuleGlobalDefinition) error

	// Compile will compile the defined HostModule, and return a ModuleInstance
	Compile() (ModuleInstance, error)
}

type HostModuleDefinitions struct {
	Functions []*HostModuleFunctionDefinition
	Memories  []*HostModuleMemoryDefinition
	Globals   []*HostModuleGlobalDefinition
}
type Instance interface {
	// Context returns the context of the function Instance
	Context() Context

	// Close will close the Instance
	Close() error

	// Load will Load the runtime with the host module.
	Load(*HostModuleDefinitions) error

	// Attach will attach plugins to the module instance
	Attach(Plugin) (PluginInstance, ModuleInstance, error)

	// Module will instantiate the module instance
	Module(name string) (ModuleInstance, error)

	// Expose returns a HostModule with the given name
	Expose(name string) (HostModule, error)

	// Filesystem returns the filesystem used by the given Instance.
	Filesystem() afero.Fs

	// Stdout returns the Reader interface of stdout
	Stdout() io.Reader

	// Stderr returns the Reader interface of stderr
	Stderr() io.Reader
}

// FunctionDefinition is a WebAssembly function exported in a module.
type FunctionDefinition interface {
	// Name is the module-defined name of the function, which is not necessarily
	// the same as its export name.
	Name() string

	// ParamTypes are the possibly empty sequence of value types accepted by a
	// function with this signature.
	ParamTypes() []ValueType

	// ResultTypes are the results of the function.
	ResultTypes() []ValueType
}

// Function is a WebAssembly function exported from an instantiated module.
type Function interface {
	// Definition is metadata about this function from its defining module.
	Definition() FunctionDefinition

	// Call invokes the function with parameters encoded according to ParamTypes. Up to one result is returned.
	Call(ctx context.Context, params ...uint64) ([]uint64, error)
}

// Global is a WebAssembly 1.0 (20191205) global exported from an instantiated module.
type Global interface {
	// Type describes the numeric type of the global.
	Type() ValueType

	// Get returns the last known value of this global. When the context is nil, it defaults to context.Background.
	Get() uint64
}

// MutableGlobal is a Global whose value can be updated at runtime (variable).
type MutableGlobal interface {
	Global

	// Set updates the value of this global. When the context is nil, it defaults to context.Background.
	Set(v uint64)
}
