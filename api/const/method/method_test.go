package method

import "testing"

func TestMethod_String(t *testing.T) {
	tests := []struct {
		name string
		m    Method
		want string
	}{
		{name: "post", m: Method(0), want: "POST"},
		{name: "get", m: Method(1), want: "GET"},
		{name: "unknown", m: Method(2), want: "unknown"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.String(); got != tt.want {
				t.Errorf("Method.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMethod_Validate(t *testing.T) {
	tests := []struct {
		name    string
		m       Method
		wantErr bool
	}{
		{name: "valid", m: Method(0), wantErr: false},
		{name: "valid", m: Method(1), wantErr: false},
		{name: "valid", m: Method(2), wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Method.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
