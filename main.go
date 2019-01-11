package main

import "github.com/gin-gonic/gin"
import c "gin/config"
import "github.com/gin-gonic/contrib/sessions"

func main() {
  r := gin.Default()

  store := sessions.NewCookieStore([]byte("secret"))
  r.Use(sessions.Sessions("gilangtaslimsession", store))

  c.Router(r)

  r.Run() // listen and serve on 0.0.0.0:8080
}
