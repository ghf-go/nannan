package secret

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"os"
)

// MD5String
func MD5String(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// MD5Byte
func MD5Byte(data []byte) []byte {
	h := md5.New()
	h.Write(data)
	return h.Sum(nil)
}

// MD5File 文件md5
func MD5File(path string) string {
	data, e := ioutil.ReadFile(path)
	if e != nil {
		return ""
	}
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
func MD5OsFile(f *os.File) string {
	data, e := ioutil.ReadAll(f)
	if e != nil {
		return ""
	}
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}
