package cat_handler

import (
	"github.com/amitizle/thecatapi_client"
	// "github.com/spf13/viper"
	"gopkg.in/telegram-bot-api.v4"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

var (
	api                *thecatapi.Client
	NewAnimationUpload string
	pollForGifsTime    = time.Second * 5
	gifCacheSize       = 50
	gifFilesDir        = "/tmp/cat_gifs"
	catRegex           = regexp.MustCompile("(?:\\A|\\s)(חתולה|ירון)(?:\\s|\\z)")
)

func Callback() func(*tgbotapi.BotAPI, *tgbotapi.Message) {
	return func(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
		if !catRegex.MatchString(message.Text) {
			return
		}
		gifFilePath, err := getCachedGif()
		if err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Cannot find a gif, tell Amit to pff")
			msg.ReplyToMessageID = message.MessageID
			bot.Send(msg)
			return
		}
		defer os.Remove(gifFilePath) // clean up
		msg := tgbotapi.NewDocumentUpload(message.Chat.ID, gifFilePath)
		msg.ReplyToMessageID = message.MessageID
		bot.Send(msg)
	}
}

func init() {
	// gifFilesDir = viper.GetString("GIF_DOWNLOAD_DIR") // TODO default temp
	log.Printf("File path %s", gifFilesDir)
	api, err := thecatapi.NewClient() // TODO
	if err != nil {
		log.Fatalf("Error creating thecatapi client: %v", err)
	}
	// api.ApiKey = viper.GetString("THE_CAT_API_KEY")
	go pollGifs(gifFilesDir, pollForGifsTime, gifCacheSize, api)
}

func getCachedGif() (string, error) {
	files, err := ioutil.ReadDir(gifFilesDir)
	if err != nil {
		return "", err
	}
	if len(files) == 0 {
		populateCache(gifFilesDir, gifCacheSize, api)
	}
	filesAfterPopulated, err := ioutil.ReadDir(gifFilesDir)
	fileName := filesAfterPopulated[rand.Intn(len(filesAfterPopulated))].Name()
	return filepath.Join(gifFilesDir, fileName), nil
}

func populateCache(downloadDir string, cacheSize int, api *thecatapi.Client) {
	files, err := ioutil.ReadDir(downloadDir)
	if err != nil {
		log.Printf("Cannot read dir %s", downloadDir)
		return
	}
	if len(files) < cacheSize {
		needToDownload := cacheSize - len(files)
		log.Printf("Downloading %d gifs", needToDownload)
		searchResponses, err := api.Images.Search([]string{"gif"}, "json", needToDownload)
		if err != nil {
			log.Printf("Error %v", err)
			return
		}
		for _, searchResponse := range searchResponses {
			go downloadGifCat(searchResponse.Url, downloadDir) // duplicates are possible ¯\_(ツ)_/¯
		}
	}

}

func pollGifs(downloadDir string, sleepTime time.Duration, cacheSize int, api *thecatapi.Client) {
	err := os.MkdirAll(downloadDir, 0755)
	if err != nil {
		log.Fatalf("Error mkdir -p: %v", err)
	}
	for {
		time.Sleep(sleepTime)
		populateCache(downloadDir, cacheSize, api)
	}
}

func downloadGifCat(url string, downloadDir string) error {
	log.Printf("Downloading file %s", url)

	// // Create the file
	tmpfile, err := ioutil.TempFile(downloadDir, "*.gif")
	if err != nil {
		return err
	}
	defer tmpfile.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(tmpfile, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
