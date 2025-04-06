package main

import (
	application "accounting-app-api/internal/application/book"
	infrastructure "accounting-app-api/internal/infrastructure/mysql"
	"accounting-app-api/internal/interface/handler"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World !")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数を取得
	dbUser := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbName := os.Getenv("MYSQL_DATABASE")

	// MySQL接続文字列を環境変数から取得
	dsn := fmt.Sprintf("%s:%s@tcp(go-mysql:3306)/%s?parseTime=true", dbUser, dbPassword, dbName)
	fmt.Println(dsn)


	// MySQLデータベースに接続
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	bookRepo := infrastructure.NewBookRepository(db)
	bookService := application.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	http.HandleFunc("/books", bookHandler.GetAllBooks)

	// サーバー起動
	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "hello world",
	// 	})
	// })


	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}