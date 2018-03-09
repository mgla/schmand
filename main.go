package main

import (
	"net/url"
	"os"
	"fmt"
	"math/rand"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/Sirupsen/logrus"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")

	log = &logger{logrus.New()}
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("did you forget your keys? " + name)
	}
	return v
}

func tweetFeed() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	log.SetLevel(logrus.InfoLevel)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)
	api.SetLogger(log)
	rand.Seed(time.Now().Unix())

	units := make([]string, 0)
	units = append(units,
		
		// temperature
		"°C",
		"KHz",
		"british thermal units",
		// time
		"days",
	        "seconds",
		"ms",
                "µs",
		// distance - astronimical
		"astronomical units",
		"parsecs",
		"milliparsecs",
		"nanoparsecs",
		"picoparsecs",
		"nautical miles",
		// distance - SI
		"meters",
		"decimeters",
		// weight
		"kilograms",
		"grams",
		// electrical - charge
		"coulomb",
		// electrical - capacitance
		"farad",
		"millifarad",
		"microfarad",
		"nanofarad",
		"picofarad",
		"becquerel",
		// data size
		"gigabytes",
		"megabytes",
		"kilobytes",
		"bytes",
		// area
		"acres",
		// volume
		"oz",
		"liters",
		"deciliters"
		"milliliters",
		"°",
		// Energy
		"calories",
		"kcal"
	)

	tweet := fmt.Sprintf("%.2f %s of schmand", (rand.Float32() + float32(rand.Intn(8096))), units[rand.Intn(len(units))])
	_, err := api.PostTweet(tweet, url.Values{})
	if err != nil {
		log.Critical(err)
	}
}

func main() {
	lambda.Start(tweetFeed)
	tweetFeed()
}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }
