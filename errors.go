package main

import "log"

func checkPanic(err error) {
	if err != nil {
		log.Panicln(err)
	}
}
