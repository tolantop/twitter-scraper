package twitterscraper

import (
	"context"
	"testing"
)

func TestGetTweets(t *testing.T) {
	count := 0
	maxTweetsNbr := 300
	dupcheck := make(map[string]bool)
	for tweet := range GetTweets(context.Background(), "Twitter", maxTweetsNbr) {
		if tweet.Error != nil {
			t.Error(tweet.Error)
		} else {
			count++
			if tweet.ID == "" {
				t.Error("Expected tweet ID is not empty")
			} else {
				if dupcheck[tweet.ID] {
					t.Errorf("Detect duplicated tweet ID: %s", tweet.ID)
				} else {
					dupcheck[tweet.ID] = true
				}
			}
			if tweet.UserID == "" {
				t.Error("Expected tweet UserID is not empty")
			}
			if tweet.Username == "" {
				t.Error("Expected tweet Username is not empty")
			}
			if tweet.PermanentURL == "" {
				t.Error("Expected tweet PermanentURL is not empty")
			}
			if tweet.Text == "" {
				t.Error("Expected tweet Text is not empty")
			}
			if tweet.TimeParsed.IsZero() {
				t.Error("Expected tweet TimeParsed is not zero")
			}
			if tweet.Timestamp == 0 {
				t.Error("Expected tweet Timestamp is greater than zero")
			}
			for _, video := range tweet.Videos {
				if video.ID == "" {
					t.Error("Expected tweet video ID is not empty")
				}
				if video.Preview == "" {
					t.Error("Expected tweet video Preview is not empty")
				}
				if video.URL == "" {
					t.Error("Expected tweet video URL is not empty")
				}
			}
		}
	}
	if count != maxTweetsNbr {
		t.Errorf("Expected tweets count=%v, got: %v", maxTweetsNbr, count)
	}
}
