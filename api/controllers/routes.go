package controllers

import "teastore/api/middlewares"

func (server *Server) initRoutes() {
	server.Router.GET("/ping", Ping)
	server.Router.GET("/", RenderHome)
	server.Router.GET("/about", RenderAbout)
	server.Router.GET("/contact", RenderContact)

	userRoute := server.Router.Group("/users")
	{
		userRoute.GET("/register", RenderRegister)
		userRoute.GET("/login", RenderLogin)
		userRoute.GET("/view/:id", server.ShowUser)
		userRoute.POST("/register", middlewares.PasserMiddleware(), server.Register)
		userRoute.POST("/login", middlewares.PasserMiddleware(), server.Login)
		userRoute.POST("/edit", middlewares.AuthenticationMiddleware(""), server.UpdateUser)
		userRoute.POST("/delete", middlewares.AuthenticationMiddleware(""), server.DeleteUser)
	}

	productRoute := server.Router.Group("/products")
	{
		productRoute.GET("/view/product/:path", server.ShowProduct)
		productRoute.POST("/add", middlewares.AuthenticationMiddleware("admin"), server.AddProduct)
		productRoute.POST("/edit", middlewares.AuthenticationMiddleware("admin"), server.UpdateProduct)
		productRoute.POST("/delete", middlewares.AuthenticationMiddleware("admin"), server.DeleteProduct)
	}

	blogRoute := server.Router.Group("/blogs")
	{
		blogRoute.GET("/view/:path", server.ReadBlog)
		blogRoute.POST("/create", middlewares.AuthenticationMiddleware("admin"), server.CreateBlog)
		blogRoute.POST("/edit", middlewares.AuthenticationMiddleware("admin"), server.UpdateBlog)
		blogRoute.POST("/delete", middlewares.AuthenticationMiddleware("admin"), server.DeleteBlog)
	}
}
