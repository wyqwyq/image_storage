package dbo

import (
    "testing"
    "gopkg.in/mgo.v2/bson"
    . "github.com/wyqwyq/image_storage/image"
)

func setup_db() {
    dataBase = "test"
}

func reset_db() {
    dataBase = "image"
}

func TestDBO(t *testing.T) {
    getSession().DB("test").DropDatabase()
    setup_db()
    defer reset_db()

    image := Image {
        Name : "Vintage.png",
        C_time : "Wed Jun  8 11:41:49 CST 2016",
        M_time : "Wed Jun  8 11:41:49 CST 2016",
        F_path : "/Users/yuri/Project/Go/src/github.com/wyqwyq/image_storage/test_dir/Vintage.png",
        Ori_url : "",
        Zoom_url : "",
        Rotate_url : "",
    }
    ret := AddImage(image)
    if ret == "false" {
        t.Error("Add Image failure")
    }
    images := ListImages()
    if len(images) != 1 {
        t.Error("ListImages failure")
    }

    for _, img := range images {
        UpdateImage(bson.M{ "Name" : img.Name},
        bson.M{
            "Name" : "Vintage0" + ".png",
            "C_time" : "Wed Jun  8 11:41:49 CST 2016",
            "M_time" : "Wed Jun  8 11:41:49 CST 2016",
            "F_path" : "/Users/yuri/Project/Go/src/github.com/wyqwyq/image_storage/test_dir/Vintage.png",
            "Ori_url" : "",
            "Zoom_url" : "",
            "Rotate_url" : "",
        })
    }

    img := GetImageByFilePath("/Users/yuri/Project/Go/src/github.com/wyqwyq/image_storage/test_dir/Vintage.png")
    if img.Name != "Vintage0.png" {
        t.Error("Update failure")
    }
    DeleteImage(bson.M{})
    if len(ListImages()) != 0 {
        t.Error("Delete failure")
    }
}