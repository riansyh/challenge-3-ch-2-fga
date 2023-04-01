package controllers

import (
	"challenge-3/models"
	"challenge-3/repositories"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateBook(ctx *gin.Context) {
	var newBook models.Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := repositories.AddBook(newBook)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error_status":  "Database Error",
			"error_message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated,
		"Created",
	)
}

func UpdateBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	idInt, err := strconv.Atoi(bookId)

	var UpdatedBook models.Book

	if err := ctx.ShouldBindJSON(&UpdatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = repositories.UpdateBook(UpdatedBook, idInt)

	if err != nil {
		panic(err)
	}

	if err == sql.ErrNoRows {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	ctx.JSON(http.StatusOK,
		"Updated",
	)
}

func GetBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	idInt, err := strconv.Atoi(bookId)

	var bookData models.Book

	bookData, err = repositories.GetBook(idInt)

	if err == sql.ErrNoRows {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK,
		bookData,
	)
}

func GetAllBooks(ctx *gin.Context) {
	var results = []models.Book{}

	results, err := repositories.GetAllBooks()

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, results)
}

func DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	idInt, err := strconv.Atoi(bookId)

	err = repositories.DeleteBook(idInt)

	if err == sql.ErrNoRows {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", bookId),
		})
		return
	}

	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK,
		"Deleted",
	)
}
