package pubgopgg

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"encoding/json"
)

type Player struct {
	Username	string
	ID			string
	Region		string
	Mode		string
	Season		string
	Stats		Stats
	Ranks		Ranks
}
type Stats struct {
	Rating				int
	Matches_cnt			int
	Win_matches_cnt		int
	Topten_matches_cnt	int
	Kills_sum			int
	Kills_max			int
	Assists_sum			int
	Headshot_kills_sum	int
	Deaths_sum			int
	Longest_kill_max	int
	Rank_avg			float64
	Damage_dealt_avg	float64
	Time_survived_avg	float64
}

type Ranks struct {
	Rating int
}

func (c *Client) GetPlayer(username string, region string, mode string, season string) (player *Player, err error) {
	player = &Player{}
	resp, err := c.Get(API_ROOT+"/user/"+username+"?server="+region)
	if err != nil {
		return player, fmt.Errorf("GetPlayer:: failed to request Player page", err)
	}
	defer resp.Body.Close()
	root, err := html.Parse(resp.Body)
	if err != nil  {
		return player, fmt.Errorf("GetPlayer:: failed to parse html on Player page", err)
	}
	element, ok := getElementById("userNickname", root)
	if !ok {
		return player, fmt.Errorf("GetPlayer:: Could not locate player ID on Player page", err)
	}
	for _, a := range element.Attr {
		if a.Key == "data-user_id" {
			player.ID = a.Val
		} else if a.Key == "data-user_nickname" {
			player.Username = a.Val
		}
	}

	path := "/api/users/" + player.ID + "/ranked-stats?season="+season+"&server="+region+"&queue_size=4&mode="+mode	
	resp, err = c.Get(API_ROOT+path)
	if err !=  nil {
		return player, fmt.Errorf("GetPlayer:: failed to get stats for Player", err)
	}
	stats, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return player, fmt.Errorf("GetPlayer:: failed to read stats for Player", err)
	}

	err = json.Unmarshal(stats, &player)
	if err != nil {
		return player, fmt.Errorf("GetPlayer:: failed to parse stats of Player", err)
	}
	player.Region = region
	player.Mode = mode
	player.Season = season

	return player, nil
}

func getElementById(id string, n *html.Node) (element *html.Node, ok bool) {
	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			return n, true
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if element, ok = getElementById(id, c); ok {
			return
		}
	}
	return
}
