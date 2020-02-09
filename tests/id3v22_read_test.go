package tests

import (
	"github.com/frolovo22/tag"
	"github.com/stretchr/testify/assert"
	"image/png"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestId3v22Read(t *testing.T) {
	asrt := assert.New(t)
	id3, err := tag.ReadFile("id3v2.2.mp3")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	asrt.Equal("*tag.ID3v22", reflect.TypeOf(id3).String())

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

	genre, err := id3.GetGenre()
	asrt.NoError(err)
	asrt.Equal("Alternative", genre)

	encodedBy, err := id3.GetEncodedBy()
	asrt.NoError(err)
	asrt.Equal("iTunes v7.0.2.16", encodedBy)

	picture, err := id3.GetPicture()
	asrt.NoError(err)
	out, err := ioutil.TempFile("", "idv22Tst.jpg")
	asrt.NoError(err)
	defer os.Remove(out.Name())
	err = png.Encode(out, picture)
	asrt.NoError(err)
	cmp := compareFiles("idv22.jpg", out.Name())
	asrt.Equal(true, cmp)
}
