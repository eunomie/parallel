# Parallel

```go
import (
	"context"

	"github.com/eunomie/parallel"
)

type InputType struct {
	//...
}

type OutputType struct {
	//...
}

func test(ctx context.Context, inputs []InputType) {
	err, result := parallel.Do[InputType, OutputType](ctx, inputs, func(context.Context, InputType) (OutputType, error) {
		return OutputType{/* ... */}, nil
	})
} 
```
