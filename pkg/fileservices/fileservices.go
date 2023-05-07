package fileutils

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type FileInfo struct {
	Name         string
	Path         string
	CreatedTime  time.Time
	ModifiedTime time.Time
	Size         int64
	Type         string
	Application  string
}

type FolderInfo struct {
	Name         string
	Path         string
	CreatedTime  time.Time
	ModifiedTime time.Time
	Size         int64
	FileCount    int
	FolderCount  int
}

// CreateFile creates a new file
func CreateFile(path string) error {
    _, err := os.Create(path)
    if err != nil {
        return err
    }
    return nil
}

// ReadFile reads the contents of a file and returns its contents as a string
func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	fileContent := make([]byte, fileInfo.Size())
	_, err = file.Read(fileContent)
	if err != nil {
		return "", err
	}

	return string(fileContent), nil
}

// WriteFile writes data to a file, overwriting any existing contents
func WriteFile(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// AppendFile appends data to a file, preserving existing contents
func AppendFile(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// ReplaceFile replaces the contents of an existing file with new data
func ReplaceFile(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// GetFileInfo returns information about a file
func GetFileInfo(path string) (FileInfo, error) {
	fileInfo := FileInfo{}

	file, err := os.Stat(path)
	if err != nil {
		return fileInfo, err
	}

	fileInfo.Name = file.Name()
	fileInfo.Path = path
	fileInfo.CreatedTime = file.ModTime()
	fileInfo.ModifiedTime = file.ModTime()
	fileInfo.Size = file.Size()

	// Определяем тип файла
	fileType, err := GetFileType(path)
	if err == nil {
		fileInfo.Type = fileType
	}

	// Определяем приложение, которое открывает такие файлы
	application, err := GetApplicationForFileType(fileType)
	if err == nil {
		fileInfo.Application = application
	}

	return fileInfo, nil
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CopyFile(srcPath string, destPath string) (FileInfo, error) {
	// Открываем исходный файл для чтения
	src, err := os.Open(srcPath)
	if err != nil {
		return FileInfo{}, err
	}
	defer src.Close()

	// Создаем целевой файл для записи
	dest, err := os.Create(destPath)
	if err != nil {
		return FileInfo{}, err
	}
	defer dest.Close()

	// Копируем содержимое исходного файла в целевой файл
	if _, err := io.Copy(dest, src); err != nil {
		return FileInfo{}, err
	}

	// Получаем информацию о скопированном файле
	fileInfo, err := GetFileInfo(destPath)
	if err != nil {
		return FileInfo{}, err
	}

	return fileInfo, nil
}

// Переименование файла
func RenameFile(path string, newName string) (FileInfo, error) {
	// Получаем информацию о файле
	fileInfo, err := GetFileInfo(path)
	if err != nil {
		return FileInfo{}, err
	}

	// Вычисляем новый путь к файлу с новым именем
	newPath := filepath.Join(filepath.Dir(path), newName)

	// Переименовываем файл
	if err := os.Rename(path, newPath); err != nil {
		return FileInfo{}, err
	}

	// Обновляем информацию о файле с новым именем
	fileInfo.Name = newName
	fileInfo.Path = newPath
	fileInfo.ModifiedTime = time.Now()

	return fileInfo, nil
}

// Перемещение файла
func MoveFile(srcPath string, destPath string) (FileInfo, error) {
	// Получаем информацию о файле
	fileInfo, err := GetFileInfo(srcPath)
	if err != nil {
		return FileInfo{}, err
	}

	// Перемещаем файл
	if err := os.Rename(srcPath, destPath); err != nil {
		return FileInfo{}, err
	}

	// Обновляем информацию о файле с новым путем
	fileInfo.Path = destPath
	fileInfo.ModifiedTime = time.Now()

	return fileInfo, nil
}

// Удаление файла
func DeleteFile(path string) error {
	// Удаляем файл
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

func GetFileHash(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func CompareFilesByHash(path1 string, path2 string) (bool, error) {
	hash1, err := GetFileHash(path1)
	if err != nil {
		return false, err
	}

	hash2, err := GetFileHash(path2)
	if err != nil {
		return false, err
	}

	if hash1 == hash2 {
		return true, nil
	} else {
		return false, nil
	}
}

func GetFileType(path string) (string, error) {
	// Открываем файл в бинарном режиме для чтения первых нескольких байт
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Читаем первые несколько байт
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Определяем MIME-тип файла на основе первых нескольких байт
	fileType := http.DetectContentType(buffer)

	return fileType, nil
}

func GetApplicationForFileType(fileType string) (string, error) {
	// Определяем приложение, которое открывает файлы данного типа, на основе MIME-типа
	switch fileType {
	case "application/pdf":
		return "Adobe Acrobat Reader", nil
	case "image/jpeg", "image/png":
		return "Windows Photo Viewer", nil
	case "text/plain":
		return "Notepad", nil
	default:
		return "", fmt.Errorf("Application not found for file type: %s", fileType)
	}
}

func FindFileByName(path string, fileName string) (string, error) {
	var foundPath string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.Name() == fileName {
			foundPath = path
			return nil
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	if foundPath == "" {
		return "", fmt.Errorf("File not found")
	}

	return foundPath, nil
}

func FindFileByHash(path string, contentHash string) (string, error) {
	var foundPath string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			hash, err := GetFileHash(path)
			if err != nil {
				return err
			}
			if hash == contentHash {
				foundPath = path
				return nil
			}
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	if foundPath == "" {
		return "", fmt.Errorf("File not found")
	}

	return foundPath, nil
}

func FolderExists(path string) bool {
	fileInfo, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	}
	return fileInfo.IsDir()
}

func CreateFolder(path string) error {
	err := os.Mkdir(path, 0755)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFolder(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

func MoveFolder(oldPath string, newPath string) error {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return err
	}
	return nil
}

func RenameFolder(path string, newName string) error {
	dir := filepath.Dir(path)
	newPath := filepath.Join(dir, newName)
	err := os.Rename(path, newPath)
	if err != nil {
		return err
	}
	return nil
}

func GetFolderInfo(path string) (FolderInfo, error) {
	folderInfo := FolderInfo{}

	folder, err := os.Stat(path)
	if err != nil {
		return folderInfo, err
	}

	folderInfo.Name = folder.Name()
	folderInfo.Path = path
	folderInfo.CreatedTime = folder.ModTime()
	folderInfo.ModifiedTime = folder.ModTime()
	folderInfo.Size, folderInfo.FileCount, folderInfo.FolderCount = calculateFolderSize(path)

	return folderInfo, nil
}

func calculateFolderSize(path string) (int64, int, int) {
	var size int64
	var fileCount int
	var folderCount int

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			folderCount++
		} else {
			size += info.Size()
			fileCount++
		}

		return nil
	})

	if err != nil {
		return 0, 0, 0
	}

	return size, fileCount, folderCount
}

func GetFolderContents(path string) ([]string, []string, error) {
	var files []string
	var folders []string

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			folders = append(folders, path)
		} else {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return files, folders, nil
}
