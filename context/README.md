# context

Exemplos simples de contexto

```
import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)

	defer cancel()

	go func(ctx context.Context, cancel context.CancelFunc) {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}(ctx, cancel)

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("timeout")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
```
