package tests

import (
	"github.com/frolovo22/tag"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestId3v1Change(t *testing.T) {
	asrt := assert.New(t)

	id3v1, err := tag.ReadFile("id3v1.mp3")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	err = id3v1.SetTitle("My new")
	asrt.NoError(err)

	err = id3v1.SetArtist("Dogs")
	asrt.NoError(err)

	err = id3v1.SetAlbum("best pets")
	asrt.NoError(err)

	err = id3v1.SetYear(1987)
	asrt.NoError(err)

	err = id3v1.SetComment("Tra la la")
	asrt.NoError(err)

	err = id3v1.DeleteTrackNumber()
	asrt.NoError(err)

	err = id3v1.SetGenre("Jazz")
	asrt.NoError(err)

	err = id3v1.SaveFile("id3v1.change.mp3")
	asrt.NoError(err, "save")
	if err != nil {
		return
	}

	change, err := tag.ReadFile("id3v1.change.mp3")
	asrt.NoError(err)
	if err != nil {
		return
	}

	title, err := change.GetTitle()
	asrt.NoError(err)
	asrt.Equal("My new", title)

	artist, err := change.GetArtist()
	asrt.NoError(err)
	asrt.Equal("Dogs", artist)

	album, err := change.GetAlbum()
	asrt.NoError(err)
	asrt.Equal("best pets", album)

	year, err := change.GetYear()
	asrt.NoError(err)
	asrt.Equal(1987, year)

	comment, err := change.GetComment()
	asrt.NoError(err)
	asrt.Equal("Tra la la", comment)

	genre, err := change.GetGenre()
	asrt.NoError(err)
	asrt.Equal("Jazz", genre)

	trackNumber, totalTracks, err := change.GetTrackNumber()
	asrt.Equal(tag.ErrTagNotFound, err)
	asrt.Equal(0, trackNumber)
	asrt.Equal(0, totalTracks)

	err = os.Remove("id3v1.change.mp3")
	asrt.NoError(err, "remove")
	if err != nil {
		return
	}
}
