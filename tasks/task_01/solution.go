package main

import (
	"fmt"
	"strings"
)

func greet(name string) string {
	res := strings.TrimSpace(name)
	if res == "" {
		res = "World"
	}
	return fmt.Sprintf("Hello, %s!", res)
}
