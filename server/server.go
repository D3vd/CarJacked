package server

import "log"

// Init : Initialize the Server
func Init() {

	if ok := VerifyIfAllENVsAreSet(); !ok {
		log.Fatal("Please Ensure that all environment variables are set")
	}

	r := NewRouter()
	r.Run()
}
