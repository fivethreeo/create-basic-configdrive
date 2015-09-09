// +build windows

package main

import (
    "fmt"
    ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)


func mkisofs(dir string, isofile string){

    ole.CoInitialize(0)
    
	unknown, _ := oleutil.CreateObject("IMAPI2FS.MsftFileSystemImage")
	
	image, _ := unknown.QueryInterface(ole.IID_IDispatch)
	
	oleutil.PutProperty(image, "VolumeName", "config-2")
    
	oleutil.MustCallMethod(image,
        "ChooseImageDefaultsForMediaType", 12).ToIDispatch()
        
	root := oleutil.MustGetProperty(image, "Root").ToIDispatch()
    
	oleutil.MustCallMethod(root, "AddTree", dir, true)
        
    result := oleutil.MustCallMethod(image, "CreateResultImage").ToIDispatch()
    
    //stream := oleutil.MustGetProperty(result, "ImageStream").ToIDispatch()
    blocksize := int(oleutil.MustGetProperty(result, "BlockSize").Val)
    totalblocks := int(oleutil.MustGetProperty(result, "TotalBlocks").Val)

    fmt.Println("%s %s", blocksize, totalblocks)
    
    image.Release()
    root.Release()
    result.Release()
    //stream.Release()
    	
    ole.CoUninitialize()

    return

}