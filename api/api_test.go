package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/munjeli/testAPI/service"
)

var (
	countURL = os.Getenv("COUNT_URL")
)

func TestCount(t *testing.T) {
	tests := []struct {
		desc    string
		data    string
		ct      string
		status  int
		wantErr bool
	}{
		{
			desc:    "test with valid file",
			data:    "testdata/testdata.txt",
			ct:      "text/plain",
			status:  http.StatusOK,
			wantErr: false,
		},
		{
			desc:    "test wrong Content-Type",
			data:    "testdata/testdata.txt",
			ct:      "application/json",
			status:  http.StatusBadRequest,
			wantErr: false,
		},
		{
			// there's a few exceptions in this file you can
			// see but I haven't gone down the rabbit hole to
			// make it perfect.
			desc:    "test wacky messy file",
			data:    "testdata/rroo.txt",
			ct:      "text/plain",
			status:  http.StatusOK,
			wantErr: false,
		},
		{
			desc:    "test file too large",
			data:    "testdata/warandpeace.txt",
			ct:      "text/plain",
			status:  http.StatusBadRequest,
			wantErr: false,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			client := &http.Client{}
			dat, err := ioutil.ReadFile(test.data)
			if err != nil {
				t.Error(err.Error())
			}
			// note that the
			req, err := http.NewRequest(http.MethodPost, countURL, bytes.NewReader(dat))
			if err != nil {
				t.Fatal(err.Error())
			}
			req.Header.Set("Content-Type", test.ct)
			resp, err := client.Do(req)
			if err != nil {
				t.Fatal(err.Error())
			}
			// I could do some checks here for content.
			if test.status == http.StatusOK {
				var w service.WordCount
				b, err := ioutil.ReadAll(resp.Body)
				if err = json.Unmarshal(b, &w); err != nil {
					t.Error(err.Error())
				}
			}
			if resp.StatusCode != test.status {
				t.Errorf("wrong status: got: %v, want: %v", resp.StatusCode, test.status)
			}
		})
	}
}
