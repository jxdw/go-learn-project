package registercenter

import "time"

type ServiceDescInfo struct {
	ServiceName string
	Host string
	Post int
    IntervalTime time.Duration
}
