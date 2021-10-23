package helpers

import (
	"log"
	"os"
)

func CheckAndExit(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}
