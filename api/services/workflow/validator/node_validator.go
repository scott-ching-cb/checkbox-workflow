package validator

import (
	"fmt"
	"workflow-code-test/api/binding/workflow"
)

func ValidateNodes(nodes []*workflow.Node) error {
	containsStartNode := false
	containsEndNode := false
	nodeIdSet := map[string]struct{}{}
	for _, node := range nodes {
		if _, ok := nodeIdSet[node.Id]; ok {
			return fmt.Errorf("duplicate nodes with id %s", node.Id)
		}

		if _, ok := ValidNodeTypes[node.Type]; !ok {
			return fmt.Errorf("invalid node type %s for node %s", node.Type, node.Id)
		}

		// Perform sanity check on node metadata
		if node.Data == nil {
			return fmt.Errorf("node %s data is empty", node.Id)
		} else if node.Data.Metadata == nil {
			return fmt.Errorf("node %s metadata is empty", node.Id)
		}

		nodeHandles := node.Data.Metadata.HasHandles
		if nodeHandles == nil {
			return fmt.Errorf("node %s is missing handles", node.Id)
		}

		// Validate start node uniqueness and handles
		if node.Type == "start" {
			if containsStartNode {
				return fmt.Errorf("multiple start nodes within workflow")
			} else if nodeHandles.Target.GetBoolValue() != false {
				return fmt.Errorf("start node cannot be the target of another node")
			}
			containsStartNode = true
		}

		// Validate end node uniqueness and handles
		if node.Type == "end" {
			if containsEndNode {
				return fmt.Errorf("multiple end nodes within workflow")
			} else if nodeHandles.Source.GetBoolValue() != false {
				return fmt.Errorf("end node cannot be the source of another node")
			}
			containsEndNode = true
		}
		nodeIdSet[node.Id] = struct{}{}
	}

	if !containsStartNode || !containsEndNode {
		return fmt.Errorf("missing start or end nodes")
	}
	return nil
}
