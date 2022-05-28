package tests

import (
	"github.com/frolovo22/tag"
	"github.com/stretchr/testify/assert"
	"image/jpeg"
	"io/ioutil"
	"os"
	"testing"
)

func TestMp4Read(t *testing.T) {
	asrt := assert.New(t)
	mp4, err := tag.ReadFile("cat_walking.mp4")
	asrt.NoError(err, "open")
	if err != nil {
		return
	}

	title, err := mp4.GetTitle()
	asrt.NoError(err)
	asrt.Equal("Cat Walking", title)

	artist, err := mp4.GetArtist()
	asrt.NoError(err)
	asrt.Equal("Red Cat", artist)

	album, err := mp4.GetAlbum()
	asrt.NoError(err)
	asrt.Equal("Travel", album)

	albumArtist, err := mp4.GetAlbumArtist()
	asrt.NoError(err)
	asrt.Equal("Album Artist Cat", albumArtist)

	year, err := mp4.GetYear()
	asrt.NoError(err)
	asrt.Equal(2019, year)

	genre, err := mp4.GetGenre()
	asrt.NoError(err)
	asrt.Equal("New Wave", genre)

	track, totalNumber, err := mp4.GetTrackNumber()
	asrt.NoError(err)
	asrt.Equal(10, track)
	asrt.Equal(0, totalNumber)

	composer, err := mp4.GetComposer()
	asrt.NoError(err)
	asrt.Equal("Composercat", composer)

	encoder, err := mp4.GetEncodedBy()
	asrt.NoError(err)
	asrt.Equal("Lavf58.29.100", encoder)

	copyright, err := mp4.GetCopyright()
	asrt.NoError(err)
	asrt.Equal("Cat", copyright)

	picture, err := mp4.GetPicture()
	asrt.NoError(err)
	out, err := ioutil.TempFile("", "mp4_test.jpg")
	asrt.NoError(err)
	defer os.Remove(out.Name())
	err = jpeg.Encode(out, picture, &jpeg.Options{
		Quality: 95,
	})
	asrt.NoError(err)
	cmp := compareFiles("cat_walking_cover.jpg", out.Name())
	asrt.Equal(true, cmp)
}
