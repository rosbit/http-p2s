package p2s

import (
	"net/http"
	"log"
)

type qElem struct {
	w  http.ResponseWriter
	r *http.Request
	next http.HandlerFunc
	done chan bool
}

var (
	q chan *qElem
)

func serialThread(i int) {
	log.Printf("[p2s #%d] is working...\n", i)
	for {
		elem := <-q
		w, r, next, done := elem.w, elem.r, elem.next, elem.done
		next(w, r)
		done <- true
	}
}

func NarrowHttpRequest(workerNum int) func(http.ResponseWriter, *http.Request, http.HandlerFunc) {
	q = make(chan *qElem, workerNum)
	for i:=0; i<workerNum; i++ {
		go serialThread(i)
	}

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		done := make(chan bool)
		defer close(done)
		q <- &qElem{w, r, next, done}
		<-done
	}
}
