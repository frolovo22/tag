package tag

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type FLAC struct {
	Blocks []*FlacMeatadataBlock
	Tags   map[string]string
	Data   []byte
}

func (F *FLAC) GetAllTagNames() []string {
	panic("implement me")
}

func (flac *FLAC) GetVersion() TagVersion {
	return TagVersionFLAC
}

func (flac *FLAC) GetFileData() []byte {
	return flac.Data
}

func (flac *FLAC) GetTitle() (string, error) {
	return flac.GetVorbisComment("TITLE")
}

func (flac *FLAC) GetArtist() (string, error) {
	return flac.GetVorbisComment("ARTIST")
}

func (flac *FLAC) GetAlbum() (string, error) {
	return flac.GetVorbisComment("ALBUM")
}

func (flac *FLAC) GetYear() (int, error) {
	return flac.GetVorbisCommentInt("YEAR")
}

func (flac *FLAC) GetComment() (string, error) {
	return flac.GetVorbisComment("COMMENT")
}

func (flac *FLAC) GetGenre() (string, error) {
	return flac.GetVorbisComment("GENRE")
}

func (flac *FLAC) GetAlbumArtist() (string, error) {
	return flac.GetVorbisComment("ALBUMARTIST")
}

func (flac *FLAC) GetDate() (time.Time, error) {
	return flac.GetVorbisCommentTime("DATE")
}

func (flac *FLAC) GetArranger() (string, error) {
	return flac.GetVorbisComment("ARRANGER")
}

func (flac *FLAC) GetAuthor() (string, error) {
	return flac.GetVorbisComment("AUTHOR")
}

func (flac *FLAC) GetBMP() (int, error) {
	return flac.GetVorbisCommentInt("BPM")
}

func (flac *FLAC) GetCatalogNumber() (string, error) {
	return flac.GetVorbisComment("CATALOGNUMBER")
}

func (flac *FLAC) GetCompilation() (string, error) {
	return flac.GetVorbisComment("COMPILATION")
}

func (flac *FLAC) GetComposer() (string, error) {
	return flac.GetVorbisComment("COMPOSER")
}

func (flac *FLAC) GetConductor() (string, error) {
	return flac.GetVorbisComment("CONDUCTOR")
}

func (flac *FLAC) GetCopyright() (string, error) {
	return flac.GetVorbisComment("COPYRIGHT")
}

func (flac *FLAC) GetDescription() (string, error) {
	return flac.GetVorbisComment("DESCRIPTION")
}

func (flac *FLAC) GetDiscNumber() (int, int, error) {
	number, err := flac.GetVorbisCommentInt("DISCNUMBER")
	if err != nil {
		return 0, 0, err
	}
	total, err := flac.GetVorbisCommentInt("DISCTOTAL")
	if err != nil {
		return 0, 0, err
	}
	return number, total, nil
}

func (flac *FLAC) GetEncodedBy() (string, error) {
	return flac.GetVorbisComment("ENCODED-BY")
}

func (flac *FLAC) GetTrackNumber() (int, int, error) {
	number, err := flac.GetVorbisCommentInt("TRACKNUMBER")
	if err != nil {
		return 0, 0, err
	}
	total, err := flac.GetVorbisCommentInt("TRACKTOTAL")
	if err != nil {
		return 0, 0, err
	}
	return number, total, nil
}

func (flac *FLAC) GetPicture() (image.Image, error) {
	pictureBlock, err := flac.GetMetadataBlockPicture()
	if err != nil {
		return nil, err
	}
	switch pictureBlock.MIME {
	case "image/jpeg":
		return jpeg.Decode(bytes.NewReader(pictureBlock.PictureData))
	case "image/png":
		return png.Decode(bytes.NewReader(pictureBlock.PictureData))
	case "-->":
		return downloadImage(string(pictureBlock.PictureData))
	}

	return nil, ErrorIncorrectTag
}

func (F *FLAC) SetTitle(title string) error {
	panic("implement me")
}

func (F *FLAC) SetArtist(artist string) error {
	panic("implement me")
}

func (F *FLAC) SetAlbum(album string) error {
	panic("implement me")
}

func (F *FLAC) SetYear(year int) error {
	panic("implement me")
}

func (F *FLAC) SetComment(comment string) error {
	panic("implement me")
}

func (F *FLAC) SetGenre(genre string) error {
	panic("implement me")
}

func (F *FLAC) SetAlbumArtist(albumArtist string) error {
	panic("implement me")
}

func (F *FLAC) SetDate(date time.Time) error {
	panic("implement me")
}

func (F *FLAC) SetArranger(arranger string) error {
	panic("implement me")
}

func (F *FLAC) SetAuthor(author string) error {
	panic("implement me")
}

func (F *FLAC) SetBMP(bmp int) error {
	panic("implement me")
}

func (F *FLAC) SetCatalogNumber(catalogNumber string) error {
	panic("implement me")
}

func (F *FLAC) SetCompilation(compilation string) error {
	panic("implement me")
}

func (F *FLAC) SetComposer(composer string) error {
	panic("implement me")
}

func (F *FLAC) SetConductor(conductor string) error {
	panic("implement me")
}

func (F *FLAC) SetCopyright(copyright string) error {
	panic("implement me")
}

func (F *FLAC) SetDescription(description string) error {
	panic("implement me")
}

func (F *FLAC) SetDiscNumber(number int, total int) error {
	panic("implement me")
}

func (F *FLAC) SetEncodedBy(encodedBy string) error {
	panic("implement me")
}

func (F *FLAC) SetTrackNumber(number int, total int) error {
	panic("implement me")
}

func (F *FLAC) SetPicture(picture image.Image) error {
	panic("implement me")
}

func (F *FLAC) DeleteAll() error {
	panic("implement me")
}

func (F *FLAC) DeleteTitle() error {
	panic("implement me")
}

func (F *FLAC) DeleteArtist() error {
	panic("implement me")
}

func (F *FLAC) DeleteAlbum() error {
	panic("implement me")
}

func (F *FLAC) DeleteYear() error {
	panic("implement me")
}

func (F *FLAC) DeleteComment() error {
	panic("implement me")
}

func (F *FLAC) DeleteGenre() error {
	panic("implement me")
}

func (F *FLAC) DeleteAlbumArtist() error {
	panic("implement me")
}

func (F *FLAC) DeleteDate() error {
	panic("implement me")
}

func (F *FLAC) DeleteArranger() error {
	panic("implement me")
}

func (F *FLAC) DeleteAuthor() error {
	panic("implement me")
}

func (F *FLAC) DeleteBMP() error {
	panic("implement me")
}

func (F *FLAC) DeleteCatalogNumber() error {
	panic("implement me")
}

func (F *FLAC) DeleteCompilation() error {
	panic("implement me")
}

func (F *FLAC) DeleteComposer() error {
	panic("implement me")
}

func (F *FLAC) DeleteConductor() error {
	panic("implement me")
}

func (F *FLAC) DeleteCopyright() error {
	panic("implement me")
}

func (F *FLAC) DeleteDescription() error {
	panic("implement me")
}

func (F *FLAC) DeleteDiscNumber() error {
	panic("implement me")
}

func (F *FLAC) DeleteEncodedBy() error {
	panic("implement me")
}

func (F *FLAC) DeleteTrackNumber() error {
	panic("implement me")
}

func (F *FLAC) DeletePicture() error {
	panic("implement me")
}

func (F *FLAC) SaveFile(path string) error {
	panic("implement me")
}

func (F *FLAC) Save(input io.WriteSeeker) error {
	panic("implement me")
}

func checkFLAC(input io.ReadSeeker) bool {
	data, err := seekAndRead(input, 0, io.SeekStart, 4)
	if err != nil {
		return false
	}

	if string(data) != "fLaC" {
		return false
	}

	return true
}

type FlacMetadataBlockType byte

/*
BLOCK_TYPE:
	0 : STREAMINFO
	1 : PADDING
	2 : APPLICATION
	3 : SEEKTABLE
	4 : VORBIS_COMMENT
	5 : CUESHEET
	6 : PICTURE
	7-126 : reserved
	127 : invalid, to avoid confusion with a frame sync code
*/
const (
	FlacStreamInfo    FlacMetadataBlockType = 0
	FlacPadding       FlacMetadataBlockType = 1
	FlacApplication   FlacMetadataBlockType = 2
	FlacSeekTable     FlacMetadataBlockType = 3
	FlacVorbisComment FlacMetadataBlockType = 4
	FlacCueSheet      FlacMetadataBlockType = 5
	FlacPicture       FlacMetadataBlockType = 6
)

type FlacMeatadataBlock struct {
	IsLast bool // Last-metadata-block flag: '1' if this block is the last metadata block before the audio blocks, '0' otherwise.
	Type   FlacMetadataBlockType
	Size   int
	Data   []byte
}

func ReadFLAC(input io.ReadSeeker) (*FLAC, error) {
	flac := FLAC{
		Tags: map[string]string{},
	}

	// FLAC identifier
	data, err := seekAndRead(input, 0, io.SeekStart, 4)
	if err != nil {
		return nil, err
	}

	if string(data) != "fLaC" {
		return nil, ErrorFileMarker
	}

	// read blocks
	for {
		block, err := readMeatadataBlock(input)
		if err != nil {
			return nil, err
		}

		// last block before audio frame
		if block.IsLast {
			break
		}

		if block.Type == FlacVorbisComment {
			comments, err := readVorbisComments(bytes.NewReader(block.Data))
			if err != nil {
				return nil, err
			}
			for _, comment := range comments {
				flac.Tags[comment.Name] = comment.Value
			}
		} else {
			flac.Blocks = append(flac.Blocks, block)
		}
	}

	// file data
	flac.Data, err = ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}

	return &flac, nil
}

func readMeatadataBlock(input io.Reader) (*FlacMeatadataBlock, error) {
	header := FlacMeatadataBlock{}

	// 4 - header size
	headerBytes, err := readBytes(input, 4)
	if err != nil {
		return nil, err
	}

	// first bit
	isLastBit := GetBit(headerBytes[0], 7)
	if isLastBit == 1 {
		header.IsLast = true
	}

	// Only 0-6 types
	// 1-7 bits
	blockType := headerBytes[0] & 0x7F
	if blockType > 6 {
		return nil, ErrorReadFile
	}
	header.Type = FlacMetadataBlockType(blockType)

	// 3 bytes size
	header.Size = ByteToInt(headerBytes[1:])

	// block data
	header.Data, err = readBytes(input, header.Size)
	if err != nil {
		return nil, err
	}
	return &header, nil
}

type VorbisComment struct {
	Name  string
	Value string
}

func (flac *FLAC) GetVorbisComment(key string) (string, error) {
	val, ok := flac.Tags[key]
	if !ok {
		return "", ErrorTagNotFound
	}
	return val, nil
}

//The comment header is decoded as follows:
//
//	1) [vendor_length] = read an unsigned integer of 32 bits
//	2) [vendor_string] = read a UTF-8 vector as [vendor_length] octets
//	3) [user_comment_list_length] = read an unsigned integer of 32 bits
//	4) iterate [user_comment_list_length] times {
//
//		5) [length] = read an unsigned integer of 32 bits
//		6) this iteration's user comment = read a UTF-8 vector as [length] octets
//
//	}
//
//	7) [framing_bit] = read a single bit as boolean
func readVorbisComments(input io.Reader) ([]VorbisComment, error) {
	result := []VorbisComment{}

	// vendor
	_, err := readLengthData(input, binary.LittleEndian)
	if err != nil {
		return nil, err
	}

	// user_comment_list_length
	var length uint32
	err = binary.Read(input, binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}

	// iterate
	for i := 0; i < int(length); i++ {
		data, err := readLengthData(input, binary.LittleEndian)
		if err != nil {
			return nil, err
		}

		// Parse data
		vorbis := strings.SplitN(string(data), "=", 2)
		if len(vorbis) != 2 {
			return nil, ErrorIncorrectTag
		}

		comment := VorbisComment{
			Name:  vorbis[0],
			Value: vorbis[1],
		}
		result = append(result, comment)
	}
	return result, nil
}

func (flac *FLAC) GetVorbisCommentInt(key string) (int, error) {
	comment, err := flac.GetVorbisComment(key)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(comment)
}

func (flac *FLAC) GetVorbisCommentTime(key string) (time.Time, error) {
	comment, err := flac.GetVorbisComment(key)
	if err != nil {
		return time.Now(), err
	}
	result, err := time.Parse("2006-01-02T15:04:05", comment)
	if err != nil {
		return time.Now(), err
	}
	return result, nil
}

func writeVorbisComment(output io.Writer, value []VorbisComment) error {
	return nil
}

type FlacMetadataBlockPicture struct {
	Type           int32
	MIME           string
	Description    string
	Width          int32
	Height         int32
	BitsPerPixel   int32
	NumberOfColors int32
	PictureData    []byte
}

func (flac *FLAC) GetMetadataBlockPicture() (*FlacMetadataBlockPicture, error) {
	for _, block := range flac.Blocks {
		if block.Type == FlacPicture {
			return readFlacPicture(bytes.NewReader(block.Data))
		}
	}

	return nil, ErrorTagNotFound
}

func readFlacPicture(input io.Reader) (*FlacMetadataBlockPicture, error) {
	var picture FlacMetadataBlockPicture

	// Picture type
	err := binary.Read(input, binary.BigEndian, &picture.Type)
	if err != nil {
		return nil, err
	}

	// MIME
	MIMEBytes, err := readLengthData(input, binary.BigEndian)
	if err != nil {
		return nil, err
	}
	picture.MIME = string(MIMEBytes)

	// Description
	DescriptionBytes, err := readLengthData(input, binary.BigEndian)
	if err != nil {
		return nil, err
	}
	picture.Description = string(DescriptionBytes)

	// Width
	err = binary.Read(input, binary.BigEndian, &picture.Width)
	if err != nil {
		return nil, err
	}

	// Height
	err = binary.Read(input, binary.BigEndian, &picture.Height)
	if err != nil {
		return nil, err
	}

	// Bits per pixel
	err = binary.Read(input, binary.BigEndian, &picture.BitsPerPixel)
	if err != nil {
		return nil, err
	}

	// Number of colors
	err = binary.Read(input, binary.BigEndian, &picture.NumberOfColors)
	if err != nil {
		return nil, err
	}

	// Picture data
	picture.PictureData, err = readLengthData(input, binary.BigEndian)
	if err != nil {
		return nil, err
	}

	return &picture, nil
}
