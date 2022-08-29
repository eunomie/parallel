package parallel

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type (
	Input  any
	Result any
)

func Do[I Input, R Result](ctx context.Context, inputs []I, transform func(context.Context, I) (R, error)) ([]R, error) {
	var (
		outputs = make([]R, len(inputs))
		g       *errgroup.Group
	)

	g, ctx = errgroup.WithContext(ctx)
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
