package utilities

import "log"

func Logger(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
