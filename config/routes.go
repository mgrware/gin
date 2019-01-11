package config

import "github.com/gin-gonic/gin"
import ctrl "gin/controllers"

func Router(r *gin.Engine) {

  v1 := r.Group("/v1")
  session := r.Group("/sessions")

  v1.Use(ctrl.AuthRequired())

  {
    v1.GET("/dashboards", ctrl.Index)
    session.POST("/create", ctrl.Login)
  }

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "pong",
    })
  })

}


