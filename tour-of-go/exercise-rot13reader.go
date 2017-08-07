package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// 中間のストリーム処理を実装する
// 読み込んだ一文字に対して都度処理を行なっていく
func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, e := rot.r.Read(p)
	if e == nil {
		for i, v := range p {
			switch {
			case v >= 'A' && v <= 'Z':
				p[i] = (v-'A'+13)%26 + 'A'
			case v >= 'a' && v <= 'z':
				p[i] = (v-'a'+13)%26 + 'a'
			}
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
