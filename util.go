package go_fire

import "log"

func CheckErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
