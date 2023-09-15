package steam

import (
	"encoding/json"
	"net/http"
)

type App struct {
	AppID int    `json:"appid"`
	Name  string `json:"name"`
}

type AppList struct {
	Apps []*App `json:"apps"`
}

type GetAppListResponse struct {
	AppList *AppList `json:"applist"`
}

type Client struct {
}

func (c *Client) RequestGetAppList() (*GetAppListResponse, error) {
	resp, err := http.Get("https://api.steampowered.com/ISteamApps/GetAppList/v2/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	p := &GetAppListResponse{}
	err = json.NewDecoder(resp.Body).Decode(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
