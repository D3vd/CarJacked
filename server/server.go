package server

// Init : Initialize the Server
func Init() {
	r := NewRouter()
	r.Run()
}
