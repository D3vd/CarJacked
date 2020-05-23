package server

import (
	"log"
	"os"
)

// VerifyIfAllENVsAreSet : Check if all necessary Envs are set
func VerifyIfAllENVsAreSet() bool {

	//	SECRET - For JWT Authentication
	if secret := os.Getenv("SECRET"); secret == "" {
		log.Fatal("SECRET env is not set")
		return false
	}

	return true
}
