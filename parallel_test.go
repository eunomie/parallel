package parallel

import (
	"context"
	"fmt"
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

func TestError(t *testing.T) {
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
			return 0, fmt.Errorf("error for el %s", el)
		}
	)

	_, err := Do[string, int](ctx, listOfThings, do, WithLimit(2))

	assert.ErrorContains(t, err, "error for el")
}
