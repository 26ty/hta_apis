package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
    // TO allow CORS
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "accept, authorization, content-type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, PATCH, GET, PUT, DELETE")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}

func Default() *gin.Engine {

	router := gin.New()
    router.Use(CORS())

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin", "Authorization"},
	}))

	return router
}
