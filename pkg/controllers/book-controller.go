package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nchandel2420/BookStore/pkg/models"
	"github.com/nchandel2420/BookStore/pkg/utils"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId := vars["bookId"] // models.GetBookById()
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	bookDetail, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()

	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId := vars["bookId"] // models.GetBookById()
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	bookDetail := models.DeleteBook(ID)
	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var UpdateBook = &models.Book{}
	utils.ParseBody(r, UpdateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println(err)
	}

	bookDetail, db := models.GetBookById(ID)

	if UpdateBook.Name != "" {
		bookDetail.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		bookDetail.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		bookDetail.Publication = UpdateBook.Publication
	}

	db.Save(&bookDetail)

	res, _ := json.Marshal(bookDetail)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
