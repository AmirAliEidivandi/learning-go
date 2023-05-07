package main

import "regexp"

func isValidEmail(addr string) bool {
	re, ok := regexp.Compile(`.+@.+\..+`)
	if ok != nil {
		panic("failed to compile regex")
	} else {
		return re.Match([]byte(addr))
	}
}