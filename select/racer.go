package racer

import (
	"net/http"
	"time"
)

func Racer(playerone, playertwo string) string {
	pOneTime := raceWithStopWatch(playerone)
	pTwoTime := raceWithStopWatch(playertwo)

	if pOneTime < pTwoTime {
		return playerone
	} else {
		return playertwo
	}
}

func raceWithStopWatch(p string) (time.Duration) {
	startTime := time.Now()
	http.Get(p)
	return time.Since(startTime)
}
