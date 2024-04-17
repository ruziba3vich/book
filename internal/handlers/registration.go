package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/book/internal/app"
	"github.com/ruziba3vich/book/internal/service"
)

func Register(c *gin.Context, db *sql.DB) {
	var user app.User

	c.ShouldBindJSON(&user)
	token, err := service.Register(user, db)
	if err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{"message": token})
}
