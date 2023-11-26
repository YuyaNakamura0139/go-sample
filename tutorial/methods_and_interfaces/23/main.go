package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (int, error) {
	// rot13Reader構造体のrフィールドのReadメソッドを呼び出す。
	// バイトスライスで読み込んだバイト数nとエラーerrを取得
	n, err := reader.r.Read(b)
	// バイトスライスに対して13個ずらしていく
	for i := 0; i < n; i++ {
		switch {
		// A~Z
		case 'A' <= b[i] && b[i] <= 'Z':
			b[i] = 'A' + (b[i]-'A'+13)%26
		case 'a' <= b[i] && b[i] <= 'z':
			b[i] = 'a' + (b[i]-'a'+13)%26
		}
	}
	return n, err
}

func main() {
	// Reader型へキャスト
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	// rot13Reader型へ
	r := rot13Reader{s}
	// https://pkg.go.dev/io#Copy
	// func Copy(dest Write, src Reader) (written int64, err error)
	// srcでEORに達するかエラーが発生するまでsrcからdstにコピーする
	// コピーされたバイト数と、コピー中に発生した最初のエラーを返す
	io.Copy(os.Stdout, &r)
}
