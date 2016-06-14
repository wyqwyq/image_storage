package dbo


import (
    "fmt"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    . "github.com/wyqwyq/image_storage/image"
)

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
    err := witchCollection("image", query)
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
    witchCollection("image", query)
    return img
}

func GetImageByFilePath(filepath string) *Image {
    img := new(Image)
    // Why not get a nil img after c.Find(bson.M{"F_path" : filepath}).One(&img)
    // if there should no matched img ?
    query := func(c *mgo.Collection) error {
        ret := c.Find(bson.M{"F_path" : filepath})
        if n, _ := ret.Count(); n == 1 {
            ret.One(&img)
        }else {
            img = nil
        }
        return nil
    }
    witchCollection("image", query)
    return img
}

func GetImageByName(name string) *Image {
    img := new(Image)
    query := func(c *mgo.Collection) error {
        ret := c.Find(bson.M{"Name" : name})
        if n, _ := ret.Count(); n == 1 {
            ret.One(&img)
        }else {
            img = nil
        }
        return nil
    }
    witchCollection("image", query)
    return img
}

func ListImages() []Image {
    var images []Image
    query := func(c *mgo.Collection) error {
        return c.Find(nil).All(&images)
    }
    err := witchCollection("image", query)
    if err != nil {
        fmt.Println(err.Error())
    }
    return images
}

func UpdateImage(query bson.M, change bson.M) string {
    exop := func(c *mgo.Collection) error {
        return c.Update(query, change)
    }
    err := witchCollection("image", exop)
    if err != nil {
        return "true"
    }
    return "false"
}

func DeleteImage(query bson.M) string {
    del := func(c *mgo.Collection) error {
        return c.Remove(query)
    }
    err := witchCollection("image", del)
    if err != nil {
        return "true"
    }
    return "false"
}
