package steam

import (
	"net/url"
	"testing"
)

func Test_buildPublishedFileServiceRequest(t *testing.T) {
	tests := []struct {
		name          string
		steamAPIKey   string
		workshopIDs   []string
		expectedQuery url.Values
		wantErr       bool
	}{
		{
			name:        "Valid request with multiple workshop items",
			steamAPIKey: "testkey",
			workshopIDs: []string{
				"909647363",
				"1510950729",
				"1516836158",
				"2122265954",
			},
			wantErr: false,
		},
		{
			name:        "Valid request with no workshop items",
			steamAPIKey: "testkey",
			workshopIDs: []string{},
			wantErr:     false,
		},
		{
			name:        "Invalid API key",
			steamAPIKey: "",
			workshopIDs: []string{
				"909647363",
				"1510950729",
				"1516836158",
				"2122265954",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildPublishedFileServiceRequest(tt.steamAPIKey, tt.workshopIDs)
			if (err != nil) != tt.wantErr {
				t.Errorf("buildPublishedFileServiceRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if got.Scheme != "https" {
					t.Errorf("Expected https scheme, got %s", got.Scheme)
				}

				if got.Host != "api.steampowered.com" {
					t.Errorf("Expected host api.steampowered.com, got %s", got.Host)
				}

				if got.Path != "/IPublishedFileService/GetDetails/v1/" {
					t.Errorf("Expected path /IPublishedFileService/GetDetails/v1/, got %s", got.Path)
				}

				query := got.Query()
				if query.Get("key") != tt.steamAPIKey {
					t.Errorf("Expected key %s, got %s", tt.steamAPIKey, query.Get("key"))
				}

				if query.Get("appid") != ZomboidSteamID {
					t.Errorf("Expected appid %s, got %s", ZomboidSteamID, query.Get("appid"))
				}
			}
		})
	}
}
