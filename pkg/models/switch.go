package models

type Switch struct {
	Base
	Name         string      `json:"name" binding:"required"`
	State        SwitchState `json:"state" binding:"required,oneof=CLOSED OPEN"`
	DesiredState SwitchState `json:"desired-state" binding:"required,oneof=CLOSED OPEN"`
}
type SwitchState string

const (
	CLOSED SwitchState = "CLOSED"
	OPEN   SwitchState = "OPEN"
)

var SwitchStates = map[string]SwitchState {
	"CLOSED": CLOSED,
	"OPEN": OPEN,
}