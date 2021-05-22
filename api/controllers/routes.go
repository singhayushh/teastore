package controllers

import "teastore/api/middlewares"

// initialize endpoints
func (server *Server) initRoutes() {

	// Index Routes
	server.Router.GET("/", RenderHome)
	server.Router.GET("/teas", server.RenderProducts)
	server.Router.GET("/login", RenderLogin)
	server.Router.GET("/register", RenderRegister)
	server.Router.GET("/about", RenderAbout)
	server.Router.GET("/contact", RenderContact)
	server.Router.POST("/login", middlewares.PasserMiddleware(), server.Login)
	server.Router.POST("/register", middlewares.PasserMiddleware(), server.Register)
	server.Router.GET("/logout", middlewares.AuthenticationMiddleware(""), server.Logout)

	dashboardRoute := server.Router.Group("/dashboard")
	dashboardRoute.Use(middlewares.AuthenticationMiddleware("Admin"))
	{
		dashboardRoute.GET("/", server.RenderDashboard)
		dashboardRoute.GET("/users", server.RenderUserDashboard)
		dashboardRoute.GET("/blogs", server.RenderBlogDashboard)
		dashboardRoute.GET("/products", server.RenderProductDashboard)
	}

	cartRoute := server.Router.Group("/cart")
	cartRoute.Use(middlewares.AuthenticationMiddleware(""))
	{
		cartRoute.GET("/", server.RenderCart)
		cartRoute.GET("/checkout", server.RenderCheckout)
		cartRoute.POST("/add", server.AddtoCart)
		cartRoute.POST("/remove", server.RemovefromCart)
		cartRoute.POST("/checkout", server.Checkout)
	}

	userRoute := server.Router.Group("/user")
	{
		userRoute.GET("/edit", server.RenderEditUser)
		userRoute.GET("/view/:id", server.RenderUserByID)
		userRoute.GET("/delete/:id", middlewares.AuthenticationMiddleware("Admin"), server.DeleteUserByID)
		userRoute.POST("/edit/:id", middlewares.AuthenticationMiddleware(""), server.UpdateUserByID)
	}

	productRoute := server.Router.Group("/product")
	{
		productRoute.GET("/view/:path", server.RenderProductByPath)

		productRoute.GET("/add", middlewares.AuthenticationMiddleware("Admin"), server.RenderAddProduct)
		productRoute.GET("/edit/:path", middlewares.AuthenticationMiddleware("Admin"), server.RenderEditProduct)

		productRoute.POST("/add", middlewares.AuthenticationMiddleware("Admin"), server.AddProduct)
		productRoute.POST("/edit/:id", middlewares.AuthenticationMiddleware("Admin"), server.UpdateProductByID)
		productRoute.GET("/delete/:id", middlewares.AuthenticationMiddleware("Admin"), server.DeleteProductByID)
	}

	blogRoute := server.Router.Group("/blogs")
	{
		blogRoute.GET("/", server.RenderBlogs)
		blogRoute.GET("/view/:path", server.RenderBlogByPath)

		blogRoute.GET("/add", middlewares.AuthenticationMiddleware("Admin"), server.RenderAddBlog)
		blogRoute.GET("/edit/:path", middlewares.AuthenticationMiddleware("Admin"), server.RenderEditBlog)

		blogRoute.POST("/add", middlewares.AuthenticationMiddleware("Admin"), server.CreateBlog)
		blogRoute.POST("/edit/:id", middlewares.AuthenticationMiddleware("Admin"), server.UpdateBlogByID)
		blogRoute.GET("/delete/:id", middlewares.AuthenticationMiddleware("Admin"), server.DeleteBlogByID)
	}
}
