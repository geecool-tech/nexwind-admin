package common

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"nexwind.net/admin/server/app"
	"nexwind.net/admin/server/constant"
	"nexwind.net/admin/server/initialize/resp"
	"os"
	"path/filepath"
	"time"
)

// Common 公共 api
type Common struct{}

// Upload 上传文件
func (u *Common) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_FAIL, Msg: "获取上传文件失败"})
		return
	}
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_FAIL, Msg: "获取上传文件失败"})
		return
	}
	ext := filepath.Ext(file.Filename)
	randomFilename := generateRandomFilename(ext)
	today := time.Now().Format("20060102")
	saveDir := filepath.Join("uploads", today)
	err = os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_FAIL, Msg: err.Error()})
		return
	}
	savePath := filepath.Join(saveDir, randomFilename)
	if err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_FAIL, Msg: err.Error()})
		return
	}
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		app.Resp.ErrorRes(c, resp.ErrorRes{Code: constant.CODE_FAIL, Msg: err.Error()})
		return
	}

	app.Resp.SuccessRes(c, map[string]string{
		"message": "文件上传成功",
		"file":    file.Filename,
		"url":     fmt.Sprintf("/uploads/%s/%s", today, randomFilename),
	})
}

// generateRandomFilename 生成随机文件名
func generateRandomFilename(ext string) string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Println("Failed to generate random filename:", err)
		return ""
	}
	return fmt.Sprintf("%x%s", b, ext)
}
