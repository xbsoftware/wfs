package wfs

import (
	"path/filepath"
	"strings"
)

var types map[string]string

func init() {
	types = map[string]string{
		"docx": "document",
		"doc":  "document",
		"odt":  "document",
		"xls":  "document",
		"xlsx": "document",
		"ods":  "document",
		"pdf":  "document",
		"djvu": "document",
		"djv":  "document",
		"pptx": "document",
		"ppt":  "document",

		"html":    "code",
		"htm":     "code",
		"js":      "code",
		"ts":      "code",
		"mjs":     "code",
		"json":    "code",
		"css":     "code",
		"scss":    "code",
		"sass":    "code",
		"less":    "code",
		"php":     "code",
		"phtml":   "code",
		"php3":    "code",
		"php4":    "code",
		"php5":    "code",
		"php7":    "code",
		"php-s":   "code",
		"pht":     "code",
		"phar":    "code",
		"sh":      "code",
		"coffee":  "code",
		"txt":     "code",
		"md":      "code",
		"go":      "code",
		"yml":     "code",
		"yaml":    "code",
		"xml":     "code",
		"sql":     "code",
		"sqlite3": "code",
		"sqlite":  "code",
		"db":      "code",
		"py":      "code",
		"pyc":     "code",
		"pyd":     "code",
		"pyo":     "code",
		"pyw":     "code",
		"pyz":     "code",
		"ini":     "code",
		"conf":    "code",

		"mpg": "video",
		"mp4": "video",
		"avi": "video",
		"mkv": "video",
		"ogv": "video",

		"png":  "image",
		"jpg":  "image",
		"jpeg": "image",
		"webp": "image",
		"gif":  "image",
		"tiff": "image",
		"tif":  "image",
		"svg":  "image",

		"mp3":  "audio",
		"ogg":  "audio",
		"flac": "audio",
		"wav":  "audio",

		"zip": "archive",
		"rar": "archive",
		"7z":  "archive",
		"tar": "archive",
		"gz":  "archive",
	}
}

func GetType(name string, isFolder bool) string {
	if isFolder {
		return "folder"
	}

	ext := filepath.Ext(name)
	if ext == "" {
		return "file"
	}

	ftype, ok := types[strings.ToLower(ext[1:])]
	if !ok {
		return "file"
	}

	return ftype
}
