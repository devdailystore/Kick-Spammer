package utils

import (
	"encoding/json"
	"net/http"
)

type ChatroomResponse struct {
	Id            int `json:"id"`
	FollowersMode struct {
		Enabled bool `json:"enabled"`
	} `json:"followers_mode"`
}

type UserSpamInfo struct {
	UserId  int  `json:"user_id"`
	CanSpam bool `json:"can_spam"`
}

func GetUserId(channelName string) UserSpamInfo {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://kick.com/api/v2/channels/"+channelName+"/chatroom", nil)
	if err != nil {
		return UserSpamInfo{UserId: -1, CanSpam: false}
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return UserSpamInfo{UserId: -1, CanSpam: false}
	}
	defer resp.Body.Close()

	var chatroom ChatroomResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatroom); err != nil {
		return UserSpamInfo{UserId: -1, CanSpam: false}
	}

	canSpam := !chatroom.FollowersMode.Enabled

	return UserSpamInfo{
		UserId:  chatroom.Id,
		CanSpam: canSpam,
	}
}
