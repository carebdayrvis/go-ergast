package ergast

import (
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLatest(t *testing.T) {

	// read e
	b, err := ioutil.ReadFile("./testdata/latest.xml")
	assert.Nil(t, err)

	// parse xml
	r := mrdata{}
	err = xml.Unmarshal(b, &r)
	assert.Nil(t, err)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(b)
	}))

	host = ts.URL

	a, err := Latest()
	assert.Nil(t, err)

	assert.Equal(t, r.RaceTable.Race, a)
}

func TestSpecificResult(t *testing.T) {
	// read e
	b, err := ioutil.ReadFile("./testdata/specificresult.xml")
	assert.Nil(t, err)

	// parse xml
	r := mrdata{}
	err = xml.Unmarshal(b, &r)
	assert.Nil(t, err)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(b)
	}))

	host = ts.URL

	a, err := SpecificResult(2017, 1)
	assert.Nil(t, err)

	assert.Equal(t, r.RaceTable.Race, a)
}
