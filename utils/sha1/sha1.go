// SHA1
package sha1

import (
	"crypto/sha1"
	"encoding/hex"
	log "gopkg.in/clog.v1"
	"io"
	"os"
	"wxapp/utils/md5"
)

// 将任意类型的变量进行SHA摘要(注意map等非排序变量造成的不同结果)
// 内部使用了md5计算，因此效率会稍微差一些，更多情况请使用 EncodeString
func Encode(v interface{}) string {
	return EncodeString(md5.Encode(v))
}

// 对字符串行SHA1摘要计算
func EncodeString(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

// 对文件内容进行SHA1摘要计算
func EncodeFile(path string) string {
	f, e := os.Open(path)
	if e != nil {
		log.Fatal(0, "file is not exist : v%", e)
	}
	h := sha1.New()
	_, e = io.Copy(h, f)
	if e != nil {
		log.Fatal(0, "file is copy false : v%", e)
	}
	return hex.EncodeToString(h.Sum(nil))
}
