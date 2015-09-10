// +build !windows

func mkisofs(workdir string, adddir string, isofile string){
    
    DownloadUnzip(uri, workdir)
    mkisofs, err := exec.LookPath("mkisofs")
    if err != nil {
        fmt.Println("mkisofs tool is required to create image.")
        os.Exit(1)
    }
    cmd := exec.Command(mkisofs, "-R", "-V", "config-2", "-o", isofile, adddir)
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    return

}