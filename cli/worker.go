package cli

// import (
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"

// 	repositories "srating/infrastructure/repositories"
// 	"srating/mail"
// 	"srating/tasks"
// 	"srating/services"
// 	"srating/utils"

// 	"github.com/hibiken/asynq"
// 	"github.com/spf13/cobra"
// )

// func worker() *cobra.Command {
// 	return &cobra.Command{
// 		Use: "worker",
// 		Run: func(cmd *cobra.Command, args []string) {
// 			initWorker()
// 		},
// 	}
// }

// func initWorker() {
// 	// Configure Redis connection options.
// 	redisOpts := asynq.RedisClientOpt{
// 		Addr:     env.RedisURL,
// 		Password: env.RedisPassword,
// 		DB:       env.RedisDB,
// 	}

// 	// Create an instance of the worker.
// 	worker := asynq.NewServer(redisOptsq.Config{})

// 	// Create a multiplexer.
// 	mux := asynq.NewServeMux()
// 	// Configure timeout.
// 	timeout := time.Duration(env.RequestTimeout) * time.Second
// 	// Create repositories and use cases.
// 	var (
// 		mailer = mail.NewGmailSender(env.EmailSenderName, env.EmailSenderAddress, env.EmailSenderPassword)
// 		ur     = repositories.NewUserRepository(db)
// 		uu     = services.NewUserService(ur, timeout)
// 		btkr   = repositories.NewBookingTrackingRepository(db)
// 		btku   = services.NewBookingTrackingService(btkr, timeout)
// 	)
// 	// Register task handlers.
// 	mux.Handle(tasks.TaskSendVerifyEmail, tasks.NewProcessTaskSendVerifyEmail(uu, mailer))
// 	mux.Handle(tasks.TaskCreateBookingTracking, tasks.NewProcessTaskCreateBookingTracking(btku))
// 	mux.Handle(tasks.TaskUpdateBookingTracking, tasks.NewProcessTaskUpdateBookingTracking(btku))
// 	mux.Handle(tasks.TaskCancelBookingTracking, tasks.NewProcessTaskCancelBookingTracking(btku))
// 	// Handle termination signals (SIGTERM and SIGINT).
// 	sigCh := make(chan os.Signal, 1)
// 	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
// 	go func() {
// 		if err := worker.Run(mux); err != nil {
// 			utils.LogError(err, "Worker could not run")
// 		}
// 	}()

// 	<-sigCh
// 	worker.Shutdown()
// }
