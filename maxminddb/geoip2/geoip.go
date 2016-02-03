package geoip2

import (
	"fmt"

	"github.com/ruizu/geoip/maxminddb"
)

type City struct {
	City               TypeCity
	Continent          TypeContinent
	Country            TypeCountry
	Location           TypeLocation
	RegisteredCountry  TypeCountry
	RepresentedCountry TypeCountry
}

type Country struct {
	Continent          TypeContinent
	Country            TypeCountry
	RegisteredCountry  TypeCountry
	RepresentedCountry TypeCountry
}

type TypeCity struct {
	GeoNameID int64
	Names     map[string]string
}

type TypeContinent struct {
	GeoNameID int64
	Code      string
	Names     map[string]string
}

type TypeCountry struct {
	GeoNameID int64
	ISOCode   string
	Names     map[string]string
}

type TypeLocation struct {
	Latitude  float64
	Longitude float64
	TimeZone  string
}

type TypeSubdivision struct {
	GeoNameID int64
	ISOCode   string
	Names     map[string]string
}

type CityDB struct {
	maxminddb.DB
}

func OpenCityDB(f string) (*CityDB, error) {
	println(maxminddb.Version())
	db, err := maxminddb.Open(f, maxminddb.ModeMMap)
	if err != nil {
		return nil, err
	}
	return &CityDB{*db}, nil
}

func (db *CityDB) Lookup(ip string) (City, error) {
	result, err := db.DB.Lookup(ip)
	if err != nil {
		return City{}, err
	}

	result.Dump()
	result.Free()

	return City{}, nil
}
