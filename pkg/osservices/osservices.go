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