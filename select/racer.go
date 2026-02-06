package racer

import (
	"net/http"
)

func Racer(playerone, playertwo string) string {
	select {
	case <- ping(playerone):
		return playerone
	case <- ping(playertwo):
		return playertwo
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
