package models

//Switch basic on/off switch control
type Switch struct {
	Base
	Name         string      `json:"name" binding:"required"`
	State        SwitchState `json:"state" binding:"required,oneof=CLOSED OPEN"`
	DesiredState SwitchState `json:"desired-state" binding:"required,oneof=CLOSED OPEN"`
}

//SwitchState represents the state of a switch
type SwitchState string

const (
	//CLOSED represents a closed switch state
	CLOSED SwitchState = "CLOSED"
	//OPEN represents an open switch state
	OPEN   SwitchState = "OPEN"
)

//SwitchStates map of avaliable switch state
var SwitchStates = map[string]SwitchState {
	"CLOSED": CLOSED,
	"OPEN": OPEN,
}