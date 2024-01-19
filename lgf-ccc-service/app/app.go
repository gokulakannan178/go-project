package app

import (
	"context"
	"lgf-ccc-service/daos"
	"lgf-ccc-service/models"
)

//GetApp :""
func GetApp(c context.Context, d *daos.Daos) *models.Context {
	ctx := new(models.Context)
	ctx.CTX = c
	ctx.DB, ctx.Session, ctx.Client = daos.GetDB(c, d)
	return ctx
}
