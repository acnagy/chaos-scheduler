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

func weather_priorities(policy string, thpool []threads.Thread) []threads.Thread {

	var lat_long string

	for i := 0; i < len(thpool); i++ {

		switch policy {
		case "weather - static":
			lat_long = static_lat_long()
		case "weather - variable":
			lat_long = variable_lat_long(thpool[i].Id)

		}

		url := "http://api.wunderground.com/api/" + os.Getenv("WUNDERGROUND_KEY") + "/geolookup/conditions/q/" + lat_long + ".json"
		log.Printf("[%s] - Request URL: %s", policy, url)

		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		status := resp.Status
		log.Printf("[%s] - response status for thread id: %d - %s\n", policy, thpool[i].Id, status)
		if err != nil {
			log.Printf("[%s] - error retrieving conditions: %s\n", policy, err)
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
			log.Printf("[%s] - error unmarshalling conditions for %s: %s\n", policy, lat_long, err)
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

	log.Printf("[%s] thread batch prioritized\n")
	return thpool
}

func static_lat_long() string {
	// Boston
	lat := 42.3601
	long := -71.0589

	return concatLatLong(lat, long)
}

func variable_lat_long(threadId uint16) string {

	division_factor := 5.1 // magic number fo' fun
	var mask_upper uint16 = 0xFF00
	var mask_lower uint16 = 0x00FF

	lat := float64(threadId&mask_upper) / division_factor
	long := float64(threadId&mask_lower) / division_factor

	return concatLatLong(lat, long)
}

func concatLatLong(lat float64, long float64) string {
	return strconv.FormatFloat(lat, 'f', -1, 64) + "," + strconv.FormatFloat(long, 'f', -1, 64)
}
