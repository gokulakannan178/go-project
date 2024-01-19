package app

import (
	"context"
	"municipalproduct1-service/daos"
	"municipalproduct1-service/models"
)

//GetApp :""
func GetApp(c context.Context, d *daos.Daos) *models.Context {
	ctx := new(models.Context)
	ctx.CTX = c
	ctx.DB, ctx.Session, ctx.Client = daos.GetDB(c, d)
	pc, err := d.GetSingleProductConfiguration(ctx, "1")
	if err != nil {
		panic(err)
	}
	ctx.ProductConfig = *pc
	return ctx
}
