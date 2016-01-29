package main

import (
	"testing"

	"github.com/ruizu/geoip/maxminddb"
)

func TestOpen(t *testing.T) {
	db, err := maxminddb.Open("sample-data/test-data/GeoIP2-City-Test.mmdb", maxminddb.ModeMMap)
	if err != nil {
		t.Errorf("%v", err)
	}
	if db != nil {
		defer db.Close()
	}
}
