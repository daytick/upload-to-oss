package main

import (
    "fmt"
    "os"
    "flag"
	"time"
    "path/filepath"
    "github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {
	// 解析入参
	param := parseParam()
    // 创建OSSClient实例
    client, err := oss.New(param.endpoint, param.accessKeyId, param.accessKeySecret)
    handleError(err)
    // 获取存储空间
    bucket, err := client.Bucket(param.bucketName)
    handleError(err)
    // 批量上传
	baseDir := fmt.Sprintf("%d%s%d%s", time.Now().Year(), "/", time.Now().Month(), "/")
    urlList := make([]string, len(param.files))
    for i, localFile := range param.files {
        // 上传文件到OSS时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg
        fileName := baseDir + parseFileName(localFile)
        // 上传单个文件
        // 由本地文件路径加文件名包括后缀组成，例如/users/local/myfile.txt
        err = bucket.PutObjectFromFile(fileName, localFile)
        handleError(err)
        // 存储文件的 url
        // https://personal-image-repo.oss-cn-shanghai.aliyuncs.com/2021%E5%B9%B4/1%E6%9C%88/1NF-1-1.png
        urlList[i] = (fmt.Sprintf("https://%s.%s/%s", param.bucketName, param.endpoint, fileName))
    }
    // 输出已上传文件的 url
    for _, url := range urlList {
        fmt.Println(url)
    }
}

// 解析入参
func parseParam() *Param {
    param := Param{}
    flag.StringVar(&param.endpoint, "e", "", "阿里云服务器所在地域")
    flag.StringVar(&param.accessKeyId, "k", "", "阿里云accessKeyId")
    flag.StringVar(&param.accessKeySecret, "s", "", "阿里云accessKeySecret")
    flag.StringVar(&param.bucketName, "b", "", "用来存放上传文件的桶")
    flag.Parse()
    param.files = flag.Args()
    return &param
}

// 处理错误
func handleError(err error) {
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(-1)    
    }
}

// 从全路径中提取文件名同时加入当前时间戳
func parseFileName(file string) string {
    fileName := filepath.Base(file)
    return fmt.Sprintf("%s-%d%s", fileName[:len(fileName) - len(filepath.Ext(file))], time.Now().Unix(), filepath.Ext(fileName))
}

// 入参
type Param struct {
    endpoint         string   // 阿里云服务器所在地域
    accessKeyId      string   // 阿里云accessKeyId 
    accessKeySecret  string   // 阿里云accessKeySecret 
    bucketName       string   // 用来存放上传文件的桶
    files            []string // 所有要上传文件的完整路径
}
