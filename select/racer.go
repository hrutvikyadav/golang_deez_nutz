package racer

import (
	"fmt"
	"net/http"
	"time"
)

const tenSecTimeOut = 10 * time.Second

func Racer(playerone, playertwo string) (winner string, err error ) {
	return ConfigurableRacer(playerone, playertwo, tenSecTimeOut)
}

func ConfigurableRacer(playerone, playertwo string, timeout time.Duration) (winner string, err error) {
	select {
	case <- ping(playerone):
		return playerone, nil
	case <- ping(playertwo):
		return playertwo, nil
	case <- time.After(timeout):
		return "", fmt.Errorf("timedout after waiting %v for %s, and %s", timeout, playerone, playertwo)
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
