package GolangStandardLibraryTesting

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestWriter(t *testing.T) {
	w := bufio.NewWriter(nil)
	w.WriteByte('a')
	w.WriteRune('世')
	w.WriteString("hello")
	buf := []byte{'a', 'b', 'c'}
	w.Write(buf)
	fmt.Println("Buffered:", w.Buffered())               // Buffered: 12
	fmt.Println("Available:", w.Available())             // Available 4084
	fmt.Println("AvailableBuffer:", w.AvailableBuffer()) // AvailableBuffer: []
	fmt.Println("Size:", w.Size())                       // Size: 4096
}

func TestReader(t *testing.T) {
	reader := strings.NewReader("hello你好")
	r := bufio.NewReader(reader)
	buf := make([]byte, 4)
	n, err := r.Read(buf)
	fmt.Println("Read:", n, err) // Read: 4 <nil>
	fmt.Println("buf:", buf)     // buf: [104 101 108 108]
	b, err := r.ReadByte()
	fmt.Println("ReadByte:", b, err) // ReadByte: 111 <nil>
	r.UnreadByte()                   // cancel previous read
	b, err = r.ReadByte()
	fmt.Println("ReadByte:", b, err) // ReadByte: 111 <nil>
	ru, n, err := r.ReadRune()
	fmt.Println("ReadRune:", ru, n, err) // ReadRune: 20320 3 <nil>

	newReader := strings.NewReader("你好,世界!\nhello world")
	r.Reset(newReader)

	line, isPrefix, err := r.ReadLine()
	fmt.Println("ReadSlice:", string(line), isPrefix, err) // ReadSlice: 你好,世界! false <nil>

	s, err := r.ReadString('e')
	fmt.Println("ReadString:", s, err) // ReadString: he <nil>

	fmt.Println("After ReadString Buffered:", r.Buffered()) // After ReadString Buffered: 9
	peakStr, err := r.Peek(5)
	fmt.Println("Peek:", string(peakStr), err)        // Peek: llo w <nil>
	fmt.Println("After Peak Buffered:", r.Buffered()) // After Peak Buffered: 9
	r.Discard(1)
	fmt.Println("After Discard Buffered:", r.Buffered()) // After Discard Buffered: 8
	r.UnreadByte()
	fmt.Println("After UnreadByte Buffered:", r.Buffered()) // After UnreadByte Buffered: 8

	writer := bufio.NewWriter(nil)
	r.WriteTo(writer)
	fmt.Println("Writer Buffered:", writer.Buffered())   // Writer Buffered: 8
	fmt.Println("Writer Available:", writer.Available()) // Writer Available: 4088

}
