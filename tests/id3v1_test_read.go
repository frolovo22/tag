package tests

import (
	"github.com/stretchr/testify/assert"
	"go-tag"
	"testing"
)

func TestId3v1Read(t *testing.T) {
	asrt := assert.New(t)

	id3v1, err := tag.ReadFile("tests/id3v1.mp3")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	title, err := id3v1.GetTitle()
	asrt.NoError(err)
	asrt.Equal("TITLE1234567890123456789012345", title)

	album, err := id3v1.GetAlbum()
	asrt.NoError(err)
	asrt.Equal("ALBUM1234567890123456789012345", album)

	artist, err := id3v1.GetArtist()
	asrt.NoError(err)
	asrt.Equal("ARTIST123456789012345678901234", artist)

	year, err := id3v1.GetYear()
	asrt.NoError(err)
	asrt.Equal(2001, year)

	comment, err := id3v1.GetComment()
	asrt.NoError(err)
	asrt.Equal("COMMENT123456789012345678901", comment)

	genre, err := id3v1.GetGenre()
	asrt.NoError(err)
	asrt.Equal("Pop", genre)

	trackNumber, totalTracks, err := id3v1.GetTrackNumber()
	asrt.NoError(err)
	asrt.Equal(1, trackNumber)
	asrt.Equal(1, totalTracks)
}
