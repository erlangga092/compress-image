package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"strconv"

	"github.com/alecthomas/kingpin"
	"github.com/nfnt/resize"
)

var (
	app      = kingpin.New("App", "Compress Image")
	argSrc   = app.Flag("src", "Image source").Short('i').Required().String()
	argDest  = app.Flag("dest", "Image destination").Short('o').Required().String()
	argWidth = app.Flag("width", "Image width").Short('w').Required().String()
	argHeigt = app.Flag("height", "Image height").Short('h').Required().String()
)

func compressImage(src, dst string, width, height uint) error {
	// Membuka file sumber
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	// Membaca gambar dari file sumber
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// Mengubah ukuran gambar
	img = resize.Resize(width, height, img, resize.Lanczos3)

	// Membuat file tujuan
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	// Menyimpan gambar kedalam file tujuan
	jpeg.Encode(out, img, nil)
	return nil
}

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	// Mendefinisikan file sumber dan tujuan
	src := *argSrc
	dst := *argDest

	// Mendefinisikan ukuran gambar setelah di kompresi
	widthStr, err := strconv.Atoi(*argWidth)
	if err != nil {
		fmt.Println(err.Error())
	}

	heightStr, err := strconv.Atoi(*argHeigt)
	if err != nil {
		fmt.Println(err.Error())
	}

	width := uint(widthStr)
	height := uint(heightStr)

	// Mengompres gambar
	err = compressImage(src, dst, width, height)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Menampilkan informasi ukuran file sumber dan tujuan
	originalSize := fileSize(src)
	compressedSize := fileSize(dst)

	fmt.Printf("Ukuran file asli: %s\n", originalSize)
	fmt.Printf("Ukuran file setelah di kompresi: %s\n", compressedSize)
}

func fileSize(file string) string {
	// Mendapatkan ukuran file
	info, err := os.Stat(file)
	if err != nil {
		return "unknown"
	}

	// Mengkonversi ukuran file ke string
	size := info.Size()
	if size < 1024 {
		return fmt.Sprintf("%d B", size)
	} else if size < 1024*1024 {
		return fmt.Sprintf("%.2f KB", float64(size)/1024)
	} else {
		return fmt.Sprintf("%.2f MB", float64(size)/1024/1024)
	}
}
