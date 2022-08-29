package parallel

import (
	"context"
	"testing"

	"gotest.tools/assert"
)

func TestDo(t *testing.T) {
	var (
		ctx = context.Background()

		listOfThings = []string{
			"lorem",
			"ipsum",
			"dolor",
			"sit",
			"amet",
		}

		do = func(_ context.Context, el string) (int, error) {
			return len(el), nil
		}
	)

	result, err := Do[string, int](ctx, listOfThings, do, WithLimit(2))

	assert.NilError(t, err)
	assert.DeepEqual(t, []int{5, 5, 5, 3, 4}, result)
}
