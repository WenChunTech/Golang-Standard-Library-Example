# Bufio

## 主要包含的功能

### Reader

- `func (b *Reader) Read(p []byte) (n int, err error)`: 读取n个字节到p中
- `func (b *Reader) ReadByte() (byte, error)`: 读取一个字节
- `func (b *Reader) ReadRune() (r rune, size int, err error)`: 读取一个size大小的rune
- `func (b *Reader) UnreadByte() error`: 撤销上一次读取的字节
- `func (b *Reader) UnreadRune() error`: 撤销上一次读取的rune
- `func (b *Reader) ReadSlice(delim byte) (line []byte, err error)`: 读取到delim为止的字节
- `func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)`: 读取一行
- `func (b *Reader) ReadBytes(delim byte) ([]byte, error)`: 读取到delim为止的字节
- `func (b *Reader) ReadString(delim byte) (string, error)`: 读取到delim为止的字符串
- `func (b *Reader) Buffered() int`: 返回缓冲区中未读取的字节数
- `func (b *Reader) Reset(r io.Reader)`: 重置Reader, 之后的数据读取都会从新的数据源中来读
- `func (b *Reader) Discard(n int) (discarded int, err error)`: 丢弃n个字节
- `func (b *Reader) Peek(n int) ([]byte, error)`: 预读n个字节, 但不移动读取位置
- `func (b *Reader) WriteTo(w io.Writer) (n int64, err error)`: 将未读取的字节写入w中

### Writer

- `func (b *Writer) Write(p []byte) (nn int, err error)`: 写入p中的字节
- `func (b *Writer) WriteByte(c byte) error`: 写入一个字节
- `func (b *Writer) WriteRune(r rune) (size int, err error)`: 写入一个rune
- `func (b *Writer) WriteString(s string) (int, error)`: 写入一个字符串
- `func (b *Writer) Available() int`: 返回缓冲区中未使用的字节数
- `func (b *Writer) Buffered() int`: 返回缓冲区中已使用的字节数
- `func (b *Writer) Flush() error`: 刷新缓冲区
- `func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)`: 从r中读取数据写入Writer
- `func (b *Writer) Reset(w io.Writer)`: 重置Writer