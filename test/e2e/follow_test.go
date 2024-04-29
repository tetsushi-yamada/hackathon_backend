package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestCreateFollowHandler(t *testing.T) {
	type args struct {
		UserID   string
		FollowID string
	}
	type Follow struct {
		UserID   string `json:"user_id"`
		FollowID string `json:"follow_id"`
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
				FollowID: "2",
				UserID:   "1",
			},
			want: want{
				statusCode: http.StatusCreated,
			},
		},
		{
			name: "invalid case",
			args: args{
				FollowID: "1",
				UserID:   "1",
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			data := Follow{
				UserID:   tt.args.UserID,
				FollowID: tt.args.FollowID,
			}
			payloadBytes, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error marshaling data:", err)
				return
			}
			body := bytes.NewReader(payloadBytes)

			resp, err := http.Post("http://localhost:8000/v1/follows", "application/json", body)
			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}
			defer resp.Body.Close()

			if ok := assert.Equal(t, tt.want.statusCode, resp.StatusCode); !ok {
				t.Fatalf("invalid status code %d", resp.StatusCode)
				return
			}
		})
	}
}

func TestGetFollowsHandler(t *testing.T) {
	type args struct {
		UserID string
	}
	type Follow struct {
		UserID    string `json:"user_id"`
		FollowID  string `json:"follow_id"`
		CreatedAt string `json:"created_at"`
	}
	type Follows struct {
		Follows []Follow `json:"follows"`
		Count   int      `json:"count"`
	}
	type want struct {
		Follows    Follows
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
				Follows: Follows{
					Follows: []Follow{
						{
							UserID:   "2",
							FollowID: "100",
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

			resp, err := http.Get(fmt.Sprintf("http://localhost:8000/v1/follows?user_id=%s", tt.args.UserID))
			if err != nil {
				t.Fatalf("Error making GET request: %v", err)
			}
			defer resp.Body.Close()

			var follows Follows
			if err := json.NewDecoder(resp.Body).Decode(&follows); err != nil {
				t.Fatalf("Error decoding response body: %v", err)
			}

			for i := 0; i < len(follows.Follows); i++ {
				createdAt, err := time.Parse(time.RFC3339, follows.Follows[i].CreatedAt)
				if err != nil {
					t.Fatalf("Error parsing CreatedAt: %v", err)
				}
				if now.Before(createdAt) || now.After(createdAt.Add(time.Hour)) {
					t.Errorf("now is not within one hour of CreatedAt: %v %v", now, createdAt)
				}

				assert.Equal(t, tt.want.Follows.Follows[i].UserID, follows.Follows[i].UserID, "UserID does not match")
				assert.Equal(t, tt.want.Follows.Follows[i].FollowID, follows.Follows[i].FollowID, "FollowID does not match")
			}
			assert.Equal(t, tt.want.Follows.Count, follows.Count, "Count does not match")
			assert.Equal(t, tt.want.statusCode, resp.StatusCode, "Status code does not match")
		})
	}
}

func TestDeleteFollowHandler(t *testing.T) {
	type args struct {
		UserID   string
		FollowID string
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
				UserID:   "2",
				FollowID: "100",
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8000/v1/follows?user_id=%s&follow_id=%s", tt.args.UserID, tt.args.FollowID), nil)
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

func TestGetFollowersHandler(t *testing.T) {
	type args struct {
		FollowID string
	}
	type Follow struct {
		UserID    string `json:"user_id"`
		FollowID  string `json:"follow_id"`
		CreatedAt string `json:"created_at"`
	}
	type Follows struct {
		Follows []Follow `json:"follows"`
		Count   int      `json:"count"`
	}
	type want struct {
		Follows    Follows
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
				FollowID: "100",
			},
			want: want{
				Follows: Follows{
					Follows: []Follow{
						{
							UserID:   "101",
							FollowID: "100",
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

			resp, err := http.Get(fmt.Sprintf("http://localhost:8000/v1/followers?follow_id=%s", tt.args.FollowID))
			if err != nil {
				t.Fatalf("Error making GET request: %v", err)
			}
			defer resp.Body.Close()

			var follows Follows
			if err := json.NewDecoder(resp.Body).Decode(&follows); err != nil {
				t.Fatalf("Error decoding response body: %v", err)
			}

			for i := 0; i < len(follows.Follows); i++ {
				createdAt, err := time.Parse(time.RFC3339, follows.Follows[i].CreatedAt)
				if err != nil {
					t.Fatalf("Error parsing CreatedAt: %v", err)
				}
				if now.Before(createdAt) || now.After(createdAt.Add(time.Hour)) {
					t.Errorf("now is not within one hour of CreatedAt: %v %v", now, createdAt)
				}

				assert.Equal(t, tt.want.Follows.Follows[i].UserID, follows.Follows[i].UserID, "UserID does not match")
				assert.Equal(t, tt.want.Follows.Follows[i].FollowID, follows.Follows[i].FollowID, "FollowID does not match")
			}
			assert.Equal(t, tt.want.Follows.Count, follows.Count, "Count does not match")
			assert.Equal(t, tt.want.statusCode, resp.StatusCode, "Status code does not match")
		})
	}
}

func TestGetFollowOrNotHandler(t *testing.T) {
	type args struct {
		UserID   string
		FollowID string
	}
	type FollowOrNot struct {
		Bool bool `json:"bool"`
	}
	type want struct {
		FollowOrNot FollowOrNot
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful case",
			args: args{
				UserID:   "101",
				FollowID: "100",
			},
			want: want{
				FollowOrNot: FollowOrNot{
					Bool: true,
				},
			},
		},
		{
			name: "invalid case",
			args: args{
				UserID:   "100",
				FollowID: "101",
			},
			want: want{
				FollowOrNot: FollowOrNot{
					Bool: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf("http://localhost:8000/v1/follow_or_not?user_id=%s&follow_id=%s", tt.args.UserID, tt.args.FollowID))
			if err != nil {
				t.Fatalf("Error making GET request: %v", err)
			}
			defer resp.Body.Close()

			var followOrNot FollowOrNot
			if err := json.NewDecoder(resp.Body).Decode(&followOrNot); err != nil {
				t.Fatalf("Error decoding response body: %v", err)
			}

			assert.Equal(t, tt.want.FollowOrNot.Bool, followOrNot.Bool, "Bool does not match")
		})
	}
}
