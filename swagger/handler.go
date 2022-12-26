package swagger

import (
	_ "github.com/glebnaz/witcher/statik"
	"github.com/labstack/echo/v4"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	"net/http"
	ioFs "io/fs"
)

func AddSwagger(g *echo.Echo, swaggerJSON string) error {
	filesystem, err := fs.New()
	if err != nil {
		log.Debugf("Error creating statik filesystem for swagger: %s", err)
		return err
	}
	staticServer := http.FileServer(filesystem)

	g.GET("/swagger.json", func(c echo.Context) error {
		return c.File(swaggerJSON)
	})

	filesystem.Open()

	sh := http.StripPrefix("/swaggerui/", staticServer)
	go http.ListenAndServe(":8080", sh)
	eh := echo.WrapHandler(sh)
	g.StaticFS("/swaggerui/", )

	g.
	return nil
}
