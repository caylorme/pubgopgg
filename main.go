package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
)

var USERNAME = "JSon_"

func main() {
	getPlayerStats(USERNAME)
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

func getPlayerID(username string) (id string, err error) {
	resp, err := http.Get("https://pubg.op.gg/user/" + username + "?server=na")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	root, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	element, ok := getElementById("userNickname", root)
	if !ok {
		return "", fmt.Errorf("Could not find player.")
	}
	for _, a := range element.Attr {
		if a.Key == "data-user_id" {
			id = a.Val
			return
		}
	}
	return "", fmt.Errorf("Could not find player token.")
}

func getPlayerStats(username string) error {
	id, err := getPlayerID(username)
	resp, err := http.Get("https://pubg.op.gg/api/users/" + id + "/ranked-stats?season=2018-01&server=na&queue_size=4&mode=fpp")
	if err != nil {
		return err
	}
	stats, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}
	fmt.Printf("%s", stats)
	return nil
}

//
