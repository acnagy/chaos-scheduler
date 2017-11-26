package scheduling

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func Retrieve() {

	type geolocation struct {
		city    string `json: "city"`
		state   string `json: "state"`
		country string `json: "country_name"`
	}
	var lat float64
	var long float64

	lat = 45.123
	long = 70.00

	lat_long := strconv.FormatFloat(lat, 'f', -1, 64) + strconv.FormatFloat(long, 'f', -1, 64)
	url := "http://api.wunderground.com/api/" + os.Getenv("WUNDERGROUND_KEY") + "/geolookup/conditions/q/" + lat_long + ".json"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	status := resp.Status
	conditions, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s - %s", status, string(conditions))

}
