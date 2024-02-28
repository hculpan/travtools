package cache

var cache map[string]interface{} = make(map[string]interface{})

func ResetCache() {
	cache = make(map[string]interface{})
}

func SetCache(key string, value interface{}) {
	cache[key] = value
}

func HasCache(key string) bool {
	_, result := cache[key]
	return result
}

func GetCache(key string) interface{} {
	return cache[key]
}
