package main

import (
    "fmt"
    "gopkg.in/mgo.v2/bson"
    "net/http"
    "os"
    "path"
    "html/template"
    
    . "github.com/wyqwyq/image_storage/image"
    "github.com/wyqwyq/image_storage/dbo"
    cloud "github.com/wyqwyq/image_storage/cloud_storage"
)

// gloal variables
var Cloud_base_url = "http:////o8cpu8afd.bkt.clouddn.com" // why '/'' ?
var Base_dir =  "/Users/yuri/Project/Go/src/github.com/wyqwyq/image_storage"
var Template_dir = path.Join(Base_dir, "templates")
var Template_name = "index.html"


func Handler_Index(w http.ResponseWriter, r *http.Request) {
    images_fps, _ := GetImagePath(path.Join(Base_dir, "test_dir"), []string{"png", "jpg"})
    // fmt.Println("images count : ", len(images_fps))
    for _, fp := range images_fps {
        img_ptr := dbo.GetImageByFilePath(fp)
        if img_ptr == nil {
            i_ptr, _ := GetLocalImageByFilePath(fp)
            fmt.Println("add image (", i_ptr.Name, ")")
            // put file
            cloud.UploadFile(fp, i_ptr.Name)
            // then update db
            i_ptr.Ori_url = path.Join(Cloud_base_url, i_ptr.Name)
            dbo.AddImage(*i_ptr)
        }
    }
    images := dbo.ListImages()
    t, err := template.ParseFiles(path.Join(Template_dir, Template_name))
    checkError(err)
    err = t.Execute(w, images)
    checkError(err)
}

func Handler_Fop(w http.ResponseWriter, r *http.Request) {
    fmt.Println("fop...")
    if r.Method == "POST"{
        name, fop := r.FormValue("Name"), r.FormValue("Fop")
        i_ptr := dbo.GetImageByName(name)
        if i_ptr == nil {
            goto failure
        }
        fmt.Println(fop + " " + name)
        var new_url string
        if fop == "Zoom" {
            dbo.UpdateImage(bson.M{"F_path" : i_ptr.F_path}, bson.M{"$set": bson.M{"Zoom_url" : path.Join(Cloud_base_url, name + "_zoom")}})
            new_url = cloud.GetZoomImageURL(path.Join("o8cpu8afd.bkt.clouddn.com", name), name + "_zoom")
        }else if fop == "Rotate"{
            dbo.UpdateImage(bson.M{"F_path" : i_ptr.F_path}, bson.M{"$set": bson.M{"Rotate_url" : path.Join(Cloud_base_url, name + "_rotate")}})
            new_url = cloud.GetRotateImageURL(path.Join("o8cpu8afd.bkt.clouddn.com", name), name + "_rotate")
        }else {
            goto failure
        }
        fmt.Println(new_url)
        _, err := cloud.HttpGet(new_url)
        if err != nil {
            fmt.Println("http.get failed:", err.Error())
                goto failure
        }else {
            w.Write([]byte("success"))
            return
        }
    }else {
        http.NotFound(w, r)
        return
    }

failure:
    w.Write([]byte("failed"))
}

func checkError(err error) {  
    if err != nil {  
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)  
    } 
}

func main() {
    http.HandleFunc("/list", Handler_Index)  
    http.HandleFunc("/fop", Handler_Fop)  
    http.ListenAndServe(":8888", nil) 
}