package main

import (
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/pustaka_Api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB Connection error")
	}

	//db.AutoMigrate(book.Book{})

	bookRepository := book.NewReposiotry(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	// books, err := bookRepository.FindAll()
	// for _, book := range books{
	// 	fmt.Println("Title : ", book.Title)
	// }

	// book, err := bookRepository.FIndByID(1)
	// fmt.Println("Title : ", book.Title)

	//Create/Insert
	// book := book.BookInput{}
	// book.Title = "Main Title3"
	// book.Price = "970000"

	// bookService.Create(book)

	// err = db.Create(&book).Error
	// if err != nil{
	// 	fmt.Println("=========================")
	// 	fmt.Println("Error creating book record")
	// 	fmt.Println("=========================")
	// }

	// var book book.Book
	//err = db.Debug().First(&book).Error
	//err = db.Debug().Last(&book).Error

	// err = db.First(&book).Error
	// if err != nil {
	// 	fmt.Println("=========================")
	// 	fmt.Println("Error finding book record")
	// 	fmt.Println("=========================")
	// }

	// fmt.Println("title :", book.Title)
	// fmt.Println("book object %v", book)

	router := gin.Default()
	// router.GET("/", bookHandler.RootHandler)
	// router.GET("/hello", bookHandler.HelloHandler)
	// router.GET("/books/:id/:title", bookHandler.BooksHandler)
	// router.GET("/query", bookHandler.QueryHandler)

	router.GET("/books", bookHandler.GetBooks)
	router.GET("/books/:id", bookHandler.GetBook)
	router.POST("/books", bookHandler.CreateBook)
	router.PUT("/books/:id", bookHandler.UpdateBook)
	router.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run(":8888")
}

//main
//handler
//service
//repository
//db
//mySql
