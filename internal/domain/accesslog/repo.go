package accesslog

import "context"

type Repository interface {
	SaveBatch(ctx context.Context, logs []LogEntry) error
}
