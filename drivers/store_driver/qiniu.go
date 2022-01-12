package store_driver

import "io"
import "github.com/qiniu/go-sdk/v7/auth/qbox"

type Qiniu struct {
	AccessKey string
	SecretKey string
	Bucket    string
	mac       *qbox.Mac
}

func (s *Qiniu) get() *qbox.Mac {
	if s.mac == nil {
		s.mac = qbox.NewMac(s.AccessKey, s.SecretKey)
	}
	return s.mac
}

func (s *Qiniu) UploadReader(string, io.Reader) error {
	return nil
}
func (s *Qiniu) UploadFile(string, string) error {
	return nil
}
func (s *Qiniu) CdnUrl(string) string {
	return ""
}
