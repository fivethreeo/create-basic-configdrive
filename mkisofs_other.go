// +build !windows


func mkisofs(workdir string, adddir string, destdir string, isofile string){
    
    addirfilepath := filepath.Join(workdir, adddir)
    isofilepath := filepath.Join(destdir, isofile)
    
    mkisofs, err := exec.LookPath("mkisofs")
    if err != nil {
        fmt.Println("mkisofs tool is required to create image.")
        os.Exit(1)
    }
    cmd := exec.Command(mkisofs, "-R", "-V", "config-2", "-o", isofilepath, addirfilepath)
    fmt.Printf("Running mkisofs %s %s %s.\n", strings.Join(cmd.Args[1:len(cmd.Args)-2], " "), isofile, adddir)
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    return

}