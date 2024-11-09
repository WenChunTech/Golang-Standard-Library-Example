package GolangStandardLibraryTesting

import (
	"archive/tar"
	"fmt"
	"math/rand"
	"os"
	"testing"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandFile(filename string, fileSize uint64) {
	buf := make([]byte, fileSize)
	for i := range buf {
		buf[i] = letters[rand.Intn(len(letters))]
	}
	os.WriteFile(filename, buf, 0644)
}

func TestArchive(t *testing.T) {
	// log.Println("starting test")
	fmt.Println("starting test")
	filename := "test_tar.txt"
	RandFile(filename, 1024)
	// Create a new tar archive.
	file, _ := os.Open(filename)
	defer file.Close()
	tar := tar.NewWriter(file)
	defer tar.Close()

	err := os.Remove(filename)
	fmt.Println(err)
	// log.Println(err)
	// log.Println("File removed")
}

func TestRemove(t *testing.T) {
	filename := "test_tar.txt"
	err := os.Remove(filename)
	fmt.Println(err)
}
