package ratelimit

type Limiter interface {
	BlockingGet() //Blocking until get a ticket
	Get() bool    //return true if get a ticket
}
