package testdata

import (
	"workflow-code-test/api/binding/workflow"

	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/structpb"
)

func GetBoolPointer(value bool) *bool {
	return &value
}

func GetStringPointer(value string) *string {
	return &value
}

var (
	SampleWorkflowId = uuid.New()

	SampleWorkflowEdges = &workflow.Edges{
		Edges: []*workflow.Edge{
			{
				Id:       "e1",
				Source:   "start",
				Target:   "form",
				Type:     "smoothstep",
				Animated: GetBoolPointer(true),
				Style: &workflow.Edge_Style{
					Stroke:      "#10b981",
					StrokeWidth: 3,
				},
				Label: GetStringPointer("Initialize"),
			},
			{
				Id:       "e2",
				Source:   "form",
				Target:   "weather-api",
				Type:     "smoothstep",
				Animated: GetBoolPointer(true),
				Style: &workflow.Edge_Style{
					Stroke:      "#3b82f6",
					StrokeWidth: 3,
				},
				Label: GetStringPointer("Submit Data"),
			},
			{
				Id:       "e3",
				Source:   "weather-api",
				Target:   "condition",
				Type:     "smoothstep",
				Animated: GetBoolPointer(true),
				Style: &workflow.Edge_Style{
					Stroke:      "#f97316",
					StrokeWidth: 3,
				},
				Label: GetStringPointer("Temperature Data"),
			},
			{
				Id:           "e4",
				Source:       "condition",
				Target:       "email",
				Type:         "smoothstep",
				SourceHandle: GetStringPointer("true"),
				Animated:     GetBoolPointer(true),
				Style: &workflow.Edge_Style{
					Stroke:      "#10b981",
					StrokeWidth: 3,
				},
				Label: GetStringPointer("✓ Condition Met"),
				LabelStyle: &workflow.Edge_LabelStyle{
					Fill:       "#10b981",
					FontWeight: "bold",
				},
			},
			{
				Id:           "e5",
				Source:       "condition",
				Target:       "end",
				Type:         "smoothstep",
				SourceHandle: GetStringPointer("false"),
				Animated:     GetBoolPointer(true),
				Style: &workflow.Edge_Style{
					Stroke:      "#6b7280",
					StrokeWidth: 3,
				},
				Label: GetStringPointer("✗ No Alert Needed"),
				LabelStyle: &workflow.Edge_LabelStyle{
					Fill:       "#6b7280",
					FontWeight: "bold",
				},
			},
			{
				Id:           "e6",
				Source:       "email",
				Target:       "end",
				Type:         "smoothstep",
				SourceHandle: GetStringPointer("false"),
				Animated:     GetBoolPointer(true),
				Style: &workflow.Edge_Style{
					Stroke:      "#ef4444",
					StrokeWidth: 2,
				},
				Label: GetStringPointer("Alert Sent"),
				LabelStyle: &workflow.Edge_LabelStyle{
					Fill:       "#ef4444",
					FontWeight: "bold",
				},
			},
		},
	}

	SampleWorkflowNodes = &workflow.Nodes{
		Nodes: []*workflow.Node{
			{
				Id:   "start",
				Type: "start",
				Position: &workflow.Node_Position{
					X: -160,
					Y: 300,
				},
				Data: &workflow.Node_Data{
					Label:       "Start",
					Description: "Begin weather check workflow",
					Metadata: &workflow.MetaData{
						HasHandles: &workflow.MetaData_HasHandles{
							Source: &structpb.Value{
								Kind: &structpb.Value_BoolValue{
									BoolValue: true,
								},
							},
							Target: &structpb.Value{
								Kind: &structpb.Value_BoolValue{
									BoolValue: false,
								},
							},
						},
					},
				},
			},
			{
				Id:   "form",
				Type: "form",
				Position: &workflow.Node_Position{
					X: 152,
					Y: 304,
				},
				Data: &workflow.Node_Data{
					Label:       "User Input",
					Description: "Process collected data - name, email, location",
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
						InputFields:     []string{"name", "email", "city"},
						OutputVariables: []string{"name", "email", "city"},
					},
				},
			},
			{
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
						ApiEndpoint:    GetStringPointer("https://api.open-meteo.com/v1/forecast?latitude={lat}&longitude={lon}&current_weather=true"),
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
			},
			{
				Id:   "condition",
				Type: "condition",
				Position: &workflow.Node_Position{
					X: 794,
					Y: 304,
				},
				Data: &workflow.Node_Data{
					Label:       "Check Condition",
					Description: "Evaluate temperature threshold",
					Metadata: &workflow.MetaData{
						HasHandles: &workflow.MetaData_HasHandles{
							Source: &structpb.Value{
								Kind: &structpb.Value_ListValue{
									ListValue: &structpb.ListValue{
										Values: []*structpb.Value{
											{
												Kind: &structpb.Value_StringValue{
													StringValue: "true",
												},
											},
											{
												Kind: &structpb.Value_StringValue{
													StringValue: "false",
												},
											},
										},
									},
								},
							},
							Target: &structpb.Value{
								Kind: &structpb.Value_BoolValue{
									BoolValue: true,
								},
							},
						},
						ConditionExpression: GetStringPointer("temperature {{operator}} {{threshold}}"),
						OutputVariables:     []string{"conditionMet"},
					},
				},
			},
			{
				Id:   "email",
				Type: "email",
				Position: &workflow.Node_Position{
					X: 1096,
					Y: 88,
				},
				Data: &workflow.Node_Data{
					Label:       "Send Alert",
					Description: "Email weather alert notification",
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
						InputVariables: []string{"name", "city", "temperature"},
						EmailTemplate: &workflow.MetaData_EmailTemplate{
							Body:    "Weather alert for {{city}}! Temperature is {{temperature}}°C!",
							Subject: "Weather Alert",
						},
						OutputVariables: []string{"emailSent"},
					},
				},
			},
			{
				Id:   "end",
				Type: "end",
				Position: &workflow.Node_Position{
					X: 1360,
					Y: 302,
				},
				Data: &workflow.Node_Data{
					Label:       "Complete",
					Description: "Workflow execution finished",
					Metadata: &workflow.MetaData{
						HasHandles: &workflow.MetaData_HasHandles{
							Source: &structpb.Value{
								Kind: &structpb.Value_BoolValue{
									BoolValue: false,
								},
							},
							Target: &structpb.Value{
								Kind: &structpb.Value_BoolValue{
									BoolValue: true,
								},
							},
						},
					},
				},
			},
		},
	}
)
