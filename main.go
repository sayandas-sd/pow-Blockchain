package main

import (
	"net/http"
)

type Block struct {
}

type Blockchain struct {
	blocks []*Block
}

var blockchain *Blockchain

func main() {

}

func run() error {

}

func router() http.Handler {
	HandleFunc("/", getBlocks).Methods("GET")
	Handlefunc("/", writeBlock).Methos("POST")
}

func getBlocks() {

}

func writeBlock() {

}

func responseJson() {

}

func validBlock() bool {

}

func createHash() string {

}

func generateBlock() {

}

func validHash() bool {

}
