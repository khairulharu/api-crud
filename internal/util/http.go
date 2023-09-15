package util

import (
	"context"
	"time"

	"github.com/khairulharu/miniapps/domain"
)

func ResponseInterceptor(ctx context.Context, resp *domain.ApiResponse) {
	traceIdInf := ctx.Value("requestid")
	traceId := ""
	if traceIdInf != nil {
		traceId = traceIdInf.(string)
	}
	resp.TimeStamp = time.Now()
	resp.TraceID = traceId
}
