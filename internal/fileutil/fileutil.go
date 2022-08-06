package fileutil

import (
	"strings"
)

const defaultContentType = "text/plain"
const downloadContentType = "application/octet-stream"

var contentTypeMap = map[string]string{
	// 图片
	"png":  "image/png",
	"jpg":  "application/x-jpg",
	"jpeg": "image/jpeg",
	"tif":  "image/tiff",
	"gif":  "image/gif",
	"ico":  "image/x-icon",
	"pdf":  "application/pdf",
	"svg":  "text/xml",

	// 音视频
	"mp3": "audio/mp3",
	"mp4": "video/mpeg4",
	"mpv": "video/mpg",
	"avi": "video/avi",
	"wmv": "video/x-ms-wmv",
	"flv": "video/x-flv",

	// 文本
	"html": "text/html;charset=utf-8",
	"htm":  "text/html;charset=utf-8",
	"xml":  "application/xml;charset=utf-8",
	"json": "application/json;charset=utf-8",
	"txt":  "text/plain",
}

// SplitPathAndExt 切分路径和最终文件名
func SplitPathAndExt(src string) (string, string) {
	clips := strings.Split(src, "/")
	len := len(clips)
	return strings.Join(clips[0:len-1], "/"), clips[len-1]
}

// SplitNameAndExt 切分文件名和拓展名
func SplitNameAndExt(fileName string) (string, string) {
	var (
		name string
		ext  string
	)
	clips := strings.Split(fileName, ".")
	if len(clips) > 1 {
		name = strings.Join(clips[0:len(clips)-1], ".")
		ext = clips[len(clips)-1]
	} else {
		name = fileName
		ext = ""
	}
	return name, ext
}

// GetFileName 简单获取路径最后的文件名
func GetFileName(path string) string {
	clips := strings.Split(path, "/")
	return clips[len(clips)-1]
}

// GetContentTypeByExt 根据拓展名映射为content-type
func GetContentTypeByExt(ext string) string {
	contentType, found := contentTypeMap[ext]
	if found {
		return contentType
	}
	return defaultContentType
}

// IsStartWithRelativePath 判断是否是相对路径
func IsStartWithRelativePath(path string) bool {
	if strings.Index(path, "./") == 0 || strings.Index(path, "../") == 0 {
		return true
	}
	return false
}
