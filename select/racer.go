package racer

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(playerone, playertwo string) (winner string, err error ) {
	select {
	case <- ping(playerone):
		return playerone, nil
	case <- ping(playertwo):
		return playertwo, nil
	case <- time.After(10 * time.Second):
		return "", fmt.Errorf("timedout afte waiting 10sec for %s, and %s", playerone, playertwo)
	}
}

func ping(url string) chan struct{} {
	c := make(chan struct{})
	go func(){
		http.Get(url)
		close(c)
	}()
	return c
}
