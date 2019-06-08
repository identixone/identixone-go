package transport

import "testing"

func TestTransport_IsValid(t *testing.T) {
	tests := []struct {
		name    string
		t       Transport
		wantErr bool
	}{
		{name: "webhook", t: Transport(0), wantErr: false},
		{name: "websocket client", t: Transport(1), wantErr: false},
		{name: "websocket server", t: Transport(2), wantErr: false},
		{name: "unknown", t: Transport(3), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.IsValid(); (err != nil) != tt.wantErr {
				t.Errorf("Transport.IsValid() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTransport_String(t *testing.T) {
	tests := []struct {
		name string
		t    Transport
		want string
	}{
		{name: "webhook", t: Transport(0), want: "webhook"},
		{name: "websocket client", t: Transport(1), want: "websocket client"},
		{name: "websocket server", t: Transport(2), want: "websocket server"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t.Errorf("Transport.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
