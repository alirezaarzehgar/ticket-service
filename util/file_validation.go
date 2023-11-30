package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"path/filepath"
	"strings"
)

var validImageExtensions = []string{
	"jpg", "jpeg", "png",
}

var validAttachmentExtensions = []string{
	"jpg", "jpeg", "png",
	"pdf", "docx", "xlsx", "pptx", "gif", "mp3", "mp4",
	"avi", "mov", "wmv", "wav", "aiff", "aifc", "psd",
	"indd", "eps", "zip", "rar", "tar", "gz", "bz2",
	"7z", "html", "css", "js", "xml", "json", "yaml",
	"txt", "csv", "xml", "xls", "xml", "xlsx", "odt",
	"ods", "odp", "pdf", "txt", "rtf", "dwg", "dxf",
	"dwf", "dgn", "skp", "rvt", "ipt",
}

func IsValidPath(pathname string, isImage bool) bool {
	ext := strings.ToLower(filepath.Ext(pathname))
	extensions := validAttachmentExtensions
	if isImage {
		extensions = validImageExtensions
	}

	for _, extension := range extensions {
		if "."+extension == ext {
			return true
		}
	}
	return false
}

func GetUserDir(id uint) string {
	hashByte := sha256.Sum256([]byte(fmt.Sprint(id)))
	hashStr := hex.EncodeToString(hashByte[:10])
	return hashStr
}

func CreateRandomString(salt string, len uint) string {
	rData := make([]byte, 10)
	if _, err := rand.Read(rData); err != nil {
		log.Println("rand.Read(): ", err)
	}
	hashByte := sha256.Sum256(append(rData, []byte(salt)...))
	return hex.EncodeToString(hashByte[:len])
}

func GetUniqueName(name string) string {
	return CreateRandomString(name, 5) + "+" + name
}
