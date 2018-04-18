package postgres

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestNewConnPool(t *testing.T) {

	tests := []struct {
		name           string
		postgresDBConf Config
		wantErr        bool
	}{
		{
			name: "Working connection postgres",
			postgresDBConf: Config{
				Host:     "127.0.0.1",
				Port:     "5432",
				User:     "internal",
				Password: "dev",
			},
			wantErr: false,
		},
		{
			name: "Working connection cockroachdb",
			postgresDBConf: Config{
				Host: "127.0.0.1",
				Port: "26257",
				User: "root",
			},
			wantErr: false,
		},
		{
			name: "Non existing db",
			postgresDBConf: Config{
				Host:     "none",
				Port:     "5432",
				User:     "internal",
				Password: "dev",
				DbName:   "test",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewConnPool(tt.postgresDBConf)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"NewConnPool() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
		})
	}
}
