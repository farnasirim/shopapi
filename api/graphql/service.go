package graphql

import (
	"context"
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

	originalHandler := &relay.Handler{
		Schema: graphql.MustParseSchema(schema, rootResolver),
	}
	serviceToReturn := &GraphqlService{
		dataService: dataService,
	}

	serviceToReturn.GraphqlHTTPHandler = &ContextWrapperHandler{
		handlerFunc: func(w http.ResponseWriter, r *http.Request) {
			contextWithDataService := context.WithValue(r.Context(), shopapi.DataServiceStr, serviceToReturn.dataService)
			contextWithDataService = context.WithValue(contextWithDataService, "a", "b")

			originalHandler.ServeHTTP(w, r.WithContext(contextWithDataService))
		},
	}

	return serviceToReturn
}

type ContextWrapperHandler struct {
	handlerFunc http.HandlerFunc
}

func (h *ContextWrapperHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handlerFunc(w, r)
}
