package redis

import (
	"testing"
)

func TestNewConnPool(t *testing.T) {
	tests := []struct {
		name      string
		redisConf Config
		wantErr   bool
	}{
		{
			name: "Working connection",
			redisConf: Config{
				Host:                 "127.0.0.1",
				Port:                 "6379",
				Password:             "dev",
				MaxActiveConnections: 256,
			},
			wantErr: false,
		},
		{
			name: "Non working password",
			redisConf: Config{
				Host:                 "127.0.0.1",
				Port:                 "6379",
				Password:             "deva",
				MaxActiveConnections: 256,
			},
			wantErr: true,
		},
		{
			name: "Non working connection",
			redisConf: Config{
				Host:                 "127.0.0.1",
				Port:                 "6378",
				Password:             "dev",
				MaxActiveConnections: 256,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewConnPool(tt.redisConf)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConnPool() = %v, want %v", err, tt.wantErr)
			}
		})
	}
}
