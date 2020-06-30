package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

func HelloFullCycleHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(map[string]string{"message": "hello full cycle developers!"}))
}
