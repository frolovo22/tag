package tag

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"strconv"
	"strings"
	"time"
)

type Id3v2Flags byte

func (flags Id3v2Flags) String() string {
	return strconv.Itoa(int(flags))
}

func (flags Id3v2Flags) IsUnsynchronisation() bool {
	return GetBit(byte(flags), 7) == 1
}

func (flags Id3v2Flags) SetUnsynchronisation(data bool) {
	SetBit((*byte)(&flags), data, 7)
}

func (flags Id3v2Flags) HasExtendedHeader() bool {
	return GetBit(byte(flags), 6) == 1
}

func (flags Id3v2Flags) SetExtendedHeader(data bool) {
	SetBit((*byte)(&flags), data, 7)
}

func (flags Id3v2Flags) IsExperimentalIndicator() bool {
	return GetBit(byte(flags), 5) == 1
}

func (flags Id3v2Flags) SetExperimentalIndicator(data bool) {
	SetBit((*byte)(&flags), data, 7)
}

type ID3v2Frame struct {
	Key   string
	Value []byte
}

type ID3v2 struct {
	Marker     string // Always 'ID3'
	Version    TagVersion
	SubVersion int
	Flags      Id3v2Flags
	Length     int
	Frames     []ID3v2Frame

	Data []byte
}

type AttachedPicture struct {
	MIME        string
	PictureType byte
	Description string
	Data        []byte
}

func (id3v2 *ID3v2) GetAllTagNames() []string {
	var result []string
	for _, value := range id3v2.Frames {
		result = append(result, value.Key)
	}
	return result
}

func (id3v2 *ID3v2) GetVersion() TagVersion {
	return id3v2.Version
}

func (id3v2 *ID3v2) GetFileData() []byte {
	return id3v2.Data
}

func (id3v2 *ID3v2) GetTitle() (string, error) {
	return id3v2.GetString("TIT2")
}

func (id3v2 *ID3v2) GetArtist() (string, error) {
	return id3v2.GetString("TPE1")
}

func (id3v2 *ID3v2) GetAlbum() (string, error) {
	return id3v2.GetString("TALB")
}

func (id3v2 *ID3v2) GetYear() (int, error) {
	date, err := id3v2.GetTimestamp("TDOR")
	return date.Year(), err
}

func (id3v2 *ID3v2) GetComment() (string, error) {
	// id3v2
	// Comment struct must be greater than 4
	// [lang \x00 text] - comment format
	// lang - 3 symbols
	// \x00 - const, delimeter
	// text - all after
	commentStr, err := id3v2.GetString("COMM")
	if err != nil {
		return "", err
	}

	if len(commentStr) < 4 {
		return "", ErrorIncorrectLength
	}

	return commentStr[4:], nil
}

func (id3v2 *ID3v2) GetGenre() (string, error) {
	return id3v2.GetString("TCON")
}

func (id3v2 *ID3v2) GetAlbumArtist() (string, error) {
	return id3v2.GetString("TPE2")
}

func (id3v2 *ID3v2) GetDate() (time.Time, error) {
	return id3v2.GetTimestamp("TDRC")
}

func (id3v2 *ID3v2) GetArranger() (string, error) {
	return id3v2.GetString("TIPL")
}

func (id3v2 *ID3v2) GetAuthor() (string, error) {
	return id3v2.GetString("TOLY")
}

func (id3v2 *ID3v2) GetBMP() (int, error) {
	return id3v2.GetInt("TBPM")
}

func (id3v2 *ID3v2) GetCatalogNumber() (string, error) {
	return id3v2.GetStringTXXX("CATALOGNUMBER")
}

func (id3v2 *ID3v2) GetCompilation() (string, error) {
	return id3v2.GetString("TCMP")
}

func (id3v2 *ID3v2) GetComposer() (string, error) {
	return id3v2.GetString("TCOM")
}

func (id3v2 *ID3v2) GetConductor() (string, error) {
	return id3v2.GetString("TPE3")
}

func (id3v2 *ID3v2) GetCopyright() (string, error) {
	return id3v2.GetString("TCOP")
}

func (id3v2 *ID3v2) GetDescription() (string, error) {
	return id3v2.GetString("TIT3")
}

func (id3v2 *ID3v2) GetDiscNumber() (int, int, error) {
	dickNumber, err := id3v2.GetString("TPOS")
	if err != nil {
		return 0, 0, err
	}
	numbers := strings.Split(dickNumber, "/")
	if len(numbers) != 2 {
		return 0, 0, ErrorIncorrectLength
	}
	number, err := strconv.Atoi(numbers[0])
	if err != nil {
		return 0, 0, err
	}
	total, err := strconv.Atoi(numbers[1])
	if err != nil {
		return 0, 0, err
	}
	return number, total, nil
}

func (id3v2 *ID3v2) GetEncodedBy() (string, error) {
	return id3v2.GetString("TENC")
}

func (id3v2 *ID3v2) GetTrackNumber() (int, int, error) {
	track, err := id3v2.GetInt("TRCK")
	return track, track, err
}

func (id3v2 *ID3v2) GetPicture() (image.Image, error) {
	pic, err := id3v2.GetAttachedPicture()
	if err != nil {
		return nil, err
	}
	switch pic.MIME {
	case "image/jpeg":
		return jpeg.Decode(bytes.NewReader(pic.Data))
	case "image/png":
		return png.Decode(bytes.NewReader(pic.Data))
	default:
		return nil, ErrorIncorrectTag
	}
}

func (id3v2 *ID3v2) SetTitle(title string) error {
	return id3v2.SetString("TIT2", title)
}

func (id3v2 *ID3v2) SetArtist(artist string) error {
	return id3v2.SetString("TPE1", artist)
}

func (id3v2 *ID3v2) SetAlbum(album string) error {
	return id3v2.SetString("TALB", album)
}

func (id3v2 *ID3v2) SetYear(year int) error {
	curDate, err := id3v2.GetTimestamp("TDOR")
	if err != nil {
		// set only year
		return id3v2.SetTimestamp("TDOR", time.Date(year, 0, 0, 0, 0, 0, 0, time.Local))
	}
	return id3v2.SetTimestamp("TDOR", time.Date(year, curDate.Month(), curDate.Day(), curDate.Hour(), curDate.Minute(), curDate.Second(), curDate.Nanosecond(), curDate.Location()))
}

func (id3v2 *ID3v2) SetComment(comment string) error {
	return id3v2.SetString("COMM", comment)
}

func (id3v2 *ID3v2) SetGenre(genre string) error {
	return id3v2.SetString("TCON", genre)
}

func (id3v2 *ID3v2) SetAlbumArtist(albumArtist string) error {
	return id3v2.SetString("TPE2", albumArtist)
}

func (id3v2 *ID3v2) SetDate(date time.Time) error {
	return id3v2.SetTimestamp("TDRC", date)
}

func (id3v2 *ID3v2) SetArranger(arranger string) error {
	return id3v2.SetString("IPLS", arranger)
}

func (id3v2 *ID3v2) SetAuthor(author string) error {
	return id3v2.SetString("TOLY", author)
}

func (id3v2 *ID3v2) SetBMP(bmp int) error {
	return id3v2.SetInt("TBMP", bmp)
}

func (id3v2 *ID3v2) SetCatalogNumber(catalogNumber string) error {
	return id3v2.SetString("TXXX", catalogNumber)
}

func (id3v2 *ID3v2) SetCompilation(compilation string) error {
	return id3v2.SetString("TCMP", compilation)
}

func (id3v2 *ID3v2) SetComposer(composer string) error {
	return id3v2.SetString("TCOM", composer)
}

func (id3v2 *ID3v2) SetConductor(conductor string) error {
	return id3v2.SetString("TPE3", conductor)
}

func (id3v2 *ID3v2) SetCopyright(copyright string) error {
	return id3v2.SetString("TCOP", copyright)
}

func (id3v2 *ID3v2) SetDescription(description string) error {
	return id3v2.SetString("TIT3", description)
}

func (id3v2 *ID3v2) SetDiscNumber(number int, total int) error {
	return id3v2.SetString("TPOS", fmt.Sprintf("%d/%d", number, total))
}

func (id3v2 *ID3v2) SetEncodedBy(encodedBy string) error {
	return id3v2.SetString("TENC", encodedBy)
}

func (id3v2 *ID3v2) SetTrackNumber(number int, total int) error {
	// only number
	return id3v2.SetInt("TRCK", number)
}

func (id3v2 *ID3v2) SetPicture(picture image.Image) error {
	// Only PNG
	buf := new(bytes.Buffer)
	err := png.Encode(buf, picture)
	if err != nil {
		return err
	}

	attacheched, err := id3v2.GetAttachedPicture()
	if err != nil {
		// Set default params
		newPicture := AttachedPicture{
			MIME:        "image/png",
			PictureType: 2, // Other file info
			Description: "",
			Data:        buf.Bytes(),
		}
		return id3v2.SetAttachedPicture(&newPicture)
	}
	// save metainfo
	attacheched.MIME = "image/png"
	attacheched.Data = buf.Bytes()

	return id3v2.SetAttachedPicture(attacheched)
}

func (id3v2 *ID3v2) DeleteAll() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteTitle() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteArtist() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteAlbum() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteYear() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteComment() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteGenre() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteAlbumArtist() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteDate() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteArranger() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteAuthor() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteBMP() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteCatalogNumber() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteCompilation() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteComposer() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteConductor() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteCopyright() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteDescription() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteDiscNumber() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteEncodedBy() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeleteTrackNumber() error {
	panic("implement me")
}

func (id3v2 *ID3v2) DeletePicture() error {
	panic("implement me")
}

func (id3v2 *ID3v2) SaveFile(path string) error {
	panic("implement me")
}

func (id3v2 *ID3v2) Save(input io.WriteSeeker) error {
	panic("implement me")
}

func (id3v2 *ID3v2) String() string {
	result := "Marker: " + id3v2.Marker + "\n" +
		"Version: " + id3v2.Version.String() + "\n" +
		"Subversion: " + strconv.Itoa(id3v2.SubVersion) + "\n" +
		"Flags: " + id3v2.Flags.String() + "\n" +
		"Length: " + strconv.Itoa(id3v2.Length) + "\n"

	for _, frame := range id3v2.Frames {
		result += frame.Key + ": " + string(frame.Value) + "\n"
	}

	return result
}

func checkID3v2(input io.ReadSeeker) TagVersion {
	if input == nil {
		return TagVersionUndefined
	}

	// read marker (3 bytes) and version (1 byte) for ID3v2
	data, err := seekAndRead(input, 0, io.SeekStart, 4)
	if err != nil {
		return TagVersionUndefined
	}
	marker := string(data[0:3])

	// id3v2
	if marker == "ID3" {
		versionByte := data[3]
		switch versionByte {
		case 2:
			return TagVersionID3v22
		case 3:
			return TagVersionID3v23
		case 4:
			return TagVersionID3v24
		}
	}

	return TagVersionUndefined
}

func ReadID3v2(input io.ReadSeeker) (*ID3v2, error) {
	header := ID3v2{}
	if input == nil {
		return nil, errors.New("empty file")
	}

	// Seek to file start
	startIndex, err := input.Seek(0, io.SeekStart)
	if startIndex != 0 {
		return nil, errors.New("error seek file")
	}

	if err != nil {
		return nil, err
	}

	// Header size
	headerByte := make([]byte, 10)
	nReaded, err := input.Read(headerByte)
	if err != nil {
		return nil, err
	}
	if nReaded != 10 {
		return nil, errors.New("error header length")
	}

	// Marker
	marker := string(headerByte[0:3])
	if marker != "ID3" {
		return nil, errors.New("error file marker")
	}
	header.Marker = marker

	// Version
	versionByte := headerByte[3]
	switch versionByte {
	case 2:
		header.Version = TagVersionID3v22
	case 3:
		header.Version = TagVersionID3v23
	case 4:
		header.Version = TagVersionID3v24
	default:
		return nil, errors.New("error file version")
	}

	// Sub version
	subVersionByte := headerByte[4]
	header.SubVersion = int(subVersionByte)

	// Flags
	header.Flags = Id3v2Flags(headerByte[5])

	// Length
	length := ByteToIntSynchsafe(headerByte[6:10])
	header.Length = length

	// Extended headers
	header.Frames = []ID3v2Frame{}
	curRead := 0
	for curRead < length {
		bytesExtendedHeader := make([]byte, 10)
		nReaded, err = input.Read(bytesExtendedHeader)
		if err != nil {
			return nil, err
		}
		if nReaded != 10 {
			return nil, errors.New("error extended header length")
		}
		// Frame identifier
		key := string(bytesExtendedHeader[0:4])

		/*if bytesExtendedHeader[0] == 0 && bytesExtendedHeader[1] == 0 && bytesExtendedHeader[2] == 0 && bytesExtendedHeader[3] == 0 {
			break
		}*/

		// Frame data size
		size := ByteToInt(bytesExtendedHeader[4:8])

		bytesExtendedValue := make([]byte, size)
		nReaded, err = input.Read(bytesExtendedValue)
		if err != nil {
			return nil, err
		}
		if nReaded != size {
			return nil, errors.New("error extended value length")
		}

		header.Frames = append(header.Frames, ID3v2Frame{
			key,
			bytesExtendedValue,
		})

		curRead += 10 + size
	}

	// TODO
	if curRead != length {
		return nil, errors.New("error extended frames")
	}
	return &header, nil
}

func (id3v2 *ID3v2) GetString(name string) (string, error) {
	for _, val := range id3v2.Frames {
		if val.Key == name {
			return GetString(val.Value)
		}
	}
	return "", ErrorTagNotFound
}

func (id3v2 *ID3v2) SetString(name string, value string) error {
	frame := ID3v2Frame{
		Key:   name,
		Value: SetString(value),
	}

	for i, val := range id3v2.Frames {
		if val.Key == name {
			id3v2.Frames[i] = frame
			return nil
		}
	}
	id3v2.Frames = append(id3v2.Frames, frame)
	return nil
}

func (id3v2 *ID3v2) GetTimestamp(name string) (time.Time, error) {
	str, err := id3v2.GetString(name)
	if err != nil {
		return time.Now(), err
	}
	result, err := time.Parse("2006-01-02T15:04:05", str)
	if err != nil {
		return time.Now(), err
	}
	return result, nil
}

func (id3v2 *ID3v2) SetTimestamp(name string, value time.Time) error {
	str := value.Format("2006-01-02T15:04:05")
	return id3v2.SetString(name, str)
}

func (id3v2 *ID3v2) GetInt(name string) (int, error) {
	intStr, err := id3v2.GetString(name)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(intStr)
}

func (id3v2 *ID3v2) SetInt(name string, value int) error {
	return id3v2.SetString(name, strconv.Itoa(value))
}

func (id3v2 *ID3v2) GetAttachedPicture() (*AttachedPicture, error) {
	var picture AttachedPicture

	picStr, err := id3v2.GetString("APIC")
	if err != nil {
		return nil, err
	}
	values := strings.SplitN(picStr, "\x00", 3)
	if len(values) != 3 {
		return nil, ErrorIncorrectTag
	}

	// MIME
	picture.MIME = values[0]

	// Type
	if len(values[1]) == 0 {
		return nil, ErrorIncorrectTag
	}
	picture.PictureType = values[1][0]

	// Description
	picture.Description = values[1][1:]

	// Image data
	picture.Data = []byte(values[2])

	return &picture, nil
}

func (id3v2 *ID3v2) SetAttachedPicture(picture *AttachedPicture) error {
	// set UTF-8
	result := []byte{0}

	// MIME type
	result = append(result, []byte(picture.MIME)...)
	result = append(result, 0x00)

	// Picture type
	result = append(result, picture.PictureType)

	// Picture description
	result = append(result, []byte(picture.Description)...)
	result = append(result, 0x00)

	// Picture data
	result = append(result, picture.Data...)

	return id3v2.SetString("APIC", string(result))
}

func (id3v2 *ID3v2) DeleteTag(name string) error {
	index := -1
	for i, val := range id3v2.Frames {
		if val.Key == name {
			index = i
			break
		}
	}
	// already deleted
	if index == -1 {
		return nil
	}
	id3v2.Frames = append(id3v2.Frames[:index], id3v2.Frames[index+1:]...)
	return nil
}

// Header for 'User defined text information frame'
// Text encoding     $xx
// Description       <text string according to encoding> $00 (00)
// Value             <text string according to encoding>
func (id3v2 *ID3v2) GetStringTXXX(name string) (string, error) {
	for _, val := range id3v2.Frames {
		if val.Key == "TXXX" {
			str, err := GetString(val.Value)
			if err != nil {
				return "", err
			}
			info := strings.SplitN(str, "\x00", 2)
			if len(info) != 2 {
				return "", ErrorIncorrectTag
			}
			if info[0] == name {
				return info[1], nil
			}
		}
	}
	return "", ErrorTagNotFound
}

func (id3v2 *ID3v2) SetStringTXXX(name string, value string) error {
	result := ID3v2Frame{
		Key:   "TXXX",
		Value: SetString(name + "\x00" + value),
	}

	// find tag
	for i, val := range id3v2.Frames {
		if val.Key == "TXXX" {
			str, err := GetString(val.Value)
			if err != nil {
				continue
			}
			info := strings.SplitN(str, "\x00", 2)
			if len(info) != 2 {
				continue
			}
			if info[0] == name {
				id3v2.Frames[i] = result
				return nil
			}
		}
	}

	id3v2.Frames = append(id3v2.Frames, result)
	return nil
}

func (id3v2 *ID3v2) GetIntTXXX(name string) (int, error) {
	str, err := id3v2.GetStringTXXX(name)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(str)
}
