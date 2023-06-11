package unit

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/xehapa/jago/api"
	"github.com/xehapa/jago/models"
)

type mockHTTPClient struct {
	mock.Mock
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestJobAdderClient_Do(t *testing.T) {
	mockClient := new(mockHTTPClient)
	mockResponseBody := "Mock response"
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(mockResponseBody)),
	}
	mockClient.On("Do", mock.Anything).Return(mockResponse, nil)

	client := api.NewJobAdderClient()
	client.HTTPClient = mockClient
	req, err := http.NewRequest("GET", "https://test.123", nil)

	if err != nil {
		t.Fatal(err)
	}

	resp, err := client.HTTPClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	expectedBody := mockResponseBody
	if string(body) != expectedBody {
		t.Errorf("Expected response body '%s', got '%s'", expectedBody, string(body))
	}

	mockClient.AssertCalled(t, "Do", mock.Anything)
}

func TestJobAdderClient_GetPlacements(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := struct {
			Items []models.Placement `json:"items"`
		}{
			Items: []models.Placement{
				{PlacementID: 1, JobTitle: "Job 1"},
				{PlacementID: 2, JobTitle: "Job 2"},
			},
		}

		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(response)

		if err != nil {
			t.Fatal(err)
		}
	}))
	defer server.Close()

	client := api.NewJobAdderClient()
	client.ApiUrl = server.URL + "/v2/"

	placements, err := client.GetPlacements()
	if err != nil {
		t.Fatal(err)
	}

	expectedPlacements := []models.Placement{
		{PlacementID: 1, JobTitle: "Job 1"},
		{PlacementID: 2, JobTitle: "Job 2"},
	}

	if !reflect.DeepEqual(expectedPlacements, placements) {
		t.Errorf("Expected placements do not match actual placements")
	}
}
