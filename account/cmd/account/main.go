package main


import (
	"log"
	"time"
	"github.com/olujimiAdebakin/ProtoGraph/account"
	"github.com/kelseyhightower/envconfig"
	"github.com/avast/retry-go/v4"
	_"github.com/tinrab/retry"
)


type Config struct{
	DatabaseURL string `envconfig:"DATABASE_URL"`
}


func main(){
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r account.AccountRepository
	
	// CORRECT: retry.Do with options
	err = retry.Do(
		func() error {
			var err error
			r, err = account.NewPostgresRepositry(cfg.DatabaseURL)
			if err != nil {
				log.Printf("failed to connect to database: %v", err)
			}
			return err
		},
		retry.Attempts(5),          // 0 = infinite retries
		retry.Delay(2*time.Second), // Delay between retries
		retry.OnRetry(func(n uint, err error) {
			log.Printf("Retry attempt %d: %v", n, err)
		}),
	)
	
	if err != nil {
		log.Fatal("Failed to connect after retries: ", err)
	}
	
	defer r.Close()

	log.Println("Listening on port 8080......")
	s := account.NewService(r)
	log.Fatal(account.ListenGRPCServer(s, 8080))
}