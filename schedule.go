package ergast

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SeasonSchedule(season int) ([]Race, error) {

	var path string = fmt.Sprintf("/%v", season)

	resp, err := http.DefaultClient.Get(host + path)
	if err != nil {
		return []Race{}, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []Race{}, err
	}

	d := mrdata{}

	err = xml.Unmarshal(b, &d)
	if err != nil {
		return []Race{}, err
	}

	return d.Races, nil

}
