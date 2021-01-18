package eio

import (
	"bytes"
	"io"
	//"sync"
)

// IAIO async io expansion
type IEWriter interface {
	Write([]byte) (int, error)
	AsyncWrite([]byte, func(int, error))
	//AsyncMWrite([]byte, func(int, error))
}

// IAIO async io expansion
type IEReader interface {
	Read([]byte) (int, error)
	AsyncRead([]byte, func(int, error))
	//AyncMRead([]byte, func(int, error))
}

// ISIO string io expansion
type IEStringReader interface {
	ReadString(int) (string, error)
	AsyncReadString(int, func(string, error))
	//AsyncMReadString(func(string, error))
}

// ISIO string io expansion
type IEStringWriter interface {
	WriteString(string) error
	AsyncWriteString(string, func(error))
	//AsyncMWriteString(string, func(error))
}

type EReader struct {
	r io.Reader
}

type EWriter struct {
	w io.Writer
}

//// EReader

func NewEReader(r io.Reader) *EReader {
	ret := EReader{
		r: r,
	}
	return &ret
}

func (er *EReader) Read(b []byte) (int, error) {
	return er.r.Read(b)

}

func (er *EReader) AsyncRead(b []byte, cb func(int, error)) {
	go func() {
		i, err := er.r.Read(b)
		cb(i, err)
	}()
}

//// EWriter

func NewEWriter(w io.Writer) *EWriter {
	ret := EWriter{
		w: w,
	}

	return &ret
}

func (ew *EWriter) AsyncWrite(b []byte, cb func(int, error)) {
	go func() {
		i, err := ew.w.Write(b)
		if cb != nil {
			cb(i, err)
		}
	}()
}

func (ew *EWriter) Write(b []byte) (int, error) {
	return ew.w.Write(b)
}

func (ew *EWriter) WriteString(str string) error {
	buff := bytes.NewBufferString(str)
	_, err := buff.WriteTo(ew.w)
	return err
}

func (ew *EWriter) AsyncWriteString(str string, cb func(error)) {
	go func() {
		buff := bytes.NewBufferString(str)
		_, err := buff.WriteTo(ew.w)
		if cb != nil {
			cb(err)
		}
	}()
}
