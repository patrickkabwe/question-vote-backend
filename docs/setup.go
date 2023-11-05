package docs

import (
	"github.com/gofiber/contrib/swagger"
)

var SwaggerFilePath = "./docs/swagger.json"

var Config = swagger.Config{
	BasePath: "/",
	FilePath: SwaggerFilePath,
	Path:     "docs",
	Title:    "Swagger API Docs",
	CacheAge: 0,
}
