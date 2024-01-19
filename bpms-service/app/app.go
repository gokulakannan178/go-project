package app

import (
	"context"
	"bpms-service/daos"
	"bpms-service/models"
)

//GetApp :""
func GetApp(c context.Context, d *daos.Daos) *models.Context {
	ctx := new(models.Context)
	ctx.CTX = c
	ctx.DB, ctx.Session = daos.GetDB(c, d)
	return ctx
}
