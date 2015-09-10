// +build windows

package main

import (
    "fmt"
    "log"
    "io"
    "io/ioutil"
    "os"
    "os/exec"
    "path/filepath"
    "archive/zip"
    "net/http"
    "bytes"
)

func Unzip(b []byte, size int, dest string) error {
    r, err := zip.NewReader(bytes.NewReader(b), int64(size))
    if err != nil {
        return err
    }/*
    defer func() {
        if err := r.Close(); err != nil {
            panic(err)
        }
    }()
*/
    os.MkdirAll(dest, 0755)

    // Closure to address file descriptors issue with all the deferred .Close() methods
    extractAndWriteFile := func(f *zip.File) error {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer func() {
            if err := rc.Close(); err != nil {
                panic(err)
            }
        }()

        path := filepath.Join(dest, f.Name)

        if f.FileInfo().IsDir() {
            os.MkdirAll(path, f.Mode())
        } else {
            f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
            if err != nil {
                return err
            }
            defer func() {
                if err := f.Close(); err != nil {
                    panic(err)
                }
            }()

            _, err = io.Copy(f, rc)
            if err != nil {
                return err
            }
        }
        return nil
    }

    for _, f := range r.File {
        err := extractAndWriteFile(f)
        if err != nil {
            return err
        }
    }

    return nil
}

func HTTPDownload(uri string) ([]byte, error) {
    res, err := http.Get(uri)
    if err != nil {
        log.Fatal(err)
    }
    content_zipped, err := ioutil.ReadAll(res.Body)
    defer res.Body.Close()
    fmt.Printf("Size of download: %d\n", len(content_zipped))
    return content_zipped, err
}

func DownloadUnzip(uri string, dst string) {
    fmt.Printf("\nDownloading mkisofs from sourceforge.\n")
    if d, err := HTTPDownload(uri); err == nil {
        fmt.Printf("Downloaded mkisofs.\n")
        if Unzip(d, len(d), dst) == nil {
            fmt.Printf("Unzipped mkisofs.\n")
        }
        
    }
}

var url string = "http://downloads.sourceforge.net/project/mkisofs-md5/mkisofs-md5-v2.01/mkisofs-md5-2.01-Binary.zip"
var querystring string = "?r=http%3A%2F%2Fsourceforge.net%2Fprojects%2Fmkisofs-md5%2Ffiles%2Fmkisofs-md5-v2.01%2Fmkisofs-md5-2.01-Binary.zip%2Fdownload&ts=1441282840&use_mirror=netcologne"
var  uri string = url + querystring

func mkisofs(workdir string, adddir string, isofile string){
    
    DownloadUnzip(uri, workdir)
    
    cmd := exec.Command(workdir + `\Binary\MinGW\Gcc-4.4.5\mkisofs.exe`, "-R", "-V", "config-2", "-o", isofile, adddir)
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    return

}