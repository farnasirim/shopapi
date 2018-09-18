package graphql

import (
	"net/http/httptest"
	"testing"
)

func TestCreateGraphqlService(t *testing.T) {
	graphqlService := NewGrpahqlService(&struct{}{})
	server := httptest.NewServer(graphqlService.GraphqlHTTPHandler)

	defer server.Close()
}
