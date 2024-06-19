package binutils

import (
	"encoding/binary"
	"os"
)

type BinaryReader struct {
	Filep *os.File
	Order binary.ByteOrder
}

func (b *BinaryReader) Seek(offset int64, whence int) (ret int64, err error) {
	return b.Filep.Seek(offset, whence)
}

func (b *BinaryReader) Read(buf any) error {
	return binary.Read(b.Filep, b.Order, buf)
}

func (b *BinaryReader) Readat(buf any, offset uint64) error {
	_, err := b.Seek(int64(offset), os.SEEK_SET)
	if err != nil {
		return err
	}

	return binary.Read(b.Filep, b.Order, buf)
}

func NewBinaryReader(path string, order binary.ByteOrder) BinaryReader {
	f, err := os.Open(path)
	check(err)

	return BinaryReader{f, order}
}
