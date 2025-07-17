package validator_test

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
	"testing"
	"workflow-code-test/api/binding/workflow"
	"workflow-code-test/api/services/workflow/testdata"
	"workflow-code-test/api/services/workflow/validator"
)

func TestValidateNodes(t *testing.T) {

	type ValidateNodesParams struct {
		ExpectedError string
		Nodes         func() []*workflow.Node
	}

	testcases := map[string]ValidateNodesParams{
		"Should return correct error if no end node in workflow": {
			ExpectedError: "missing start or end nodes",
			Nodes: func() []*workflow.Node {
				nodes := make([]*workflow.Node, 0)
				for _, node := range testdata.SampleWorkflowNodes.Nodes {
					if node.Type != "end" {
						nodes = append(nodes, node)
					}
				}
				return nodes
			},
		},
		"Should return correct error if no start node in workflow": {
			ExpectedError: "missing start or end nodes",
			Nodes: func() []*workflow.Node {
				nodes := make([]*workflow.Node, 0)
				for _, node := range testdata.SampleWorkflowNodes.Nodes {
					if node.Type != "start" {
						nodes = append(nodes, node)
					}
				}
				return nodes
			},
		},
		"Should return correct error if there are duplicate node ids": {
			ExpectedError: "duplicate nodes with id form",
			Nodes: func() []*workflow.Node {
				existingNodes := testdata.SampleWorkflowNodes.Nodes
				existingNodes = append(existingNodes, &workflow.Node{
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
				})
				return existingNodes
			},
		},
		"Should return correct error if start node has invalid handles": {
			ExpectedError: "start node cannot be the target of another node",
			Nodes: func() []*workflow.Node {
				return []*workflow.Node{
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
											BoolValue: true,
										},
									},
								},
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
				}
			},
		},
		"Should return correct error if end node has invalid handles": {
			ExpectedError: "end node cannot be the source of another node",
			Nodes: func() []*workflow.Node {
				return []*workflow.Node{
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
											BoolValue: true,
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
				}
			},
		},
	}

	count := 1
	for description, testcase := range testcases {
		t.Run(description, func(t *testing.T) {
			err := validator.ValidateNodes(testcase.Nodes())
			if testcase.ExpectedError != "" {
				assert.EqualError(t, err, testcase.ExpectedError)
			} else {
				assert.NoError(t, err)
			}
		})
		count++
	}
}
