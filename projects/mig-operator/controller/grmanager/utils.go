package grmanager

import (
	"reflect"
	"runtime"
)

const (
	TriggerByTaskComplete = "TriggerByTaskComplete"
	TriggerByIntervention = "TriggerByIntervention"
)

//unkown
func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
