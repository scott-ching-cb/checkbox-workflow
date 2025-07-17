package workflow_engine_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
	"net/http"
	"net/http/httptest"
	"testing"
	"workflow-code-test/api/binding/workflow"
	"workflow-code-test/api/services/workflow/testdata"
	"workflow-code-test/api/services/workflow/workflow_engine"
)

func TestExecuteIntegrationNode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{
			"latitude": -33.75,
			"longitude": 151.125,
			"generationtime_ms": 0.04291534423828125,
			"utc_offset_seconds": 0,
			"timezone": "GMT",
			"timezone_abbreviation": "GMT",
			"elevation": 86.0,
			"current_weather_units": {
				"time": "iso8601",
				"interval": "seconds",
				"temperature": "°C",
				"windspeed": "km/h",
				"winddirection": "°",
				"is_day": "",
				"weathercode": "wmo code"
			},
			"current_weather": {
				"time": "2025-07-16T16:45",
				"interval": 900,
				"temperature": 2.4,
				"windspeed": 1.8,
				"winddirection": 349,
				"is_day": 0,
				"weathercode": 0
			}
		}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()
	internalServerErrorClient := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer internalServerErrorClient.Close()

	integrationNode := &workflow.Node{
		Id:   "weather-api",
		Type: "integration",
		Position: &workflow.Node_Position{
			X: 460,
			Y: 304,
		},
		Data: &workflow.Node_Data{
			Label:       "Weather API",
			Description: "Fetch current temperature for {{city}}",
			Metadata: &workflow.MetaData{
				HasHandles: &workflow.MetaData_HasHandles{
					Source: &structpb.Value{
						Kind: &structpb.Value_BoolValue{
							BoolValue: true,
						},
					},
					Target: &structpb.Value{
						Kind: &structpb.Value_BoolValue{
							BoolValue: true,
						},
					},
				},
				InputVariables: []string{"city"},
				ApiEndpoint:    testdata.GetStringPointer(ts.URL),
				Options: []*workflow.MetaData_Option{
					{
						City: "Sydney",
						Lat:  -33.8688,
						Lon:  151.2093,
					},
					{
						City: "Melbourne",
						Lat:  -37.8136,
						Lon:  144.9631,
					},
					{
						City: "Brisbane",
						Lat:  -27.4698,
						Lon:  153.0251,
					},
					{
						City: "Perth",
						Lat:  -31.9505,
						Lon:  115.8605,
					},
					{
						City: "Adelaide",
						Lat:  -34.9285,
						Lon:  138.6007,
					},
				},
				OutputVariables: []string{"temperature"},
			},
		},
	}

	type ExecuteIntegrationNodeParams struct {
		ExpectedError               string
		ExpectedResponse            *workflow.ExecutionStep_Output_ApiResponse
		ExpectedStepOutputVariables map[string]string
		Node                        func() *workflow.Node
		StoredOutputVariables       map[string]string
	}

	testcases := map[string]ExecuteIntegrationNodeParams{
		"Should return the correct integration api response": {
			Node: func() *workflow.Node {
				return integrationNode
			},
			ExpectedResponse: &workflow.ExecutionStep_Output_ApiResponse{
				Data: &structpb.Value{
					Kind: &structpb.Value_NumberValue{
						NumberValue: 2.4,
					},
				},
				Endpoint:   fmt.Sprintf("%s?latitude=-33.8688&longitude=151.2093", ts.URL),
				Method:     "GET",
				StatusCode: 200,
			},
			ExpectedStepOutputVariables: map[string]string{"temperature": "2.40"},
			StoredOutputVariables:       map[string]string{"city": "Sydney"},
		},
		"Should return the correct error if server has error": {
			Node: func() *workflow.Node {
				currentNode := integrationNode
				currentNode.Data.Metadata.ApiEndpoint = testdata.GetStringPointer(internalServerErrorClient.URL)
				return currentNode
			},
			ExpectedError: "unsuccessful request to integration api endpoint",
			ExpectedResponse: &workflow.ExecutionStep_Output_ApiResponse{
				Endpoint:   fmt.Sprintf("%s?latitude=-33.8688&longitude=151.2093", internalServerErrorClient.URL),
				Method:     "GET",
				StatusCode: 500,
			},
			StoredOutputVariables: map[string]string{"city": "Sydney"},
		},
	}

	count := 1
	for description, testcase := range testcases {
		testDescription := fmt.Sprintf("%d %s", count, description)
		t.Run(testDescription, func(t *testing.T) {
			apiResponse, stepOutputVariables, err := workflow_engine.ExecuteIntegrationNode(
				&workflow_engine.ExecuteIntegrationNodeParams{
					Node:                  testcase.Node(),
					StoredOutputVariables: testcase.StoredOutputVariables,
				},
			)

			assert.Equal(t, testcase.ExpectedResponse.StatusCode, apiResponse.StatusCode)
			assert.Equal(t, testcase.ExpectedResponse.Method, apiResponse.Method)
			assert.Equal(t, testcase.ExpectedResponse.Endpoint, apiResponse.Endpoint)

			if testcase.ExpectedError != "" {
				assert.NotNil(t, err)
				assert.EqualError(t, err, testcase.ExpectedError)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testcase.ExpectedResponse.Data, apiResponse.Data)
				for key, value := range testcase.ExpectedStepOutputVariables {
					_, hasOutputVariable := stepOutputVariables[key]
					assert.True(t, hasOutputVariable)
					if hasOutputVariable {
						assert.Equal(t, value, stepOutputVariables[key])
					}
				}
			}
		})
	}
}
