package graphql

import (
	"net/http"

	"github.com/farnasirim/shopapi"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

type GraphqlService struct {
	dataService        shopapi.DataService
	GraphqlHTTPHandler http.Handler
}

func NewGrpahqlService(dataService shopapi.DataService) *GraphqlService {
	// We obtain the schema here (as opposed to receiving it from the caller)
	// because the implementation of the resolvers is tightly coupled with the
	// schema. We won't be able to generate an all new api on the fly from any
	// supplied schema. Therefore we make sure we're only allowing for our own.
	schema := Schema()
	rootResolver := &RootResolver{}
	serviceToReturn := &GraphqlService{
		dataService:        dataService,
		GraphqlHTTPHandler: &relay.Handler{Schema: graphql.MustParseSchema(schema, rootResolver)},
	}

	return serviceToReturn
}
