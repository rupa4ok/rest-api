package Controllers

import (
	"github.com/gin-gonic/gin"
	ApiHelpers2 "main/src/ApiHelpers"
	Models2 "main/src/Models"
)

func ListBook(c *gin.Context) {
	var book []Models2.Book
	err := Models2.GetAllBook(&book)
	if err != nil {
		ApiHelpers2.RespondJSON(c, 404, book)
	} else {
		ApiHelpers2.RespondJSON(c, 200, book)
	}
}

func AddNewBook(c *gin.Context) {
	var book Models2.Book
	c.BindJSON(&book)
	err := Models2.AddNewBook(&book)
	if err != nil {
		ApiHelpers2.RespondJSON(c, 404, book)
	} else {
		ApiHelpers2.RespondJSON(c, 200, book)
	}
}

func GetOneBook(c *gin.Context) {
	id := c.Params.ByName("id")
	var book Models2.Book
	err := Models2.GetOneBook(&book, id)
	if err != nil {
		ApiHelpers2.RespondJSON(c, 404, book)
	} else {
		ApiHelpers2.RespondJSON(c, 200, book)
	}
}

func PutOneBook(c *gin.Context) {
	var book Models2.Book
	id := c.Params.ByName("id")
	err := Models2.GetOneBook(&book, id)
	if err != nil {
		ApiHelpers2.RespondJSON(c, 404, book)
	}
	c.BindJSON(&book)
	err = Models2.PutOneBook(&book, id)
	if err != nil {
		ApiHelpers2.RespondJSON(c, 404, book)
	} else {
		ApiHelpers2.RespondJSON(c, 200, book)
	}
}

func DeleteBook(c *gin.Context) {
	var book Models2.Book
	id := c.Params.ByName("id")
	err := Models2.DeleteBook(&book, id)
	if err != nil {
		ApiHelpers2.RespondJSON(c, 404, book)
	} else {
		ApiHelpers2.RespondJSON(c, 200, book)
	}
}
