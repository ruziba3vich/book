package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/ruziba3vich/book/internal/handlers"
)

// func main() {
// O'zingiz kitoblarga kamida 4 ta statik ma`lumot qo'shing

/*
	Foydalanuvchi
		1. Kitobarni bron qilishi
		2. Kitoblarni qaytarishi
		3. O'zini bron qilgan kitoblarini ko'rishi
		4. Kitoblarga yangi kitob kiritishi
		5. Kitoblar ichidan kitobni olib tashlashi
		6. Barcha kitoblarni ko'rishi
		7.
*/
// }
func main() {
	router := gin.Default()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "postgres", "Dost0n1k", "book")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		PrintError(err)
	}
	defer db.Close()

	dbs := []string{"users", "authors", "books", "borrowed_books"}
	for _, dbName := range dbs {
		name := "../internal/db/" + dbName + ".sql"
		sqlFile, err := os.ReadFile(name)
		// fmt.Println(string(sqlFile))
		if err != nil {
			PrintError(err)
		}

		_, err = db.Exec(string(sqlFile))
		if err != nil {
			PrintError(err)
		}
	}

	router.POST("/user/borrow-book", func(c *gin.Context) {
		handlers.BorrowBook(c, db)
	})
	router.POST("/user/create-author", func(c *gin.Context) {
		handlers.CreateAuthor(c, db)
	})
	router.POST("/user/register", func(c *gin.Context) {
		handlers.Register(c, db)
	})

	addr := "localhost:7777"
	log.Println("Server listening on", addr)
	if err := router.Run(addr); err != nil {
		PrintError(err)
	}
}

func PrintError(err error) {
	log.Fatal("Error :", err)
}
