package aes_ext

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

type ModeFunc func(b cipher.Block, iv []byte) cipher.BlockMode

type cryptor struct {
	key []byte //密钥
	iv  []byte //加密向量
}

func NewCryptor(key, iv []byte) *cryptor {
	return &cryptor{
		key: key,
		iv:  iv,
	}
}

func (a *cryptor) Encrypt(data []byte) []byte {
	block, err := aes.NewCipher(a.key)
	CheckErr(err)
	content := padding(data, block.BlockSize())
	encrypted := make([]byte, len(content))
	CheckErr(err)
	blockMode := cipher.NewCBCEncrypter(block, a.iv)
	blockMode.CryptBlocks(encrypted, content)
	return encrypted
}

func (a *cryptor) Decrypt(src []byte) []byte {
	decrypted := make([]byte, len(src))
	block, err := aes.NewCipher(a.key)
	CheckErr(err)
	blockMode := cipher.NewCBCDecrypter(block, a.iv)
	blockMode.CryptBlocks(decrypted, src)
	return unPadding(decrypted)
}

func padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func unPadding(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
