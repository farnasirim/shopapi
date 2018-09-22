package graphql

import (
	"context"

	"github.com/farnasirim/shopapi"
)

func dataServiceFromContext(ctx context.Context) shopapi.DataService {
	return ctx.Value(shopapi.DataServiceStr).(shopapi.DataService)
}
