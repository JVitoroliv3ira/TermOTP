package utils

import "log"

func HandleError(err error) {
	if err != nil {
		log.Fatal("erro: ", err)
	}
}
