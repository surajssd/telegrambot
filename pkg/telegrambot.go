package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"

	"github.com/Sirupsen/logrus"
)

const (
	URL             = "https://api.telegram.org/bot"
	IndiaTimeZoneID = "Asia/Kolkata"
)

var TOKEN string
var WEBHOOK_URL string
var NAMES_FILE string
var NOPINGDAYS string
var HOUR int
var MINUTE int


func (b *BotInfo) InitBotObject() error {
	var err error
	if TOKEN == "" {
		return fmt.Errorf("Please provide TOKEN for accessing the telegram server")
	}

	b.URL, err = url.Parse(URL + TOKEN)
	if err != nil {
		return fmt.Errorf("Error parsing url: %s, Error: %v", (URL + TOKEN), err)
	}

	b.WebhookURL, err = url.Parse(WEBHOOK_URL)
	if err != nil {
		return fmt.Errorf("Error parsing url: %s, Error: %v", WEBHOOK_URL, err)
	}
	b.WebhookURL.Path = path.Join(b.WebhookURL.Path, TOKEN)

	logrus.Infof("Botinfo object initialized.")
	return nil
}

func (b *BotInfo) GetUpdates() ([]Update, error) {
	updateURL, err := url.Parse(b.URL.String())
	if err != nil {
		return []Update{}, fmt.Errorf("Error parsing url: %s, Error: %v", b.URL.String(), err)
	}
	updateURL.Path = path.Join(updateURL.Path, "getUpdates")

	resp, err := http.Get(updateURL.String())
	if err != nil {
		return []Update{}, fmt.Errorf("Error getting latest updates. Error: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Update{}, fmt.Errorf("Error reading response body. Error: %v", err)
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return []Update{}, fmt.Errorf("Could not unmarshal the response. Error: %v", err)
	}
	if response.OK {
		return response.Result, nil
	} else {
		return []Update{}, fmt.Errorf("No data received")
	}
}

func (b *BotInfo) SendMessage(chatid int, message string) error {
	sendMsgURL, err := url.Parse(b.URL.String())
	if err != nil {
		return fmt.Errorf("Error parsing url: %s, Error: %v", b.URL.String(), err)
	}
	sendMsgURL.Path = path.Join(sendMsgURL.Path, "sendMessage")

	q := sendMsgURL.Query()
	q.Add("chat_id", strconv.Itoa(chatid))
	q.Add("text", message)
	sendMsgURL.RawQuery = q.Encode()

	resp, err := http.Get(sendMsgURL.String())
	if err != nil {
		return fmt.Errorf("Could not post request. Error: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Error reading response body. Error: %v", err)
	}

	var response ResponseSentMessage
	err = json.Unmarshal(body, &response)
	if err != nil {
		return fmt.Errorf("Could not unmarshal the response. Error: %v", err)
	}
	if response.OK {
		logrus.Infof("Request to server was made")
		return nil
	} else {
		return fmt.Errorf("No data received. Error Code: %d, Desc: %s", response.ErrorCode, response.Description)
	}
}

func PostMessage(message string) {
	var b BotInfo
	err := b.InitBotObject()
	if err != nil {
		logrus.Fatalf("%v", err)
	}

	logrus.Infof("Requesting for updates.")
	updates, err := b.GetUpdates()
	if err != nil {
		logrus.Warningf("Could not get updates. Error: %v", err)
		return
	}

	lastUpdate := updates[len(updates)-1]
	chatid := lastUpdate.Message.Chat.ID

	logrus.Infof("Sending message: %s", message)
	err = b.SendMessage(chatid, message)
	if err != nil {
		logrus.Warningf("Failed sending message. Error: %v", err)
	}
}

func getNames() ([]string, error) {
	type NamesList struct {
		Names []string `yaml:"names"`
	}
	nameFileContents, err := ioutil.ReadFile(NAMES_FILE)
	if err != nil {
		return []string{}, fmt.Errorf("Error while reading names file: %v", err)
	}

	var readNames NamesList
	err = yaml.Unmarshal(nameFileContents, &readNames)
	if err != nil {
		return []string{}, fmt.Errorf("Error while unmarshalling yaml: %v", err)
	}

	return readNames.Names, nil
}

func shouldIPingToday(t time.Time) bool {
	noPingDays := strings.Split(NOPINGDAYS, ",")
	for _, day := range noPingDays {
		if t.Weekday().String() == day {
			return false
		}
	}
	return true
}

func PingForLunch() {
	logrus.Debug("Calculating Indian TimeZone")
	indiaTZ, err := time.LoadLocation(IndiaTimeZoneID)
	if err != nil {
		logrus.Fatalf("Error parsing timezone. Error: %v", err)
	}

	// flag that keeps track if ping was done or not
	var pingLunch bool
	var pingFri bool

	logrus.Debug("Starting endless loop")
	for {
		t := time.Now()
		indiaTime := t.In(indiaTZ)

		if !shouldIPingToday(indiaTime) {
			continue
		}

		// ping only when time matches and if ping not done for that day
		if indiaTime.Hour() == HOUR && indiaTime.Minute() == MINUTE && !pingLunch {
			// trigger send message from here:set
			logrus.Debug("Fetching names from the names file.")
			names, err := getNames()
			if err != nil {
				logrus.Errorln(err)
				continue
			}

			var nameTags []string
			for _, name := range names {
				name = fmt.Sprintf("@%s", name)
				nameTags = append(nameTags, name)
			}

			message := fmt.Sprintf("ping for lunch %s", strings.Join(nameTags, " "))
			PostMessage(message)
			pingLunch = true
		} else if indiaTime.Hour() == 1 && pingLunch {
			// reset pingLunch flag at 1 in morning
			pingLunch = false
		}

		// friday snacks
		if indiaTime.Weekday().String() == "Friday" && indiaTime.Hour() == 16 && indiaTime.Minute() == 0 && !pingFri {
		//if indiaTime.Weekday().String() == "Monday" && indiaTime.Hour() == 11 && indiaTime.Minute() == 28 && !pingFri {
			// trigger send message from here:set
			logrus.Debug("Fetching names from the names file.")
			names, err := getNames()
			if err != nil {
				logrus.Errorln(err)
				continue
			}

			var nameTags []string
			for _, name := range names {
				name = fmt.Sprintf("@%s", name)
				nameTags = append(nameTags, name)
			}

			message := fmt.Sprintf("ping for Friday snacks %s", strings.Join(nameTags, " "))
			PostMessage(message)
			pingFri = true
		} else if indiaTime.Hour() == 1 && pingFri {
			pingFri = false
		}
	}
}

func StartBot() {
	logrus.SetLevel(logrus.DebugLevel)

	PingForLunch()
}
