package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

type TrafficLightState int

const (
	StateRed TrafficLightState = iota
	StateYellow
	StateGreen
)

var lightName = map[TrafficLightState]string{
	StateRed:    "Red",
	StateYellow: "Yellow",
	StateGreen:  "Green",
}

func nextState(t TrafficLightState) TrafficLightState {
	switch t {
	case StateRed:
		return StateYellow
	case StateYellow:
		return StateGreen
	case StateGreen:
		return StateRed
	default:
		panic(fmt.Errorf("unknown state: %s", t))
	}
}

func (tl TrafficLightState) String() string {
	return lightName[tl]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)

	currentState := nextState(StateGreen)
	fmt.Println(currentState)
	nextLight := nextState(currentState)
	fmt.Println(nextLight)

}
