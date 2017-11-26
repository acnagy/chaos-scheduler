package scheduling

import (
	"encoding/json"
	"fmt"
	"github.com/acnagy/chaos-scheduler/threads"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type conditions struct {
	city         string  `json: "city"`
	state        string  `json: "state"`
	country      string  `json: "country_name"`
	wind_gust    float32 `json: "wind_gust_mph"`
	temp         float32 `json: "temp_f"`
	precip_total float32 `json: "precip_today_in"`
	pressure     float32 `json: "pressure_in"`
}

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
		body, err := ioutil.ReadAll(resp.Body)
		log.Printf("[weather] response status for thread id: %d - %s\n", thpool[i].Id, status)
		if err != nil {
			log.Printf("[weather] error retrieving conditions: %s\n", err)
		}

		var cdtn conditions
		err = json.Unmarshal(body, &cdtn)
		if err != nil {
			log.Printf("[weather] error unmarshalling conditions for %s: %s\n", lat_long, err)
		}

		priority := (cdtn.temp / cdtn.pressure) * (cdtn.wind_gust + cdtn.precip_total)
		thpool[i].Priority = uint16(priority)
	}

	return thpool

}
