package handler

import (
	"net/http"
	"os"
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"io"
)

func ImageUpload(c echo.Context, file *multipart.FileHeader, fileName string) error {
    // 画像を保存するディレクトリパスを設定
    saveDir := "/public/uploads/"
    savePath := saveDir + fileName

    // ディレクトリを作成
    if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to create directory")
    }

    // 画像ファイルを保存
    outFile, err := os.Create(savePath)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to create image file")
    }
    defer outFile.Close()

    src, err := file.Open()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to open image file")
    }
    defer src.Close()

    _, err = io.Copy(outFile, src)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Failed to save image file")
    }

    return c.JSON(http.StatusOK, "Image uploaded and saved successfully.")
}
