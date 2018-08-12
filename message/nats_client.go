package nats

import (
	"github.com/nats-io/go-nats"
)


// Nats 용 클라이언트 
type nats_client struct {
	URL string
	Connected boolean
	Done chan int
	buf chan string	
}


func Pub()
