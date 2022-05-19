package metaweather

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Weather struct {
	Woeid       string `json:"woeid"`
	Description string `json:"WeatherDescription"`
	Temperature int    `json:"DegreeCelsiusTemperature"`
}

// GetWoeidByCityName retrieves Metaweather woeid
// (Where On Earth Identifier) for a given cityName.
func GetWoeidByCityName(cityName string) (int, error) {
	var woeid int

	url := buildWoidQuery(cityName)

	res, err := makeHTTPRequest(url)
	if err != nil {
		return woeid, err
	}

	yr, err := unmarshalWoeidResult(res)
	if err != nil {
		return woeid, err
	}

	for _, w := range yr {
		if cityName == w.Title {
			woeid = w.Woeid
			break
		}
	}

	return woeid, nil
}

// GetWeatherByWoeid get current Weather from Yahoo
// for a given woeid.
func GetWeatherByWoeid(woeid string) (Weather, error) {
	cityID, err := strconv.Atoi(woeid)
	if err != nil {
		return Weather{}, err
	}

	url := buildWeatherQuery(cityID)

	res, err := makeHTTPRequest(url)
	if err != nil {
		return Weather{}, err
	}

	yr, err := unmarshalWeatherResult(res)
	if err != nil {
		return Weather{}, err
	}

	var description string
	var temperature int

	if len(yr.Weathers) > 0 {
		description = yr.Weathers[0].Description
		temperature = int(yr.Weathers[0].Temperature)
	}

	return Weather{
		woeid,
		description,
		temperature,
	}, nil
}

type woeidQueryResult []struct {
	Title        string `json:"title"`
	Woeid        int    `json:"woeid"`
	LocationType string `json:"location_type"`
}

type weatherQueryResult struct {
	Woeid    int    `json:"woeid"`
	Location string `json:"title"`
	Weathers []struct {
		Description string  `json:"weather_state_name"`
		Temperature float64 `json:"the_temp"`
		Date        string  `json:"applicable_date"`
	} `json:"consolidated_weather"`
}

func buildWoidQuery(city string) string {

	u := url.URL{
		Scheme: "https",
		Host:   "www.metaweather.com",
		Path:   "api/location/search/",
	}

	q := url.Values{}
	q.Set("query", city)
	u.RawQuery = q.Encode()
	return u.String()
}

func buildWeatherQuery(cityID int) string {

	u, err := url.Parse(fmt.Sprintf("https://www.metaweather.com/api/location/%d", cityID))
	if err != nil {
		log.Println("Url could not be parsed certainly because a bad cityID parameter (", cityID, ")")
		return ""
	}

	return u.String()
}

func makeHTTPRequest(url string) ([]byte, error) {

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func unmarshalWoeidResult(data []byte) (woeidQueryResult, error) {

	var result woeidQueryResult
	// To complete
	return result, nil
}

func unmarshalWeatherResult(data []byte) (weatherQueryResult, error) {

	var result weatherQueryResult
	// To complete
	return result, nil
}
