// Readメソッドは、bにデータを読み込み、読み込んだバイト数とエラーを返すべきです。
// エラーは通常、読み取りに問題があった場合に返されます。
// すべてのデータが読み取られた場合、Readはio.EOFを返すべきです。
// この問題では、Readメソッドが常に'A'のASCII値（65）をbに書き込むように実装することが求められています。
// したがって、bの各バイトを'A'のASCII値に設定し、読み込んだバイト数（bの長さ）
// とnilエラーを返すようにReadメソッドを実装することがヒントとなります。

package main

type MyReader struct {
}

func (reader MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
