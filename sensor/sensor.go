package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func IntRand(min int,max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min + 1) + min
}

func sendInt(data int,url string) {
	sdata := strconv.Itoa(data)
	resp, err := http.Post(url,"application/binary", bytes.NewBuffer([]byte(sdata)))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close();

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Dispatcher response : ",string(body))
}

func main() {
	cfg := NewDispatcherCfg();
	url := cfg.Url()
	log.Println("dispatcher url : ", url);
	for {
		sendInt(IntRand(10,400),url)
		time.Sleep(time.Duration(IntRand(100,1000)) * time.Millisecond)
	}
}