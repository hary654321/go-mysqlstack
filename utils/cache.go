/*
 * @Description:
 * @Version: 2.0
 * @Autor: ABing
 * @Date: 2024-07-12 11:59:54
 * @LastEditors: lhl
 * @LastEditTime: 2024-07-12 16:20:55
 */
package utils

import (
	"sync"
)

// 创建一个全局的 sync.Map 实例作为内存缓存
var memCache sync.Map

// SetItem 向缓存中设置一个项
func SetItem(key, value interface{}) {
	memCache.Store(key, value)
}

// GetItem 从缓存中获取一个项
func GetItem(key interface{}) (interface{}, bool) {
	return memCache.Load(key)
}

// GetItem 从缓存中获取一个项
func GetItemString(key interface{}) string {

	if filename, have := memCache.Load(key); have {

		str, ok := filename.(string)
		if ok {
			return str
		}
	}

	return ""
}

// DeleteItem 从缓存中删除一个项
func DeleteItem(key interface{}) {
	memCache.Delete(key)
}

// ClearCache 清空整个缓存
func ClearCache() {
	// sync.Map 目前没有直接的清空方法，需要遍历并删除
	memCache.Range(func(key, value interface{}) bool {
		memCache.Delete(key)
		return true
	})
}
