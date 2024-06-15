package api

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kn-kraken/hackwarsaw-fintech/lib/db"
)

func AddRoutes(r *gin.Engine, database *db.Database) {
	r.GET("/real-estates", func(c *gin.Context) {
    var req RealEstateScoresRequest
    
    err := c.Bind(&req)
		if err != nil {
			slog.Error("validating request", "error", err)
			return
		}

		businesses, err := database.ListBusinessesInArea(req.BusinessType)
		if err != nil {
			slog.Error("listing businesses in area", "error", err)
			return
		}

		c.JSON(http.StatusOK, NewRealEstateScores(businesses))
	})
}
