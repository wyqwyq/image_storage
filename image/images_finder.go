package image


import (
    // "fmt"
    "io/ioutil"
    "os"
    "strings"
    // "time"
    // "reflect"
 )

var (
    path_sep = string(os.PathSeparator)
)

func GetImagePath(dir_path string, suffixes []string) ([] string, error) {
    ret := [] string{}
    dir, err := ioutil.ReadDir(dir_path)
    if err != nil {
        return nil, err
    }

    for _, fi := range dir {
        if fi.IsDir() {
            sub_ret, sub_err := GetImagePath(dir_path + path_sep + fi.Name(), suffixes)
            if sub_err != nil {
                return nil, sub_err
            }
            ret = append(ret, sub_ret...)
            continue
        } 
        
        for _, suf := range suffixes {
            if strings.HasSuffix(strings.ToUpper(fi.Name()), strings.ToUpper(suf)) {
                ret = append(ret, dir_path + path_sep + fi.Name())
                break
            }
        }
    }
    return ret, nil
}


func GetLocalImageByFilePath(filepath string) (*Image, error) {
    img := new(Image)
    if fi, err := os.Stat(filepath); err == nil {
        img.Name = fi.Name()
        img.M_time = fi.ModTime().Format("2006-01-02 15:04:05")
        // stat := fi.Sys().(*syscall.Stat_t)
        // The following code is not compiled correctly due to Mac OS X (Darwin)
        // The OS does support Linux stat struct (see https://github.com/restic/restic/issues/85)
        // ctime := time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
        
        // sec := reflect.ValueOf(fi.Sys()).Elem().FieldByName("Ctim").Field(0).Int()
        // nsec := reflect.ValueOf(fi.Sys()).Elem().FieldByName("Ctim").Field(1).Int()
        // ctime := time.Unix(sec, nsec)
        img.C_time = img.M_time
        img.F_path = filepath
    }else {
        panic(err)
    }
    return img, nil
}

