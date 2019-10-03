package tag

import (
	"bytes"
	"image"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

// fix size - 128 bytes
type ID3v1 struct {
	Type     string // Always 'TAG'
	Title    string // length 30. 30 characters of the title
	Artist   string // length 30. 30 characters of the artist name
	Album    string // length 30. 30 characters of the album name
	Year     int    // length 4. A four-digit year.
	Comment  string // length 28 or 30. The comment.
	ZeroByte byte   // length 1. If a track number is stored, this byte contains a binary 0.
	Track    byte   // length 1. The number of the track on the album, or 0. Invalid, if previous byte is not a binary 0.
	Genre    byte   // length 1. Index in a list of genres, or 255

	// another file data
	Data []byte
}

func (id3v1 *ID3v1) GetFileData() []byte {
	return id3v1.Data
}

func (id3v1 *ID3v1) String() string {
	var trackNumber string
	if id3v1.ZeroByte == 0 {
		trackNumber = "TrackNumber: " + strconv.Itoa(int(id3v1.Track)) + "\n"
	}

	return "Type: " + id3v1.Type + "\n" +
		"Title: " + id3v1.Title + "\n" +
		"Artist: " + id3v1.Artist + "\n" +
		"Album: " + id3v1.Album + "\n" +
		"Year: " + strconv.Itoa(id3v1.Year) + "\n" +
		"Comment: " + id3v1.Comment + "\n" +
		trackNumber
}

func checkID3v1(input io.ReadSeeker) TagVersion {
	// id3v1
	data, err := seekAndRead(input, -128, io.SeekEnd, 3)
	if err != nil {
		return TagVersionUndefined
	}

	marker := string(data)
	if marker == "TAG" {
		return TagVersionID3v1
	}

	return TagVersionUndefined
}

func ReadID3v1(input io.ReadSeeker) (*ID3v1, error) {
	header := ID3v1{}

	// 128 byte - Header size
	headerByte, err := seekAndRead(input, -128, io.SeekEnd, 128)
	if err != nil {
		return nil, err
	}

	// Type
	marker := string(headerByte[0:3])
	if marker != "TAG" {
		return nil, ErrorFileMarker
	}
	header.Type = marker

	// Title
	header.Title = stringBeforeZero(headerByte[3:33])

	// Artist
	header.Artist = stringBeforeZero(headerByte[33:63])

	// Album
	header.Album = stringBeforeZero(headerByte[63:93])

	// Year
	header.Year, err = strconv.Atoi(string(headerByte[93:97]))
	if err != nil {
		return nil, ErrorReadFile
	}

	// Comment
	// The track number is stored in the last two bytes of the comment field. If the comment is 29 or 30 characters long, no track number can be stored
	if headerByte[125] == 0 {
		header.Comment = stringBeforeZero(headerByte[97:125])
		header.ZeroByte = 0
		header.Track = headerByte[126]
	} else {
		header.Comment = stringBeforeZero(headerByte[97:127])
		header.ZeroByte = headerByte[125]
		header.Track = 0
	}

	// Genre
	// Index in a list of genres, or 255
	header.Genre = headerByte[127]

	// Read another file data
	_, err = input.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}
	// without header
	header.Data = data[:len(data)-128]

	return &header, nil
}

// Return string without zero characters
func stringBeforeZero(data []byte) string {
	n := bytes.IndexByte(data, 0)
	if n == -1 {
		return string(data)
	}
	return string(data[:n])
}

func (id3v1 *ID3v1) SaveFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return id3v1.Save(file)
}

func (id3v1 *ID3v1) Save(input io.WriteSeeker) error {
	_, err := input.Write(id3v1.Data)
	if err != nil {
		return err
	}

	// id3v1 marker
	_, err = input.Write([]byte("TAG"))
	if err != nil {
		return err
	}

	// Title
	err = writeString(input, id3v1.Title, 30)
	if err != nil {
		return err
	}

	// Artist
	err = writeString(input, id3v1.Artist, 30)
	if err != nil {
		return err
	}

	// Album
	err = writeString(input, id3v1.Album, 30)
	if err != nil {
		return err
	}

	// Year
	err = writeString(input, strconv.Itoa(id3v1.Year), 4)
	if err != nil {
		return err
	}

	// Track number
	if id3v1.ZeroByte != 0 {
		if len(id3v1.Comment) > 28 {
			err = writeString(input, id3v1.Comment, 30)
			if err != nil {
				return err
			}
		} else {
			// for fill track number
			err = writeString(input, id3v1.Comment, 28)
			_, err = input.Write([]byte{1, 0})
			if err != nil {
				return err
			}
		}
	} else {
		err = writeString(input, id3v1.Comment, 28)
		_, err = input.Write([]byte{0, id3v1.Track})
		if err != nil {
			return err
		}
	}

	_, err = input.Write([]byte{id3v1.Genre})
	if err != nil {
		return err
	}

	return nil
}

func writeString(input io.Writer, data string, size int) error {
	if len(data) > size {
		return ErrorWriteFile
	}

	bytesStr := make([]byte, size)
	for i, val := range data {
		bytesStr[i] = byte(val)
	}
	n, err := input.Write(bytesStr)
	if err != nil {
		return err
	}
	if n != size {
		return ErrorWriteFile
	}

	return nil
}

func (id3v1 *ID3v1) GetAllTagNames() []string {
	result := []string{"Title", "Artist", "Album", "Year", "Comment"}
	if id3v1.ZeroByte == 0 {
		result = append(result, "TrackNumber")
	}
	return result
}

func (id3v1 *ID3v1) GetVersion() TagVersion {
	return TagVersionID3v1
}

func (id3v1 *ID3v1) GetTitle() (string, error) {
	return id3v1.Title, nil
}

func (id3v1 *ID3v1) GetArtist() (string, error) {
	return id3v1.Artist, nil
}

func (id3v1 *ID3v1) GetAlbum() (string, error) {
	return id3v1.Album, nil
}

func (id3v1 *ID3v1) GetYear() (int, error) {
	return id3v1.Year, nil
}

func (id3v1 *ID3v1) GetComment() (string, error) {
	return id3v1.Comment, nil
}

func (id3v1 *ID3v1) GetGenre() (string, error) {
	genre, ok := genres[int(id3v1.Genre)]
	if !ok {
		return "", nil
	}
	return genre, nil
}

func (id3v1 *ID3v1) GetAlbumArtist() (string, error) {
	return "", ErrorUnsupportedTag
}

func (id3v1 *ID3v1) GetDate() (time.Time, error) {
	return time.Now(), ErrorUnsupportedTag
}

func (id3v1 *ID3v1) GetArranger() (string, error) {
	return "", ErrorUnsupportedTag
}

func (id3v1 *ID3v1) GetAuthor() (string, error) {
	return "", ErrorUnsupportedTag
}

func (id3v1 *ID3v1) GetBMP() (int, error) {
	return 0, ErrorUnsupportedTag
}

func (id3v1 *ID3v1) GetCatalogNumber() (int, error) {
	return 0, ErrorUnsupportedTag
}

func (id3v1 *ID3v1) GetCompilation() (string, error) {
	return "", ErrorUnsupportedTag
}

func (id3v1 *ID3v1) GetComposer() (string, error) {
	return "", ErrorUnsupportedTag
}

func (id3v1 *ID3v1) GetConductor() (string, error) {
	return "", ErrorUnsupportedTag
}

func (id3v1 *ID3v1) GetCopyright() (string, error) {
	return "", ErrorUnsupportedTag
}

func (id3v1 *ID3v1) GetDescription() (string, error) {
	return "", ErrorUnsupportedTag
}

func (id3v1 *ID3v1) GetDiscNumber() (int, int, error) {
	return 0, 0, ErrorUnsupportedTag
}

func (id3v1 *ID3v1) GetEncodedBy() (string, error) {
	return "", ErrorUnsupportedTag
}

func (id3v1 *ID3v1) GetTrackNumber() (int, int, error) {
	if id3v1.ZeroByte == 0 {
		return int(id3v1.Track), int(id3v1.Track), nil
	}
	return 0, 0, ErrorTagNotFound
}

func (id3v1 *ID3v1) GetPicture() (image.Image, error) {
	return nil, ErrorUnsupportedTag
}

func (id3v1 *ID3v1) SetTitle(title string) error {
	if len(title) > 30 {
		return ErrorIncorrectLength
	}
	id3v1.Title = title
	return nil
}

func (id3v1 *ID3v1) SetArtist(artist string) error {
	if len(artist) > 30 {
		return ErrorIncorrectLength
	}
	id3v1.Artist = artist
	return nil
}

func (id3v1 *ID3v1) SetAlbum(album string) error {
	if len(album) > 30 {
		return ErrorIncorrectLength
	}
	id3v1.Album = album
	return nil
}

func (id3v1 *ID3v1) SetYear(year int) error {
	id3v1.Year = year
	return nil
}

func (id3v1 *ID3v1) SetComment(comment string) error {
	if len(comment) > 30 {
		return ErrorIncorrectLength
	}
	if id3v1.ZeroByte == 0 && len(comment) > 28 {
		return ErrorIncorrectLength
	}
	id3v1.Comment = comment
	return nil
}

func (id3v1 *ID3v1) SetGenre(genre string) error {
	for key, val := range genres {
		if val == genre {
			id3v1.Genre = byte(key)
			return nil
		}
	}
	return ErrorIncorrectGenre
}

func (id3v1 *ID3v1) SetAlbumArtist(albumArtist string) error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) SetDate(date time.Time) error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) SetArranger(arranger string) error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) SetAuthor(author string) error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) SetBMP(bmp int) error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) SetCatalogNumber(catalogNumber int) error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) SetCompilation(compilation string) error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) SetComposer(composer string) error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) SetConductor(conductor string) error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) SetCopyright(copyright string) error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) SetDescription(description string) error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) SetDiscNumber(number int, total int) error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) SetEncodedBy(encodedBy string) error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) SetTrackNumber(number int, total int) error {
	if len(id3v1.Comment) > 28 {
		return ErrorIncorrectLength
	}
	id3v1.ZeroByte = 0
	id3v1.Track = byte(number)
	return nil
}

func (id3v1 *ID3v1) SetPicture(picture image.Image) error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) DeleteAll() error {
	id3v1.Title = ""
	id3v1.Artist = ""
	id3v1.Album = ""
	id3v1.Year = 0
	id3v1.Comment = ""
	id3v1.ZeroByte = 1 // without track number
	id3v1.Track = 0
	id3v1.Genre = 255
	return nil
}

func (id3v1 *ID3v1) DeleteTitle() error {
	id3v1.Title = ""
	return nil
}

func (id3v1 *ID3v1) DeleteArtist() error {
	id3v1.Artist = ""
	return nil
}

func (id3v1 *ID3v1) DeleteAlbum() error {
	id3v1.Album = ""
	return nil
}

func (id3v1 *ID3v1) DeleteYear() error {
	id3v1.Year = 0
	return nil
}

func (id3v1 *ID3v1) DeleteComment() error {
	id3v1.Comment = ""
	return nil
}

func (id3v1 *ID3v1) DeleteGenre() error {
	id3v1.Genre = 255
	return nil
}

func (id3v1 *ID3v1) DeleteAlbumArtist() error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) DeleteDate() error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) DeleteArranger() error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) DeleteAuthor() error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) DeleteBMP() error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) DeleteCatalogNumber() error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) DeleteCompilation() error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) DeleteComposer() error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) DeleteConductor() error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) DeleteCopyright() error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) DeleteDescription() error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) DeleteDiscNumber() error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) DeleteEncodedBy() error {
	return ErrorUnsupportedTag
}

func (id3v1 *ID3v1) DeleteTrackNumber() error {
	id3v1.ZeroByte = 1
	id3v1.Track = 0
	return nil
}

func (id3v1 *ID3v1) DeletePicture() error {
	return ErrorUnsupportedTag
}
