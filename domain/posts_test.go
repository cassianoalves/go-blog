package domain

import (
	"fmt"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestPost_ToJsonReader(t *testing.T) {
	assert := assert2.New(t)
	postIn := Post{
		Title:   "Some Subject",
		Author:  "John Smith",
		Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
	}

	gotReader := postIn.ToJsonReader()
	gotJsonBytes := make([]byte, 1024)
	read, err := gotReader.Read(gotJsonBytes)
	if err != nil {
		return
	}

	assert.Equal(115, read)
	jsonString := fmt.Sprintf("%.*s", read, string(gotJsonBytes))
	assert.JSONEq("{\"title\":\"Some Subject\",\"content\":\"Lorem ipsum dolor sit amet, consectetur adipiscing elit.\",\"author\":\"John Smith\"}", jsonString)

}
