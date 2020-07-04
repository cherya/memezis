package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/cherya/memezis/config"
	"github.com/cherya/memezis/internal/app/auth"
	fs "github.com/cherya/memezis/internal/app/filestore"
	"github.com/cherya/memezis/internal/app/memezis"
	s "github.com/cherya/memezis/internal/app/store"
	"github.com/cherya/memezis/pkg/logger"
	_ "github.com/cherya/memezis/pkg/logger"
	desc "github.com/cherya/memezis/pkg/memezis"
	q "github.com/cherya/memezis/pkg/queue"
	_ "github.com/cherya/memezis/web/statik"

	"github.com/go-chi/chi"
	"github.com/gomodule/redigo/redis"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	statik "github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	clog "github.com/utrack/clay/v2/log"
	"github.com/utrack/clay/v2/transport/middlewares/mwgrpc"
	"github.com/utrack/clay/v2/transport/server"
	"google.golang.org/grpc"
)

const logDateFormat = "02-01-2006 15:04:05"

func main() {
	initEnv()
	logger.Init(log.DebugLevel, logDateFormat)
	l := logger.NewLogger(log.DebugLevel, logDateFormat)

	db, err := sqlx.Connect("postgres", config.GetValue(config.DatabaseDsn))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := s.NewStore(db)

	redisPool := initRedisPool()

	queue := q.NewManager(redisPool, "memezis")
	f := fs.NewStore(
		config.GetValue(config.S3BucketURL),
		config.GetValue(config.S3Endpoint),
		config.GetValue(config.S3Key),
		config.GetValue(config.S3Secret),
		config.GetValue(config.S3Region),
		config.GetValue(config.S3BucketName),
	)
	router := chi.NewRouter()
	serveSwaggerui(router)

	memezis := memezis.NewMemezis(
		store,
		queue,
		f,
	)

	httpPort := config.GetInt(config.HttpPort)
	httpServer := server.NewServer(
		httpPort,
		server.WithHTTPMux(router),
		server.WithGRPCUnaryMiddlewares(
			mwgrpc.UnaryPanicHandler(clog.Default),
			grpc_logrus.UnaryServerInterceptor(log.NewEntry(l)),
			grpc_auth.UnaryServerInterceptor(auth.NewAuthenticator(getClients()).AuthMiddleware),
		),
	)

	grpcPort := config.GetInt(config.GrpcPort)
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			grpc_logrus.UnaryServerInterceptor(log.NewEntry(l)),
			grpc_auth.UnaryServerInterceptor(auth.NewAuthenticator(getClients()).AuthMiddleware),
		),
	)
	desc.RegisterMemezisServer(grpcServer, memezis)

	go func() {
		err = httpServer.Run(memezis)
		if err != nil {
			log.Error(err)
		}
	}()

	go func() {
		err = grpcServer.Serve(listen)
		if err != nil {
			log.Error(err)
		}
	}()

	log.Infof("app running on ports: http=%d, grpc=%d", httpPort, grpcPort)

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan

	httpServer.Stop()
	grpcServer.Stop()
}

func initEnv() {
	env := flag.String("env", "local.env", "env file with config values")
	flag.Parse()
	log.Info("Loading env from", *env)
	err := godotenv.Load(*env)

	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	logEnv(env)
}

func logEnv(env *string) {
	envMap, _ := godotenv.Read(".env", *env)
	for key, val := range envMap {
		log.Printf("%s = %s", key, val)
	}
}

func serveSwaggerui(mux *chi.Mux) {
	statikFS, err := statik.New()
	if err != nil {
		panic(err)
	}
	fileServer := http.StripPrefix("/swaggerui/", http.FileServer(statikFS))
	mux.Get("/swaggerui/*", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Vary", "Accept-Encoding")
		//w.Header().Set("Cache-Control", "public, max-age=7776000")
		fileServer.ServeHTTP(w, r)
	})
}

func getClients() map[string]*auth.Client {
	//TODO: remove hardcode
	return map[string]*auth.Client{
		config.GetValue(config.BotClientKey): {Name: "memezis_bot"},
	}
}

func initRedisPool() *redis.Pool {
	pool := &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp",
				config.GetValue(config.RedisAddress),
				redis.DialPassword(config.GetValue(config.RedisPassword)))
		},
	}
	conn := pool.Get()
	defer conn.Close()

	_, err := conn.Do("PING")
	if err != nil {
		log.Fatal("can't connect to redis. PING failed: ", err)
	}

	return pool
}