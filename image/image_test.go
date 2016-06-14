package image

import (
    "testing"
)


func TestImageFinder(t *testing.T) {
    fps, err := GetImagePath("/Users/yuri/Project/Go/src/github.com/wyqwyq/image_storage/test_dir",
            [] string{"png", "jpeg"})
    if err != nil {
        t.Error(err.Error())
    }
    if len(fps) != 1 {
        t.Error("We expect to find one image here, but it's %d(s) images.", len(fps))
    }
    for _, fp := range fps {
        img, _ := GetLocalImageByFilePath(fp)
        t.Log(img.Name, " " ,img.M_time, " ", img.C_time)
    }
}