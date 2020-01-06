package tests

import (
	"github.com/frolovo22/tag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestId3v22Read(t *testing.T) {
	asrt := assert.New(t)
	id3, err := tag.ReadFile("id3v2.2.mp3")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	title, err := id3.GetTitle()
	asrt.NoError(err)
	asrt.Equal("You Are The One", title)
}
