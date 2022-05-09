package main

import (
	"log"
)

type ConnLimiter struct {
	connCurrentCount int
	bucket           chan int
}

func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		connCurrentCount: cc,
		bucket:           make(chan int, cc),
	}
}

func (cl *ConnLimiter) GetConn() bool {
	if len(cl.bucket) >= cl.connCurrentCount {
		log.Println("Reached the rate limitation")
		return false
	}
	cl.bucket <- 1
	return true
}

func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	log.Printf("New connection coming %d \n", c)
}
