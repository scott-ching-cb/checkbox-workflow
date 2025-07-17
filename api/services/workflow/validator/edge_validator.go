package validator

import (
	"fmt"
	"workflow-code-test/api/binding/workflow"
)

type NodeMetadata struct {
	StartNodeId string
	EndNodeId   string
	NodeIdSet   map[string]struct{}
}

// getNodeMetadata returns the start and end nodes, and the node ids as a set
func getNodeMetadata(nodes []*workflow.Node) NodeMetadata {
	nodeMetadata := &NodeMetadata{
		NodeIdSet: make(map[string]struct{}),
	}
	for _, node := range nodes {
		if node.Type == "start" {
			nodeMetadata.StartNodeId = node.Id
		} else if node.Type == "end" {
			nodeMetadata.EndNodeId = node.Id
		}
		nodeMetadata.NodeIdSet[node.Id] = struct{}{}
	}
	return *nodeMetadata
}

// ValidateEdges validates the edges and ensures no duplicate edge id, and that the source and target nodes are valid ids
func ValidateEdges(edges []*workflow.Edge, nodes []*workflow.Node) error {
	nodeMetadata := getNodeMetadata(nodes)
	edgeIdSet := map[string]struct{}{}
	for _, edge := range edges {
		if _, ok := edgeIdSet[edge.Id]; ok {
			return fmt.Errorf("duplicate edges with id %s", edge.Id)
		}
		_, sourceNodeExists := nodeMetadata.NodeIdSet[edge.Source]
		_, targetNodeExists := nodeMetadata.NodeIdSet[edge.Target]
		isHandlesInWorkflow := sourceNodeExists && targetNodeExists
		isValidHandles := edge.Source != nodeMetadata.EndNodeId && edge.Target != nodeMetadata.StartNodeId
		if !isHandlesInWorkflow || !isValidHandles {
			return fmt.Errorf("invalid target or source node for edge %s", edge.Id)
		}
		edgeIdSet[edge.Id] = struct{}{}
	}
	return nil
}
