package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBook_ToJSON(t *testing.T) {
	book := Book{Title: "The Art of Heckie", Author: "Hector Rios", ISBN: "0123445566"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"The Art of Heckie","author":"Hector Rios","isbn":"0123445566"}`, string(json), "Book JSON Marshalling wrong")
}

func TestFromJSON(t *testing.T) {
	dataBytes := []byte(`{"title":"The Art of Heckie","author":"Hector Rios","isbn":"0123445566"}`)
	book := FromJSON(dataBytes)

	assert.Equal(t, Book{Title: "The Art of Heckie", Author: "Hector Rios", ISBN: "0123445566"}, book, "Book unmarshalling is wrong")
}

