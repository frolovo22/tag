package tag

import (
	"errors"
	"image"
	"io"
	"time"
)

type TagVersion int

const (
	TagVersionUndefined TagVersion = 0
	TagVersionID3v1     TagVersion = 1
	TagVersionID3v22    TagVersion = 2
	TagVersionID3v23    TagVersion = 3
	TagVersionID3v24    TagVersion = 4
)

var id3VersionMap = map[TagVersion]string{
	TagVersionUndefined: "",
	TagVersionID3v1:     "id3v1",
	TagVersionID3v22:    "id3v2.2",
	TagVersionID3v23:    "id3v2.3",
	TagVersionID3v24:    "id3v2.4",
}

func (v TagVersion) String() string {
	return id3VersionMap[v]
}

var (
	ErrorUnsupportedFormat = errors.New("unsupported format")
	ErrorIncorrectLength   = errors.New("tag incorrect length")
	ErrorUnsupportedTag    = errors.New("unsupported tag")
	ErrorTagNotFound       = errors.New("tag not found")
	ErrorEmptyFile         = errors.New("empty file")
	ErrorFileMarker        = errors.New("error file marker")
	ErrorReadFile          = errors.New("error read file")
	ErrorIncorrectGenre    = errors.New("incorrect genre")
	ErrorSeekFile          = errors.New("error seek file")
	ErrorWriteFile         = errors.New("error write file")
	ErrorIncorrectTag      = errors.New("incorrect tag")
)

type GetMetadata interface {
	GetAllTagNames() []string
	GetVersion() TagVersion
	GetFileData() []byte // all another file data

	GetTitle() (string, error)
	GetArtist() (string, error)
	GetAlbum() (string, error)
	GetYear() (int, error)
	GetComment() (string, error)
	GetGenre() (string, error)
	GetAlbumArtist() (string, error)
	GetDate() (time.Time, error)
	GetArranger() (string, error)
	GetAuthor() (string, error)
	GetBMP() (int, error)
	GetCatalogNumber() (string, error)
	GetCompilation() (string, error)
	GetComposer() (string, error)
	GetConductor() (string, error)
	GetCopyright() (string, error)
	GetDescription() (string, error)
	GetDiscNumber() (int, int, error) // number, total
	GetEncodedBy() (string, error)
	GetTrackNumber() (int, int, error) // number, total
	GetPicture() (image.Image, error)
}

type SetMetadata interface {
	SetTitle(title string) error
	SetArtist(artist string) error
	SetAlbum(album string) error
	SetYear(year int) error
	SetComment(comment string) error
	SetGenre(genre string) error
	SetAlbumArtist(albumArtist string) error
	SetDate(date time.Time) error
	SetArranger(arranger string) error
	SetAuthor(author string) error
	SetBMP(bmp int) error
	SetCatalogNumber(catalogNumber string) error
	SetCompilation(compilation string) error
	SetComposer(composer string) error
	SetConductor(conductor string) error
	SetCopyright(copyright string) error
	SetDescription(description string) error
	SetDiscNumber(number int, total int) error
	SetEncodedBy(encodedBy string) error
	SetTrackNumber(number int, total int) error
	SetPicture(picture image.Image) error
}

type DeleteMetadata interface {
	DeleteAll() error

	DeleteTitle() error
	DeleteArtist() error
	DeleteAlbum() error
	DeleteYear() error
	DeleteComment() error
	DeleteGenre() error
	DeleteAlbumArtist() error
	DeleteDate() error
	DeleteArranger() error
	DeleteAuthor() error
	DeleteBMP() error
	DeleteCatalogNumber() error
	DeleteCompilation() error
	DeleteComposer() error
	DeleteConductor() error
	DeleteCopyright() error
	DeleteDescription() error
	DeleteDiscNumber() error
	DeleteEncodedBy() error
	DeleteTrackNumber() error
	DeletePicture() error
}

type SaveMetadata interface {
	SaveFile(path string) error
	Save(input io.WriteSeeker) error
}

type Metadata interface {
	GetMetadata
	SetMetadata
	DeleteMetadata
	SaveMetadata
}
