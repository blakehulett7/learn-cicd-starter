package main

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAuth(t *testing.T) {
	tests := map[string]struct {
		headers          http.Header
		wantedAuthString string
		wantedErr        error
	}{
		"simple":              {headers: http.Header{"Authorization": []string{"ApiKey testkey"}}, wantedAuthString: "testkey", wantedErr: nil},
		"noAuthHeader":        {headers: http.Header{}, wantedAuthString: "", wantedErr: auth.ErrNoAuthHeaderIncluded},
		"malformedAuthHeader": {headers: http.Header{"Authorization": []string{"testkey"}}, wantedAuthString: "", wantedErr: auth.ErrMalformedAuthHeader},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			gotAuthString, gotErr := auth.GetAPIKey(test.headers)
			if !reflect.DeepEqual(test.wantedAuthString, gotAuthString) {
				t.Fatalf("%s: ExpectedAuthString %v, Got %v", name, test.wantedAuthString, gotAuthString)
			}
			if test.wantedErr != gotErr {
				t.Fatalf("%s: ExpectedErr %v, Got %v", name, test.wantedErr, gotErr)
			}
			fmt.Println(name, "test passed!")
		})
	}
}
