package logger

import (
	"context"

	"github.com/rs/zerolog"
)

func FromCtx(ctx context.Context, layer, op string) zerolog.Logger {
	return zerolog.Ctx(ctx).With().
		Str("layer", layer).
		Str("op", op).
		Logger()
}
