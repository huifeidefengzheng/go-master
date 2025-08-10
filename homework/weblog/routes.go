package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

func setupRoutes(db *gorm.DB) error {
	r := gin.Default()

	// 设置信任网络 - 在生产环境中应该根据实际情况设置
	r.SetTrustedProxies(nil)

	// 公开路由（无需认证）
	public := r.Group("/api")
	{
		public.POST("/register", func(c *gin.Context) {
			RegisterUser(c, db)
		})
		public.POST("/login", func(c *gin.Context) {
			LoginUser(c, db)
		})
		public.GET("/posts", func(c *gin.Context) {
			GetPosts(c, db)
		})
		public.GET("/posts/:id", func(c *gin.Context) {
			GetPost(c, db)
		})
		public.GET("/posts/:post_id/comments", func(c *gin.Context) {
			GetComments(c, db)
		})
	}

	// 受保护的路由（需要认证）
	protected := r.Group("/api")
	protected.Use(AuthMiddleware())
	{
		protected.POST("/posts", func(c *gin.Context) {
			CreatePost(c, db)
		})
		protected.PUT("/posts/:id", func(c *gin.Context) {
			UpdatePost(c, db)
		})
		protected.DELETE("/posts/:id", func(c *gin.Context) {
			DeletePost(c, db)
		})
		protected.POST("/posts/:post_id/comments", func(c *gin.Context) {
			CreateComment(c, db)
		})
	}

	log.Println("服务器启动在端口 8080")
	return r.Run(":8080")
}
