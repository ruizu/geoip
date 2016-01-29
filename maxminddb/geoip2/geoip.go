package geoip2

import (
	"github.com/ruizu/geoip/maxminddb"
)

type CityDB struct {
	City               CityType
	Continent          ContinentType
	Country            CountryType
	Location           LocationType
	RegisteredCountry  CountryType
	RepresentedCountry CountryType
}

type CountryDB struct {
	Continent          ContinentType
	Country            CountryType
	RegisteredCountry  CountryType
	RepresentedCountry CountryType
}

type CityType struct {
	GeoNameID int64
	Names     map[string]string
}

type ContinentType struct {
	GeoNameID int64
	Code      string
	Names     map[string]string
}

type CountryType struct {
	GeoNameID int64
	ISOCode   string
	Names     map[string]string
}

type LocationType struct {
	Latitude  float64
	Longitude float64
	TimeZone  string
}

type SubdivisionType struct {
	GeoNameID int64
	ISOCode   string
	Names     map[string]string
}

type DB struct {
	maxminddb.DB
}

func Open(f string, m uint32) (*DB, error) {
	db, err := maxminddb.Open(f, m)
	if err != nil {
		return nil, err
	}
	return &DB{*db}, nil
}

func (db *DB) CityLookup() {

}

func (db *DB) CountryLookup() {

}
