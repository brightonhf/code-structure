package cmd

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"code_structure/internals/ports"
)

type HTTPCommand interface {
	Run() error
}

type httpCommand struct {
	BaseCommand
	api ports.API
}

func NewHTTPCommand(command BaseCommand, api ports.API) HTTPCommand {
	return &httpCommand{BaseCommand: command, api: api}
}

func (h httpCommand) Run() error {

	r := gin.Default()

	r.GET("selection", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"count": h.api.GetSelection(),
		})
	})

	return r.Run(":7000")

}
