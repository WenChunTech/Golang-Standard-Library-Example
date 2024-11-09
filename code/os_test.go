package GolangStandardLibraryTesting

import (
	"fmt"
	"os"
	"testing"
)

func TestOSEnv(t *testing.T) {
	// 获取环境变量的值
	value := os.Getenv("PATH")
	fmt.Println("PATH:", value)

	// 设置环境变量的值
	os.Setenv("TEST", "Hello, World")
	value = os.Getenv("TEST")
	fmt.Println("TEST:", value)

	// 删除环境变量
	os.Unsetenv("TEST")
	value = os.Getenv("TEST")
	fmt.Println("TEST:", value)

	// 获取所有环境变量
	for _, env := range os.Environ() {
		fmt.Println(env)
	}

}

func TestOsProcess(t *testing.T) {
	// 获取当前进程PID
	fmt.Println("PID:", os.Getpid())

	// 获取当前进程的父进程PID
	fmt.Println("PPID:", os.Getppid())

	// 获取当前进程的用户ID
	fmt.Println("UID:", os.Getuid())

	// 获取当前进程的组ID
	fmt.Println("GID:", os.Getgid())

	// 获取当前进程的所有组ID
	groups, err := os.Getgroups()
	if err != nil {
		fmt.Println("ERROR: os.Getgroups(), err is:", err)
	}
	fmt.Println("GIDs:", groups)

	// 获取系统内存页大小
	fmt.Println("PageSize:", os.Getpagesize())

	// 获取主机名
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("ERROR: os.Hostname(), err is:", err)
	}
	fmt.Println("Hostname:", hostname)

	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("ERROR: os.Getwd(), err is:", err)
	}
	fmt.Println("Working Directory:", wd)

	// 修改当前工作目录
	err = os.Chdir("/tmp")
	if err != nil {
		fmt.Println("ERROR: os.Chdir(), err is:", err)
	}

	// 清空所有环境变量
	os.Clearenv()

	// 替换字符串中的变量
	fmt.Println(os.Expand("Hello, $USER", func(s string) string {
		return "World"
	}))

	// 获取环境变量的值
	value, exist := os.LookupEnv("PATH")
	if exist {
		fmt.Println("PATH:", value)
	} else {
		fmt.Println("PATH env is not set")
	}

	// 启动一个新进程
	// process, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, &os.ProcAttr{})
	// if err != nil {
	// 	fmt.Println("ERROR: os.StartProcess(), err is:", err)
	// }
	// fmt.Println("Process:", process.Pid)

	// 查找一个进程
	process, err := os.FindProcess(1)
	if err != nil {
		fmt.Println("ERROR: os.FindProcess(), err is:", err)
	} else {
		fmt.Println("Process:", process.Pid)
	}

	// 获取用户缓存目录
	dir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println("ERROR: os.UserCacheDir(), err is:", err)
	} else {
		fmt.Println("UserCacheDir", dir)
	}

	// 获取用户配置目录
	dir, err = os.UserConfigDir()
	if err != nil {
		fmt.Println("ERROR: os.UserConfigDir(), err is:", err)
	} else {
		fmt.Println("UserConfigDir", dir)
	}

	// 获取用户主目录
	dir, err = os.UserHomeDir()
	if err != nil {
		fmt.Println("ERROR: os.UserHomeDir(), err is:", err)
	} else {
		fmt.Println("UserHomeDir", dir)
	}
}

func TestOsFile(t *testing.T) {
	// 创建一个文件, 如果文件已经存在, 则会清空文件内容, 可读写
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("ERROR: os.Create(), err is:", err)
	}
	file.WriteString("hello world")
	file.Close()

	// 打开一个文件， 只读
	file, err = os.Open("test.txt")
	if err != nil {
		fmt.Println("ERROR: os.Open(), err is:", err)
	}
	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil {
		fmt.Println("ERROR: file.Read(), err is:", err)
	}
	fmt.Println("Read:", string(buf[:n]))
	file.Close()

	// 重命名一个文件
	err = os.Rename("test.txt", "test1.txt")
	if err != nil {
		fmt.Println("ERROR: os.Rename(), err is:", err)
	}

	// 获取文件信息
	fileInfo, err := os.Stat("test1.txt")
	if err != nil {
		fmt.Println("ERROR: os.Stat(), err is:", err)
	}
	fmt.Println("file info:", fileInfo)

	// 创建一个文件夹
	err = os.Mkdir("test", 0755)
	if err != nil {
		fmt.Println("ERROR: os.Mkdir(), err is:", err)
	}

	// 移动一个文件
	err = os.Rename("test1.txt", "test/test1.txt")
	if err != nil {
		fmt.Println("ERROR: os.Rename(), err is:", err)
	}

	// 创建一个硬链接
	err = os.Link("test/test1.txt", "test/test2.txt")
	if err != nil {
		fmt.Println("ERROR: os.Link(), err is:", err)
	}

	// 创建一个软链接
	err = os.Symlink("test/test1.txt", "test/test3.txt")
	if err != nil {
		fmt.Println("ERROR: os.Symlink(), err is:", err)
	}

	// 读取软链接
	link, err := os.Readlink("test/test3.txt")
	if err != nil {
		fmt.Println("ERROR: os.Readlink(), err is:", err)
	}
	fmt.Println("Readlink:", link)

	// 删除一个文件
	err = os.Remove("test/test1.txt")
	if err != nil {
		fmt.Println("ERROR: os.Remove(), err is:", err)
	}

	// 删除一个不为空的文件夹
	err = os.Remove("test")
	if err != nil {
		fmt.Println("ERROR: os.Remove(), err is:", err)
	}

	// 删除一个文件夹，如果文件夹不为空， 则会递归删除
	err = os.RemoveAll("test")
	if err != nil {
		fmt.Println("ERROR: os.RemoveAll(), err is:", err)
	}

	// 判断文件是否存在
	_, err = os.Stat("test.txt")
	if os.IsNotExist(err) {
		fmt.Println("File not exist")
	} else {
		fmt.Println("File exist")
	}
}
