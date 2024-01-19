package main

import (
	"blockchain/database"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Inventory struct {
	ID              string     `json:"id,omitempty" bson:"id,omitempty"`
	From            time.Time  `json:"from,omitempty" bson:"from,omitempty"`
	To              *time.Time `json:"to,omitempty" bson:"to,omitempty"`
	BeforeInventory string     `json:"beforeInventory,omitempty" bson:"beforeInventory,omitempty"`
	AfterInventory  string     `json:"afterInventory,omitempty" bson:"afterInventory,omitempty"`
}
type InventoryV2 struct {
	ID   string `json:"id,omitempty" bson:"id,omitempty"`
	From struct {
		ID              string  `json:"id,omitempty" bson:"id,omitempty"`
		Name            string  `json:"name,omitempty" bson:"name,omitempty"`
		Type            string  `json:"type,omitempty" bson:"type,omitempty"`
		BeforeInventory float64 `json:"beforeInventory,omitempty" bson:"beforeInventory,omitempty"`
		AfterInventory  float64 `json:"afterInventory,omitempty" bson:"afterInventory,omitempty"`
	} `json:"from,omitempty" bson:"from,omitempty"`

	To struct {
		ID              string  `json:"id,omitempty" bson:"id,omitempty"`
		Name            string  `json:"name,omitempty" bson:"name,omitempty"`
		Type            string  `json:"type,omitempty" bson:"type,omitempty"`
		BeforeInventory float64 `json:"beforeInventory,omitempty" bson:"beforeInventory,omitempty"`
		AfterInventory  float64 `json:"afterInventory,omitempty" bson:"afterInventory,omitempty"`
	} `json:"to,omitempty" bson:"to,omitempty"`
	Quantity   string    `json:"beforeInventory,omitempty" bson:"beforeInventory,omitempty"`
	Price      string    `json:"price,omitempty" bson:"price,omitempty"`
	TimeStramp time.Time `json:"timeStramp,omitempty" bson:"timeStramp,omitempty"`
}
type BlockChain struct {
	TimeStramp *time.Time  `json:"timeStramp,omitempty" bson:"timeStramp,omitempty"`
	UniqueId   string      `json:"uniqueId,omitempty" bson:"uniqueId,omitempty"`
	Tran       InventoryV2 `json:"transaction,omitempty" bson:"transaction,omitempty"`
	PrevHash   []byte      `json:"prevHash" bson:"prevHash"`
	Hash       []byte      `json:"hash,omitempty" bson:"hash,omitempty"`
	Blocks     *BlockChain `json:"blocks" bson:"blocks"`
}

type Block struct {
	TimeStramp time.Time `json:"timeStramp,omitempty" bson:"timeStramp,omitempty"`
	Tran       Inventory `json:"transaction,omitempty" bson:"transaction,omitempty"`
	PrevHash   []byte    `json:"prevHash" bson:"prevHash"`
	Hash       []byte    `json:"hash,omitempty" bson:"hash,omitempty"`
	Remarks    string    `json:"remarks,omitempty" bson:"remarks,omitempty"`
}

func main() {
	route := mux.NewRouter()
	//route.HandleFunc("/api/stundent", GetBlockChain).Methods("POST")
	route.HandleFunc("/api/blockchain", SaveBlockChain).Methods("POST")
	route.HandleFunc("/api/getstundent/{id}", getBlockChains).Methods("GET")
	fmt.Println("Go run")
	log.Fatal(http.ListenAndServe(":9000", route))
	route.Use(AllowCors)
	route.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		w.WriteHeader(http.StatusNoContent)
		return
	})
}
func AllowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
		return
	})
}

func Blocks(transaction InventoryV2, prevHash []byte) *BlockChain {

	currentTime := time.Now()
	return &BlockChain{
		TimeStramp: &currentTime,
		Tran:       transaction,
		PrevHash:   prevHash,
		Hash:       NewHash(currentTime, transaction, prevHash),
	}
}
func NewHash(time time.Time, transaction InventoryV2, prevHash []byte) []byte {

	input := append(prevHash, time.String()...)
	for transaction := range transaction.ID {
		input = append(input, string(rune(transaction))...)
	}
	hash := sha256.Sum256(input)
	return hash[:]

}
func print(block *Block) {

	fmt.Printf("\ttime:%s\n", block.TimeStramp)
	fmt.Printf("\tprevhash:%s\n", block.PrevHash)
	fmt.Printf("\thash:%s\n", block.Hash)
	Trancation(block)
}
func Trancation(block *Block) {
	fmt.Println("\tTrancation")
	for i, transaction := range block.Tran.BeforeInventory {
		fmt.Println("\t\t%v:%q\n", i, transaction)
	}
}

func SaveBlockChain(W http.ResponseWriter, r *http.Request) {

	W.Header().Add("Content-Type", "application/json")

	invet := new(BlockChain)
	_ = json.NewDecoder(r.Body).Decode(&invet)
	//invet.UniqueId = "1"
	invets, err := getBlocks()
	if err != nil {
		fmt.Println(err)
		return
	}
	if invets != nil {
		if invets.Hash != nil {
			invet.PrevHash = invets.Hash
			invet = Blocks(invet.Tran, invets.Hash)
			collection := database.ConnectDB("blockchainlog")
			_, err := collection.InsertOne(context.TODO(), invet)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("invets===>", invets)
			invet.Blocks = invets
		}
	} else {
		invet = Blocks(invet.Tran, []byte{})
		collection := database.ConnectDB("blockchainlog")
		_, err := collection.InsertOne(context.TODO(), invet)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	collection := database.ConnectDB("blockchain")
	// insert our book model.
	opts := options.Update().SetUpsert(true)
	updateQuery := bson.M{"uniqueId": ""}
	updateData := bson.M{"$set": invet}
	_, err = collection.UpdateOne(context.TODO(), updateQuery, updateData, opts)
	if err != nil {
		fmt.Println(err)
	}
	m := make(map[string]interface{})
	m["data"] = invet

	dataB, err := json.Marshal(invet)
	if err != nil {
		W.WriteHeader(422)
		fmt.Fprintf(W, "Invalid Data")
		return
	}
	W.Header().Set("Content-Type", "application/json")
	W.WriteHeader(200)
	W.Write(dataB)
}
func getBlocks() (*BlockChain, error) {
	// set header.
	//w.Header().Set("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var stundent []BlockChain
	var stundents *BlockChain
	filter := []bson.M{}
	filter = append(filter, bson.M{"$sort": bson.M{"_id": -1}})

	collection := database.ConnectDB("blockchain")

	fmt.Println("collection===>", collection)
	fmt.Println("filter===>", filter)
	cursor, err := collection.Aggregate(context.TODO(), filter, nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	// cursor, err := context.DB.Collection(constants.COLLECTIONBILLCLAIMLOG).Aggregate(ctx.CTX, mainPipeline, nil)
	// if err != nil {
	// 	return nil, err
	// }
	// if err = err.All(ctx.CTX, &billClaimLogs); err != nil {
	// 	return nil, err
	// }
	fmt.Println("student ====>", stundent)

	if err := cursor.All(ctx, &stundent); err != nil {
		return nil, err
	}
	fmt.Println("student ====>", stundent)
	fmt.Println("students ====>", stundents)

	if len(stundent) > 0 {
		stundents = &stundent[0]
	}
	fmt.Println("students ====>", stundents)
	return stundents, nil
}
func getBlockChains(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var BlockChains []BlockChain
	var BlockChain *BlockChain
	// we get params with mux.
	var params = mux.Vars(r)

	s := params["id"]

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := []bson.M{}
	if s != "" {
		filter = append(filter, bson.M{"transaction.id": s})

	}
	collection := database.ConnectDB("blockchain")
	cursor, err := collection.Aggregate(context.TODO(), filter, nil)
	if err != nil {
		log.Panicln(err)
	}
	defer cursor.Close(ctx)
	fmt.Println("student ====>", BlockChains)

	if err := cursor.All(ctx, &BlockChains); err != nil {
		log.Panicln(err)
	}
	if len(BlockChains) > 0 {
		BlockChain = &BlockChains[0]
	}
	json.NewEncoder(w).Encode(BlockChain)
}
