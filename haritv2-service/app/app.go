package app

import (
	"context"
	"fmt"
	"haritv2-service/daos"
	"haritv2-service/models"
)

//GetApp :""
func GetApp(c context.Context, d *daos.Daos) *models.Context {
	ctx := new(models.Context)
	ctx.CTX = c
	ctx.DB, ctx.Session, ctx.Client = daos.GetDB(c, d)
	auth, ok := c.Value("Authorization").(models.Authentication)
	if !ok {
		fmt.Println("Not Ok")
	} else {
		ctx.Auth = auth
		fmt.Println(ctx.Auth)
	}
	return ctx
}
