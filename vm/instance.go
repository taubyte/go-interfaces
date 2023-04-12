package vm

import "time"

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
