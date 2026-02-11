package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	tests := []struct {
		name         string
		inputHeaders http.Header
		wantKey      string
		wantErr      bool
	}{
		{
			name: "Valid header",
			inputHeaders: http.Header{
				"Authorization": []string{"ApiKey valid_api_key"},
			},
			wantKey: "valid_api_key",
			wantErr: false,
		},
		{
			name: "Missing header",
			inputHeaders: http.Header{
				"Some-header": []string{"test"},
			},
			wantKey: "",
			wantErr: true,
		},
		{
			name: "Empty header",
			inputHeaders: http.Header{
				"Authorization": []string{""},
			},
			wantKey: "",
			wantErr: true,
		},
		{
			name: "Invalid header",
			inputHeaders: http.Header{
				"Authorization": []string{"valid_api_key"},
			},
			wantKey: "",
			wantErr: true,
		},
		{
			name: "Invalid header value prefix",
			inputHeaders: http.Header{
				"Authorization": []string{"Test valid_api_key"},
			},
			wantKey: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.inputHeaders)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKey != tt.wantKey {
				t.Errorf("GetAPIKey() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}
