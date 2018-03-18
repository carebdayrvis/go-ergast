package ergast

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

var host string = "https://ergast.com/api/f1"

type mrdata struct {
	XMLName xml.Name `xml:"MRData"`
	Races   []Race   `xml:"RaceTable>Race"`
}

type Race struct {
	Circuit           Circuit
	Date              ErgastDate
	Time              ErgastTime
	RaceName          string
	Season            int                `xml:"season,attr"`
	Round             int                `xml:"round,attr"`
	Results           []Result           `xml:"ResultsList>Result"`
	QualifyingResults []QualifyingResult `xml:"QualifyingList>QualifyingResult"`
}

type QualifyingResult struct {
	Driver      Driver
	Constructor Constructor
	Q1          ErgastDuration
	Q2          ErgastDuration
	Q3          ErgastDuration
	Position    int `xml:"position,attr"`
}

type Circuit struct {
	CircuitName string
}

type Result struct {
	Constructor Constructor
	Driver      Driver
	Laps        int
	Grid        int
	StatusID    int `xml:"statusId"`
	Status      string
	FastestLap  Lap
	Number      int `xml:"number,attr"`
	Position    int `xml:"position,attr"`
	Points      int `xml:"points,attr"`
}

type Lap struct {
	Time ErgastDuration
	Rank int `xml:"rank,attr"`
	Lap  int `xml:"lap,attr"`

	// There seems to be a bug with https://ergast.com/api/f1/2017/20/results which makes the averagespeed field match the time field.
	// Leaving this field unimplemented for now. A bug report has been submitted to the ergast API
	//AverageSpeed float64
	//AverageSpeedUnits string `xml:"units,attr"`
}

type Driver struct {
	DriverID        string `xml:"driverId,attr"`
	Code            string `xml:"code,attr"`
	PermanentNumber int
	GivenName       string
	FamilyName      string
	Nationality     string
	DateOfBirth     ErgastDate
}

type Constructor struct {
	ConstructorID string `xml:"constructorId,attr"`
	Name          string
	Nationality   string
}

// Latest returns the results from the latest round in the current season: https://ergast.com/api/f1/current/last/results
func Latest() (Race, error) {
	const path string = "/current/last/results"

	resp, err := http.DefaultClient.Get(host + path)
	if err != nil {
		return Race{}, err
	}

	defer resp.Body.Close()

	d := mrdata{}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Race{}, err
	}

	err = xml.Unmarshal(b, &d)
	if err != nil {
		return Race{}, err
	}

	return d.Races[0], nil
}

// SpecificResult returns the results for a specific round in a specific season, e.g.: https://ergast.com/api/f1/2017/1/results
func SpecificResult(season int, round int) (Race, error) {
	var path string = fmt.Sprintf("/%v/%v/results", season, round)

	resp, err := http.DefaultClient.Get(host + path)
	if err != nil {
		return Race{}, err
	}

	defer resp.Body.Close()

	d := mrdata{}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Race{}, err
	}

	err = xml.Unmarshal(b, &d)
	if err != nil {
		return Race{}, err
	}

	return d.Races[0], nil
}

func SpecificQualifying(season int, round int) (Race, error) {
	var path string = fmt.Sprintf("/%v/%v/qualifying", season, round)

	resp, err := http.DefaultClient.Get(host + path)
	if err != nil {
		return Race{}, err
	}

	defer resp.Body.Close()

	d := mrdata{}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Race{}, err
	}

	err = xml.Unmarshal(b, &d)
	if err != nil {
		return Race{}, err
	}

	return d.Races[0], nil

}
