package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zohaibsoomro/go-blockchain/models"
	"github.com/zohaibsoomro/go-blockchain/utils"
)

var bc = models.NewBlockChain()

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/all", getAllBlocks).Methods("GET")
	r.HandleFunc("/new", writeBlock).Methods("POST")
	r.HandleFunc("/delete/{id}", deleteBlock).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":300", r))
}

func getAllBlocks(w http.ResponseWriter, r *http.Request) {
	utils.GetAllBlocks(bc, w, r)
}
func deleteBlock(w http.ResponseWriter, r *http.Request) {
	utils.DeleteBlock(bc, w, r)
}
func writeBlock(w http.ResponseWriter, r *http.Request) {
	utils.WriteBlock(bc, w, r)
}
