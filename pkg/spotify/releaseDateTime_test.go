package spotify

import (
	"net/http"
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
		{
			name: "Case #1",
			fields: fields{
				http:        new(http.Client),
				id:          "1",
				secret:      "secret",
				accessToken: "token",
			},
			args: args{
				ReleaseDatePrecision: "day",
				ReleaseDate:          "2006-01-02",
			},
			want: time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "Case #2",
			fields: fields{
				http:        new(http.Client),
				id:          "1",
				secret:      "secret",
				accessToken: "token",
			},
			args: args{
				ReleaseDatePrecision: "day",
				ReleaseDate:          "2018-08-21",
			},
			want: time.Date(2018, time.August, 21, 0, 0, 0, 0, time.UTC),
		},
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
