package controller

import (
	"github.com/goadesign/goa"
	"github.com/jarifibrahim/fabric8-foo/app"
)

// StatusController implements the status resource.
type StatusController struct {
	*goa.Controller
}

// NewStatusController creates a status controller.
func NewStatusController(service *goa.Service) *StatusController {
	return &StatusController{Controller: service.NewController("StatusController")}
}

// Show runs the show action.
func (c *StatusController) Show(ctx *app.ShowStatusContext) error {
	res := &app.Status{
		BuildTime: app.BuildTime,
		Commit:    app.Commit,
		StartTime: app.StartTime,
	}
	return ctx.OK(res)
}
