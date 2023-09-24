package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/zohaibsoomro/go-blockchain/models"
)

func GetAllBlocks(bc *models.BlockChain, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(bc.Blocks)
	if err != nil {
		log.Fatal(err)
	}
}
func DeleteBlock(bc *models.BlockChain, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]
	w.Header().Set("Content-type", "application/json")
	index, book := bc.GetBlockByBookId(bookID)
	if book != nil {
		deleted := bc.DeleteBlockAtIndex(index)
		if deleted {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Book deleted with id %q\n", bookID)
			json.NewEncoder(w).Encode(bc.Blocks)
		} else {
			fmt.Fprintf(w, "Some error occurred while deleting book with id %q", bookID)

		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No book found!")
	}
}
func WriteBlock(bc *models.BlockChain, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	b := &models.Book{}
	err := json.NewDecoder(r.Body).Decode(b)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		w.Write([]byte(err.Error()))
		return
	}
	id := b.GenerateID()
	block := models.NewBlock(models.NewBookCheckout(b, time.Now().String()), bc)
	bc.WriteBlock(block)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Book created with id %q", id)
}
