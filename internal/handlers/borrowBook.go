package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/book/internal/app"
	"github.com/ruziba3vich/book/internal/service"
)

func BorrowBook(c *gin.Context, db *sql.DB) {
	var bbr app.BookBorrowRequest
	if err := c.ShouldBindJSON(&bbr); err != nil {
		log.Fatal(err)
	}
	returnedObj, err := service.BorrowBook(bbr, db)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, returnedObj)
}

func ReturnBook(c *gin.Context, db *sql.DB) {
	var brr app.BookBorrowRequest
	if err := c.ShouldBindJSON(&brr); err != nil {
		log.Fatal(err)
		return
	}
	err := service.ReturnBook(brr, db)
	if err != nil {
		log.Fatal(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"messgae": "Book has been returned",
	})
}
