package image

import (
    // "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

type Image struct {  
    Id bson.ObjectId `bson:"_id"`
    Name    string `bson:"Name"`
    C_time string `bson:"C_tame"`
    M_time string `bson:"M_time"`
    F_path string `bson:"F_path"`
    Ori_url string `bson:"Ori_url"`
    Zoom_url string `bson:"Zoom_url"`
    Rotate_url string `bson:"Rotate_url"`
}


/*
image1 := Image {
        Name : "Vintage.png",
        C_time : "Wed Jun  8 11:41:49 CST 2016",
        M_time : "Wed Jun  8 11:41:49 CST 2016",
        F_path : "/Users/yuri/Project/Go/src/github.com/wyqwyq/image_storage/test_dir/Vintage.png",
        Ori_url : "",
        Zoom_url : "",
        Rotate_url : "",
    }
    image2 := Image {
        Name : "Vintage.png",
        C_time : "Wed Jun  8 11:41:49 CST 2016",
        M_time : "Wed Jun  8 11:41:49 CST 2016",
        F_path : "/Users/yuri/Project/Go/src/github.com/wyqwyq/image_storage/test_dir/Vintage.png",
        Ori_url : "",
        Zoom_url : "",
        Rotate_url : "",
    }
*/
