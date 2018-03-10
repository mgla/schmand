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
		// distance - obscure
		"nautical miles",
		"german miles",
		"inches",
		// distance - SI
		"meters",
		"decimeters",
		// weight
		"kilograms",
		"grams",
		"ounces",
		// electrical - current
		"amperes",
		// electrical - charge
		"coulomb",
		// electrical - capacitance
		"farad",
		"millifarad",
		"microfarad",
		"nanofarad",
		"picofarad",
		"becquerel",
		// electrical - resistance
		"Ω",
		// magnetic - flux
		"weber",
		"microweber",
		// digital information
		"petabytes",
		"terabytes",
		"gigabytes",
		"megabytes",
		"kilobytes",
		"bytes",
		// area
		"acres",
		// volume
		"oz",
		"liters",
		"deciliters",
		"milliliters",
		"°",
		// Energy
		"calories",
		"kcal",
		// torque
		"dynemetres",
		// luminence
		"candelas per square metre",
	)
	choice := rand.Intn(8)
	number := "0"
	if choice < 4 {
		number = fmt.Sprintf("%d", rand.Intn(512))
	} else if choice < 6 {
		number = fmt.Sprintf("%.2f", rand.Float32() + float32(rand.Intn(4096)))
	} else {
		fractions := make([]string, 0)
		fractions = append(fractions,
			"½",
			"⅓",
			"⅔",
			"¼",
			"¾",
			"⅕",
			"⅖",
			"⅗",
			"⅘",
			"⅙",
			"⅚",
			"⅛",
			"⅜",
			"⅝",
			"⅞",
		)
		number = fmt.Sprint("%d %s", (rand.Intn(9) + 1) ,fractions[rand.Intn(len(fractions))])
	}
	tweet := fmt.Sprintf("%s %s of schmand", number, units[rand.Intn(len(units))])
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
