package GolangStandardLibraryTesting

import (
	"syscall"
	"testing"
)

func TestSocket(t *testing.T) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		t.Fatalf("Socket failed: %v", err)
	}
	defer syscall.Close(fd)

	t.Logf("Socket created with fd: %d", fd)
}

func TestSocketPair(t *testing.T) {
	fds, err := syscall.Socketpair(syscall.AF_UNIX, syscall.SOCK_DGRAM, 0)
	if err != nil {
		t.Fatalf("Socketpair failed: %v", err)
	}
	defer syscall.Close(fds[0])
	defer syscall.Close(fds[1])

	t.Logf("Socketpair created with fds: %v", fds)
}

func TestClose(t *testing.T) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		t.Fatalf("Socket failed: %v", err)
	}

	err = syscall.Close(fd)
	if err != nil {
		t.Fatalf("Close failed: %v", err)
	}
}

func TestStat(t *testing.T) {
	var stat syscall.Stat_t
	err := syscall.Stat(".", &stat)
	if err != nil {
		t.Fatalf("Stat failed: %v", err)
	}

	t.Logf("Stat successful: %v", stat)
}

func TestStatfs(t *testing.T) {
	var statfs syscall.Statfs_t
	err := syscall.Statfs(".", &statfs)
	if err != nil {
		t.Fatalf("Statfs failed: %v", err)
	}

	t.Logf("Statfs successful: %v", statfs)
}

func TestTimespecToNsec(t *testing.T) {
	var ts syscall.Timespec
	ts.Sec = 1
	ts.Nsec = 100
	nsec := syscall.TimespecToNsec(ts)
	t.Logf("TimespecToNsec: %d", nsec)
}

func TestTimevalToNsec(t *testing.T) {
	var tv syscall.Timeval
	tv.Sec = 1
	tv.Usec = 100
	nsec := syscall.TimevalToNsec(tv)
	t.Logf("TimevalToNsec: %d", nsec)
}
