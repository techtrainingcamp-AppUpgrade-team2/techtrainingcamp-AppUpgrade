package model

import (
	"sync"
)

type DeviceIDCache struct {
	DeviceIDCache map[RuleIdType]map[DeviceIdType]struct{}
}

var instance *DeviceIDCache
var once sync.Once

func GetInstance() DeviceIDCache {
	once.Do(func() {
		instance = new(DeviceIDCache)
		instance.DeviceIDCache = make(map[RuleIdType]map[DeviceIdType]struct{})
	})
	return *instance
}
