package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
)

func Upload(ctx *gin.Context) {
	uploadType := ctx.Param("type")
	log.Println(uploadType)

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(400, gin.H{"error": err})
		return
	}

	wd, _ := os.Getwd()
	err = ctx.SaveUploadedFile(file, filepath.Join(wd, "upload",fmt.Sprintf("%s.jpeg",ctx.GetHeader("userId"))))
	if err != nil {
		log.Println(err)
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	ctx.String(200, fmt.Sprintf("%s upload success!", file.Filename))
}
