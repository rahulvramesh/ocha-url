package bigcache

import (
	"github.com/allegro/bigcache/v3"
	"testing"
)

func TestGetDS(t *testing.T) {
	InitializeDataSource()

	tests := []struct {
		name string
		want *bigcache.BigCache
	}{
		{
			name: "Positive Shared Connection Return",
			want: nil, // !nil
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDS(); got == tt.want {
				t.Errorf("GetDS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitializeDataSource(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Positive Connection Testing",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitializeDataSource()

			if DS == nil{
				// Throw Error
				t.Errorf("InitializeDataSource() = %v, want %v", "NIL", "Non Nil")
			}
		})
	}
}