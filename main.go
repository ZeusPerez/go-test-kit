// This server is just an example for the tests and that's why it is not validating, logging, dividding in packages and structs, etc.

package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

var cfg Config

type Config struct {
	RandomURL string `default:"https://www.random.org/integers/?num=1&min=1&max=10000&col=1&base=10&format=plain&rnd=new"`
}

func main() {
	http.HandleFunc("/random-sum", randomSumHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func randomSumHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	randomNumberProvider := NewRandomNumber(cfg)
	sum, err := randomSum(ctx, randomNumberProvider)
	if err != nil {
		http.Error(w, "500 error", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "%d", sum)
}

func randomSum(ctx context.Context, numProvider Random) (int, error) {
	var x, y int
	var err error
	var mu sync.Mutex

	errGroup, _ := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		logrus.Info("Making request to random number API")

		mu.Lock()
		defer mu.Unlock()

		x, err = numProvider.GetRandomNum()
		logrus.Infof("Obtained random number: %d", x)
		return err
	})

	errGroup.Go(func() error {
		logrus.Info("Making request to random number API")

		mu.Lock()
		defer mu.Unlock()

		y, err = numProvider.GetRandomNum()
		logrus.Infof("Obtained random number: %d", y)
		return err
	})

	err = errGroup.Wait()
	if err != nil {
		fmt.Printf("Error getting randomg numbers: %v", err)
	}

	return sum(x, y), nil
}

func sum(x, y int) int {
	return x + y
}

//go:generate mockery --case underscore --inpackage --name Random
type Random interface {
	GetRandomNum() (int, error)
}

type randomNumber struct {
	cfg Config
}

func NewRandomNumber(cfg Config) Random {
	return &randomNumber{cfg: cfg}
}

func (rn *randomNumber) GetRandomNum() (int, error) {
	resp, err := http.Get(rn.cfg.RandomURL)
	if err != nil {
		logrus.Errorf("Error making the request to %v: %v", rn.cfg.RandomURL, err)
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Error parsing the response: %v", err)
		return 0, err
	}

	strInt := string(string(body))
	myInt, err := strconv.Atoi(strings.TrimSpace(strInt))

	if err != nil {
		logrus.Errorf("Error converting the response: %v", err)
		return 0, err
	}

	return myInt, nil
}

func init() {
	err := envconfig.Process("testkit", &cfg)
	if err != nil {
		logrus.Fatal(err.Error())
	}
	logrus.Infof("Config: %+v", cfg)
}
