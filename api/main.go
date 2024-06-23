package main

import (
  "github.com/Salz-Team/salz/api/controllers"
  "github.com/Salz-Team/salz/api/middlewares"
  "github.com/Salz-Team/salz/api/config"
  "github.com/gin-gonic/gin"
)

func main() {
  // TODO Dependency injection?
  // TODO config stuff
  cfg := config.NewConfig()
  ctrl := controllers.NewController(cfg)

  r := gin.Default()

  // All routes should start with /api

  // Public routes
  r.GET("/health", controllers.Ping)
  r.GET("/login", ctrl.OAuthLoginHandler)
  r.GET("/login/callback", func (c *gin.Context) {
    ctrl.OAuthCallbackHandler(c)
  })

  // Protected routes
  protectedRoutes := r.Group("")
  protectedRoutes.Use(middlewares.AuthMiddleware(cfg.AuthDBHandler))
  {
    protectedRoutes.GET("/user/me", func (c *gin.Context) {
      controllers.GetUser(c, cfg.ApiDBHandler, cfg.AuthDBHandler)
    })
    protectedRoutes.POST("/bot/upload", func (c *gin.Context) {
      ctrl.BotUploadHandler(c)
    })
  }

  r.Run()
}
