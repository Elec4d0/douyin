package ffmepg

import (
	"bytes"
	"fmt"
	"github.com/u2takey/ffmpeg-go"
	"image"
	"image/jpeg"
	"log"
	"os"
)

func GetVideoFirstFrameBytes(videoUrl string) ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	err := ffmpeg_go.Input(videoUrl).Filter("select", ffmpeg_go.Args{fmt.Sprintf("gte(n, %d)", 1)}).
		Output("pipe:", ffmpeg_go.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjepg"}).
		WithOutput(buf, os.Stdout).
		Run()

	if err != nil {
		log.Println("ffmepg解码失败")
		log.Println(err)
		return nil, err
	}

	img, _, err := image.Decode(buf)
	if err != nil {
		log.Println("img解码ffpemg视频帧失败")
		log.Println(err)
		return nil, err
	}

	jpegBuf := new(bytes.Buffer)
	err = jpeg.Encode(buf, img, nil)
	if err != nil {
		log.Println("jepg解码image对象失败")
		log.Println(err)
		return nil, err
	}

	return jpegBuf.Bytes(), nil
	/*
		img, err := imaging.Decode(buf)
		if err != nil {
			log.Println("视频帧解码为img失败")
			log.Println(err)
			return "", err
		}

		//jepg保存
		jepgSavePath := "./tmp/"
		err = imaging.Save(img, jepgSavePath+jepgName+".jpg")

		//对于main包的相对路径
		jepgPath := "./tools/ffmepg/tmp/" + jepgSavePath + jepgName + ".jpg"
	*/
}
