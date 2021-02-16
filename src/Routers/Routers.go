package Routers

import (
	"github.com/gin-gonic/gin"
	Controllers2 "main/src/Controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("book", Controllers2.ListBook)
		v1.POST("book", Controllers2.AddNewBook)
		v1.GET("book/:id", Controllers2.GetOneBook)
		v1.PUT("book/:id", Controllers2.PutOneBook)
		v1.DELETE("book/:id", Controllers2.DeleteBook)
	}

	return r
}
