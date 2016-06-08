package main

import (
    "fmt"
    "gopkg.in/mgo.v2"
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

const URL = "localhost:27017"

var (
    mgoSession *mgo.Session
    dataBase   = "images_meta"
)

func getSession() *mgo.Session {
    if mgoSession == nil {
        var err error
        mgoSession, err = mgo.Dial(URL)
        if err != nil {
            panic(err)
        }
    }
    //最大连接池默认为4096
    return mgoSession.Clone()
}

func witchCollection(collection string, s func(*mgo.Collection) error) error {
    session := getSession()
    defer session.Close()
    c := session.DB(dataBase).C(collection)
    return s(c)
}

func AddImage(img Image) string {
    img.Id = bson.NewObjectId()
    query := func(c *mgo.Collection) error {
        return c.Insert(img)
    }
    err := witchCollection("images", query)
    if err != nil {
        return "false"
    }
    return img.Id.Hex()
}


func GetImageById(id string) *Image {
    objid := bson.ObjectIdHex(id)
    img := new(Image)
    query := func(c *mgo.Collection) error {
        return c.FindId(objid).One(&img)
    }
    witchCollection("images", query)
    return img
}

func ListImages() []Image {
    var images []Image
    query := func(c *mgo.Collection) error {
        return c.Find(nil).All(&images)
    }
    err := witchCollection("images", query)
    if err != nil {
        fmt.Println(err.Error())
    }
    return images
}

func UpdateImage(query bson.M, change bson.M) string {
    exop := func(c *mgo.Collection) error {
        return c.Update(query, change)
    }
    err := witchCollection("images", exop)
    if err != nil {
        return "true"
    }
    return "false"
}

func DeleteImage(query bson.M) string{
    del := func(c *mgo.Collection) error {
        return c.Remove(query)
    }
    err := witchCollection("images", del)
    if err != nil {
        return "true"
    }
    return "false"
}

func main() {
    image1 := Image {
        Name : "Vintage.png",
        C_time : "Wed Jun  8 11:41:49 CST 2016",
        M_time : "Wed Jun  8 11:41:49 CST 2016",
        F_path : "/Users/yuri/Project/Go/src/github.com/wyqwyq/image_storage/test_dir/Vintage.png",
        Ori_url : "",
        Zoom_url : "",
        Rotate_url : "",
    }
    AddImage(image1)
    images := ListImages()
    for i, img := range images {
        UpdateImage(bson.M{ "Name" : img.Name},
        bson.M{
            "Name" : "Vintage" + string(i) + ".png",
            "C_time" : "Wed Jun  8 11:41:49 CST 2016",
            "M_time" : "Wed Jun  8 11:41:49 CST 2016",
            "F_path" : "/Users/yuri/Project/Go/src/github.com/wyqwyq/image_storage/test_dir/Vintage.png",
            "Ori_url" : "",
            "Zoom_url" : "",
            "Rotate_url" : "",
        })
    }

    
}
