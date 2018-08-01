package redis

import (
	"testing"
)

func TestExists(t *testing.T) {

	redPool, _ := NewConnPool(Config{
		Host:                 "127.0.0.1",
		Port:                 "6379",
		Password:             "dev",
		MaxActiveConnections: 256,
	})
	if _, err := redPool.Get().Do("SET", "123", ""); err != nil {
		t.Errorf("Do(SET(): %v", err)
		return
	}
	if _, err := redPool.Get().Do("SET", "12", "-"); err != nil {
		t.Errorf("Do(SET(): %v", err)
		return
	}

	tests := []struct {
		name      string
		key       string
		shdExists bool
		wantErr   bool
	}{
		{
			name:      "exist with empty val",
			key:       "123",
			shdExists: true,
			wantErr:   false,
		},
		{
			name:      "exists",
			key:       "12",
			shdExists: true,
			wantErr:   false,
		},
		{
			name:      "not exist",
			key:       "1234",
			shdExists: false,
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			exists, err := Exists(redPool, tt.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exists error %v unexpected", err)
				return
			}
			if exists != tt.shdExists {
				t.Errorf("Exists() shd exists but didnt")
				return
			}
		})
	}

}
