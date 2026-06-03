package main

import "fmt"

// Go does not have a built-in enum type.
//
// The common pattern is:
//   1. Create a custom integer type.
//   2. Use iota to generate sequential values.
//
// ServerState represents all possible states
// a server can be in.
type ServerState int

const (

	// iota starts at 0 and increments automatically.
	//
	// StateIdle      = 0
	// StateConnected = 1
	// StateError     = 2
	// StateRetrying  = 3
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

// Internally states are integers:
//
// StateIdle      -> 0
// StateConnected -> 1
// StateError     -> 2
// StateRetrying  -> 3
//
// This map converts enum values into readable strings.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

// String implements fmt.Stringer:
//
//	type Stringer interface {
//	    String() string
//	}
//
// Because ServerState implements String(),
// fmt.Println automatically prints:
//
//    idle
//
// instead of:
//
//    0
//
// This is one of the most common interfaces
// in the Go standard library.
func (ss ServerState) String() string {
	return stateName[ss]
}

// transition models a simple state machine.
//
// Current State ---> Next State
//
// Idle       ---> Connected
// Connected  ---> Idle
// Retrying   ---> Idle
// Error      ---> Error
//
// State machines are common in:
//
//   - network connections
//   - distributed systems
//   - workflow engines
//   - protocol implementations
func transition(s ServerState) ServerState {

	switch s {

	// Idle -> Connected
	case StateIdle:
		return StateConnected

	// Multiple cases can share logic.
	//
	// Connected -> Idle
	// Retrying  -> Idle
	case StateConnected, StateRetrying:
		return StateIdle

	// Error remains Error.
	case StateError:
		return StateError

	// Defensive programming:
	//
	// If a new state is added but not handled,
	// fail immediately instead of silently
	// producing incorrect behavior.
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

func main() {

	// Initial state:
	//
	// StateIdle
	//     ↓
	// transition()
	//     ↓
	// StateConnected
	ns := transition(StateIdle)

	// Because ServerState implements String(),
	// fmt.Println automatically calls:
	//
	//    ns.String()
	//
	// Output:
	//    connected
	fmt.Println(ns)

	// StateConnected
	//      ↓
	// transition()
	//      ↓
	// StateIdle
	ns2 := transition(ns)

	// Output:
	//    idle
	fmt.Println(ns2)
}
