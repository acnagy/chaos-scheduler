package scheduling

import (
	"encoding/json"
	"github.com/acnagy/chaos-scheduler/threads"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var lat float64
var long float64

func weather_priorities(thpool []threads.Thread) []threads.Thread {

	// Boston
	lat = 42.3601
	long = -71.0589

	lat_long := strconv.FormatFloat(lat, 'f', -1, 64) + "," + strconv.FormatFloat(long, 'f', -1, 64)
	url := "http://api.wunderground.com/api/" + os.Getenv("WUNDERGROUND_KEY") + "/geolookup/conditions/q/" + lat_long + ".json"
	log.Printf("[weather] Request URL: %s", url)

	for i := 0; i < len(thpool); i++ {

		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		status := resp.Status
		log.Printf("[weather] response status for thread id: %d - %s\n", thpool[i].Id, status)
		if err != nil {
			log.Printf("[weather] error retrieving conditions: %s\n", err)
		}

		cdtn, err := ioutil.ReadAll(resp.Body)

		type Conditions struct {
			City         string  `json:"city"`
			State        string  `json:"state"`
			Country      string  `json:"country_name"`
			Wind_gust    string  `json:"wind_gust_mph"`
			Temp         float64 `json:"temp_f"`
			Precip_total string  `json:"precip_today_in"`
			Pressure     string  `json:"pressure_in"`
			Station      string  `json:"station_id"`
		}

		type currentObservation struct {
			Data Conditions `json:"current_observation"`
		}

		var current currentObservation
		if err := json.Unmarshal(cdtn, &current); err != nil {
			log.Printf("[weather] error unmarshalling conditions for %s: %s\n", lat_long, err)
		}

		temp := current.Data.Temp
		pressure := current.Data.Pressure
		gust := current.Data.Wind_gust
		precip := current.Data.Precip_total

		prs, _ := strconv.ParseFloat(pressure, 64)
		g, _ := strconv.ParseFloat(gust, 64)
		prc, _ := strconv.ParseFloat(precip, 64)

		priority := (temp / prs) * (g + prc)
		thpool[i].Priority = uint16(priority)
	}

	return thpool

}
