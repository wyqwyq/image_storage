package cloud_storage

import (
    "qiniupkg.com/api.v7/kodo"
    "qiniupkg.com/api.v7/conf"
    "qiniupkg.com/api.v7/kodocli"
    "fmt"
    "encoding/base64"
    "crypto/hmac"
    "crypto/sha1"
    "net/url"
    "net/http"
    "io"
)

var (
    //设置上传到的空间
    bucket = "test"
)

//构造返回值字段
type PutRet struct {
    Hash    string `json:"hash"`
    Key     string `json:"key"`
}

func init() {
    //初始化AK，SK
    conf.ACCESS_KEY = "iVel2tyTiv0Zv4LEqDsOe3y-9LtDHvqeduRCJP79"
    conf.SECRET_KEY = "g0oZ6OLV9Y420Ris80VdmeJBIHCWm33yZRTWbj7u"
}

func UploadFile(filepath string, key string) {
    //创建一个Client
    c := kodo.New(0, nil)

    //设置上传的策略
    policy := &kodo.PutPolicy{
        Scope:   bucket + ":" + key,
        //设置Token过期时间
        Expires: 3600,
        InsertOnly: 1,
    }

    //生成一个上传token
    token := c.MakeUptoken(policy)

    //构建一个uploader
    zone := 0
    uploader := kodocli.NewUploader(zone, nil)

    var ret PutRet
    res := uploader.PutFile(nil, &ret, token, key, filepath, nil)
    if res != nil {
        fmt.Println("io.Put failed:", res)
        return
    }
}

func GetZoomImageURL(URL string, key string) string{
    url := URL + "?imageMogr2/thumbnail/!50p"
    new_url := makeSaveasUrl(url, conf.ACCESS_KEY, []byte(conf.SECRET_KEY), bucket, key)
    new_url = "http://" + new_url
    return new_url
}

func GetRotateImageURL(URL string, key string) string{
    url := URL + "?imageMogr2/rotate/180"
    new_url := makeSaveasUrl(url, conf.ACCESS_KEY, []byte(conf.SECRET_KEY), bucket, key)
    new_url = "http://" + new_url
    return new_url
}


func makeSaveasUrl(URL, accessKey string, secretKey []byte, saveBucket, saveKey string) string {
      encodedEntryURI := base64.URLEncoding.EncodeToString([]byte(saveBucket +":"+saveKey))
      URL += "|saveas/" + encodedEntryURI
      h := hmac.New(sha1.New, secretKey)
      // 签名内容不包括Scheme，仅含如下部分：
      // <Domain>/<Path>?<Query>
      u, _ := url.Parse(URL)
      io.WriteString(h, u.Host + u.RequestURI())
      d := h.Sum(nil)
      sign := accessKey + ":" + base64.URLEncoding.EncodeToString(d)
      return URL + "/sign/" + sign
}

func HttpGet(URL string) (*http.Response, error){
    resp, err := http.Get(URL)
    return resp, err
}

func ClearStorage() {
    // func (p Bucket) List( ctx Context, prefix, delimiter, marker string, limit int) (entries []ListItem, commonPrefixes []string, markerOut string, err error)
    c := kodo.New(0, nil) // 用默认配置创建 Client
    bkt := c.Bucket(bucket)
    entries, _, _, _ := bkt.List(nil, "", "", "", (1 << 31))
    fmt.Println("files count : ", len(entries))
    keys := []string{}
    for _, k := range entries {
        fmt.Println("key: ", k.Key)
        keys = append(keys, k.Key)
    }
    bkt.BatchDelete(nil, keys...)
}


