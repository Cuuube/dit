package httpio

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/Cuuube/dit/pkg/fileio"
)

// Read一个页面，返回body的reader。注意需要手动close掉
func Read(uri string) (io.ReadCloser, error) {
	res, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	return res.Body, err
}

// Fetch 下载一个页面，存储为文件
func Fetch(uri string, dst string) error {
	reader, err := Read(uri)
	if err != nil {
		return nil
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	return fileio.WriteFile(fileio.AbsPath(dst), data)
}
