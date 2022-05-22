package main

var cache map[string][]string

//init cache
func newCache() {
	cache = make(map[string][]string)
}

//make cache empty
func flushCache() {
	cache = make(map[string][]string)
}

//add table and region server to the cache
func setCache(table string, server []string) {
	cache[table] = server
}

func deleteKey(tableName string) {
	delete(cache, tableName)
}

//find out if the table and its region server is in the cache
func getCache(table string) []string {
	result, ok := cache[table]
	if ok {
		return result
	} else {
		return nil
	}
}
