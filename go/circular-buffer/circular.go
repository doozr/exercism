// Package circular contains a circular buffer implementation
package circular

import "fmt"

var testVersion = 4

// Buffer represents a circular buffer
//
// The `read` value specifies the index of the "start" of the ring,
// which is to say, the oldest data not yet read.
//
// The `write` value is an offset from the read value to indicate
// how much data has been written ahead of the `read` index.
//
// If the `write` count hits the length of the buffer, no more
// data can be written until some is read.
//
// If the `write` count is zero, no more data can be read until
// some is written.
type Buffer struct {
	buffer []byte
	read   int
	write  int
}

// NewBuffer creates a Buffer instance of the specified size
func NewBuffer(size int) *Buffer {
	buffer := Buffer{
		buffer: make([]byte, size, size),
	}
	return &buffer
}

func (buf *Buffer) mapIndex(ix int) int {
	return ix % len(buf.buffer)
}

// IsFull returns true if the buffer is full
func (buf *Buffer) IsFull() bool {
	return buf.write >= len(buf.buffer)
}

// ReadByte returns the next byte in the buffer or errors if exhausted
func (buf *Buffer) ReadByte() (b byte, err error) {
	if buf.write <= 0 {
		err = fmt.Errorf("Buffer exhausted")
		return
	}
	b = buf.buffer[buf.read]
	buf.read = buf.mapIndex(buf.read + 1)
	buf.write--
	return
}

// WriteByte writes a byte to the buffer or errors if full
func (buf *Buffer) WriteByte(c byte) (err error) {
	if buf.IsFull() {
		err = fmt.Errorf("Buffer overrun")
		return
	}
	buf.buffer[buf.mapIndex(buf.read+buf.write)] = c
	buf.write++
	return
}

// Overwrite the oldest data in the buffer if full, or write a byte if not
// Does not change the `write` counter if full as the buffer remains
// full, but with the oldest data replaced with the newest data.
func (buf *Buffer) Overwrite(c byte) {
	if !buf.IsFull() {
		buf.WriteByte(c)
		return
	}
	buf.buffer[buf.read] = c
	buf.read = buf.mapIndex(buf.read + 1)
}

// Reset the buffer to empty
func (buf *Buffer) Reset() {
	*buf = *NewBuffer(len(buf.buffer))
}
