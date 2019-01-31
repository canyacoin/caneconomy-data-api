package main

import (
	"context"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	logging "github.com/op/go-logging"
	"google.golang.org/api/option"
)

var (
	serviceID             = "caneconomy-data-api"
	router                *gin.Engine
	logger                = logging.MustGetLogger("main")
	startedAt             = time.Now()
	cacheUpdatedAt        time.Time
	ethNodeURL            string
	canworkEscrowContract string
	canyaCoinContract     = "0x1d462414fe14cf489c7a21cac78509f4bf8cd7c0"
	nodeConnection        *ethclient.Client
	fireStore             *firestore.Client
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
	firebaseKeyFile := mustGetenv("FIREBASE_KEY_PATH")

	sa := option.WithCredentialsFile(firebaseKeyFile)
	app, err := firebase.NewApp(context.Background(), nil, sa)
	if err != nil {
		logger.Fatalf("unable parse service account credentials: %v", err)
	}
	fireStore, err = app.Firestore(context.Background())
	if err != nil {
		logger.Fatalf("unable to establish firestore connection: %v", err)
	}

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
