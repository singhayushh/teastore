package controllers

import "teastore/api/middlewares"

func (server *Server) initRoutes() {
	server.Router.GET("/", RenderHome)
	server.Router.GET("/about", RenderAbout)
	server.Router.GET("/contact", RenderContact)
	server.Router.GET("/dashboard", middlewares.AuthenticationMiddleware("Admin"), server.RenderDashboard)

	userRoute := server.Router.Group("/user")
	{
		userRoute.GET("/", server.RenderAllUsers)
		userRoute.GET("/login", RenderLogin)
		userRoute.GET("/register", RenderRegister)
		userRoute.GET("/edit", server.RenderEditUser)
		userRoute.GET("/view/:id", server.RenderUser)

		userRoute.POST("/login", middlewares.PasserMiddleware(), server.Login)
		userRoute.POST("/register", middlewares.PasserMiddleware(), server.Register)

		userRoute.POST("/edit/:id", middlewares.AuthenticationMiddleware(""), server.UpdateUserByID)
		userRoute.GET("/delete/:id", middlewares.AuthenticationMiddleware("Admin"), server.DeleteUserByID)
	}

	productRoute := server.Router.Group("/products")
	{
		productRoute.GET("/", server.RenderAllProducts)
		productRoute.GET("/view/:id", server.RenderProduct)

		productRoute.GET("/add", middlewares.AuthenticationMiddleware("Admin"), server.RenderAddProduct)
		productRoute.GET("/edit/:id", middlewares.AuthenticationMiddleware("Admin"), server.RenderEditProduct)

		productRoute.POST("/add", middlewares.AuthenticationMiddleware("Admin"), server.AddProduct)
		productRoute.POST("/edit/:id", middlewares.AuthenticationMiddleware("Admin"), server.UpdateProductByID)
		productRoute.GET("/delete/:id", middlewares.AuthenticationMiddleware("Admin"), server.DeleteProductByID)
	}

	blogRoute := server.Router.Group("/blogs")
	{
		blogRoute.GET("/", server.RenderAllBlogs)
		blogRoute.GET("/view/:id", server.RenderBlog)

		blogRoute.GET("/add", middlewares.AuthenticationMiddleware("Admin"), server.RenderAddBlog)
		blogRoute.GET("/edit/:id", middlewares.AuthenticationMiddleware("Admin"), server.RenderEditBlog)

		blogRoute.POST("/add", middlewares.AuthenticationMiddleware("Admin"), server.CreateBlog)
		blogRoute.POST("/edit/:id", middlewares.AuthenticationMiddleware("Admin"), server.UpdateBlogByID)
		blogRoute.GET("/delete/:id", middlewares.AuthenticationMiddleware("Admin"), server.DeleteBlogByID)
	}
}
