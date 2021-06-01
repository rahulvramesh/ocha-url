package services

import (
	"github.com/rahulvramesh/ocha-url/internal/app/ocha/models"
	"github.com/rahulvramesh/ocha-url/internal/pkg/bigcache"
	"testing"
)

func TestOchaService_GetByKey(t *testing.T) {

	bigcache.InitializeDataSource()
	var tObj OchaService
	_ = tObj.StoreLink("google", "https://google.com")

	type fields struct {
		Data models.Item
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Positive Test Case",
			args:    args{key: "google"},
			wantErr: false,
		},
		{
			name:    "Negative Test Case",
			args:    args{key: "Unknown"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OchaService{
				Data: tt.fields.Data,
			}
			if err := o.GetByKey(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("GetByKey() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func TestOchaService_CheckIfKeyExists(t *testing.T) {

	bigcache.InitializeDataSource()
	var tObj OchaService
	_ = tObj.StoreLink("google", "https://google.com")

	type fields struct {
		Data models.Item
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Positive Test Case",
			args:    args{key: "google"},
			wantErr: false,
		},
		{
			name:    "Negative Test Case",
			args:    args{key: "Unknown"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OchaService{
				Data: tt.fields.Data,
			}
			if err := o.CheckIfKeyExists(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("CheckIfKeyExists() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOchaService_StoreLink(t *testing.T) {

	bigcache.InitializeDataSource()

	type fields struct {
		Data models.Item
	}
	type args struct {
		key string
		url string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		response string
		wantErr  bool
	}{
		{
			name:     "Positive Test Case",
			args:     args{key: "google", url: "https://google.com"},
			response: "https://google.com",
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &OchaService{
				Data: tt.fields.Data,
			}
			if err := o.StoreLink(tt.args.key, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("StoreLink() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
