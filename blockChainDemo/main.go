// 区块链demo
package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/gorilla/mux"
)

//Block 定义一个结构体，标识组成区块链的每一个块的数据模型
type Block struct {
	Index     int    //这个块在整个链中的位置
	Timestamp string //生成块时的时间戳
	Value     string //数据
	Hash      string //当前块通过SHA256算法生成的散列值
	PreHash   string //前一个块的散列值
}

// Message 接收post请求的body
type Message struct {
	Value string
}

var mutex = &sync.Mutex{}

// Blockchain 表示整个链
var Blockchain []Block

// 计算Hash值
func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Value + block.PreHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func main() {
	go func() {
		//创建创世块
		t := time.Now()
		genesisBlock := Block{}
		genesisBlock = Block{0, t.String(), "创世块", calculateHash(genesisBlock), ""}
		mutex.Lock()
		Blockchain = append(Blockchain, genesisBlock)
		mutex.Unlock()
	}()
	log.Fatal(run())
}

// 生成新块
func generateBlock(oldBlock Block, value string) Block {
	var newBlock Block
	t := time.Now()
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Value = value
	newBlock.PreHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)
	return newBlock
}

// 校验块是否正确，是否被篡改
func validBlock(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PreHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

// 保证当前链是最新的（始终选择最长链为原则）
func replaceChain(newBlocks []Block){
	if len(newBlocks) > len(Blockchain){
		Blockchain = newBlocks
	}
}

// 创建web服务，对外提供查询或新增块
func run() error {
	mux := makeMuxRouter()
	log.Println("Http Server Listening om port 8880")
	s := &http.Server{
		Addr:           ":8880",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/blockChain", getChainHandler).Methods("GET")
	muxRouter.HandleFunc("/blockChain", writeBlockHandler).Methods("POST")
	return muxRouter
}

func getChainHandler(w http.ResponseWriter, r *http.Request) {
	b, err := json.MarshalIndent(Blockchain, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(b))
}

func writeBlockHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var msg Message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&msg)
	if err != nil {
		responseWithJSON(w, r, 500, r.Body)
		return
	}
	defer r.Body.Close()

	mutex.Lock()
	preBlock := Blockchain[len(Blockchain)-1]
	newBlock := generateBlock(preBlock, msg.Value)
	if validBlock(newBlock, preBlock) {
		newBlockchain := append(Blockchain, newBlock)
		replaceChain(newBlockchain)
		spew.Dump(Blockchain)
	}
	mutex.Unlock()
	responseWithJSON(w, r, http.StatusCreated, newBlock)
}

func responseWithJSON(w http.ResponseWriter, r *http.Request, code int, payload interface{}) {
	response, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("system error"))
		return
	}
	w.WriteHeader(code)
	w.Write(response)
}
