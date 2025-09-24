package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Actor struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

type Repo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Event struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	Actor     Actor           `json:"actor"`
	Repo      Repo            `json:"repo"`
	Payload   json.RawMessage `json:"payload"` // <-- fleksibel
	Public    bool            `json:"public"`
	CreatedAt string          `json:"created_at"`
}

type WatchPayload struct {
	Action string `json:"action"`
}

// payload khusus PushEvent
type PushPayload struct {
	PushID       int64 `json:"push_id"`
	Size         int   `json:"size"`
	DistinctSize int   `json:"distinct_size"`
	Ref          string `json:"ref"`
	Head         string `json:"head"`
	Before       string `json:"before"`
	Commits      []struct {
		Sha     string `json:"sha"`
		Message string `json:"message"`
	} `json:"commits"`
}

func GetActivity(username string) {
	resp, err := http.Get("https://api.github.com/users/" + username + "/events?per_page=10")
	if err != nil {
		panic("Error:" + err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic("Error:" + err.Error())
	}

	var events []Event
	err = json.Unmarshal(body, &events)
	if err != nil {
		panic("Error:" + err.Error())
	}

	fmt.Println("Output:")
	for _, e := range events {
		if e.Type == "WatchEvent"  {
			var  payload WatchPayload
			err = json.Unmarshal(e.Payload, &payload)
			if err != nil {
				panic("Error:" + err.Error())
			}
			fmt.Printf("%s %s %s\n", e.Actor.Login, payload.Action, e.Repo.Name)
		}
		if e.Type == "PushEvent" {
			var payload PushPayload
			err = json.Unmarshal(e.Payload, &payload)
			if err != nil {
				panic("Error:" + err.Error())
			}
			fmt.Printf("%s pushed %d commits to %s\n", e.Actor.Login, payload.DistinctSize, e.Repo.Name)
		}
	}
}