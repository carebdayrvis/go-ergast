package ergast

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSeasonSchedule(t *testing.T) {

	// read e
	b, err := ioutil.ReadFile("./testdata/schedule.xml")
	assert.Nil(t, err)

	// parse xml
	r := mrdata{}
	err = xml.Unmarshal(b, &r)
	assert.Nil(t, err)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(b)
	}))

	host = ts.URL

	a, err := SeasonSchedule(2018)
	assert.Nil(t, err)

	assert.Equal(t, r.Races, a)
}
