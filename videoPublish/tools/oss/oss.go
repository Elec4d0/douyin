package oss

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

var ossClient *minio.Client

func InitOss() {
	var err error
	ctx := context.Background()
	endpoint := "127.0.0.1:9000"
	accessKeyID := "minioadmin"
	secretAccessKey := "minioadmin"
	//useSSL := true

	// Initialize minio client object.
	ossClient, err = minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		//Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called video
	bucketName := "video"
	location := "localServer"

	err = ossClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := ossClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	// Make a new bucket called jepg
	jepgbucketName := "jepg"
	err = ossClient.MakeBucket(ctx, jepgbucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := ossClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", jepgbucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", jepgbucketName)
	}
}

func UploadVideo(videoByte []byte, objectName string) {
	// Upload the zip file

	//对象的文件类型，视频文件类型过多，未写转码微服务，暂不指派
	//contentType := "application/zip"

	//链接OSS的用具，桶名
	ctx := context.Background()
	bucketName := "video"
	log.Println("oss接收到的videoByte数组长度：", len(videoByte))
	//接受到的byte数组包装为一组流，上传至OSS
	videoStream := bytes.NewReader(videoByte)

	//上传视频流到OSS
	info, err := ossClient.PutObject(ctx, bucketName, objectName, videoStream, int64(len(videoByte)), minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("成功上传视频至OSS %s of size %d\n", objectName, info.Size)
}

func UploadJpeg(jpegByte []byte, objectName string) {
	// Upload the zip file

	//链接OSS的用具，桶名
	ctx := context.Background()
	bucketName := "jepg"
	log.Println("oss接收到的videoByte数组长度：", len(jpegByte))
	//接受到的byte数组包装为一组流，上传至OSS
	videoStream := bytes.NewReader(jpegByte)

	//上传视频流到OSS
	info, err := ossClient.PutObject(ctx, bucketName, objectName, videoStream, int64(len(jpegByte)), minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("成功上传视频至OSS %s of size %d\n", objectName, info.Size)
}
