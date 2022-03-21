package types

import (
	"bytes"
	"crypto/subtle"
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"fmt"
)

// Bytes16 is an alias for [16]byte,
// Bytes16 implements Marshal and Unmarshal.
type Bytes16 [16]byte

// Marshal obj into your JSON variable.
func (b Bytes16) Marshal() ([]byte, error) {
	return b[:], nil
}

func (b Bytes16) MarshalTo(buf []byte) (n int, err error) {
	copy(buf, b[:])
	return len(b), nil
}

// Unmarshal your JSON variable into dest.
func (b *Bytes16) Unmarshal(buf []byte) error {
	if len(buf) != 16 {
		return fmt.Errorf("invalid bytes16 (got %d bytes)", len(buf))
	}
	copy(b[:], buf)
	return nil
}

func (b Bytes16) Compare(other Bytes16) int {
	return bytes.Compare(b[:], other[:])
}

func (b Bytes16) Equal(other Bytes16) bool {
	return subtle.ConstantTimeCompare(b[:], other[:]) == 1
}

// ParseBytes is like Parse, except it parses a byte slice instead of a string.
func ParseBytes16(buf []byte) (b Bytes16, err error) {
	if len(buf) != hex.EncodedLen(len(b)) {
		err = fmt.Errorf("invalid bytes16 length: %d", len(buf))
	}
	_, err = hex.Decode(b[:], buf)
	return
}

func (b *Bytes16) UnmarshalJSON(from []byte) error {
	quote := []byte("\"")
	quoteSize := len(quote)

	if len(from) < quoteSize*2 {
		return errors.New("invalid quote notation")
	}

	if !bytes.HasPrefix(from, quote) || !bytes.HasSuffix(from, quote) {
		return errors.New("invalid quote notation")
	} else if parsed, err := ParseBytes16(from[quoteSize : len(from)-quoteSize]); err == nil {
		*b = parsed
	}
	return nil
}

func (b Bytes16) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteRune('"')
	buffer.WriteString(hex.EncodeToString(b[:]))
	buffer.WriteRune('"')
	return buffer.Bytes(), nil
}

func (b *Bytes16) Size() int {
	if b == nil {
		return 0
	}
	return 16
}

func (b *Bytes16) FromHexString(buf []byte) error {
	hexBuf := make([]byte, hex.DecodedLen(len(buf)))
	if n, err := hex.Decode(hexBuf, buf); err != nil {
		return err
	} else {
		hexBuf = hexBuf[:n]
	}
	if err := b.Unmarshal(hexBuf); err != nil {
		return err
	}
	return nil
}

func (b Bytes16) ToHexString() string {
	return hex.EncodeToString(b[:])
}

// Scan implements the Scanner interface.
func (b *Bytes16) Scan(src any) error {
	switch src.(type) {
	case string:
		b.FromHexString([]byte(src.(string)))
	default:
		return errors.New("Incompatible type")
	}
	return nil
}

// Value implements the driver Valuer interface.
func (b Bytes16) Value() (driver.Value, error) {
	return b.ToHexString(), nil
}
