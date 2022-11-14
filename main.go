package main

import (
	"context"
	"flag"
	"fmt"
	stdlog "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IRFAN374/goRestApi/common/chttp"
	"github.com/IRFAN374/goRestApi/grocery"
	grocerymw "github.com/IRFAN374/goRestApi/grocery/service"
	grocerysvctransport "github.com/IRFAN374/goRestApi/grocery/transport"
	grocerysvctransporthttp "github.com/IRFAN374/goRestApi/grocery/transport/http"
	"github.com/oklog/oklog/pkg/group"

	"github.com/IRFAN374/goRestApi/handlers"

	"github.com/IRFAN374/goRestApi/repository/mygrocery"
	mygroceryRedisRepo "github.com/IRFAN374/goRestApi/repository/mygrocery/redis"
	mygrocerymw "github.com/IRFAN374/goRestApi/repository/mygrocery/service"

	gokitlogzap "github.com/go-kit/kit/log/zap"
	kitHttp "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var env string

func init() {
	flag.StringVar(&env, "env", "", "kube env")
}

func main() {
	fmt.Println("Hello World")

	flag.Parse()

	if env == "" {
		os.Getenv("env")
	}

	ServiceName := fmt.Sprintf("%s-grocery-rest-api", env)

	// logger
	// var logLevel zapcore.Level
	debugLogger, _, _, _ := getLogger(ServiceName, zapcore.DebugLevel)

	// redis client
	redisClient := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"127.0.0.1:7000"},
	})

	var httpServerBefore = []kitHttp.ServerOption{
		kitHttp.ServerErrorEncoder(kitHttp.ErrorEncoder(chttp.EncodeError)),
	}

	// Htpp Middleware
	var mwf []mux.MiddlewareFunc

	// router
	httpRouter := mux.NewRouter().StrictSlash(false)
	httpRouter.Use(mwf...)

	var mygroceryRepo mygrocery.Repository
	{
		mygroceryRepo = mygroceryRedisRepo.NewRepository(redisClient)
		mygroceryRepo = mygrocerymw.LoggingMiddleware(log.With(debugLogger, "repository", "mygrocery Repository"))(mygroceryRepo)
	}

	var grocerySvc grocery.Service
	{
		grocerySvc = grocery.NewService(mygroceryRepo)
		grocerySvc = grocerymw.LogginMiddleware(log.With(debugLogger, "service", "gracery service"))(grocerySvc)

		grocerySvcEndpoints := grocerysvctransport.Endpoints(grocerySvc)
		grocerySvcHandler := grocerysvctransporthttp.NewHTTPHandler(&grocerySvcEndpoints, httpServerBefore...)
		httpRouter.PathPrefix("/api/vi").Handler(grocerySvcHandler)

	}

	var server group.Group
	{
		httpServer := &http.Server{
			Addr:    ":8080",
			Handler: httpRouter,
		}

		server.Add(func() error {
			fmt.Printf("starting http server, port:%d \n", 8080)
			return httpServer.ListenAndServe()
		}, func(err error) {

			// write code here for gracefull shutDown

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
			defer cancel()

			httpServer.Shutdown(ctx)
		})

	}

	{
		cancelInterrupt := make(chan struct{})

		server.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)

			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}

		}, func(err error) {
			close(cancelInterrupt)
		})
	}

	fmt.Printf("exiting...  errors: %v\n", server.Run())

	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/allgrocery", handlers.AddGrocery).Methods("GET")
	router.HandleFunc("/grocery/{name}", handlers.GetGrocery).Methods("GET")
	router.HandleFunc("/grocery", handlers.AddGrocery).Methods("POST")
	router.HandleFunc("/grocery/{name}", handlers.UpdateGrocery).Methods("PUT")
	router.HandleFunc("/grocery/{name}", handlers.DeleteGrocery).Methods("DELETE")

	stdlog.Fatal(http.ListenAndServe(":8080", router))
}

func getLogger(serviceName string, level zapcore.Level) (debugL, infoL, errorL log.Logger, zapLogger *zap.Logger) {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	ec.EncodeDuration = zapcore.StringDurationEncoder
	ec.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(ec)

	fw, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}

	core := zapcore.NewCore(encoder, fw, level)
	zapLogger = zap.New(core)

	debugL = gokitlogzap.NewZapSugarLogger(zapLogger, zap.DebugLevel)
	debugL = log.With(debugL, "service", serviceName)

	infoL = gokitlogzap.NewZapSugarLogger(zapLogger, zap.InfoLevel)
	infoL = log.With(infoL, "service", serviceName)

	errorL = gokitlogzap.NewZapSugarLogger(zap.New(zapcore.NewCore(encoder, os.Stderr, level)), zap.ErrorLevel)
	errorL = log.With(errorL, "service", serviceName)

	return

}
