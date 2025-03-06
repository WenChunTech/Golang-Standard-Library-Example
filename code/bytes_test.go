package GolangStandardLibraryTesting

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBytes(t *testing.T) {
	buf := []byte("hello, world")

	// 复制一个Buffer
	buf2 := bytes.Clone(buf)
	buf2[0] = 'H'
	fmt.Println(buf)                          // hello, world
	fmt.Println(buf2)                         // Hello, world
	fmt.Printf("buf=%p buf2=%p\n", buf, buf2) // buf=0xc00000a350 buf2=0xc00000a360

	// 比较两个Buffer
	fmt.Println(bytes.Compare(buf, buf2)) // 1

	// 判断buf中是否包含world
	fmt.Println(bytes.Contains(buf, []byte("world"))) // true

	// 判断buf中是否包含by中任一字符
	fmt.Println(bytes.ContainsAny([]byte("by"), string(buf)))  // false
	fmt.Println(bytes.ContainsAny([]byte("bye"), string(buf))) // true

	// 判断b是否包含满足f函数的Rune
	fmt.Println(bytes.ContainsFunc([]byte("aA"), func(r rune) bool {
		return r == 97
	})) // true

	// 判断b是否包含Rune
	fmt.Println(bytes.ContainsRune([]byte("Hello"), 'e')) // true
	fmt.Println(bytes.ContainsRune([]byte("Hello"), 'z')) // false

}

func TestReadBytes(t *testing.T) {
	s := []byte("Hello, World!")

	// 查找子序列在s中第一次出现的位置
	fmt.Println(bytes.Index(s, []byte("World"))) // 7

	// 查找chars中任意一个字符在s中第一次出现的位置
	fmt.Println(bytes.IndexAny(s, "abc")) // -1

	// 查找字节c在s中第一次出现的位置
	fmt.Println(bytes.IndexByte(s, 'W')) // 7

	// 查找满足f函数的Rune在s中第一次出现的位置
	fmt.Println(bytes.IndexFunc(s, func(r rune) bool {
		return r == 'W'
	})) // 7

	// 查找Rune在s中第一次出现的位置
	fmt.Println(bytes.IndexRune(s, 'W')) // 7

	// 计算sep在s中出现的次数
	fmt.Println(bytes.Count(s, []byte("l"))) // 3

	// 按分隔符切割字节切片
	fmt.Println(bytes.Split([]byte("a,b,c"), []byte(","))) // [[97] [98] [99]]

	// 按指定分隔符切割为两部分
	parts, after, found := bytes.Cut([]byte("Hello, World!"), []byte(", "))
	fmt.Println(string(parts), string(after), found) // Hello World! true

	// 检查前缀和后缀
	fmt.Println(bytes.HasPrefix(s, []byte("Hello"))) // true
	fmt.Println(bytes.HasSuffix(s, []byte("!")))     // true
}

func TestWriteBytes(t *testing.T) {
	// 转换为大写
	fmt.Println(string(bytes.ToUpper([]byte("hello")))) // HELLO

	// 转换为小写
	fmt.Println(string(bytes.ToLower([]byte("HELLO")))) // hello

	// 连接byte切片
	slices := [][]byte{[]byte("hello"), []byte("world")}
	fmt.Println(string(bytes.Join(slices, []byte(", ")))) // hello, world

	// 替换操作
	s := []byte("hello hello hello")
	// 替换前两个hello为hi
	fmt.Println(string(bytes.Replace(s, []byte("hello"), []byte("hi"), 2))) // hi hi hello
	// 替换所有hello为hi
	fmt.Println(string(bytes.ReplaceAll(s, []byte("hello"), []byte("hi")))) // hi hi hi

	// 修剪操作
	// 去除两端指定字符
	fmt.Println(string(bytes.Trim([]byte("  hello  "), " "))) // hello
	// 去除两端空白字符
	fmt.Println(string(bytes.TrimSpace([]byte("  hello  ")))) // hello
	// 去除前缀
	fmt.Println(string(bytes.TrimPrefix([]byte("hello world"), []byte("hello ")))) // world
	// 去除后缀
	fmt.Println(string(bytes.TrimSuffix([]byte("hello world"), []byte(" world")))) // hello

	// 重复字节序列
	fmt.Println(string(bytes.Repeat([]byte("hi"), 3))) // hihihi

	// 对每个rune应用映射函数
	mapFunc := func(r rune) rune {
		if r == 'e' {
			return 'E'
		}
		return r
	}
	fmt.Println(string(bytes.Map(mapFunc, []byte("hello")))) // hEllo
}
