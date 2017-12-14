package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wildnature/macaque/pkg/server/api"
)

const (
	server = "http://localhost:9090"
)

// TestMain before each test
func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

func skipUnitTest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}
}

func TestGetScheduler(t *testing.T) {
	skipUnitTest(t)
	url := server + "/api/v0/schedulers/%s"
	cases := []struct {
		description        string
		schedulerID        string
		expectedHTTPStatus int
		expectedResponse   *api.Scheduler
	}{
		{
			description:        "Test I - Success",
			schedulerID:        "1",
			expectedHTTPStatus: http.StatusOK,
			expectedResponse: &api.Scheduler{
				Name: "my-scheduler",
			},
		},
		{
			description:        "Test II - Scheduler not found",
			schedulerID:        "00",
			expectedHTTPStatus: http.StatusNotFound,
			expectedResponse:   nil,
		},
	}
	for _, c := range cases {
		resp, err := http.Get(fmt.Sprintf(url, c.schedulerID))
		if err != nil {
			t.Error(err.Error())
		}
		defer resp.Body.Close()
		assert.Equal(t, c.expectedHTTPStatus, resp.StatusCode)
		body, err := ioutil.ReadAll(resp.Body)
		if c.expectedResponse != nil {
			res := &api.Scheduler{}
			err := json.Unmarshal(body, res)
			if err != nil {
				t.Error(err.Error())
			}
			assert.Equal(t, c.expectedResponse.Name, res.Name)
		}
	}

}
