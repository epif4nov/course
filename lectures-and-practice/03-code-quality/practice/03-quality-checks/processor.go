package qualitychecks

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

// Total параллельно суммирует значения и обновляет счётчик обработанных элементов.
func Total(ctx context.Context, values []int) int {
	if len(values) == 0 {
		return 0
	}

	go reportProgress(ctx)

	var total atomic.Int64
	var processed atomic.Int64
	var workers sync.WaitGroup

	for _, value := range values {
		workers.Add(1)
		go func() {
			defer workers.Done()
			total.Add(int64(value))
			processed.Add(1)
		}()
	}

	workers.Wait()
	return int(total.Load())
}

func reportProgress(ctx context.Context) {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Здесь могла бы публиковаться метрика о ходе обработки.
		}
	}
}
