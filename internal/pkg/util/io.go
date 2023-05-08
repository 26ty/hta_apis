package util

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"eirc.app/internal/v1/structure/file"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	ss "eirc.app/internal/pkg/amazon/s3"
	"eirc.app/internal/pkg/log"
)

// fudn 範例: 檔案(專案)路徑 + 檔名, ex: `dindin/images/` + `uber.png`
func FileToS3(input *file.Created, s3BucketName string) (url string) {
	log.Info("into FileToS3")
	log.Info("do NewAmazonStorage")
	sg := ss.NewAmazonStorage(s3BucketName)
	log.Info("do transfer to *strings.Reader")
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(input.Base64))
	o := &s3.PutObjectInput{
		Key:  aws.String(input.FilePath),
		Body: reader,
	}
	log.Info("do Upload")
	info, err := sg.Upload(o)
	if err != nil {
		log.Error(err)
		return ""
	}
	marshal, _ := json.Marshal(info)
	log.Info(string(marshal))
	url = info.Location
	return url
}

func Base64ToByteArray(str string) []byte {
	b, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("error:", err)
	}
	return b
}

func Base64ToFile(input string, filename string) {
	b := Base64ToByteArray(input)
	ioutil.WriteFile(filename, b, 0644)
	return
}
