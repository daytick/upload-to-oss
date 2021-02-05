## 一、简介

上传文件到阿里云 OSS。

## 二、使用

### 查看帮助
```bash
go run upload_to_oss.go -h
```

### 上传文件
```bash
go run upload_to_oss.go -e <end_point> -k <access_key_id> -s <access_key_secret> -b <bucket_name> <file_absolute_path>
```
