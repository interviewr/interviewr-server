package organization

import (
	"context"
	"time"
	"github.com/go-kit/kit/log"
)

type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return &loggingMiddleware{
			next: next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next Service
	logger log.Logger
}

func (mw loggingMiddleware) GetOrganization(ctx context.Context, id string) (org Organization, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetOrganization", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())

	return mw.next.GetOrganization(ctx, id)
}