package graphql

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadSchema(t *testing.T) {
	schemaContentFromBindata := Schema()

	// Doesn't look good doing this in a unit test
	actual, err := ioutil.ReadFile("./schema.graphql")

	if err != nil {
		assert.FailNow(t, "error reading the schema file")
	}

	assert.Equal(t, string(actual), schemaContentFromBindata)
}
