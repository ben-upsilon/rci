package main

import "rci/log"

func PanicIf(err error) {
	if err != nil {
		log.Error(err)
		return
	}
}
