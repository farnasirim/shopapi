package graphql

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadSchema(t *testing.T) {
	schemaContentFromBindata := Schema()

	// Doesn't look good doing this in a unit test
	// edit: Turns out to be really helpful. Reminds you if you've changed
	// the schema and haven't generated the binary data.
	actual, err := ioutil.ReadFile("./schema.graphql")

	if err != nil {
		assert.FailNow(t, "error reading the schema file")
	}

	assert.Equal(t, string(actual), schemaContentFromBindata)
}
