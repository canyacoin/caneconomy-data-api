package main

import (
	"context"
	"fmt"
	"net/http"

	humanize "github.com/dustin/go-humanize"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"gitlab.com/canya-com/shared/data-structures/canwork"
)

func summaryHandler(c *gin.Context) {
	var (
		userCount          int64
		jobCount           int64
		totalCanTransacted int64
		ctx                = context.Background()
		err                error
	)

	// Count users:
	ui := fireStore.Collection("users").Documents(ctx)
	for {
		d, err := ui.Next()
		if err != nil {
			break
		}
		var user canwork.UserDocument
		d.DataTo(&user)

		if user.Type == "User" {
			userCount++
		} else if user.Type == "Provider" && user.WhiteListed == "true" {
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
		jobCount++

		for _, al := range job.ActionLog {
			totalCanTransacted = (totalCanTransacted + int64(al.AmountCan))
		}
	}

	nodeConnection, err = ethclient.Dial(ethNodeURL)
	if err != nil {
		logger.Fatalf("error connecting to node at: %s error was: %s", ethNodeURL, err.Error())
	}

	canyaCoinContractInstance, err := NewCanYaCoin(common.HexToAddress(canyaCoinContract), nodeConnection)
	if err != nil {
		m := fmt.Sprintf("error loading contract: %s error was: %s", canyaCoinContract, err.Error())
		logger.Errorf(m)
		c.JSON(http.StatusInternalServerError, gin.H{"message": m})
		return
	}

	address := common.HexToAddress(canworkEscrowContract)
	logger.Debugf("checking hash: %s", canworkEscrowContract)

	balance, err := canyaCoinContractInstance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		m := fmt.Sprintf("unable to get balance of contract: %s error was: %s", canyaCoinContract, err.Error())
		logger.Errorf(m)
		c.JSON(http.StatusInternalServerError, gin.H{"message": m})
		return
	}

	bal := canAmountToFloat(balance)

	sr := SummaryResponse{
		TotalUsers:            userCount,
		TotalJobsCompleted:    jobCount,
		TotalCanTransacted:    totalCanTransacted,
		TotalCanFees:          (0.01 * float64(totalCanTransacted)),
		EscrowContractAddress: canworkEscrowContract,
		TotalCanInEscrow:      bal,
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
