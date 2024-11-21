package steam

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

const (
	ZomboidSteamID       = "108600"
	api_url              = "https://api.steampowered.com/IPublishedFileService/GetDetails/v1/"
	publishedFieldFormat = "publishedfileids%5B%d%5D=%s"
	flags                = "&includetags=true&short_description=true&appid=%s"
)

func CheckForUpdates(steamAPIKey string, acfFile string) int {
	acf, err := ParseACF(acfFile)
	if err != nil {
		panic(err)
	}

	workshopMods, err := GetModInfoFromWorkshop(steamAPIKey, acf.WorkshopItemsInstalled)
	if err != nil {
		panic(err)
	}

	count := 0
	for _, modInfo := range workshopMods.Response.PublishedFileDetails {
		localModInfo := acf.WorkshopItemsInstalled[modInfo.PublishedFileID]

		localTimeUpdated, err := strconv.ParseInt(localModInfo.TimeUpdated, 10, 64)
		if err != nil {
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

// GetModInfoFromWorkshop
func GetModInfoFromWorkshop(steamAPIKey string, workshops map[string]WorkshopItem) (*SteamApiResponse, error) {
	client := &http.Client{}

	// build the full url to be sent to steam's API. for each published file, add a new query parameter
	workshopItemsParam := ""
	pos := 0
	for workshopItemID := range workshops {
		workshopItemsParam += "&publishedfileids%5B" + strconv.Itoa(pos) + "%5D=" + workshopItemID
		pos++
	}

	url := fmt.Sprintf("%s?key=%s%s&includetags=true&short_description=true&appid=%s", api_url, steamAPIKey, workshopItemsParam, ZomboidSteamID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var apiResponse SteamApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	return &apiResponse, nil
}
