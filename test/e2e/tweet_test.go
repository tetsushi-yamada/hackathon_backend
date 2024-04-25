package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestCreateTweetHandler(t *testing.T) {
	type args struct {
		UserID    string
		TweetText string
	}
	type Tweet struct {
		TweetID   string
		UserID    string `json:"user_id"`
		TweetText string `json:"tweet_text"`
	}
	type want struct {
		statusCode int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful case",
			args: args{
				UserID:    "1",
				TweetText: "test_tweet",
			},
			want: want{
				statusCode: http.StatusCreated,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			data := Tweet{
				UserID:    tt.args.UserID,
				TweetText: tt.args.TweetText,
			}
			payloadBytes, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error marshaling data:", err)
				return
			}
			body := bytes.NewReader(payloadBytes)

			req, err := http.NewRequest("PUT", "http://localhost:8000/v1/tweets", body)
			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}
			req.Header.Set("Content-Type", "application/json")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			if ok := assert.Equal(t, tt.want.statusCode, resp.StatusCode); !ok {
				t.Fatalf("invalid status code %d", resp.StatusCode)
				return
			}
		})
	}
}

func TestGetTweetsHandler(t *testing.T) {
	type args struct {
		UserID string
	}
	type Tweet struct {
		TweetID   string `json:"tweet_id"`
		UserID    string `json:"user_id"`
		TweetText string `json:"tweet_text"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	type Tweets struct {
		Tweets []Tweet `json:"tweets"`
		Count  int     `json:"count"`
	}
	type want struct {
		Tweets     Tweets
		statusCode int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful case",
			args: args{
				UserID: "2",
			},
			want: want{
				Tweets: Tweets{
					Tweets: []Tweet{
						{
							TweetID:   "3",
							TweetText: "I am Te!",
						},
					},
					Count: 1,
				},
				statusCode: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Now().UTC().Add(time.Hour * 9) // 現在の時刻を取得

			resp, err := http.Get(fmt.Sprintf("http://localhost:8000/v1/tweets?user_id=%s", tt.args.UserID))
			if err != nil {
				t.Fatalf("Error making GET request: %v", err)
			}
			defer resp.Body.Close()

			var tweets Tweets
			if err := json.NewDecoder(resp.Body).Decode(&tweets); err != nil {
				t.Fatalf("Error decoding response body: %v", err)
			}

			for i := 0; i < len(tweets.Tweets); i++ {
				createdAt, err := time.Parse(time.RFC3339, tweets.Tweets[i].CreatedAt)
				if err != nil {
					t.Fatalf("Error parsing CreatedAt: %v", err)
				}
				if now.Before(createdAt) || now.After(createdAt.Add(time.Hour)) {
					t.Errorf("now is not within one hour of CreatedAt: %v %v", now, createdAt)
				}

				assert.Equal(t, tt.want.Tweets.Tweets[i].TweetID, tweets.Tweets[i].TweetID, "TweetID does not match")
				assert.Equal(t, tt.want.Tweets.Tweets[i].TweetText, tweets.Tweets[i].TweetText, "TweetText does not match")
			}
			assert.Equal(t, tt.want.Tweets.Count, tweets.Count, "Count does not match")
			assert.Equal(t, tt.want.statusCode, resp.StatusCode, "Status code does not match")
		})
	}
}

func TestDeleteTweetHandler(t *testing.T) {
	type args struct {
		TweetID string
	}
	type want struct {
		statusCode int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful case",
			args: args{
				TweetID: "2",
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8000/v1/tweets?tweets_id=%s", tt.args.TweetID), nil)
			if err != nil {
				t.Fatalf("Error creating DELETE request: %v", err)
			}

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("Error making DELETE request: %v", err)
			}
			defer resp.Body.Close()

			if ok := assert.Equal(t, tt.want.statusCode, resp.StatusCode); !ok {
				t.Fatalf("invalid status code %d", resp.StatusCode)
			}
		})
	}
}