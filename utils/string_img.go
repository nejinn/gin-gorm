package utils

import (
	"errors"
	"fmt"
	"github.com/golang/freetype"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"note/global"
	"os"
)

// 生成默认头像
func CreateDefaultAvatar(username string) (path string, err error) {
	//图片的宽度
	dx := 200
	//图片的高度
	dy := 200
	fileDir := fmt.Sprintf("%s/%s/%s", global.NLY_CONFIG.Static.Path, "avatar", username)
	err = CreateDir(fileDir)
	if err != nil {
		return "", errors.New("创建默认用户头像失败")
	}
	filePath := fmt.Sprintf("%s/%s", fileDir, "default_avatar.png")
	imgfile, _ := os.Create(filePath)

	defer imgfile.Close()
	img := image.NewNRGBA(image.Rect(0, 0, dx, dy))
	//设置每个点的 RGBA (Red,Green,Blue,Alpha(设置透明度))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			//设置一块 白色(255,255,255)不透明的背景
			img.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}
	//读取字体数据
	fontBytes, err := ioutil.ReadFile("fonts.ttf")
	if err != nil {
		return "", errors.New("创建默认用户头像读取字体失败")
	}
	//载入字体数据
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return "", errors.New("创建默认用户头像加载字体失败")
	}
	f := freetype.NewContext()
	//设置分辨率
	f.SetDPI(72)
	//设置字体
	f.SetFont(font)
	//设置尺寸
	f.SetFontSize(150)
	f.SetClip(img.Bounds())
	//设置输出的图片
	f.SetDst(img)
	//设置字体颜色
	f.SetSrc(image.NewUniform(color.RGBA{0, 0, 0, 255}))

	//截取username第一个字符串
	avatarStr := "姚"

	//设置字体的位置
	pt := freetype.Pt(25, 150+int(f.PointToFixed(26))>>8)
	_, err = f.DrawString(avatarStr, pt)
	if err != nil {
		return "", errors.New("创建默认用户头像写入文字失败")
	}
	//以png 格式写入文件
	err = png.Encode(imgfile, img)
	if err != nil {
		return "", errors.New("创建默认用户头像写入文件失败")
	}
	return filePath, nil
}
