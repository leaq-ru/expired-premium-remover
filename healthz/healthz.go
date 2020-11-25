package healthz

import "github.com/rs/zerolog"

type Healthz struct {
	listenPort string
	logger     zerolog.Logger
}
