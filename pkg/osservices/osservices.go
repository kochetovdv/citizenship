package osservices

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
//	"time"
)

// Check if the path already exists, Return true if it exists, false if it does not
func CheckDir(path string) (bool, error) {
	// Create the directory if it does not exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return false, fmt.Errorf("error with creating folder:%s", path)
		}
	}
	return true, nil
}

func SaveToFile(path string, data []byte) error {
	// Create the directory if it does not exist
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error with creating file:%s", path)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error with writing file:%s", path)
	}
	return nil
}

func DownloadFile(path, fileName, url string) error {
	// Create the directory if it does not exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}

	// Check if the file already exists in the path
	filePath := filepath.Join(path, fileName)
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		log.Printf("File %s already exists\n", fileName)
		return nil
	}

	// Download the file from the URL
	log.Printf("Starting download file from %s\n", url)
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Create a new file and save the response body to it
	log.Printf("Saving file %s\n", fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	log.Printf("File %s downloaded to %s\n", fileName, filePath)
	return nil
}



/*
package main

import (
    "crypto/md5"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "path/filepath"
    "strings"
    "time"
)

type FileInfo struct {
    Path         string
    Name         string
    Size         int64
    Extension    string
    App          string
    CreationTime time.Time
    ModTime      time.Time
}

type FolderInfo struct {
    Path            string
    Name            string
    NumFiles        int
    NumSubfolders   int
    CreationTime    time.Time
    ModTime         time.Time
}

type FolderItemInfo struct {
    Path         string
    Name         string
    IsFile       bool
    Size         int64
    Extension    string
    App          string
    CreationTime time.Time
    ModTime      time.Time
}

func ReadFile(path string) ([]byte, error) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    return data, nil
}

func WriteFile(path string, data []byte) error {
    err := ioutil.WriteFile(path, data, 0644)
    if err != nil {
        return err
    }
    return nil
}

func CreateFile(path string) (*os.File, error) {
    file, err := os.Create(path)
    if err != nil {
        return nil, err
    }
    return file, nil
}

func AppendToFile(path string, data []byte) error {
    file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
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

func DeleteFile(path string) error {
    err := os.Remove(path)
    if err != nil {
        return err
    }
    return nil
}

func MoveFile(srcPath, dstPath string) error {
    err := os.Rename(srcPath, dstPath)
    if err != nil {
        return err
    }
    return nil
}

func RenameFile(oldPath, newPath string) error {
    err := os.Rename(oldPath, newPath)
    if err != nil {
        return err
    }
    return nil
}

func GetFileInfo(path string) (FileInfo, error) {
    file, err := os.Stat(path)
    if err != nil {
        return FileInfo{}, err
    }
    name := file.Name()
    size := file.Size()
    modTime := file.ModTime()
    extension := filepath.Ext(name)
    app := GetAppForFile(name)
    return FileInfo{
        Path:         path,
        Name:         name,
        Size:         size,
        Extension:    extension,
        App:          app,
        CreationTime: modTime,
        ModTime:      modTime,
    }, nil
}

func GetFileHash(path string) (string, error) {
    file, err := os.Open(path)
    if err != nil {
        return "", err
    }
    defer file.Close()
    hash := md5.New()
    if _, err := io.Copy(hash, file); err != nil {
        return "", err
    }
    return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func GetFilesFromFolder(path string) ([]os.FileInfo, error) {
    files, err := ioutil.ReadDir(path)
    if err != nil {
        return nil, err
    }
    return files, nil
}

func CreateFolder(path string) error {
    err := os.MkdirAll(path, 0755)
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

func MoveFolder(srcPath, dstPath string) error {
    err := os.Rename(srcPath, dstPath)
    if err != nil {
        return err
    }
    return nil
}

func RenameFolder(oldPath, newPath string) error {
    err := os.Rename(oldPath, newPath)
    if err != nil {
        return err
    }
    return nil
}

func GetFolderInfo(path string) (FolderInfo, error) {
    folder, err := os.Stat(path)
    if err != nil {
        return FolderInfo{}, err
    }
    name := folder.Name()
    modTime := folder.ModTime()

    // Get number of files and subfolders in folder
    numFiles := 0
    numSubfolders := 0
    files, err := GetFilesFromFolder(path)
    if err != nil {
        return FolderInfo{}, err
    }
    for _, file := range files {
        if file.IsDir() {
            numSubfolders++
        } else {
            numFiles++
        }
    }

    return FolderInfo{
        Path:            path,
        Name:            name,
        NumFiles:        numFiles,
        NumSubfolders:   numSubfolders,
        CreationTime:    modTime,
        ModTime:         modTime,
    }, nil
}

func GetFilesAndFoldersFromFolder(path string) ([]FolderItemInfo, error) {
    var items []FolderItemInfo
    files, err := ioutil.ReadDir(path)
    if err != nil {
        return nil, err
    }
    for _, file := range files {
        itemPath := filepath.Join(path, file.Name())

        if file.IsDir() {
            folderInfo, err := GetFolderInfo(itemPath)
            if err != nil {
                return nil, err
            }
            item := FolderItemInfo{
                Path:         folderInfo.Path,
                Name:         folderInfo.Name,
                IsFile:       false,
                Size:         0,
                Extension:    "",
                App:          "",
                CreationTime: folderInfo.CreationTime,
                ModTime:      folderInfo.ModTime,
            }
            items = append(items, item)
        } else {
            fileInfo, err := GetFileInfo(itemPath)
            if err != nil {
                return nil, err
            }
            item := FolderItemInfo{
                Path:         fileInfo.Path,
                Name:         fileInfo.Name,
                IsFile:       true,
                Size:         fileInfo.Size,
                Extension:    fileInfo.Extension,
                App:          fileInfo.App,
                CreationTime: fileInfo.CreationTime,
                ModTime:      fileInfo.ModTime,
            }
            items = append(items, item)
        }
    }
    return items, nil
}

func FileExists(path string) bool {
    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        return false
    }
    return true
}

func GetAppForFile(path string) (string, error) {
    file, err := os.Open(path)
    if err != nil {
        return "", err
    }
    defer file.Close()

    // Read the first 512 bytes of the file to determine the content type
    buffer := make([]byte, 512)
    _, err = file.Read(buffer)
    if err != nil {
        return "", err
    }

    // Determine the MIME type based on the file contents
    contentType := http.DetectContentType(buffer)

    // Map the MIME type to an application name
    switch contentType {
    case "application/pdf":
        return "PDF document", nil
    case "application/msword":
        return "Microsoft Word document", nil
    case "application/vnd.ms-excel":
        return "Microsoft Excel spreadsheet", nil
    case "text/html":
        return "HTML document", nil
    default:
        return "", nil
    }
}


*/


/*

// Я хочу создать пакет, который смогу использовать во многих своих проектах. Этот пакет должен упросить работу с файлами и папками.
// С файлами он должен:
// - проверять, существует ли файл
// - создавать, читать, дозаписывать, удалять, перемещать, переименовывать файл
// - выдавать хэш содержимого файла
// - сравнивать файлы по хэшам содержимого
// - предоставлять информацию о файле (имя, расширение, путь, дата создания, дата обновления, размер файла, наименование приложения, связанного с файлом).
// Для папок:
// - проверять, существует ли папка
// - создавать, удалять, перемещать, переименовывать папку
// - предоставлять информацию о папке (имя, путь, дата создания, дата обновления, размер папки)
// - предоставлять информацию о содержимом папки (количество вложенных папок и файлов)
// - предоставлять список вложенных файлов и папок.
// Создай модуль на Go 

package fileutil

import (
    "crypto/md5"
    "fmt"
    "io"
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

func FileExists(path string) bool {
    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        return false
    }
    return true
}

func CreateFile(path string) error {
    _, err := os.Create(path)
    if err != nil {
        return err
    }
    return nil
}

func ReadFile(path string) (string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

func AppendFile(path string, text string) error {
    f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        return err
    }
    defer f.Close()

    if _, err := f.WriteString(text); err != nil {
        return err
    }
    return nil
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

*/