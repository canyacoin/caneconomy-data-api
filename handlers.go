package main

import (
	"net/http"

	humanize "github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
)

func summaryHandler(c *gin.Context) {
	sr := SummaryResponse{
		TotalCanFees:          111,
		EscrowContractAddress: canworkEscrowContract,
	}
	c.JSON(http.StatusOK, sr)
}

func statusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":         "OK",
		"serviceID":      serviceID,
		"serviceStarted": humanize.Time(startedAt),
	})
}
