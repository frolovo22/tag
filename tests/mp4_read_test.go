package tests

import (
	"github.com/frolovo22/tag"
	"github.com/stretchr/testify/assert"
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
}
