package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/yaegaki/itunes-app-interface"
)

func main() {
	err := NowPlaying()
	if err != nil {
		log.Fatal(err)
	}
}

const ConsumerKey = "Your Consumer Key (API Key)"
const ConsumerSecret = "Your Consumer Secret (API Secret)"
const AccessToken = "Your Access Token"
const AccessTokenSecret = "Your Access Token Secret"

func NowPlaying() error {
	err := itunes.Init()
	if err != nil {
		return err
	}
	defer itunes.UnInit()

	it, err := itunes.CreateItunes()
	if err != nil {
		return err
	}
	defer it.Close()

	t, err := it.CurrentTrack()
	if err != nil {
		return errors.New("not nowplaying.")
	}
	defer t.Close()

	dir, err := ioutil.TempDir("", "go-nowplaying")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	artworks, err := t.GetArtworks()
	if err != nil {
		return err
	}

	api := initTwitterAPI()

	artwork := <-artworks
	var data url.Values
	if artwork != nil {
		err = func() error {
			defer artwork.Close()
			path, err := artwork.SaveToFile(dir, "artwork")
			if err != nil {
				return err
			}

			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()

			b, err := ioutil.ReadAll(f)
			if err != nil {
				return err
			}

			str := base64.StdEncoding.EncodeToString(b)
			media, err := api.UploadMedia(str)
			data = url.Values{}
			data.Add("media_ids", media.MediaIDString)

			return nil
		}()

		if err != nil {
			return err
		}
	}

	tweetStr := fmt.Sprintf("#nowplaying %v", t.Name)
	if t.Artist != "" {
		tweetStr = fmt.Sprintf("%v by %v", tweetStr, t.Artist)
	}

	tweet, err := api.PostTweet(tweetStr, data)
	if err != nil {
		return err
	}
	fmt.Println("Tweet Success.")
	fmt.Println("Posted:")
	fmt.Println(tweet.Text)

	return nil
}

func initTwitterAPI() *anaconda.TwitterApi {
	consumerKey := os.Getenv("GN_CONSUMER_KEY")
	if consumerKey == "" {
		consumerKey = ConsumerKey
	}
	consumerSecret := os.Getenv("GN_CONSUMER_SECRET")
	if consumerSecret == "" {
		consumerSecret = ConsumerSecret
	}
	accessToken := os.Getenv("GN_ACCESS_TOKEN")
	if accessToken == "" {
		accessToken = AccessToken
	}
	accessTokenSecret := os.Getenv("GN_ACCESS_TOKEN_SECRET")
	if accessTokenSecret == "" {
		accessTokenSecret = AccessTokenSecret
	}

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	return anaconda.NewTwitterApi(accessToken, accessTokenSecret)
}
