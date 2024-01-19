package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type Blockchain struct {
	genesisBlock Block
	chain        []Block
	difficulty   int
}

func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.data)
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow)
	blockHash := sha256.Sum256([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
		b.pow++
		b.hash = b.calculateHash()
	}
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
		hash:      "0",
		timestamp: time.Now(),
	}
	return Blockchain{
		genesisBlock,
		[]Block{genesisBlock},
		difficulty,
	}
}

func (b *Blockchain) addBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
		"from":   from,
		"to":     to,
		"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
		data:         blockData,
		previousHash: lastBlock.hash,
		timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

func (b Blockchain) isValid() bool {
	for i := range b.chain[1:] {
		previousBlock := b.chain[i]
		currentBlock := b.chain[i+1]
		if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
			return false
		}
	}
	return true
}

// func main() {
// 	// create a new blockchain instance with a mining difficulty of 2
// 	blockchain := CreateBlockchain(2)

// 	// record transactions on the blockchain for Alice, Bob, and John
// 	blockchain.addBlock("Alice", "Bob", 5)
// 	blockchain.addBlock("John", "Bob", 2)

// 	// check if the blockchain is valid; expecting true
// 	fmt.Println(blockchain.isValid())
// }
func timer(d time.Duration) <-chan int {
	c := make(chan int)
	go func() {
		time.Sleep(d)
		c <- 1
	}()
	return c
}
func player(table chan int) {
	fmt.Println("cccc", table)
	for {
		ball := <-table
		ball++
		fmt.Println("111table", table)

		time.Sleep(10 * time.Second)
		table <- ball
		fmt.Println("ball", ball)

	}

}

func producer(ch chan int, d time.Duration) {
	var i int
	for {
		ch <- i
		i++
		time.Sleep(d)
		fmt.Println("ch1111", ch)

	}
}
func reader(out chan int) {
	for x := range out {
		fmt.Println(x)
	}
}

func main() {
	// for i := 0; i < 24; i++ {
	// 	c := timer(1 * time.Second)
	// 	fmt.Println("cccc", c, i)
	// 	<-c
	// }

	// var Ball int
	// table := make(chan int)

	// go player(table)
	// go player(table)

	// table <- Ball
	// fmt.Println("cssssccc", Ball)

	// time.Sleep(1 * time.Minute)
	// fmt.Println("table", table)

	// <-table
	// ch := make(chan int)
	// out := make(chan int)
	// go producer(ch, 10*time.Second)
	// // go producer(ch, 250*time.Millisecond)
	// go reader(out)
	// for i := range ch {
	// 	out <- i
	// 	// outa : <-out
	// 	fmt.Println("out1111", i)
	// }
	var wg sync.WaitGroup
	wg.Add(36)
	go pool(&wg, 36, 50)
	wg.Wait()
}

func worker(tasksCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasksCh
		if !ok {
			return
		}
		d := time.Duration(task) * time.Millisecond
		time.Sleep(d)
		fmt.Println("processing task", task)
	}
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	tasksCh := make(chan int)

	for i := 0; i < workers; i++ {
		fmt.Println("vvvv task", tasksCh)
		fmt.Println("i task", i)

		go worker(tasksCh, wg)
	}
	fmt.Println("gggtasksCh task", tasksCh)

	for i := 0; i < tasks; i++ {
		fmt.Println("ddd task", tasksCh)
		fmt.Println("iss task", i)
		tasksCh <- i
	}
	fmt.Println("bbbbbbtasksCh task", tasksCh)

	close(tasksCh)
}
