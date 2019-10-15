package tests

import (
	"github.com/frolovo22/tag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFLACRead(t *testing.T) {
	asrt := assert.New(t)

	flac, err := tag.ReadFile("BeeMoved.flac")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	title, err := flac.GetTitle()
	asrt.NoError(err)
	asrt.Equal("Bee Moved", title)

	artist, err := flac.GetArtist()
	asrt.NoError(err)
	asrt.Equal("Blue Monday FM", artist)

	album, err := flac.GetAlbum()
	asrt.NoError(err)
	asrt.Equal("Bee Moved", album)

	albumArtist, err := flac.GetAlbumArtist()
	asrt.NoError(err)
	asrt.Equal("Blue Monday FM", albumArtist)
}
