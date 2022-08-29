package parallel

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type (
	Input  any
	Result any
	Option func(*Config)
	Config struct {
		limit int
	}
)

func WithLimit(limit int) Option {
	return func(config *Config) {
		config.limit = limit
	}
}

func Do[I Input, R Result](ctx context.Context, inputs []I, transform func(context.Context, I) (R, error), options ...Option) ([]R, error) {
	var (
		outputs = make([]R, len(inputs))
		g       *errgroup.Group
		config  = Config{
			limit: -1,
		}
	)

	for _, opt := range options {
		opt(&config)
	}

	g, ctx = errgroup.WithContext(ctx)
	g.SetLimit(config.limit)

	for i, input := range inputs {
		i, input := i, input
		g.Go(func() error {
			output, err := transform(ctx, input)
			if err != nil {
				return err
			}
			outputs[i] = output
			return nil
		})
	}

	return outputs, nil
}
