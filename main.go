package main

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

const dificulty = 1

var mutex = &sync.Mutex{}

type Block struct {
	Index      int
	Timestamp  string
	Data       int
	PrevHash   string
	Hash       string
	Difficulty int
	Nonce      string
}

var blockchain []Block

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	go func() {
		t := time.Now()
		genesisBlock := &Block{}
		genesisBlock = &Block{0, t.String(), 0, "", createHash(genesisBlock), dificulty, ""}
		spew.Dump(genesisBlock)

		mutex.Lock()
		blockchain = append(blockchain, genesisBlock)
		mutex.Unlock()
	}()

	log.Fatal(run())
}

func run() error {
	router := mux.NewRouter()

	port := os.Getenv("PORT")
	log.Println("server is running on port:", port)
	s := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
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
