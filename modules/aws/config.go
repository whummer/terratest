package aws

import (
	"sync"
)

var once sync.Once
var GlobalCustomEndpoints *map[string]string

func SetAwsEndpointsOverrides(endpoints *map[string]string) {
	once.Do(func() {
		GlobalCustomEndpoints = endpoints
	})
}