package graphql

import (
	"net/http/httptest"
	"testing"
)

func TestCreateGraphqlService(t *testing.T) {
	graphqlService := NewGrpahqlService(nil)
	server := httptest.NewServer(graphqlService.GraphqlHTTPHandler)

	defer server.Close()
}
