package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func sendMessage(channelId int, message string, proxy string, token string) error {
	client := &http.Client{}
	if proxy == "" {
		return fmt.Errorf("no proxy provided")
	}

	if !strings.HasPrefix(proxy, "http://") && !strings.HasPrefix(proxy, "https://") {
		proxy = "http://" + proxy
	}

	proxyURL, err := url.Parse(proxy)
	if err != nil {
		return fmt.Errorf("invalid proxy URL: %v", err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client = &http.Client{Transport: transport}
	data := strings.NewReader(`{"chatroom_id": ` + fmt.Sprintf("%d", channelId) + `, "message": "` + message + `"}`)

	// Old tokens thats why i use old api
	req, err := http.NewRequest("POST", "https://kick.com/api/v1/chat-messages", data)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("accept", "application/json")
	req.Header.Set("accept-encoding", "gzip")
	req.Header.Set("accept-language", "pl_PL")
	req.Header.Set("authorization", "Bearer "+token)
	req.Header.Set("connection", "Keep-Alive")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("host", "kick.com")
	req.Header.Set("user-agent", "okhttp/4.9.2")

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %v", err)
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)

	if err != nil {
		return fmt.Errorf("failed to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server error (status %d): %s", resp.StatusCode, bodyText)
	}
	return nil
}
