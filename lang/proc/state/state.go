package state

//go:generate stringer -type=FunctionStates

// FunctionStates is what the point along the murex pipeline a proc.Process is at
type FunctionStates int

// The different states available to FunctionStates:
const (
	Undefined FunctionStates = iota
	MemAllocated
	Assigned
	Starting
	Executing
	Executed
	Terminating
	AwaitingGC
)
