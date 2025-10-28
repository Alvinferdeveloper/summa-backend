package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type Coords struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type NominatimResponse struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

const cacheTTL = 720 * time.Hour

func GetCoordsForLocation(city, country string) (*Coords, error) {
	query := fmt.Sprintf("%s, %s", city, country)
	cacheKey := "geo:" + query

	cachedVal, err := GetCache(cacheKey)
	if err != nil {
		fmt.Printf("Redis Get error: %v\n", err)
	}
	if cachedVal != "" {
		var coords Coords
		if err := json.Unmarshal([]byte(cachedVal), &coords); err == nil {
			return &coords, nil
		}
	}

	coords, err := fetchCoordsFromNominatim(query)
	if err != nil {
		return nil, err
	}

	jsonCoords, err := json.Marshal(coords)
	if err == nil {
		if err := SetCache(cacheKey, jsonCoords, cacheTTL); err != nil {
			fmt.Printf("Redis Set error: %v\n", err)
		}
	}

	return coords, nil
}

func fetchCoordsFromNominatim(query string) (*Coords, error) {
	baseURL := "https://nominatim.openstreetmap.org/search"

	params := url.Values{}
	params.Add("q", query)
	params.Add("format", "json")
	params.Add("limit", "1")

	req, err := http.NewRequest("GET", baseURL+"?"+params.Encode(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Summa-App/1.0 (alvinfer67@gmail.com)")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("nominatim API returned status: %s", resp.Status)
	}

	var results []NominatimResponse
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("no results found for query: %s", query)
	}

	lat, _ := strconv.ParseFloat(results[0].Lat, 64)
	lon, _ := strconv.ParseFloat(results[0].Lon, 64)

	return &Coords{Latitude: lat, Longitude: lon}, nil
}
