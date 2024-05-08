package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestCreateGoodHandler(t *testing.T) {
	type args struct {
		TweetID string
		UserID  string
	}
	type Good struct {
		TweetID string `json:"tweet_id"`
		UserID  string `json:"user_id"`
	}
	type want struct {
		code int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful case",
			args: args{
				TweetID: "1",
				UserID:  "1",
			},
			want: want{
				code: http.StatusCreated,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			data := Good{
				TweetID: tt.args.TweetID,
				UserID:  tt.args.UserID,
			}
			payloadBytes, err := json.Marshal(data)
			if err != nil {
				t.Fatalf("failed to marshal data: %v", err)
			}
			reader := bytes.NewReader(payloadBytes)
			resp, err := http.Post("http://localhost:8000/v1/goods", "application/json", reader)
			if err != nil {
				t.Fatalf("failed to send request: %v", err)
			}
			if ok := assert.Equal(t, tt.want.code, resp.StatusCode); !ok {
				t.Fatalf("invalid status code %d", resp.StatusCode)
			}
		})
	}
}

func TestGetGoodHandler(t *testing.T) {
	type args struct {
		TweetID string
		UserID  string
	}
	type Good struct {
		TweetID string `json:"tweet_id"`
		UserID  string `json:"user_id"`
	}
	type Goods struct {
		Goods []Good `json:"goods"`
		Count int    `json:"count"`
	}
	type want struct {
		goods Goods
		code  int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful TweetID case",
			args: args{
				TweetID: "2",
				UserID:  "",
			},
			want: want{
				goods: Goods{
					Goods: []Good{
						{
							TweetID: "2",
							UserID:  "1000",
						},
					},
					Count: 1,
				},
				code: http.StatusOK,
			},
		},
		{
			name: "successful UserID case",
			args: args{
				TweetID: "",
				UserID:  "1000",
			},
			want: want{
				goods: Goods{
					Goods: []Good{
						{
							TweetID: "2",
							UserID:  "1000",
						},
						{
							TweetID: "3",
							UserID:  "1000",
						},
					},
					Count: 2,
				},
				code: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.TweetID == "" {
				resp, err := http.Get(fmt.Sprintf("http://localhost:8000/v1/goods?user_id=%s", tt.args.UserID))
				if err != nil {
					t.Fatalf("failed to send request: %v", err)
				}
				if ok := assert.Equal(t, tt.want.code, resp.StatusCode); !ok {
					t.Fatalf("invalid status code %d", resp.StatusCode)
				}
				var goods Goods
				if err := json.NewDecoder(resp.Body).Decode(&goods); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				for i := 0; i < tt.want.goods.Count; i++ {
					assert.Equal(t, tt.want.goods.Goods[i].TweetID, goods.Goods[i].TweetID, "TweetID does not match")
					assert.Equal(t, tt.want.goods.Goods[i].UserID, goods.Goods[i].UserID, "UserID does not match")
				}
			} else if tt.args.UserID == "" {
				resp, err := http.Get("http://localhost:8000/v1/goods?tweet_id=" + tt.args.TweetID)
				if err != nil {
					t.Fatalf("failed to send request: %v", err)
				}
				if ok := assert.Equal(t, tt.want.code, resp.StatusCode); !ok {
					t.Fatalf("invalid status code %d", resp.StatusCode)
				}
				var goods Goods
				if err := json.NewDecoder(resp.Body).Decode(&goods); err != nil {
					t.Fatalf("failed to decode response: %v", err)
				}
				for i := 0; i < tt.want.goods.Count; i++ {
					assert.Equal(t, tt.want.goods.Goods[i].TweetID, goods.Goods[i].TweetID, "TweetID does not match")
					assert.Equal(t, tt.want.goods.Goods[i].UserID, goods.Goods[i].UserID, "UserID does not match")
				}
			} else {
				t.Fatalf("invalid args")
			}
		})
	}
}

func TestDeleteGoodHandler(t *testing.T) {
	type args struct {
		TweetID string
		UserID  string
	}
	type Good struct {
		TweetID string `json:"tweet_id"`
		UserID  string `json:"user_id"`
	}
	type want struct {
		code int
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful case",
			args: args{
				TweetID: "1",
				UserID:  "1",
			},
			want: want{
				code: http.StatusNoContent,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", "http://localhost:8000/v1/goods/"+tt.args.TweetID+"/"+tt.args.UserID, nil)
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("failed to send request: %v", err)
			}
			if ok := assert.Equal(t, tt.want.code, resp.StatusCode); !ok {
				t.Fatalf("invalid status code %d", resp.StatusCode)
			}
		})
	}
}
