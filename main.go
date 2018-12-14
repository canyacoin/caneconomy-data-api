package main

import (
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	logging "github.com/op/go-logging"
)

var (
	serviceID             = "caneconomy-data-api"
	router                *gin.Engine
	logger                = logging.MustGetLogger("main")
	startedAt             = time.Now()
	cacheUpdatedAt        time.Time
	ethNodeURL            string
	canworkEscrowContract string
	nodeConnection        *ethclient.Client
	// contractInstance      *CanWork
)

func init() {
	logFormatter := logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.10s} %{id:03x}%{color:reset} %{message}`)
	logging.SetFormatter(logFormatter)
	consoleBackend := logging.NewLogBackend(os.Stdout, "", 0)
	consoleBackend.Color = true
	logging.SetLevel(logging.DEBUG, "main")

	ethNodeURL = mustGetenv("ETH_NODE_URL")
	canworkEscrowContract = mustGetenv("CANWORK_ESCROW_HASH")

	router = gin.Default()
	router.Use(gin.Logger())

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}
	router.Use(cors.New(config))

	router.GET("/status", statusHandler)
	router.GET("/summary", summaryHandler)
}

func main() {
	router.Run()
}
