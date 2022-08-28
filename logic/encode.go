package logic

import (
	"bytes"
	"io"

	"golang.org/x/xerrors"
)

const PlantUMLAlphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-_"

func EncodeBytes(b []byte) (string, error) {

	var err error
	buf := bytes.NewBuffer(make([]byte, len(b)))

	leng := len(b)

	for idx := 0; idx < leng; idx = idx + 3 {
		b1 := b[idx]
		b2 := b[idx+1]
		b3 := b[idx+2]
		err = writeBytes(buf, b1, b2, b3)
		if err != nil {
			return "", xerrors.Errorf("error: %w", err)
		}
	}

	switch leng % 3 {
	case 1:
		err = writeBytes(buf, b[leng-1], 0, 0)
	case 2:
		err = writeBytes(buf, b[leng-2], b[leng-1], 0)
	}
	if err != nil {
		return "", xerrors.Errorf("error: %w", err)
	}

	return changeString(buf.Bytes())
}

func changeString(data []byte) (string, error) {

	buf := bytes.NewBuffer(make([]byte, len(data)))
	for _, b := range data {
		if b >= 64 {
			return "", xerrors.Errorf("byte over 64")
		}
		buf.WriteByte(PlantUMLAlphabet[b])
	}

	return buf.String(), nil
}

func writeBytes(w io.ByteWriter, b1, b2, b3 byte) error {

	w1 := 0x3F & (b1 >> 2)
	w2 := 0x3F&((0x3&b1)<<4) | (b2 >> 4)
	w3 := 0x3F&((0xF&b2)<<2) | (b3 >> 6)
	w4 := 0x3F & (0x3f & b3)

	err := w.WriteByte(w1)
	if err != nil {
		return xerrors.Errorf("WriteByte(1) error: %w", err)
	}
	err = w.WriteByte(w2)
	if err != nil {
		return xerrors.Errorf("WriteByte(2) error: %w", err)
	}
	err = w.WriteByte(w3)
	if err != nil {
		return xerrors.Errorf("WriteByte(3) error: %w", err)
	}
	err = w.WriteByte(w4)
	if err != nil {
		return xerrors.Errorf("WriteByte(4) error: %w", err)
	}
	return nil
}
