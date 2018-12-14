package main

import (
	"fmt"
	"os"
)

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		panic(fmt.Sprintf("%s environment variable not set.", k))
	}
	return v
}
