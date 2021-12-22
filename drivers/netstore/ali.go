package netstore

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/ghf-go/nannan/gerr"
	"github.com/ghf-go/nannan/secret"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type AliOss struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
	CdnDomain       string
	CdnType         string
	CdnSecret       string
	CdnExpire       int64
}

func (ali AliOss) getBucket() *oss.Bucket {
	c, e := oss.New(ali.Endpoint, ali.AccessKeyId, ali.AccessKeySecret)
	if e != nil {
		gerr.Error(123, e.Error())
	}
	b, e := c.Bucket(ali.BucketName)
	if e != nil {
		gerr.Error(123, e.Error())
	}
	return b
}
func (ali AliOss) CdnUrl(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	switch ali.CdnType {
	case "A":
		return ali.a(path)
	case "B":
		return ali.b(path)
	case "C":
		return ali.c(path)

	}
	return ""
}
func (ali AliOss) UploadReader(k string, data io.Reader) error {
	return ali.getBucket().PutObject(k, data)
}
func (ali AliOss) UploadFile(k string, file string) error {
	return ali.getBucket().PutObjectFromFile(k, file)
}
func (ali AliOss) a(path string) string {
	baseStr := fmt.Sprintf("%d-%d-0-", time.Now().Unix()+ali.CdnExpire, rand.Int())
	k := secret.MD5String(fmt.Sprintf("%s-%s%s", path, baseStr, ali.CdnSecret))
	return fmt.Sprintf("%s%s?auth_key=%s%s", ali.CdnDomain, path, baseStr, k)
}
func (ali AliOss) b(path string) string {
	ext := time.Now().Add(time.Duration(ali.CdnExpire) * time.Second).Format("200605041502")
	k := secret.MD5String(fmt.Sprintf("%s%s%s", ali.CdnSecret, ext, path))

	return fmt.Sprintf("%s/%s/%s%s", ali.CdnDomain, ext, k, path)
}
func (ali AliOss) c(path string) string {
	t := time.Now().Unix() + ali.CdnExpire
	th := strconv.FormatInt(t, 16)
	k := secret.MD5String(fmt.Sprintf("%s%s%s", ali.CdnSecret, path, th))
	return fmt.Sprintf("%s/%s/%s%s", ali.CdnDomain, k, th, path)
}
