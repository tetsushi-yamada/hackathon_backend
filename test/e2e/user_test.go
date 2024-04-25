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

func TestCreateUserHandler(t *testing.T) {
	type args struct {
		UserName string
		Email    string
	}
	type User struct {
		UserId   string `json:"user_id"`
		UserName string `json:"user_name"`
		Email    string `json:"email"`
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
				UserName: "test_user",
				Email:    "test@hello.jp",
			},
			want: want{
				statusCode: http.StatusCreated,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			data := User{
				UserName: tt.args.UserName,
				Email:    tt.args.Email,
			}
			payloadBytes, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error marshaling data:", err)
				return
			}
			body := bytes.NewReader(payloadBytes)

			req, err := http.NewRequest("PUT", "http://localhost:8000/v1/users", body)
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

func TestGetUserHandler(t *testing.T) {
	type args struct {
		UserID string
	}
	type want struct {
		UserName   string
		Email      string
		statusCode int
	}
	type User struct {
		UserId    string `json:"user_id"`
		UserName  string `json:"user_name"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "successful case",
			args: args{
				UserID: "1",
			},
			want: want{
				UserName:   "JohnDoe",
				Email:      "john@example.com",
				statusCode: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Now().UTC().Add(time.Hour * 9) // 現在の時刻を取得

			resp, err := http.Get(fmt.Sprintf("http://localhost:8000/v1/users?user_id=%s", tt.args.UserID))
			if err != nil {
				t.Fatalf("Error making GET request: %v", err)
			}
			defer resp.Body.Close()

			var user User
			if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
				t.Fatalf("Error decoding response body: %v", err)
			}

			createdAt, err := time.Parse(time.RFC3339, user.CreatedAt)
			if now.Before(createdAt) || now.After(createdAt.Add(time.Hour)) {
				t.Errorf("now is not within one hour of CreatedAt: %v %v", now, createdAt)
			}

			assert.Equal(t, tt.want.UserName, user.UserName, "UserName does not match")
			assert.Equal(t, tt.want.Email, user.Email, "Email does not match")
			assert.Equal(t, tt.want.statusCode, resp.StatusCode, "Status code does not match")
		})
	}
}

func TestDeleteUserHandler(t *testing.T) {
	type args struct {
		UserID string
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
				UserID: "1",
			},
			want: want{
				statusCode: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8000/v1/users?user_id=%s", tt.args.UserID), nil)
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
