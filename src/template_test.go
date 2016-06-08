package main

import (
    "fmt"
    "net/http"
    "os"
    "path"
    "html/template"
)

type Image struct {  
    Name    string
    C_time string
    M_time string
    F_path string
    Ori_url string
    Zoom_url string
    Rotate_url string
}  

// gloal variables
var Base_url, _ = os.Getwd()
var Template_dir = path.Join(Base_url, "templates")
var Template_name = "index.html"

func Handler(w http.ResponseWriter, r *http.Request) {  
    images := [] Image{}
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
    images = append(images, image1, image2)
    t, err := template.ParseFiles(path.Join(Template_dir, Template_name))
    checkError(err)
    err = t.Execute(w, images)
    checkError(err)
}

func checkError(err error) {  
    if err != nil {  
        fmt.Println("Fatal error ", err.Error())  
        os.Exit(1)  
    } 
}

func main() {
    http.HandleFunc("/", Handler)  
    http.ListenAndServe(":8888", nil) 
}
