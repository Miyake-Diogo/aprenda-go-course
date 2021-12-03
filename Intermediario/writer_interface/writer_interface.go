package main

import (
	"fmt"
)

type Writer interface {
	Write(p []byte) (n int, err error)
}

type MyWriter struct {
}

func (w *MyWriter) Write(p []byte) (n int, err error) {
	fmt.Println("MyWriter.Write")
	return len(p), nil
}

func main() {
	var writer Writer
	writer = new(MyWriter)
	writer.Write([]byte("Hello, World!"))
	fmt.Println(writer)
}
/*
- A interface writer do pacote io.
- type Writer interface { Write(p []byte) (n int, err error) }
    - pkg os:   func (f *File) Write(b []byte) (n int, err error)
    - pkg json: func NewEncoder(w io.Writer) *Encoder
- "Println [...] writes to standard output."
    - func Println [...] return Fprintln(os.Stdout, a...)
    - func Fprintln(w io.Writer, a ...interface{}) (n int, err error)
    - Stdout: NewFile(uintptr(syscall.Stdout), "/dev/stdout") (Google: Standard streams)
    - func NewFile(fd uintptr, name string) *File
    - func (f *File) Write(b []byte) (n int, err error)
- Exemplo:
    - Println
    - Fprintln os.Stdout
    - io.WriteString os.Stdout
    - Ou:
        - func Dial(network, address string) (Conn, error)
        - type Conn interface { [...] Write(b []byte) (n int, err error) [...] }

*/