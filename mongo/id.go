package mongo

import (
	"encoding/hex"
	"encoding/json"
	"errors"

	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx/bsoncore"
)

type ID string

var NilID = ID(primitive.NilObjectID.Hex())

var EmptyID = ID("")

func NewID() ID {
	return ID(primitive.NewObjectID().Hex())
}

func IDFromString(id string) (ID, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ID(""), err
	}
	return ID(objectID.Hex()), nil
}

func IDFromInterface(id interface{}) (ID, error) {
	switch id.(type) {
	default:
		return NilID, errors.New("IDFromInterface: unhandled type")
	case primitive.ObjectID:
		return IDFromString(id.(primitive.ObjectID).Hex())
	case string:
		return IDFromString(id.(string))
	}
}

func (id ID) String() string {
	return string(id)
}

func (id ID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	b, err := hex.DecodeString(string(id))
	if err != nil {
		return bsontype.ObjectID, bsoncore.AppendObjectID(nil, primitive.NilObjectID), err
	}
	if len(b) != 12 {
		return bsontype.ObjectID, bsoncore.AppendObjectID(nil, primitive.NilObjectID), primitive.ErrInvalidHex
	}
	var oid [12]byte
	copy(oid[:], b[:])
	return bsontype.ObjectID, bsoncore.AppendObjectID(nil, oid), nil
}

func (id *ID) UnmarshalBSONValue(t bsontype.Type, val []byte) error {
	var oid [12]byte
	copy(oid[:], val[:])

	*id = ID(primitive.ObjectID(oid).Hex())
	return nil
}

func (id ID) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(id))
}

func (i *ID) UnmarshalJSON(val []byte) error {
	var idStr string
	err := json.Unmarshal(val, &idStr)
	if err != nil {
		return err
	}
	*i = ID(idStr)
	return nil
}
