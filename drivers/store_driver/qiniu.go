package store_driver

import (
	"context"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
)

type Qiniu struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Domain    string
	Zone      string
	mac       *qbox.Mac
	conf      *storage.Config
}

func (s *Qiniu) getMac() *qbox.Mac {
	if s.mac == nil {
		s.mac = qbox.NewMac(s.AccessKey, s.SecretKey)
		s.conf = &storage.Config{
			UseHTTPS:      true,
			UseCdnDomains: false,
		}
		switch s.Zone {
		case "hd":
			s.conf.Zone = &storage.ZoneHuadong
		case "hb":
			s.conf.Zone = &storage.ZoneHuabei
		case "hn":
			s.conf.Zone = &storage.ZoneHuanan
		case "bm":
			s.conf.Zone = &storage.ZoneBeimei
		case "xjp":
			s.conf.Zone = &storage.ZoneXinjiapo

		}
	}
	return s.mac
}

func (s *Qiniu) UploadReader(key string, file io.Reader) error {
	putPolicy := storage.PutPolicy{
		Scope: s.Bucket,
	}
	upToken := putPolicy.UploadToken(s.getMac())
	formUploader := storage.NewFormUploader(s.conf)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	data := []byte("hello, this is qiniu cloud")
	dataLen := int64(len(data))
	err := formUploader.Put(context.Background(), &ret, upToken, key, file, dataLen, &putExtra)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret.Key, ret.Hash)

	return nil
}
func (s *Qiniu) UploadFile(string, string) error {
	return nil
}
func (s *Qiniu) CdnUrl(string) string {
	return ""
}
