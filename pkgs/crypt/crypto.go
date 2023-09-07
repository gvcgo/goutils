package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/moqsien/goutils/pkgs/gtui"
)

type Crypt struct {
	key []byte
}

func NewCrypt(password string) (c *Crypt) {
	has := md5.Sum([]byte(password))
	c = &Crypt{
		key: []byte(fmt.Sprintf("%x", has)),
	}
	return
}

func NewCrptWithKey(key []byte) (c *Crypt) {
	c = &Crypt{
		key: key,
	}
	return
}

var DefaultCrypt = &Crypt{
	key: []byte("x^)dixf&*1$free]"),
}

func (that *Crypt) pKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (that *Crypt) pKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func (that *Crypt) AesEncrypt(origData []byte) ([]byte, error) {
	block, err := aes.NewCipher(that.key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = that.pKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, that.key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func (that *Crypt) AesDecrypt(crypted []byte) ([]byte, error) {
	block, err := aes.NewCipher(that.key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, that.key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = that.pKCS7UnPadding(origData)
	return origData, nil
}

func EncodeBase64(str string) (res string) {
	if str == "" {
		return
	}
	res = base64.StdEncoding.EncodeToString([]byte(str))
	res = strings.ReplaceAll(res, "+", "-")
	res = strings.ReplaceAll(res, "/", "_")
	res = strings.ReplaceAll(res, "=", "")
	return
}

func DecodeBase64(rawStr string) (res string) {
	rawStr = strings.ReplaceAll(rawStr, "-", "+")
	rawStr = strings.ReplaceAll(rawStr, "_", "/")
	rawStr = strings.TrimSpace(rawStr)
	count := (4 - len(rawStr)%4)
	if count > 0 && strings.HasSuffix(rawStr, "=") {
		rawStr = strings.TrimSuffix(rawStr, "=")
		count = (4 - len(rawStr)%4)
	}

	if count < 4 {
		for i := 0; i < count; i++ {
			rawStr += "="
		}
	}
	if s, err := base64.StdEncoding.DecodeString(rawStr); err == nil {
		res = string(s)
	} else {
		s, err = base64.RawStdEncoding.DecodeString(rawStr)
		res = string(s)
		if err == nil {
			return
		}
		gtui.PrintError(err)
		if len(rawStr) > 5 {
			fmt.Println(rawStr[len(rawStr)-5:])
		} else {
			fmt.Println(rawStr)
		}
	}
	return
}
