package surly_test

import (
	"encoding/json"
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/lukasbob/surly"
)

type Test struct {
	XMLName xml.Name  `json:"-" xml:"test"`
	URL     surly.URL `json:"url" xml:"url"`
}

var u, _ = surly.New("http://example.com")

func TestURL_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		t       *Test
		want    string
		wantErr bool
	}{
		{
			name: "Simple",
			t:    &Test{URL: u},
			want: `{"url":"http://example.com"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := json.Marshal(tt.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("URL.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(string(got), tt.want) {
				t.Errorf("URL.MarshalJSON() = %s, want %v", got, tt.want)
			}
		})
	}
}

func TestURL_UnmarshalJSON(t *testing.T) {
	type args struct {
		b string
	}
	tests := []struct {
		name    string
		u       *Test
		args    args
		wantErr bool
	}{
		{
			name: "Valid URL",
			u:    &Test{URL: u},
			args: args{b: `{"url":"http://example.com"}`},
		},
		{
			name:    "Invalid URL",
			u:       nil,
			args:    args{b: `{"url":"[foul] http://example.com"}`},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var test Test
			if err := json.Unmarshal([]byte(tt.args.b), &test); (err != nil) != tt.wantErr {
				t.Errorf("URL.UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestURL_UnmarshalXML(t *testing.T) {
	type args struct {
		b string
	}
	tests := []struct {
		name    string
		u       *Test
		args    args
		wantErr bool
	}{
		{
			name: "Valid URL",
			u:    &Test{URL: u},
			args: args{b: `<test><url>http://example.com</url></test>`},
		},
		{
			name: "With CDATA",
			u:    &Test{URL: u},
			args: args{b: `<test><url><![CDATA[http://example.com]]></url></test>`},
		},
		{
			name:    "Invalid URL",
			u:       nil,
			args:    args{b: `<test><url>[foul] http://example.com</url></test>`},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var test Test
			if err := xml.Unmarshal([]byte(tt.args.b), &test); (err != nil) != tt.wantErr {
				t.Errorf("URL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
