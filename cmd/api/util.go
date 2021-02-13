package main

import "os"

func get(key string, fallbacks ...string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		if len(fallbacks) > 0 {
			return fallbacks[0]
		}
		panic("env variable '" + key + "' not found.")
	}
	return v
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
