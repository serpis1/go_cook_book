package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type inputData struct {
	id string
	ts int64
}

type externalData struct {
	inputData
	relatedIds []string
}

type saveResult struct {
	ids []string
	ts int64
}

func generate(filename string) <-chan string {
	c := make(chan string)
	go func() {
		file, _ := os.Open(filename)
		defer file.Close()

		reader := bufio.NewReader(file)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF { break }
			line = strings.TrimSuffix(line, "\n")
			c <- line
		}
		close(c)
	}()
	return c
}

func prepare(inp <-chan string) <-chan inputData {
	out := make(chan inputData)
	go func() {
		for id := range inp {
			input := inputData{id: id, ts: time.Now().UnixNano(), }
			log.Printf("Ready data for processing: %+v \n", input)
			out <- input
		}
		close(out)
	}()
	return out
}

func fetch(in <-chan inputData) <-chan externalData {
	out := make(chan externalData)
	go func() {
		wg := &sync.WaitGroup{}

		for input := range in {
			wg.Add(1)
			go fetchFromExternalService(input, out, wg)
		}
		wg.Wait()
		close(out)
	}()
	return out
}

func fetchFromExternalService(in inputData, out chan<- externalData, wg *sync.WaitGroup) {
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	related := make([]string, 0)
	for i := 0; i < rand.Intn(10); i++ {
		related = append(related, strconv.Itoa(i))
	}
	out <- externalData{in, related,}
	fmt.Println("Here")
	wg.Done()
}

func save(in <-chan externalData) <-chan saveResult {
	out := make(chan saveResult)
	go func() {
		const batchSize = 7
		batch := make([]string, 0)
		for input := range in {
			if len(batch) < batchSize {
				batch = append(batch, input.inputData.id)
			} else {
				out <- persistBatch(batch)
				batch = []string{input.inputData.id}
			}
		}

		if len(batch) > 0 {
			out <- persistBatch(batch)
		}

		close(out)
	}()
	return out
}

func persistBatch(batch []string) saveResult {
	return saveResult{
		ids: batch,
		ts:  time.Now().UnixNano(),
	}
}

func main()  {
	c :=
					prepare(
							generate("uuids.txt"),
						)
	for data := range c {
		log.Printf("Items saved: %v", data)
	}

}
