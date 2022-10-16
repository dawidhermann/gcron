package trigger

import "time"

func GetTicker() <-chan time.Time {
	ticker := time.NewTicker(time.Minute)
	return ticker.C
}
