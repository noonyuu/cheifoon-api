package handler

import (
	"path/filepath"
	"net/http"
	"github.com/labstack/echo/v4"

	"fmt"
)

func GetImage(c echo.Context) error {
    imageName := c.Param("get")
    if imageName == "" {
        return c.String(http.StatusBadRequest, "Image name is missing")
    }

    imagePath := filepath.Join("/public/uploads/", imageName)
    c.Response().Header().Set("Content-Type", "image/jpeg")  // 画像のContent-Typeを指定
		fmt.Printf("%T\n", "wa----",c.File(imagePath))
		// fmt.Println("siiiiiii",c.File(imagePath))
    return c.File(imagePath)
}