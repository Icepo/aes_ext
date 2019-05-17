package aes_ext

import (
	"encoding/base64"
	"testing"
)

func TestCryptor_Encrypt(t *testing.T) {
	str := "哈哈哈哈"
	key := "12345678901234561234567890123456"
	iv := "1234567890123456"
	cryptor := NewCryptor([]byte(key), []byte(iv))
	enc := cryptor.Encrypt([]byte(str))
	t.Log(base64.StdEncoding.EncodeToString(enc))
	newStr := cryptor.Decrypt(enc)
	t.Log(string(newStr))
}

func TestCryptor_EncryptECB(t *testing.T) {
	str := "哈哈哈哈"
	key := "1234567890123456"
	cryptor := NewCryptor([]byte(key), nil)
	enc := cryptor.EncryptECB([]byte(str))
	t.Log(base64.StdEncoding.EncodeToString(enc))
	newStr := cryptor.DecryptECB(enc)
	t.Log(string(newStr))
}
