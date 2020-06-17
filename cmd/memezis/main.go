package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/cherya/memezis/config"
	fs "github.com/cherya/memezis/internal/filestore"
	"github.com/cherya/memezis/internal/memezis"
	p "github.com/cherya/memezis/internal/memezis"
	s "github.com/cherya/memezis/internal/store"
	q "github.com/cherya/memezis/pkg/queue"

	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func initEnv() {
	env := flag.String("env", "local.env", "env file with config values")
	flag.Parse()
	log.Printf("Loading env from %s", *env)
	err := godotenv.Load(*env)

	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	logEnv(env)
}

func logEnv(env *string) {
	envMap, _ := godotenv.Read(".env", *env)
	for key, val := range envMap {
		fmt.Printf("[godotenv] %s = %s\n", key, val)
	}
}

func main() {
	initEnv()

	db, err := sqlx.Connect("postgres", config.GetValue(config.DatabaseDsn))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := s.NewStore(db)

	var redisPool = &redis.Pool{
		MaxActive: 5,
		MaxIdle:   5,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", config.GetValue(config.RedisAddress), redis.DialPassword(config.GetValue(config.RedisPassword)))
		},
	}

	queue := q.NewManager(redisPool)
	f := fs.NewStore(
		config.GetValue(config.S3BucketURL),
		config.GetValue(config.S3Endpoint),
		config.GetValue(config.S3Key),
		config.GetValue(config.S3Secret),
		config.GetValue(config.S3Region),
		config.GetValue(config.S3BucketName),
	)
	app := p.NewMemezis(
		store,
		queue,
		f,
		//TODO: remove hardcode
		map[string]*memezis.Client{
			config.GetValue(config.BotClientKey): {Name: memezis.SourceMemezisBot},
		},
	)

	if err := app.Run(8080); err != nil {
		log.Fatal(err)
	}
}
