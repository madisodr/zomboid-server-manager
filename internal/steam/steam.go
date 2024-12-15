package steam

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

const (
	ZomboidSteamID = "108600"
)

// CheckForUpdates fetches the mod info from the worshop that are declared in the configuration file
// and compares the result with the local acf file, checking whether the mod's last updated time is
// newer than the local mod's last updated time.
func CheckForUpdates(steamAPIKey string, workshopItems []string, acfFile string) int {
	workshopMods, err := GetModInfoFromWorkshop(steamAPIKey, workshopItems)
	if err != nil {
		panic(err)
	}

	acf, err := ParseACF(acfFile)
	if err != nil {
		panic(err)
	}

	count := 0
	for _, modInfo := range workshopMods.Response.PublishedFileDetails {

		localModInfo := acf.WorkshopItemsInstalled[modInfo.PublishedFileID]
		localTimeUpdated, err := strconv.ParseInt(localModInfo.TimeUpdated, 10, 64)
		if err != nil {
			log.Println("Checking mod: \n", modInfo)
			log.Printf("Error parsing mod last updated time: %v\n", err)
			continue
		}

		if modInfo.TimeUpdated > localTimeUpdated {
			log.Printf("Title:      %s\n", modInfo.Title)
			log.Printf("Update available. %s\n\n", modInfo.PublishedFileID)
			count++
		}
	}

	return count
}

// buildPublishedFileServiceRequest formats the full url to be sent to steam's API.
// Each published file needs its own query parameter in the form of publishedfileids[i]=1234567890
func buildPublishedFileServiceRequest(steamAPIKey string, workshops []string) (*url.URL, error) {
	apiUrl := "https://api.steampowered.com/IPublishedFileService/GetDetails/v1/"

	baseUrl, err := url.Parse(apiUrl)
	if err != nil {
		return nil, err
	}

	if steamAPIKey == "" {
		return nil, fmt.Errorf("steamAPIKey is required")
	}

	params := url.Values{}
	params.Add("key", steamAPIKey)
	params.Add("appid", ZomboidSteamID)
	params.Add("includetags", "true")
	params.Add("short_description", "true")

	for i, key := range workshops {
		params.Add(fmt.Sprintf("publishedfileids[%d]", i), key)
	}

	baseUrl.RawQuery = params.Encode()
	return baseUrl, nil
}

// GetModInfoFromWorkshop
func GetModInfoFromWorkshop(steamAPIKey string, workshopItems []string) (*SteamApiResponse, error) {
	url, err := buildPublishedFileServiceRequest(steamAPIKey, workshopItems)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var parsedApiResponse SteamApiResponse
	err = json.Unmarshal(body, &parsedApiResponse)
	if err != nil {
		return nil, err
	}

	return &parsedApiResponse, nil
}
