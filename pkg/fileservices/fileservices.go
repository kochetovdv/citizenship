package fileutils

import (
    "os"
)

// CreateFile creates a new file with the specified name and permissions
func CreateFile(filename string, perm os.FileMode) (*os.File, error) {
    return os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
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

// GetFileInfo returns information about a file, including its size and modification time
func GetFileInfo(filename string) (os.FileInfo, error) {
    return os.Stat(filename)
}
