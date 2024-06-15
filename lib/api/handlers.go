package api

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kn-kraken/hackwarsaw-fintech/lib/db"
	"github.com/kn-kraken/hackwarsaw-fintech/lib/utils"
)

func AddRoutes(r *gin.Engine, db *db.Database) {
	r.GET("/real-estates", func(c *gin.Context) {
		businesses, err := db.ListBusinessesInArea()
		if err != nil {
			slog.Error("listing businesses in area", "error", err)
			return
		}

		c.JSON(http.StatusOK, utils.MapRef(businesses, NewBusiness))
	})
}
