package main

import (
	"errors"
	"fmt"
)

// Reader - интерфейс для чтения файлов.
type Reader interface {
	Read(path string) error
}

// ReaderFile - реализация интерфейса Reader для чтения файлов с расширением "*.doc".
type ReaderFile struct{}

// Read - реализация метода чтения файлов.
func (r *ReaderFile) Read(path string) error {
	if path == "*.doc" {
		fmt.Println("file read")
		return nil
	}
	return errors.New("error read file")
}

// Scanner - интерфейс для сканирования файлов.
type Scanner interface {
	Scan(path string) error
}

// ScannerFile - реализация интерфейса Scanner для сканирования файлов.
type ScannerFile struct{}

// Scan - реализация метода сканирования файлов.
func (s *ScannerFile) Scan(path string) error {
	if path == "exploit.bat" {
		return errors.New("virus")
	}
	fmt.Println("ok!")
	return nil
}

// OpenScanReader - структура, которая адаптирует интерфейс Scanner к интерфейсу Reader.
type OpenScanReader struct {
	scanner Scanner
}

// Read - реализация метода чтения файлов.
func (o *OpenScanReader) Read(path string) error {
	if path == "troyan.bat" {
		err := o.scanner.Scan(path)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// Создание адаптера OpenScanReader, который адаптирует ScannerFile к интерфейсу Reader.
	adapter := OpenScanReader{scanner: &ScannerFile{}}

	// Чтение файла "troyan.bat" с использованием адаптера.
	err := adapter.Read("troyan.bat")
	if err != nil {
		return
	}
}
