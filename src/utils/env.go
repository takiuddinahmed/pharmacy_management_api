package utils

import "os"

func GetEnv(key, default_val string) string {
	val := os.Getenv(key)
	if (len(key) == 0){
		return default_val
	} 
	return val
}
