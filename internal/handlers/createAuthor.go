package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/book/internal/app"
	"github.com/ruziba3vich/book/internal/service"
)


func CreateAuthor(c *gin.Context, db *sql.DB) {
	var auth app.Author

	if err := c.ShouldBindJSON(&auth); err != nil {
		log.Fatal(err)
		return
	}
	returnedObj, err := service.CreateAuthor(auth, db)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, returnedObj)
}
