package ethereum

import "testing"

func TestNewRPCCli(t *testing.T) {

	tests := []struct {
		name    string
		ethConf Config
		wantErr bool
	}{
		{
			name: "Working connection",
			ethConf: Config{
				JSONRPCProtocol: "http",
				Host:            "127.0.0.1",
				JSONRPCPort:     "8645",
				NetworkID:       17,
			},
			wantErr: false,
		},
		{
			name: "Working connection, wrong network",
			ethConf: Config{
				JSONRPCProtocol: "http",
				Host:            "127.0.0.1",
				JSONRPCPort:     "8645",
				NetworkID:       15,
			},
			wantErr: true,
		},
		{
			name: "Non working connection",
			ethConf: Config{
				JSONRPCProtocol: "http",
				Host:            "127.0.0.1",
				JSONRPCPort:     "8888",
			},
			wantErr: true,
		},
		{
			name: "Non working Dial",
			ethConf: Config{
				JSONRPCProtocol: "zxp",
				Host:            "127.0.0.1",
				JSONRPCPort:     "8888",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewRPCCli(tt.ethConf)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"NewRPCCli() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
		})
	}
}

func TestNewWSCli(t *testing.T) {

	tests := []struct {
		name    string
		ethConf Config
		wantErr bool
	}{
		{
			name: "Working connection",
			ethConf: Config{
				WSProtocol: "ws",
				Host:       "127.0.0.1",
				WSPort:     "8646",
				NetworkID:  17,
			},
			wantErr: false,
		},
		{
			name: "Working connection, wrong network",
			ethConf: Config{
				WSProtocol: "ws",
				Host:       "127.0.0.1",
				WSPort:     "8646",
				NetworkID:  15,
			},
			wantErr: true,
		},
		{
			name: "Non working connection",
			ethConf: Config{
				WSProtocol: "ws",
				Host:       "127.0.0.1",
				WSPort:     "8888",
			},
			wantErr: true,
		},
		{
			name: "Non working Dial",
			ethConf: Config{
				WSProtocol: "zxp",
				Host:       "127.0.0.1",
				WSPort:     "8888",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewWSCli(tt.ethConf)
			if (err != nil) != tt.wantErr {
				t.Errorf(
					"NewWSCli() error = %v, wantErr %v",
					err, tt.wantErr)
				return
			}
		})
	}
}
