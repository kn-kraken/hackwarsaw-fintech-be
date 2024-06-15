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

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			slog.Error("validating request", "error", err)
			return
		}

		businesses, err := database.ListBusinessesInArea(
			req.BusinessType,
			req.Longitude,
			req.Latitude,
			req.Distance,
		)
		if err != nil {
			slog.Error("listing businesses in area", "error", err)
			return
		}

		c.JSON(http.StatusOK, NewRealEstateScores(businesses))
	})

	r.GET("/polygons", func(c *gin.Context) {
		polygons, err := database.ListPolygons()
		if err != nil {
			slog.Error("listing businesses in area", "error", err)
			return
		}

		c.JSON(http.StatusOK, polygons)
	})
}
