package tag

import (
	"io"
)

func seekAndRead(input io.ReadSeeker, offset int64, whence int, read int) ([]byte, error) {
	if input == nil {
		return nil, ErrorEmptyFile
	}

	_, err := input.Seek(offset, whence)
	if err != nil {
		return nil, ErrorSeekFile
	}

	data := make([]byte, read)
	nReaded, err := input.Read(data)
	if err != nil {
		return nil, err
	}
	if nReaded != read {
		return nil, ErrorReadFile
	}

	return data, nil
}
