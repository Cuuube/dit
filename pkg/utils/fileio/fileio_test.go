package fileio

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
)

func TestJoinPath(t *testing.T) {
	fmt.Println(JoinPath("/root/", "abc"))
	fmt.Println(JoinPath("/root", "abc"))
	fmt.Println(JoinPath("/root", "/abc"))
	fmt.Println(JoinPath("/root/", "/abc"))
	fmt.Println(JoinPath("/root/", "./abc"))
	fmt.Println(JoinPath("/root", "./abc"))
	fmt.Println(JoinPath("/root/aaa/", "../abc"))
	fmt.Println(JoinPath("/root/aaa/", "../../abc"))
	fmt.Println(JoinPath("root/aaa/", "./abc"))
}

func TestAbsPath(t *testing.T) {
	fmt.Println(AbsPath("/root"))
	fmt.Println(AbsPath("aaa"))
	fmt.Println(AbsPath("./aaa"))
	fmt.Println(AbsPath("../aaa/"))
}

func TestIsExist(t *testing.T) {
	fmt.Println(IsExist("/usr"))
	fmt.Println(IsExist("/"))
}

func TestPWD(t *testing.T) {
	fmt.Println(Pwd())
}

func TestURL(t *testing.T) {
	u, e := url.Parse("https://192.168.0.1:9999/path/to/my/file#zh?v=1")
	j, _ := json.Marshal(*u)
	fmt.Printf("%+v, e: %v\n", string(j), e)

	u, e = url.Parse("ssh://user:passwd@localhost:22")
	j, _ = json.Marshal(*u)
	fmt.Printf("%+v, e: %v\n", string(j), e)

	u, e = url.Parse("jdbc:mysql://localhost:3306/database?useUnicode=true&characterEncoding=utf8&autoReconnect=true&rewriteBatchedStatements=TRUE")
	j, _ = json.Marshal(*u)
	fmt.Printf("%+v, e: %v\n", string(j), e)

	u, e = url.Parse("mailto:someone@example.com?subject=This%20is%20the%20subject&cc=someone_else@example.com&body=This%20is%20the%20body")
	j, _ = json.Marshal(*u)
	fmt.Printf("%+v, e: %v\n", string(j), e)

}
