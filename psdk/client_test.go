package psdk

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

const (
	testAouBase = "http://127.0.0.1:8800"
	testToken   = ""
)

func getClient() (*PushdeerClient, error) {
	var got *PushdeerClient
	var err error
	if testToken != "" {
		got, err = New(testAouBase, testToken)
	} else {
		got, err = NewFaker(testAouBase)
	}
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	log.Infof("get faker token:%v", got.token)

	return got, nil
}

func TestNewFaker(t *testing.T) {
	type args struct {
		apiBase string
	}
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "new-faker",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFaker(testAouBase)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFaker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.token == "" {
				t.Errorf("get token fail")
			}
		})
	}
}

func TestUserInfo(t *testing.T) {
	type args struct {
		apiBase string
	}
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "user-info",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := getClient()
			if err != nil {
				t.Errorf("err:%v", err)
				return
			}

			got, err := client.UserInfo(&UserInfoReq{})
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFaker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			log.Infof("%+v", got)
		})
	}
}
