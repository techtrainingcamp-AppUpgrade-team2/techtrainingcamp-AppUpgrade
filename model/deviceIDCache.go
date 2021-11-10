package model

import (
	"sync"
)

// DeviceIDCache	DeviceID缓存对象，定义了规则对应的设备白名单信息
type DeviceIDCache struct {
	DeviceIDCache map[RuleIdType]map[DeviceIdType]struct{}
}

var instance *DeviceIDCache
var once sync.Once

// @title             GetInstance
// @description       返回DeviceID缓存对象类的唯一实例
// @auth              高宏宇         2021/11/10
// @return            *instance     DeviceIDCache
func GetInstance() DeviceIDCache {
	once.Do(func() {
		instance = new(DeviceIDCache)
		instance.DeviceIDCache = make(map[RuleIdType]map[DeviceIdType]struct{})
	})
	return *instance
}
