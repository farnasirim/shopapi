//go:generate go-bindata -ignore=\.go -pkg=graphql -o=bindata.go .
package graphql

func Schema() string {
	assetBytes := MustAsset("schema.graphql")
	return string(assetBytes)
}
