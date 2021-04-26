package gql

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/purefun/x/mongo"
)

func MarshalID(id mongo.ID) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(id.String()))
	})
}

// And the same for the unmarshaler
func UnmarshalID(v interface{}) (mongo.ID, error) {
	str, ok := v.(string)
	if !ok {
		return mongo.NilID, fmt.Errorf("id must be string")
	}
	return mongo.IDFromString(str)
}
