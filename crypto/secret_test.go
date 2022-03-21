package crypto

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
	"io"
	"strings"
	"testing"
)

func TestShortEncryption1(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)
	box := NewSecretBox(key)
	source := "h"
	hasher := sha1.New()
	_, err := strings.NewReader(source).WriteTo(hasher)
	if err != nil {
		t.Fatal(err)
	}
	formerHash := hasher.Sum(nil)
	t.Log("former hash ", hex.EncodeToString(formerHash))

	encryptor, err := box.NewEncryptReader(strings.NewReader(source))
	if err != nil {
		t.Fatal(err)
	}
	trsnBuf := &bytes.Buffer{}
	_, err = trsnBuf.ReadFrom(encryptor)
	if err != nil {
		t.Fatal(err)
	}
	hasher.Reset()
	decrypted, err := box.NewDecryptWriter(hasher)
	if err != nil {
		t.Fatal(err)
	}
	_, err = trsnBuf.WriteTo(decrypted)
	if err != nil {
		t.Fatal(err)
	}
	_ = decrypted.Close()
	latterHash := hasher.Sum(nil)
	t.Log("latter hash ", hex.EncodeToString(latterHash))
	if subtle.ConstantTimeCompare(formerHash, latterHash) != 1 {
		t.Fatal("hash not matched!")
	} else {
		t.Log("Hash Match!")
	}
}
func TestShortEncryption2(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)
	box := NewSecretBox(key)
	source := "h"
	hasher := sha1.New()
	_, err := strings.NewReader(source).WriteTo(hasher)
	if err != nil {
		t.Fatal(err)
	}
	formerHash := hasher.Sum(nil)
	t.Log("former hash ", hex.EncodeToString(formerHash))

	encryptor, err := box.NewEncryptReader(strings.NewReader(source))
	if err != nil {
		t.Fatal(err)
	}
	trsnBuf := &bytes.Buffer{}
	_, err = trsnBuf.ReadFrom(encryptor)
	if err != nil {
		t.Fatal(err)
	}
	hasher.Reset()
	decrypted, err := box.NewDecryptReader(trsnBuf)
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(hasher, decrypted)
	if err != nil {
		t.Fatal(err)
	}
	latterHash := hasher.Sum(nil)
	t.Log("latter hash ", hex.EncodeToString(latterHash))
	if subtle.ConstantTimeCompare(formerHash, latterHash) != 1 {
		t.Fatal("hash not matched!")
	} else {
		t.Log("Hash Match!")
	}
}

func TestShortEncryption3(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)
	box := NewSecretBox(key)
	source := "h"
	hasher := sha1.New()
	_, err := strings.NewReader(source).WriteTo(hasher)
	if err != nil {
		t.Fatal(err)
	}
	formerHash := hasher.Sum(nil)
	t.Log("former hash ", hex.EncodeToString(formerHash))

	trsnBuf := &bytes.Buffer{}
	encryptor, err := box.NewEncryptWriter(trsnBuf)
	if err != nil {
		t.Fatal(err)
	}
	_, err = strings.NewReader(source).WriteTo(encryptor)
	if err != nil {
		t.Fatal(err)
	}
	_ = encryptor.Close()
	hasher.Reset()
	decrypted, err := box.NewDecryptReader(trsnBuf)
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(hasher, decrypted)
	if err != nil {
		t.Fatal(err)
	}
	latterHash := hasher.Sum(nil)
	t.Log("latter hash ", hex.EncodeToString(latterHash))
	if subtle.ConstantTimeCompare(formerHash, latterHash) != 1 {
		t.Fatal("hash not matched!")
	} else {
		t.Log("Hash Match!")
	}
}

func TestShortEncryption4(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)
	box := NewSecretBox(key)
	source := "h"
	hasher := sha1.New()
	_, err := strings.NewReader(source).WriteTo(hasher)
	if err != nil {
		t.Fatal(err)
	}
	formerHash := hasher.Sum(nil)
	t.Log("former hash ", hex.EncodeToString(formerHash))

	trsnBuf := &bytes.Buffer{}
	encryptor, err := box.NewEncryptWriter(trsnBuf)
	if err != nil {
		t.Fatal(err)
	}
	_, err = strings.NewReader(source).WriteTo(encryptor)
	if err != nil {
		t.Fatal(err)
	}
	_ = encryptor.Close()
	hasher.Reset()
	decrypted, err := box.NewDecryptWriter(hasher)
	if err != nil {
		t.Fatal(err)
	}
	_, err = trsnBuf.WriteTo(decrypted)
	if err != nil {
		t.Fatal(err)
	}
	_ = decrypted.Close()
	latterHash := hasher.Sum(nil)
	t.Log("latter hash ", hex.EncodeToString(latterHash))
	if subtle.ConstantTimeCompare(formerHash, latterHash) != 1 {
		t.Fatal("hash not matched!")
	} else {
		t.Log("Hash Match!")
	}
}

func TestLongEncryption1(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)

	buf := make([]byte, 1024*1024*5)
	rand.Read(buf)

	box := NewSecretBox(key)
	hasher := sha1.New()
	_, err := hasher.Write(buf)
	if err != nil {
		t.Fatal(err)
	}
	formerHash := hasher.Sum(nil)
	t.Log("former hash ", hex.EncodeToString(formerHash))

	encryptor, err := box.NewEncryptReader(bytes.NewReader(buf))
	if err != nil {
		t.Fatal(err)
	}
	trsnBuf := &bytes.Buffer{}
	_, err = trsnBuf.ReadFrom(encryptor)
	if err != nil {
		t.Fatal(err)
	}
	hasher.Reset()
	decrypted, err := box.NewDecryptWriter(hasher)
	if err != nil {
		t.Fatal(err)
	}
	_, err = trsnBuf.WriteTo(decrypted)
	if err != nil {
		t.Fatal(err)
	}
	_ = decrypted.Close()
	latterHash := hasher.Sum(nil)
	t.Log("latter hash ", hex.EncodeToString(latterHash))
	if subtle.ConstantTimeCompare(formerHash, latterHash) != 1 {
		t.Fatal("hash not matched!")
	} else {
		t.Log("Hash Match!")
	}
}
func TestLongEncryption2(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)
	box := NewSecretBox(key)

	buf := make([]byte, 1024*1024*5)
	rand.Read(buf)

	hasher := sha1.New()
	_, err := hasher.Write(buf)
	if err != nil {
		t.Fatal(err)
	}
	formerHash := hasher.Sum(nil)
	t.Log("former hash ", hex.EncodeToString(formerHash))

	encryptor, err := box.NewEncryptReader(bytes.NewReader(buf))
	if err != nil {
		t.Fatal(err)
	}
	trsnBuf := &bytes.Buffer{}
	_, err = trsnBuf.ReadFrom(encryptor)
	if err != nil {
		t.Fatal(err)
	}
	hasher.Reset()
	decrypted, err := box.NewDecryptReader(trsnBuf)
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(hasher, decrypted)
	if err != nil {
		t.Fatal(err)
	}
	latterHash := hasher.Sum(nil)
	t.Log("latter hash ", hex.EncodeToString(latterHash))
	if subtle.ConstantTimeCompare(formerHash, latterHash) != 1 {
		t.Fatal("hash not matched!")
	} else {
		t.Log("Hash Match!")
	}
}

func TestLongEncryption3(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)
	box := NewSecretBox(key)

	buf := make([]byte, 1024*1024*5)
	rand.Read(buf)

	hasher := sha1.New()
	_, err := hasher.Write(buf)
	if err != nil {
		t.Fatal(err)
	}
	formerHash := hasher.Sum(nil)
	t.Log("former hash ", hex.EncodeToString(formerHash))

	trsnBuf := &bytes.Buffer{}
	encryptor, err := box.NewEncryptWriter(trsnBuf)
	if err != nil {
		t.Fatal(err)
	}
	_, err = encryptor.Write(buf)
	if err != nil {
		t.Fatal(err)
	}
	_ = encryptor.Close()
	hasher.Reset()
	decrypted, err := box.NewDecryptReader(trsnBuf)
	if err != nil {
		t.Fatal(err)
	}
	_, err = io.Copy(hasher, decrypted)
	if err != nil {
		t.Fatal(err)
	}
	latterHash := hasher.Sum(nil)
	t.Log("latter hash ", hex.EncodeToString(latterHash))
	if subtle.ConstantTimeCompare(formerHash, latterHash) != 1 {
		t.Fatal("hash not matched!")
	} else {
		t.Log("Hash Match!")
	}
}

func TestLongEncryption4(t *testing.T) {
	key := make([]byte, 32)
	rand.Read(key)
	box := NewSecretBox(key)

	buf := make([]byte, 1024*1024*5)
	rand.Read(buf)

	hasher := sha1.New()
	_, err := hasher.Write(buf)
	if err != nil {
		t.Fatal(err)
	}
	formerHash := hasher.Sum(nil)
	t.Log("former hash ", hex.EncodeToString(formerHash))

	trsnBuf := &bytes.Buffer{}
	encryptor, err := box.NewEncryptWriter(trsnBuf)
	if err != nil {
		t.Fatal(err)
	}
	_, err = encryptor.Write(buf)
	if err != nil {
		t.Fatal(err)
	}
	_ = encryptor.Close()
	hasher.Reset()
	decrypted, err := box.NewDecryptWriter(hasher)
	if err != nil {
		t.Fatal(err)
	}
	_, err = trsnBuf.WriteTo(decrypted)
	if err != nil {
		t.Fatal(err)
	}
	_ = decrypted.Close()
	latterHash := hasher.Sum(nil)
	t.Log("latter hash ", hex.EncodeToString(latterHash))
	if subtle.ConstantTimeCompare(formerHash, latterHash) != 1 {
		t.Fatal("hash not matched!")
	} else {
		t.Log("Hash Match!")
	}
}

//
//func TestLongEncryptionDescription(t *testing.T) {
//	var key [32]byte
//	rand.Read(key[:])
//
//	buf := make([]byte, 1024*1024*5)
//	rand.Read(buf)
//	hasher := sha1.New()
//	hasher.Write(buf)
//	formerHash := hasher.Sum(nil)
//	t.Log("former hash ", hex.EncodeToString(formerHash))
//
//	var encryptBuf bytes.Buffer
//	encryptor, err := NewEncryptor(&encryptBuf, &key)
//	if err != nil {
//		t.Fatal("cannot create crypto", err)
//	}
//	_, err = encryptor.Write(buf)
//	if err != nil {
//		t.Fatal("cannot create encryptorGenerator", err)
//	}
//	err = encryptor.Close()
//
//	decryptor, err := NewDecryptor(&encryptBuf, &key)
//	if err != nil {
//		t.Fatal("cannot create decryptorGenerator", err)
//	}
//	decrypted, err := ioutil.ReadAll(decryptor)
//	if err != nil {
//		t.Fatal("error occurred while reading ", err)
//	}
//	hasher.Reset()
//	hasher.Write(decrypted)
//	latterHash := hasher.Sum(nil)
//	t.Log("former hash ", hex.EncodeToString(latterHash))
//	if subtle.ConstantTimeCompare(formerHash, latterHash) != 1 {
//		t.Fatal("hash not matched!")
//	} else {
//		t.Log("Hash Match!")
//	}
//}
