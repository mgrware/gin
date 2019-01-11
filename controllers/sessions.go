package controllers

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/sessions"
  "net/http"
  "fmt"
)

type Auth struct {
  User     string `form:"username" json:"username" xml:"username"  binding:"required"`
  Password string `form:"password" json:"password" xml:"password" binding:"required"`
}


func AuthRequired() gin.HandlerFunc {
  return func(c *gin.Context) {
    session := sessions.Default(c)
    user := session.Get("user")
    fmt.Printf("%v\n", user)
    if user == nil {
      // You'd normally redirect to login page
      c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session token"})
      c.Abort()
    } else {
      // Continue down the chain to handler etc
      c.Next()
    }
  }
}

func Login(c *gin.Context) {
  session := sessions.Default(c)
  var json Auth
  if err := c.ShouldBindJSON(&json); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
      return
    }

  if json.User == "hello" && json.Password == "itsme" {
    session.Set("user", json.User) //In real world usage you'd set this to the users ID
    err := session.Save()
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session token"})
    } else {
      c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
    }
  } else {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
  }
}


