package main

import (
	"context"
	"net/http"

	humanize "github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
)

func summaryHandler(c *gin.Context) {
	var (
		userCount          int64
		jobCount           int64
		totalCanTransacted int64
		ctx                = context.Background()
	)

	// Count users:
	ui := fireStore.Collection("users").Documents(ctx)
	for {
		d, err := ui.Next()
		if err != nil {
			break
		}
		user := d.Data()
		if user["type"] == "User" {
			userCount++
		} else if user["type"] == "Provider" && user["whitelisted"] == true {
			userCount++
		}
	}

	// Count jobs and job fees:
	ji := fireStore.Collection("jobs").Documents(ctx)
	for {
		d, err := ji.Next()
		if err != nil {
			break
		}
		var job Job
		d.DataTo(&job)

		if job.ID == "" {
			continue
		}

		logger.Infof("job id: %s", job.ID)
		jobCount++
		// switch job.State {
		// case "Awaiting Escrow":
		// 	js.awaitingEscrowCount++
		// case "Offer pending":
		// 	js.offerPendingCount++
		// case "Cancelled":
		// 	js.cancelledCount++
		// case "Pending completion":
		// 	js.completionPendingCount++
		// case "Complete":
		// 	js.completedCount++
		// case "Review added":
		// 	js.completedCount++
		// 	js.reviewedCount++
		// default:
		// 	js.otherCount++
		// }

		for _, al := range job.ActionLog {
			totalCanTransacted = (totalCanTransacted + int64(al.AmountCan))
		}
	}

	sr := SummaryResponse{
		TotalUsers:            userCount,
		TotalJobsCompleted:    jobCount,
		TotalCanTransacted:    totalCanTransacted,
		TotalCanFees:          (0.01 * float64(totalCanTransacted)),
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
