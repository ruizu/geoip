package main

import (
	"testing"

	"github.com/ruizu/geoip/maxminddb"
	"github.com/ruizu/geoip/maxminddb/geoip2"
)

func TestOpen(t *testing.T) {
	db, err := geoip2.Open("sample-data/test-data/GeoIP2-City-Test.mmdb", maxminddb.ModeMMap)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	result, err := db.Lookup("81.2.69.142")
	if err != nil {
		t.Fatal(err)
	}

	t.Fatal(result)
}
