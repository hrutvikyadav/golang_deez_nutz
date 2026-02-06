package racer

import (
	"net/http"
	"time"
)

func Racer(playerone, playertwo string) string {
	t1 := time.Now()
	http.Get(playerone)
	pOneTime := time.Since(t1)

	t2 := time.Now()
	http.Get(playertwo)
	pTwoTime := time.Since(t2)

	if pOneTime < pTwoTime {
		return playerone
	} else {
		return playertwo
	}
}
