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

func stringPointer(s string) *string {
	return &s
}

func TestCreateTweetHandler(t *testing.T) {
	type args struct {
		UserID    string
		TweetText string
		ParentID  *string
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
		{
			name: "successful case with parent",
			args: args{
				UserID:    "1",
				TweetText: "test_reply",
				ParentID:  stringPointer("1"),
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

			req, err := http.NewRequest("POST", "http://localhost:8000/v1/tweets", body)
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

			resp, err := http.Get(fmt.Sprintf("http://localhost:8000/v1/tweets/%s", tt.args.UserID))
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

func TestGetTweetByTweetIDHandler(t *testing.T) {
	type args struct {
		TweetID string
	}
	type Tweet struct {
		TweetID   string `json:"tweet_id"`
		UserID    string `json:"user_id"`
		TweetText string `json:"tweet_text"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	type want struct {
		Tweet      Tweet
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
				TweetID: "7",
			},
			want: want{
				Tweet: Tweet{
					TweetID:   "7",
					UserID:    "3",
					TweetText: "TEST TWEET",
				},
				statusCode: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf("http://localhost:8000/v1/tweets/by-tweet/%s", tt.args.TweetID))
			if err != nil {
				t.Fatalf("Error making GET request: %v", err)
			}
			defer resp.Body.Close()

			var tweet Tweet
			if err := json.NewDecoder(resp.Body).Decode(&tweet); err != nil {
				t.Fatalf("Error decoding response body: %v", err)
			}

			assert.Equal(t, tt.want.Tweet.TweetID, tweet.TweetID, "TweetID does not match")
			assert.Equal(t, tt.want.Tweet.UserID, tweet.UserID, "UserID does not match")
			assert.Equal(t, tt.want.Tweet.TweetText, tweet.TweetText, "TweetText does not match")
			assert.Equal(t, tt.want.statusCode, resp.StatusCode, "Status code does not match")
		})
	}

}

func TestUpdateTweetHandler(t *testing.T) {
	type args struct {
		TweetID   string
		TweetText string
	}
	type want struct {
		statusCode int
	}
	type Tweet struct {
		TweetID   string
		UserID    string `json:"user_id"`
		TweetText string `json:"tweet_text"`
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful case",
			args: args{
				TweetID:   "4",
				TweetText: "updated_tweet",
			},
			want: want{
				statusCode: http.StatusNoContent,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := Tweet{
				TweetText: tt.args.TweetText,
			}
			payloadBytes, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error marshaling data:", err)
				return
			}
			body := bytes.NewReader(payloadBytes)

			req, err := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8000/v1/tweets/by-tweet/%s", tt.args.TweetID), body)
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
			}
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
				statusCode: http.StatusNoContent,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8000/v1/tweets/by-tweet/%s", tt.args.TweetID), nil)
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

func TestSearchTweetHandler(t *testing.T) {
	type args struct {
		SearchWord string
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
				SearchWord: "search",
			},
			want: want{
				Tweets: Tweets{
					Tweets: []Tweet{
						{
							TweetID:   "5",
							TweetText: "Will be searched!",
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
			resp, err := http.Get(fmt.Sprintf("http://localhost:8000/v1/tweets/search/%s", tt.args.SearchWord))
			if err != nil {
				t.Fatalf("Error making GET request: %v", err)
			}
			defer resp.Body.Close()

			var tweets Tweets
			if err := json.NewDecoder(resp.Body).Decode(&tweets); err != nil {
				t.Fatalf("Error decoding response body: %v", err)
			}

			for i := 0; i < len(tweets.Tweets); i++ {
				assert.Equal(t, tt.want.Tweets.Tweets[i].TweetID, tweets.Tweets[i].TweetID, "TweetID does not match")
				assert.Equal(t, tt.want.Tweets.Tweets[i].TweetText, tweets.Tweets[i].TweetText, "TweetText does not match")
			}
			assert.Equal(t, tt.want.Tweets.Count, tweets.Count, "Count does not match")
			assert.Equal(t, tt.want.statusCode, resp.StatusCode, "Status code does not match")
		})
	}
}

func TestGetRepliesHandler(t *testing.T) {
	type args struct {
		TweetID string
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
				TweetID: "5",
			},
			want: want{
				Tweets: Tweets{
					Tweets: []Tweet{
						{
							TweetID:   "6",
							TweetText: "REPLY",
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
			resp, err := http.Get(fmt.Sprintf("http://localhost:8000/v1/tweets/reply/%s", tt.args.TweetID))
			if err != nil {
				t.Fatalf("Error making GET request: %v", err)
			}
			defer resp.Body.Close()

			var tweets Tweets
			if err := json.NewDecoder(resp.Body).Decode(&tweets); err != nil {
				t.Fatalf("Error decoding response body: %v", err)
			}

			for i := 0; i < len(tweets.Tweets); i++ {
				assert.Equal(t, tt.want.Tweets.Tweets[i].TweetID, tweets.Tweets[i].TweetID, "TweetID does not match")
				assert.Equal(t, tt.want.Tweets.Tweets[i].TweetText, tweets.Tweets[i].TweetText, "TweetText does not match")
			}
			assert.Equal(t, tt.want.Tweets.Count, tweets.Count, "Count does not match")
			assert.Equal(t, tt.want.statusCode, resp.StatusCode, "Status code does not match")
		})
	}
}
