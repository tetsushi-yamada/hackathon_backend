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
		UserId   string
		UserName string
	}
	type User struct {
		UserId   string `json:"user_id"`
		UserName string `json:"user_name"`
	}
	type want struct {
		UserId     string
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
				UserId:   "8989",
				UserName: "test_user",
			},
			want: want{
				UserId:     "8989",
				statusCode: http.StatusCreated,
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			data := User{
				UserId:   tt.args.UserId,
				UserName: tt.args.UserName,
			}
			payloadBytes, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error marshaling data:", err)
				return
			}
			body := bytes.NewReader(payloadBytes)

			req, err := http.NewRequest("POST", "http://localhost:8000/v1/users", body)
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
			if ok := assert.Equal(t, tt.want.UserId, data.UserId); !ok {
				t.Fatalf("invalid user id %s", data.UserId)
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
		statusCode int
	}
	type User struct {
		UserId    string `json:"user_id"`
		UserName  string `json:"user_name"`
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
				statusCode: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Now().UTC().Add(time.Hour * 9) // 現在の時刻を取得

			resp, err := http.Get(fmt.Sprintf("http://localhost:8000/v1/users/%s", tt.args.UserID))
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
				statusCode: http.StatusNoContent,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8000/v1/users/%s", tt.args.UserID), nil)
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

func TestSearchUsersHandler(t *testing.T) {
	type args struct {
		SearchWord string
	}
	type User struct {
		UserId    string `json:"user_id"`
		UserName  string `json:"user_name"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	type Users struct {
		Users []User `json:"users"`
		Count int    `json:"count"`
	}
	type want struct {
		Users      Users
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
				SearchWord: "FOLLOW",
			},
			want: want{
				Users: Users{
					Users: []User{
						{
							UserId:   "100",
							UserName: "FOLLOW ME",
						},
						{
							UserId:   "101",
							UserName: "FOLLOW YOU",
						},
					},
					Count: 2,
				},
				statusCode: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := http.Get(fmt.Sprintf("http://localhost:8000/v1/users/search/%s", tt.args.SearchWord))
			if err != nil {
				t.Fatalf("Error making GET request: %v", err)
			}
			defer resp.Body.Close()

			assert.Equal(t, tt.want.statusCode, resp.StatusCode, "Status code does not match")

			var users Users
			if err := json.NewDecoder(resp.Body).Decode(&users); err != nil {
				t.Fatalf("Error decoding response body: %v", err)
			}

			for i := 0; i < len(users.Users); i++ {
				assert.Equal(t, tt.want.Users.Users[i].UserId, users.Users[i].UserId, "UserID does not match")
				assert.Equal(t, tt.want.Users.Users[i].UserName, users.Users[i].UserName, "UserName does not match")
			}
			assert.Equal(t, tt.want.Users.Count, users.Count, "Count does not match")
		})
	}
}
