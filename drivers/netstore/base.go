package netstore

import "io"

type NetStore interface {
	UploadReader(string, io.Reader) error
	UploadFile(string, string) error
	CdnUrl(string) string
}
