package main

import "os"

func lookupEnvOrUseDefault(key string, defaultValue string) string {
	val, found := os.LookupEnv(key)
	if found {
		return val
	}

	return defaultValue
}
