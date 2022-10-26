package digger

import "time"

type (
	Hole struct {
		bindAddr      string
		dialAddr      string
		retryInterval time.Duration
	}
)

func NewDigger(cfg *config.TokenConfig) {

}
