package x

import "testing"

func TestSplitHostPort(t *testing.T) {
	type args struct {
		addr string
	}
	tests := []struct {
		name     string
		args     args
		wantHost string
		wantPort int
		wantErr  bool
	}{
		// TODO: Add test cases.
		{"1", args{"localhost:8080"}, "localhost", 8080, false},
		{"2", args{"localhost:yellow"}, "", -1, true},
		{"3", args{"localhost"}, "", -1, true},
		{"4", args{"localhost:"}, "", -1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHost, gotPort, err := SplitHostPort(tt.args.addr)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitHostPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotHost != tt.wantHost {
				t.Errorf("SplitHostPort() gotHost = %v, want %v", gotHost, tt.wantHost)
			}
			if gotPort != tt.wantPort {
				t.Errorf("SplitHostPort() gotPort = %v, want %v", gotPort, tt.wantPort)
			}
		})
	}
}
