// +build !windows

func mkisofs(workdir string, adddir string, isofile string){
    
    DownloadUnzip(uri, workdir)
    
    cmd := exec.Command(mkisofs.exe, "-R", "-V", "config-2", "-o", isofile, adddir)
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    return

}