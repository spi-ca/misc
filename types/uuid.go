package types

import (
	"bytes"
	"crypto/rand"
	"crypto/subtle"
	"database/sql/driver"
	"encoding/hex"
	"errors"
	"fmt"
)

type UUID [16]byte

var (
	zeroUUID UUID
)

func (u UUID) Marshal() ([]byte, error) {
	return u[:], nil
}

func (u UUID) MarshalTo(buf []byte) (n int, err error) {
	if len(u) == 0 {
		return 0, nil
	}
	copy(buf, u[:])
	return len(u), nil
}
func (u *UUID) Unmarshal(buf []byte) error {
	if len(buf) != 16 {
		return fmt.Errorf("invalid UUID (got %d bytes)", len(buf))
	}
	copy(u[:], buf)
	return nil
}

func (u UUID) Compare(other UUID) int {
	return bytes.Compare(u[:], other[:])
}

func (u UUID) Equal(other UUID) bool {
	return subtle.ConstantTimeCompare(u[:], other[:]) == 1
}
func (u *UUID) UnmarshalJSON(from []byte) error {
	quote := []byte("\"")
	quoteSize := len(quote)

	if len(from) < quoteSize*2 {
		return errors.New("invalid quote notation")
	}

	if !bytes.HasPrefix(from, quote) || !bytes.HasSuffix(from, quote) {
		return errors.New("invalid quote notation")
	} else if _, err := hex.Decode(u[:], from[quoteSize:len(from)-quoteSize]); err != nil {
		return err
	}

	return nil
}

func (u UUID) MarshalJSON() ([]byte, error) {
	var buffer bytes.Buffer
	buffer.WriteRune('"')
	buffer.WriteString(hex.EncodeToString(u[:]))
	buffer.WriteRune('"')
	return buffer.Bytes(), nil
}

func (u *UUID) Size() int {
	if u == nil {
		return 0
	}
	if len(*u) == 0 {
		return 0
	}
	return 16
}

func NewUUID() (u UUID) {
	newObj := UUID{}
	newObj.Random()
	return newObj
}

func (u *UUID) UUIDFromHexString(buf []byte) error {
	hexBuf := make([]byte, hex.DecodedLen(len(buf)))
	if n, err := hex.Decode(hexBuf, buf); err != nil {
		return err
	} else {
		hexBuf = hexBuf[:n]
	}
	if err := u.Unmarshal(hexBuf); err != nil {
		return err
	}
	return nil
}

func (u UUID) ToHexString() string {
	return hex.EncodeToString(u[:])
}

// Scan implements the Scanner interface.
func (u *UUID) Scan(src any) error {
	if src == nil {
		return nil
	}

	b, ok := src.([]byte)
	if !ok {
		return errors.New("Scan source was not []bytes")
	}

	return u.UUIDFromHexString(b)
}

// Value implements the driver Valuer interface.
func (u UUID) Value() (driver.Value, error) {
	return u.ToHexString(), nil
}

func (u *UUID) Random() *UUID {
	_, _ = rand.Read(u[:])
	u[6] = (u[6] & 0x0f) | 0x40 // Version 4
	u[8] = (u[8] & 0x3f) | 0x80 // Variant is 10
	return u
}

func (u *UUID) Clear() *UUID {
	copy(u[:], zeroUUID[:])
	return u
}

func (u UUID) IsZero() bool {
	return subtle.ConstantTimeCompare(zeroUUID[:], u[:]) == 1
}
