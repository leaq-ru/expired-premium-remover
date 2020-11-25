package healthz

import "github.com/rs/zerolog"

func NewHealthz(logger zerolog.Logger, listenPort string) Healthz {
	return Healthz{
		logger:     logger,
		listenPort: listenPort,
	}
}
