package main

import (
	"context"
	"os"
	// "testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	ctx := context.Background()
	lg := zerolog.New(os.Stderr).With().Caller().Str("site", "main").Timestamp().Logger()
	lgCtx := lg.WithContext(ctx)
	// lg1 := log.Ctx(lgCtx).With().Logger()

	logFn(lgCtx)

	lg.Debug().Msg("logging in main")
}

func logFn(ctx context.Context) {
	lgFn := log.Ctx(ctx).With().Str("site", "function").Logger()

	lgFn.Debug().Msg("logging in function")
}
