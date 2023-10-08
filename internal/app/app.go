package app

import (
	"context"
	"git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/internal/app/connector"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	handlerSender "git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/internal/app/handler/sender"
	repositoriesSender "git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/internal/app/repositories/sender"
	usecaseSender "git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/internal/app/usecase/sender"

	handlerClient "git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/internal/app/handler/client"
	repoClient "git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/internal/app/repositories/client"
	usecaseClient "git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/internal/app/usecase/client"

	"git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/internal/app/configuration"

	handlerMonitoring "git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/internal/app/handler/dashboard_all_transaction"
	"git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/internal/app/middleware"
	repositoryMonitoring "git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/internal/app/repositories/dashboard_all_transaction"
	usecaseMonitoring "git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/internal/app/usecase/dashboard_all_transaction"
	"git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/pkg/logrusx"
	"git-rbi.jatismobile.com/jns-revamp/backend/tool-dashboard/pkg/logx"
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func Run() {

	cfg, err := configuration.LoadConfig("config.yaml")
	if err != nil {
		panic("Error loading config file")
	}

	ctxLog := context.Background()
	logger := logrusx.NewProvider(&ctxLog, cfg.Log)

	db := connector.InitSqlConnection(&cfg)

	router := gin.Default()
	router.Use(middleware.Cors(&cfg))
	router.Use(middleware.TimeoutMiddleware(time.Duration(cfg.App.Timeout) * time.Second))

	repoClient := repoClient.NewRepositoryClient(db, logger.GetLogger("client_repository_fetch_list"))
	useCaseClient := usecaseClient.NewClientUseCase(repoClient, logger.GetLogger("client_usecase_fetch_list"))
	handlerClient.NewHandlerClient(router, useCaseClient, logger.GetLogger("client_handler"), &cfg)

	repoSender := repositoriesSender.NewRepositorySender(db, logger.GetLogger("repo_sender"))
	useCaseSender := usecaseSender.NewSenderUseCase(repoSender, logger.GetLogger("usecase_sender"))
	handlerSender.NewHandlerSender(router, useCaseSender, logger.GetLogger("handler_sender"), &cfg)

	repoMonitoring := repositoryMonitoring.NewMonitoringRepository(logger.GetLogger("repo-monitoring"), db)
	useCaseMonitoring := usecaseMonitoring.NewMonitoringUseCase(logger.GetLogger("usecase-monitoring"), repoMonitoring, repoClient, repoSender)
	handlerMonitoring.NewHandlerMonitoring(router, useCaseMonitoring, useCaseClient, logger.GetLogger("Handler-monitoring"), &cfg)

	// Create a server with desired configurations
	server := &http.Server{
		Addr:    "0.0.0.0:" + cfg.App.Port,
		Handler: router,
	}

	// Start the server in a separate goroutine
	go func() {
		logx.GetLogger().Info("Server running at: ", cfg.App.Port)
		if err := server.ListenAndServe(); err != nil {
			logger.GetLogger("tool-dashboard").Fatal("", "Server error:", err)
		}
	}()

	// Now, set up the signal handling to catch SIGINT (Ctrl+C) and SIGTERM (kill)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	<-stop

	// Create a context with a timeout to allow for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		logger.GetLogger("tool-dashboard").Fatal("", "Error during server shutdown: %v", err)
	}

	logger.GetLogger("tool-dashboard").Info("", "Server gracefully shut down")
}
