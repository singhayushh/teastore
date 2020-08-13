package controllers

import "teastore/api/middlewares"

func (server *Server) initRoutes() {
	server.Router.GET("/ping", Ping)
	server.Router.GET("/", RenderHome)
	server.Router.POST("/u/register", server.Register)
	server.Router.POST("/u/login", server.Login)

	server.Router.GET("/profile", middlewares.AuthenticationMiddleware(), RenderProfile)
}
