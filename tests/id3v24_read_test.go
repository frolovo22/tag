package tests

import (
	"github.com/frolovo22/tag"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestId3v24Read(t *testing.T) {
	asrt := assert.New(t)
	id3, err := tag.ReadFile("tests/meow_id2.4.mp3")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	title, err := id3.GetTitle()
	asrt.NoError(err)
	asrt.Equal("MEOW", title)

	album, err := id3.GetAlbum()
	asrt.NoError(err)
	asrt.Equal("CatAlbum", album)

	artist, err := id3.GetArtist()
	asrt.NoError(err)
	asrt.Equal("Cute Kitten", artist)

	year, err := id3.GetYear()
	asrt.NoError(err)
	asrt.Equal(2008, year)

	comment, err := id3.GetComment()
	asrt.NoError(err)
	asrt.Equal("catcomment", comment)

	genre, err := id3.GetGenre()
	asrt.NoError(err)
	asrt.Equal("catmusic", genre)

	albumArtist, err := id3.GetAlbumArtist()
	asrt.NoError(err)
	asrt.Equal("CatAlbumArtist", albumArtist)

	date, err := id3.GetDate()
	asrt.NoError(err)
	asrt.Equal(time.Date(2008, time.September, 15, 15, 53, 0, 0, time.UTC), date)

	arranger, err := id3.GetArranger()
	asrt.NoError(err)
	asrt.Equal("CK", arranger)

	author, err := id3.GetAuthor()
	asrt.NoError(err)
	asrt.Equal("Kitten", author)

	bpm, err := id3.GetBMP()
	asrt.NoError(err)
	asrt.Equal(777, bpm)

	catalogNumber, err := id3.GetCatalogNumber()
	asrt.NoError(err)
	asrt.Equal("catalogcat", catalogNumber)

	compilation, err := id3.GetCompilation()
	asrt.NoError(err)
	asrt.Equal("catcomp", compilation)

	composer, err := id3.GetComposer()
	asrt.NoError(err)
	asrt.Equal("catcomposer", composer)

	conductor, err := id3.GetConductor()
	asrt.NoError(err)
	asrt.Equal("catconductor", conductor)

	copyright, err := id3.GetCopyright()
	asrt.NoError(err)
	asrt.Equal("2019", copyright)

	description, err := id3.GetDescription()
	asrt.NoError(err)
	asrt.Equal("subtitle", description)

	discNumber, discTotal, err := id3.GetDiscNumber()
	asrt.NoError(err)
	asrt.Equal(1, discNumber)
	asrt.Equal(7, discTotal)

	encodedBy, err := id3.GetEncodedBy()
	asrt.NoError(err)
	asrt.Equal("encodedbycat", encodedBy)

	trackNumber, trackTotal, err := id3.GetTrackNumber()
	asrt.NoError(err)
	asrt.Equal(12, trackNumber)
	asrt.Equal(12, trackTotal)

	/*picture, ok := id3.GetPicture()
	if asrt.True(ok) {
		asrt.Equal("image/jpeg", picture.Mime)
		asrt.Equal("mew mew", picture.Description)
	}*/
}
