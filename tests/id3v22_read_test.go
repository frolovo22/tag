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

	album, err := id3.GetAlbum()
	asrt.NoError(err)
	asrt.Equal("We Are Pilots", album)

	artist, err := id3.GetArtist()
	asrt.NoError(err)
	asrt.Equal("Shiny Toy Guns", artist)

	year, err := id3.GetYear()
	asrt.NoError(err)
	asrt.Equal(2006, year)
}
