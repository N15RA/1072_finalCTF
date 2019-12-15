package main

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
)

func main() {


    flag := "NISRA{df8y4TOfUlqoSIh1GvU0rz5gABneLnjMsh9lYLqIXEUDc1vGmdwECsK1RV3s8c3ri2A4DQDqjnSlfdSrgcTGusvbdnscpltvbRqExRTG2BQ5CwlGy4aRmuQqH8qUhHeiosLJ9sgZTUVje6NC0Nhb7YpCUFn2cr8qjJZGZ3vGLgPqcsg14n21bhuUMF0MWI2tgPZ5HkCiphWpy2gsf6hfdzZN4zpQ5f9bTW7q7xwkCiz0aGgeCojEN6hGvlmqykqh}"
    files := []string{}

    for i := 0; i < len(flag)-1; i++ {
        if(flag[i] == flag[i+1]) {
            fmt.Println("nope!")
            return
        }

    }

    for _,v := range Reverse(flag) {

        name := string(v) + ".zip"
        ZipFiles(name, files)
        for _,fv := range files {
            os.Remove(fv)
        }
        files = []string{name} 

    }
    fmt.Println("done")

}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// ZipFiles compresses one or many files into a single zip archive file.
// Param 1: filename is the output zip file's name.
// Param 2: files is a list of files to add to the zip.
func ZipFiles(filename string, files []string) error {

    newZipFile, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer newZipFile.Close()

    zipWriter := zip.NewWriter(newZipFile)
    defer zipWriter.Close()

    // Add files to zip
    for _, file := range files {
        if err = AddFileToZip(zipWriter, file); err != nil {
            return err
        }
    }
    return nil
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {

    fileToZip, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer fileToZip.Close()

    // Get the file information
    info, err := fileToZip.Stat()
    if err != nil {
        return err
    }

    header, err := zip.FileInfoHeader(info)
    if err != nil {
        return err
    }

    // Using FileInfoHeader() above only uses the basename of the file. If we want
    // to preserve the folder structure we can overwrite this with the full path.
    header.Name = filename

    // Change to deflate to gain better compression
    // see http://golang.org/pkg/archive/zip/#pkg-constants
    header.Method = zip.Deflate

    writer, err := zipWriter.CreateHeader(header)
    if err != nil {
        return err
    }
    _, err = io.Copy(writer, fileToZip)
    return err
}