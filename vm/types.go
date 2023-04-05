package vm

import (
	"io"
	"time"

	"github.com/spf13/afero"
	"golang.org/x/net/context"
)

type Context interface {
	// Context returns the go context of the function instance.
	Context() context.Context
	// Close calls the go context cancel method.
	Close() error
	// Project returns the Taubyte project id
	Project() string
	// Application returns the application, if none returns an empty string
	Application() string
	// Resource returns the id of the resource being used.
	Resource() string
	// Branch returns the branch name used by this resource execution pipeline.
	Branch() string
	// Commit returns the commit id used by this resource execution pipeline.
	Commit() string
}

type Service interface {
	New(context Context) (Instance, error)
	Source() Source
	Close() error
}

type Backend interface {
	// Returns the URI scheme the backend supports.
	Scheme() string
	// Get attempts to retrieve the WASM asset.
	Get(uri string) (io.ReadCloser, error)
	// Close will close the Backend.
	Close() error
}

type Resolver interface {
	// Lookup resolves a module name and returns the uri
	Lookup(ctx Context, module string) (string, error)
}

type Loader interface {
	// Load resolves the module, then loads the module using a Backend
	Load(ctx Context, module string) (io.ReadCloser, error)
}

type Source interface {
	// Module Loads the given module name, and returns the SourceModule
	Module(ctx Context, name string) (SourceModule, error)
}

type SourceModule interface {
	// Source returns the raw data of the source
	Source() []byte
	// Imports returns functions, and memories required for instantiation
	Imports() []string
	// Imports returns functions, and memories for the specific module.
	ImportsByModule(name string) []string
	// ImportsFunction returns a boolean based on existence of a function in given module.
	ImportsFunction(module, name string) bool
}

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
	// Function adds the function definition to the HostModule
	Function(*HostModuleFunctionDefinition) error
	// Function adds multiple function definitions to the HostModule
	Functions([]*HostModuleFunctionDefinition) error
	// Memory adds the memory definition to the HostModule
	Memory(*HostModuleMemoryDefinition) error
	// Global adds the global definition to the HostModule
	Global(*HostModuleGlobalDefinition) error
	// Globals adds multiple global definitions to the HostModule
	Globals([]*HostModuleGlobalDefinition) error
	// Compile will compile the defined HostModule, and return a ModuleInstance
	Compile() (ModuleInstance, error)
}

type Instance interface {
	// Context returns the context of the function Instance
	Context() Context
	// Close will close the Instance
	Close() error
	// Runtime returns a Runtime with  the HostModuleFunctionDefinitions
	Runtime(...*HostModuleFunctionDefinition) (Runtime, error)
	// Filesystem returns the filesystem used by the given Instance.
	Filesystem() afero.Fs
	// Stdout returns the Reader interface of stdout
	Stdout() io.Reader
	// Stderr returns the Reader interface of stderr
	Stderr() io.Reader
}

type Runtime interface {
	// Module returns the ModuleInstance of the given module name.
	// A module name is in format <type>/<name>
	Module(name string) (ModuleInstance, error)
	// Expose returns a new HostModule, with the given name
	Expose(name string) (HostModule, error)
	// Attach will return a PluginInstance, and ModuleInstance of a Runtime with attached plugins
	Attach(plugin Plugin) (PluginInstance, ModuleInstance, error)
	// Stdout returns the Reader interface of stdout
	Stdout() io.Reader
	// Stderr returns the Reader interface of stderr
	Stderr() io.Reader
	// Close will close the runtime
	Close() error
}

type ModuleInstance interface {
	// Function returns a FunctionInstance of given name from the ModuleInstance
	Function(name string) (FunctionInstance, error)
}

type FunctionInstanceCommon interface {
	// Timeout will assign a timeout the FunctionInstance
	Timeout(timeout time.Duration) FunctionInstance
	// Cancel will cancel the context of the FunctionInstance
	Cancel() error
}

type FunctionInstance interface {
	FunctionInstanceCommon
	Call(args ...interface{}) Return
}

type Return interface {
	// Returns an error
	Error() error
	// Reflect assigns the return values to the given args
	Reflect(args ...interface{}) error
}

type PluginInstance interface {
	// Load will load all Factories to the HostModule, and return the ModuleInstance
	Load(HostModule) (ModuleInstance, error)
	// Close will close the PluginInstance
	Close() error
	// LoadFactory will load a single Factory on the HostModule
	LoadFactory(factory Factory, hm HostModule) error
}

type Factory interface {
	// Load will initialize the Factory
	Load(hm HostModule) error
	// Close will close and cleanup the Factory
	Close() error
	// Name returns the name of the Factory
	Name() string
}

// TODO: New takes options for factories
type Plugin interface {
	// New creates a new PluginInstance
	New(Instance) (PluginInstance, error)
	// Name returns the name of the Plugin
	Name() string
}
