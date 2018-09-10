package postgres

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
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
		{
			name: "Secure connection",
			postgresDBConf: Config{
				Host:     "none",
				Port:     "5432",
				User:     "internal",
				Password: "dev",
				DbName:   "test",
				SSLConf: SSLConf{
					CertPath:     "/test",
					KeyPath:      "/test",
					RootCertPath: "/test",
				},
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

func Test_getDSN(t *testing.T) {
	tests := []struct {
		name string
		c    Config
		want string
	}{
		{
			name: "without ssl",
			c: Config{
				Host:     "none",
				Port:     "5432",
				User:     "internal",
				Password: "dev",
				DbName:   "test",
			},
			want: "postgres://internal:dev@none:5432/test?sslmode=disable",
		},
		{
			name: "without ssl",
			c: Config{
				Host:     "none",
				Port:     "5432",
				User:     "internal",
				Password: "dev",
				DbName:   "test",
				SSLConf: SSLConf{
					CertPath:     "/test",
					KeyPath:      "/test",
					RootCertPath: "/test",
				},
			},
			want: "postgres://internal:dev@none:5432/test?sslmode=verify-full&sslcert=/test&sslkey=/test&sslrootcert=/test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDSN(tt.c); got != tt.want {
				t.Errorf("getDSN() = %v, want %v", got, tt.want)
			}
		})
	}
}
