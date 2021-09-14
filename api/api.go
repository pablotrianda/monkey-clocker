package api

func Start(port string) {
	server := newServer(port, routes())
	server.startServer()
}
