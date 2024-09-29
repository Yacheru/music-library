package spotify

import (
	"reflect"
	"testing"
	"time"
)

func TestClient_ReleaseDateTime(t *testing.T) {
	type fields struct {
		http        *http.Client
		id          string
		secret      string
		accessToken string
	}
	type args struct {
		ReleaseDatePrecision string
		ReleaseDate          string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Time
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				http:        tt.fields.http,
				id:          tt.fields.id,
				secret:      tt.fields.secret,
				accessToken: tt.fields.accessToken,
			}
			if got := c.ReleaseDateTime(tt.args.ReleaseDatePrecision, tt.args.ReleaseDate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReleaseDateTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
