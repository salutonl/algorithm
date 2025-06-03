package examples

import (
	"fmt"
)

type ServiceState int 

const (
	ServiceIdle ServiceState = iota
	ServiceConnected
	ServiceErr
	ServiceRetry
)

var ServiceName = map[ServiceState]string {
	ServiceIdle: "idle",
	ServiceConnected: "connected",
	ServiceErr: "Error",
	ServiceRetry: "retry",
}

func (ss ServiceState) String() string {
	return ServiceName[ss]
}

func Trans(s ServiceState) ServiceState {
	switch s {
	case ServiceIdle:
		return ServiceConnected
	case ServiceConnected, ServiceRetry:
		return ServiceIdle
	case ServiceErr:
		return ServiceErr
	default:
		panic(fmt.Errorf("unknown state"))
		
	}
}