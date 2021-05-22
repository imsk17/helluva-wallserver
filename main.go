package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

var (
	WallpaperDir = "./wallpapers"
	Port = ":8080"
)

func init() {
	osWallDir := os.Getenv("WALLPATH")
	osPort := os.Getenv("PORT")
	if osWallDir != "" {
		WallpaperDir = osWallDir
	}
	if osPort != "" {
		Port = ":"+osPort
	}
}

func main() {
	app := fiber.New()

	app.Use("/monitor", monitor.New())
	app.Use(recover.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		var walls []string
		err := filepath.WalkDir(WallpaperDir, func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				walls = append(walls, "/wallpaper/"+d.Name())
			}
			return nil
		})
		if err != nil {
			return err
		}
		return ctx.JSON(map[string]interface{}{
			"wallpapers": walls,
		})
	})

	app.Get("/wallpaper/:wallpaper", func(ctx *fiber.Ctx) error {
		wallpaper := ctx.Params("wallpaper", "404")
		return ctx.SendFile(fmt.Sprintf("./wallpapers/%s", wallpaper))
	})

	log.Panic(app.Listen(Port))
}