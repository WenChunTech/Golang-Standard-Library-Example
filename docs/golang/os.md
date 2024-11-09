# OS包

## 主要包含的功能

### 1. 环境变量

- `os.Getenv(key string) string`: 获取环境变量的值
- `os.Setenv(key, value string) error`: 设置环境变量的值
- `os.Unsetenv(key string) error`: 删除环境变量的值
- `os.Environ() []string`: 获取所有环境变量的值

### 2. 进程

- `os.Exit(code int)`: 退出当前进程
- `os.Getpid() int`: 获取当前进程的PID
- `os.Getppid() int`: 获取当前进程的父进程的PID
- `os.Getuid() int`: 获取当前进程的用户ID
- `os.Getgid() int`: 获取当前进程的组ID
- `os.Geteuid() int`: 获取当前进程的有效用户ID
- `os.Getegid() int`: 获取当前进程的有效组ID
- `os.Getgroups() ([]int, error)`: 获取当前进程的所有组ID
- `os.Getpagesize() int`: 获取系统内存页的大小
- `os.Hostname() (name string, err error)`: 获取主机名
- `os.Getwd() (dir string, err error)`: 获取当前工作目录
- `os.Chdir(dir string) error`: 修改当前工作目录
- `os.Chroot(path string) error`: 修改根目录
- `os.Clearenv()`: 清空所有环境变量
- `os.ExpandEnv(s string) string`: 替换字符串中的环境变量
- `os.LookupEnv(key string) (string, bool)`: 获取环境变量的值
- `os.StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)`: 启动一个新进程
- `os.FindProcess(pid int) (*Process, error)`: 查找一个进程
- `os.UserCacheDir() (string, error)`: 获取用户缓存目录
- `os.UserConfigDir() (string, error)`: 获取用户配置目录
- `os.UserHomeDir() (string, error)`: 获取用户主目录

### 3. 文件

- `os.Open(name string) (*File, error)`: 打开一个文件
- `os.Create(name string) (*File, error)`: 创建一个文件
- `os.OpenFile(name string, flag int, perm FileMode) (*File, error)`: 打开或创建一个文件
- `os.Remove(name string) error`: 删除一个文件
- `os.RemoveAll(path string) error`: 删除一个目录及其所有子目录
- `os.Rename(oldpath, newpath string) error`: 重命名一个文件
- `os.Stat(name string) (FileInfo, error)`: 获取文件信息
- `os.Lstat(name string) (FileInfo, error)`: 获取文件信息，如果是符号链接则返回链接信息
- `os.Chmod(name string, mode FileMode) error`: 修改文件权限
- `os.Chown(name string, uid, gid int) error`: 修改文件所有者
- `os.Chtimes(name string, atime, mtime time.Time) error`: 修改文件访问时间和修改时间
- `os.Link(oldname, newname string) error`: 创建一个硬链接
- `os.Symlink(oldname, newname string) error`: 创建一个符号链接
- `os.Readlink(name string) (string, error)`: 读取一个符号链接
- `os.TempDir() string`: 获取临时目录
- `os.Mkdir(name string, perm FileMode) error`: 创建一个目录
- `os.MkdirAll(path string, perm FileMode) error`: 创建一个目录及其所有父目录
- `os.ReadDir(name string) ([]DirEntry, error)`: 读取一个目录
- `os.IsExist(err error) bool`: 返回一个布尔值说明该错误是否表示一个文件或目录已经存在。ErrExist 和一些系统调用错误会使它返回真
- `os.IsNotExist`: 返回一个布尔值说明该错误是否表示一个文件或目录不存在。ErrNotExist 和一些系统调用错误会使它返回真

### 4. 文件描述符

- `os.Stdin`: 标准输入
- `os.Stdout`: 标准输出
- `os.Stderr`: 标准错误输出
- `os.DevNull`: 空设备
- `os.NewFile(fd uintptr, name string) *File`: 创建一个文件对象
- `os.FileMode`: 文件权限
- `os.O_RDONLY`: 只读
- `os.O_WRONLY`: 只写
- `os.O_RDWR`: 读写
- `os.O_APPEND`: 追加
- `os.O_CREATE`: 创建
- `os.O_EXCL`: 排他
- `os.O_SYNC`: 同步
- `os.O_TRUNC`: 截断
- `os.SeekStart`: 起始位置
- `os.SeekCurrent`: 当前位置
- `os.SeekEnd`: 结尾位置