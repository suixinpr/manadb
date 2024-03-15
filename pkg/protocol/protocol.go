package protocol

import (
	"bufio"
	"encoding/binary"
	"github/suixinpr/manadb/internal/mana/datum"
	"io"
)

const (
	_ byte = iota
	MessageError
	MessageFinish
	MessageRowData
	MessageRowDesc
	MessageSQL
)

/********************************************************************************
*
*  Send
*
********************************************************************************/

func SendMessage(w *bufio.Writer, b byte) {
	w.WriteByte(b)
}

func SendInt8(w *bufio.Writer, i int8) {
	w.WriteByte(byte(i))
}

func SendInt16(w *bufio.Writer, i int16) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(i))
	w.Write(b)
}

func SendInt32(w *bufio.Writer, i int32) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(i))
	w.Write(b)
}

func SendInt64(w *bufio.Writer, i int64) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	w.Write(b)
}

func SendUint8(w *bufio.Writer, u uint8) {
	w.WriteByte(u)
}

func SendUint16(w *bufio.Writer, u uint16) {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, u)
	w.Write(b)
}

func SendUint32(w *bufio.Writer, u uint32) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, u)
	w.Write(b)
}

func SendUint64(w *bufio.Writer, u uint64) {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, u)
	w.Write(b)
}

func SendString(w *bufio.Writer, s string) {
	SendInt32(w, int32(len(s)))
	w.WriteString(s)
}

func SendDatum(w *bufio.Writer, d datum.Datum) {
	if d == nil {
		SendInt32(w, int32(-1))
	} else {
		SendInt32(w, int32(len(d)))
	}
	w.Write(d)
}

/********************************************************************************
*
*  Recv
*
********************************************************************************/

func RecvMessage(r *bufio.Reader) (byte, error) {
	return r.ReadByte()
}

func RecvInt8(r *bufio.Reader) (int8, error) {
	b, err := r.ReadByte()
	return int8(b), err
}

func RecvInt16(r *bufio.Reader) (int16, error) {
	b := make([]byte, 2)
	_, err := io.ReadFull(r, b)
	return int16(binary.BigEndian.Uint16(b)), err
}

func RecvInt32(r *bufio.Reader) (int32, error) {
	b := make([]byte, 4)
	_, err := io.ReadFull(r, b)
	return int32(binary.BigEndian.Uint32(b)), err
}

func RecvInt64(r *bufio.Reader) (int64, error) {
	b := make([]byte, 8)
	_, err := io.ReadFull(r, b)
	return int64(binary.BigEndian.Uint64(b)), err
}

func RecvUint8(r *bufio.Reader) (uint8, error) {
	return r.ReadByte()
}

func RecvUint16(r *bufio.Reader) (uint16, error) {
	b := make([]byte, 2)
	_, err := io.ReadFull(r, b)
	return binary.BigEndian.Uint16(b), err
}

func RecvUint32(r *bufio.Reader) (uint32, error) {
	b := make([]byte, 4)
	_, err := io.ReadFull(r, b)
	return binary.BigEndian.Uint32(b), err
}

func RecvUint64(r *bufio.Reader) (uint64, error) {
	b := make([]byte, 8)
	_, err := io.ReadFull(r, b)
	return binary.BigEndian.Uint64(b), err
}

func RecvString(r *bufio.Reader) (string, error) {
	n, err := RecvInt32(r)
	if err != nil {
		return "", err
	}
	b := make([]byte, n)
	_, err = io.ReadFull(r, b)
	return string(b), err
}

func RecvDatum(r *bufio.Reader) ([]byte, error) {
	n, err := RecvInt32(r)
	if err != nil {
		return nil, err
	}
	if n == -1 {
		return nil, nil
	}

	b := make([]byte, n)
	_, err = io.ReadFull(r, b)
	return b, err
}
