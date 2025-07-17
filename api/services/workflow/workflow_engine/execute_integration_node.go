package workflow_engine

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"workflow-code-test/api/binding/workflow"
)

type ExecuteIntegrationNodeParams struct {
	Node                  *workflow.Node
	StoredOutputVariables map[string]string
}

// fetchIntegrationApiResponse retrieves the data from the integration API endpoint and returns an ApiResponse object
func fetchIntegrationApiResponse(apiEndpoint string) (*workflow.ExecutionStep_Output_ApiResponse, error) {
	slog.Debug(fmt.Sprintf("fetch_api_integration : making integration request to %s", apiEndpoint))
	apiOutputResponse := &workflow.ExecutionStep_Output_ApiResponse{
		Endpoint: apiEndpoint,
		Method:   "GET",
	}

	// Retrieve data from the integration endpoint
	apiResponse, err := http.Get(apiEndpoint)
	if err != nil || apiResponse.StatusCode != 200 {
		slog.Error(fmt.Sprintf("fetch_api_integration : error in endpoint request %s", apiEndpoint), "error", err)
		if apiResponse != nil {
			apiOutputResponse.StatusCode = int32(apiResponse.StatusCode)
		}
		return apiOutputResponse, fmt.Errorf("unsuccessful request to integration api endpoint")
	}
	apiOutputResponse.StatusCode = int32(apiResponse.StatusCode)

	// Defer closing of the io reader
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("fetch_api_integration : failed to close api io reader", "error", err)
		}
	}(apiResponse.Body)

	// Parse and unmarshal the body of the API response
	responseBodyBytes, err := io.ReadAll(apiResponse.Body)
	if err != nil {
		slog.Error("fetch_api_integration : failed to read response body", "error", err)
		return apiOutputResponse, fmt.Errorf("invalid response body received from integration api endpoint")
	}
	var responseBody workflow.IntegrationApiResponse
	if err := protojson.Unmarshal(responseBodyBytes, &responseBody); err != nil {
		slog.Error("fetch_api_integration : failed to parse response body", "error", err)
		return apiOutputResponse, fmt.Errorf("invalid response body received from integration api endpoint")
	}

	// Check the status code in the api response
	if apiResponse.StatusCode != http.StatusOK {
		apiResponseError := fmt.Errorf("received non 200 status code: %d", apiResponse.StatusCode)
		if responseBody.GetError() {
			apiResponseError = fmt.Errorf(responseBody.GetReason())
		}
		return apiOutputResponse, apiResponseError
	}

	apiOutputResponse.Data = &structpb.Value{
		Kind: &structpb.Value_NumberValue{
			NumberValue: responseBody.CurrentWeather.Temperature,
		},
	}
	return apiOutputResponse, nil
}

// ExecuteIntegrationNode parses the API endpoint in the integration node, sets the query parameters and
// returns the response from the API
func ExecuteIntegrationNode(params *ExecuteIntegrationNodeParams) (*workflow.ExecutionStep_Output_ApiResponse, map[string]string, error) {
	// Select a single location variable from previous node outputs
	var selectedLocation string
	for _, inputVariable := range params.Node.Data.Metadata.InputVariables {
		if _, ok := params.StoredOutputVariables[inputVariable]; ok {
			selectedLocation = params.StoredOutputVariables[inputVariable]
		}
	}

	// Retrieve the location for which to query whether results
	cityToOptionMap := make(map[string]*workflow.MetaData_Option)
	for _, option := range params.Node.Data.Metadata.Options {
		cityToOptionMap[option.City] = option
	}
	selectedOption := cityToOptionMap[selectedLocation]

	// Parse integration URL and replace required query parameters
	apiEndpoint, err := url.Parse(params.Node.Data.Metadata.GetApiEndpoint())
	if err != nil {
		return &workflow.ExecutionStep_Output_ApiResponse{
			Endpoint: params.Node.Data.Metadata.GetApiEndpoint(),
			Method:   "GET",
		}, nil, fmt.Errorf("failed to parse integration endpoint to url")
	}
	queryParameters := apiEndpoint.Query()
	strconv.FormatFloat(selectedOption.Lat, 'f', -1, 64)
	queryParameters.Set("latitude", strconv.FormatFloat(selectedOption.Lat, 'f', -1, 64))
	queryParameters.Set("longitude", strconv.FormatFloat(selectedOption.Lon, 'f', -1, 64))
	apiEndpoint.RawQuery = queryParameters.Encode()

	apiResponse, err := fetchIntegrationApiResponse(apiEndpoint.String())
	outputVariables := make(map[string]string)
	if err == nil {
		for _, outputVariable := range params.Node.Data.Metadata.OutputVariables {
			outputVariables[outputVariable] = strconv.FormatFloat(apiResponse.Data.GetNumberValue(), 'f', 2, 64)
		}
	}
	return apiResponse, outputVariables, err
}
