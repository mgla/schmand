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
	"gopkg.in/yaml.v2"
)

type Config struct {
	Units []string
	Fractions []string
}

var (
	consumerKey       = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret    = getenv("TWITTER_CONSUMER_SECRET")
	accessToken       = getenv("TWITTER_ACCESS_TOKEN")
	accessTokenSecret = getenv("TWITTER_ACCESS_TOKEN_SECRET")
	environment       = getenvf("ENVIRONMENT", "dev")
	log = &logger{logrus.New()}
	config Config
)

func readConfig() {
	data := MustAsset("data/phrases.yaml")
	err := yaml.Unmarshal(data, &config)
	if err != nil {
		log.Critical(err)
	}

}

func getenvf(key, fallback string) string {
	res := os.Getenv(key)
	if res == "" {
		return fallback
	}
	return res
}

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
	readConfig()

	choice := rand.Intn(8)
	number := "0"
	if choice < 4 {
		number = fmt.Sprintf("%d", rand.Intn(512))
	} else if choice < 6 {
		number = fmt.Sprintf("%.2f", rand.Float32() + float32(rand.Intn(4096)))
	} else {
		number = fmt.Sprintf("%d %s", (rand.Intn(9) + 1), config.Fractions[rand.Intn(len(config.Fractions))])
	}
	tweet := fmt.Sprintf("%s %s Schmand", number, config.Units[rand.Intn(len( config.Units ))])
	if (environment == "prod") {
		_, err := api.PostTweet(tweet, url.Values{})
		if err != nil {
			log.Critical(err)
		}
	} else {
		log.Info(fmt.Sprintf("Tweet: %s", tweet))
	}
}

func main() {
	if (environment == "prod") {
		lambda.Start(tweetFeed)
	}
	tweetFeed()
}

type logger struct {
	*logrus.Logger
}

func (log *logger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *logger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *logger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *logger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }
