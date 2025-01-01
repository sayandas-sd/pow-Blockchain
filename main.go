package main

import (
	"encoding/json"
	"fmt"
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

type Message struct {
	Data int
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
	muxRouter := router()

	port := os.Getenv("PORT")

	log.Println("server is running on port:", port)

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        muxRouter,
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

	r := mux.NewRouter()

	r.HandleFunc("/", getBlocks).Methods("GET")
	r.HandleFunc("/", writeBlock).Methods("POST")

	return r
}

func getBlocks(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(blockchain, "", "")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func writeBlock(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "applcation/json")

	var m Message

	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		responseJson(w, r, http.StatusBadRequest, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()

	newBlock := generateBlock(blockchain[len(blockchain)-1], m.Data)

	mutex.Unlock()

	if validBlock(newBlock, blockchain[len(blockchain)-1]) {
		blockchain = append(blockchain, newBlock)
		spew.Dump(blockchain)
	}

	responseJson(w, r, http.StatusOK, newBlock)

}

func responseJson(w http.ResponseWriter, r *http.Request, statusCode int, payload interface{}) {
	w.Header().Set("Content-type", "application/json")

	res, err := json.MarshalIndent(payload, "", "")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error encoding response: %v", err)

		w.Write([]byte("Internal server error"))
		return
	}

	w.WriteHeader(statusCode)
	w.Write(res)
}

func validBlock(newBlock, oldBlock Block) bool {

	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if createHash(newBlock) != newBlock.Hash {
		return false
	}

	return true
}

func createHash() string {

}

func generateBlock(oldBlock Block, Data int) Block {
	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = Data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Difficulty = dificulty

	for i := 0; ; i++ {
		newhex := fmt.SPrintf("%x", i)
		newBlock.Nonce == newhex

		if !validHash(createHash(newBlock), newBlock.Difficulty) {
			fmt.Println(createHash(newBlock), "do more work....")
			time.Sleep(time.Second)
			continue
		} else {
			fmt.Println(createHash(newBlock), "Work Done")
			newBlock.Hash = createHash(newBlock)
			break
		}
	}
	return newBlock
}

func validHash() bool {

}
