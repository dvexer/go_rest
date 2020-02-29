package logger

import "log"

func LogErrorIfExist(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
