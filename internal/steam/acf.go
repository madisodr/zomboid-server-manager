package steam

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

type ACF struct {
	AppID                  string                  `json:"appid"`
	SizeOnDisk             string                  `json:"sizeondisk"`
	NeedsUpdate            string                  `json:"needsupdate"`
	NeedsDownload          string                  `json:"needsdownload"`
	TimeLastUpdated        string                  `json:"timelastupdated"`
	TimeLastAppRan         string                  `json:"timelastappran"`
	LastBuildID            string                  `json:"lastbuildid"`
	WorkshopItemsInstalled map[string]WorkshopItem `json:"workshopitemsinstalled"`
}

func removeSpaceAndTab(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}

func ParseACF(filename string) (*ACF, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var stack int = 0
	acf := ACF{WorkshopItemsInstalled: make(map[string]WorkshopItem)}
	var currentItem WorkshopItem
	var currentKey string

	parsingInstalled := true

	for scanner.Scan() {
		// remove spaces and tabs from the line
		line := removeSpaceAndTab(scanner.Text())

		if line == "{" {
			stack++
		} else if line == "}" {
			stack--
		} else {
			// split by double quotes "" to get key value pairs
			parts := strings.Split(line, "\"\"")

			for i, part := range parts {
				parts[i] = strings.ReplaceAll(part, "\"", "")
			}

			key := parts[0]

			if stack == 1 {
				switch key {
				case "appid":
					acf.AppID = parts[1]
					if acf.AppID != ZomboidSteamID {
						return nil, fmt.Errorf("AppID is not for Project Zomboid (%s), got %s", ZomboidSteamID, acf.AppID)
					}
				case "SizeOnDisk":
					acf.SizeOnDisk = parts[1]
				case "NeedsUpdate":
					acf.NeedsUpdate = parts[1]
				case "NeedsDownload":
					acf.NeedsDownload = parts[1]
				case "TimeLastUpdated":
					acf.TimeLastUpdated = parts[1]
				case "TimeLastAppRan":
					acf.TimeLastAppRan = parts[1]
				case "LastBuildID":
					acf.LastBuildID = parts[1]
				case "WorkshopItemsInstalled":
					log.Println("WorkshopItemsInstalled")
					parsingInstalled = true
				case "WorkshopItemDetails":
					log.Println("WorkshopItemDetails")
					parsingInstalled = false
				}
			} else if stack == 2 {
				if parsingInstalled {
					currentKey = key
					currentItem = WorkshopItem{}
				}
			} else if stack == 3 {
				if parsingInstalled {
					switch key {
					case "size":
						currentItem.Size = parts[1]
					case "timeupdated":
						currentItem.TimeUpdated = parts[1]
					case "manifest":
						currentItem.Manifest = parts[1]
					}
					acf.WorkshopItemsInstalled[currentKey] = currentItem
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	log.Println("Parsed ACF File")
	return &acf, nil
}
