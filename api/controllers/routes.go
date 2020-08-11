package controllers

func (server *Server) initRoutes() {
	server.Router.GET("/ping", Ping)
	server.Router.GET("/", RenderHome)
}
