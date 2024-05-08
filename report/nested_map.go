package report

import (
	"fmt"

	ginkgotypes "github.com/onsi/ginkgo/v2/types"
)

type NestedContents struct {
	HasSkipped bool
	HasFailed  bool
	Contents   map[string]NestedContents
}

func newNestedContents() NestedContents {
	return NestedContents{
		HasSkipped: false,
		HasFailed:  false,
		Contents:   make(map[string]NestedContents),
	}
}

// BuildNestedContents recursively builds a map of nested maps from the given Ginkgo reports.
//
// This is used to convert the typed structure of the Ginkgo reports
// to a nested Markdown or HTML structure.
func BuildNestedContents(reports []ginkgotypes.Report) (NestedContents, error) {
	spec := newNestedContents()

	for _, report := range reports {
		for _, specReport := range report.SpecReports {
			containerHierarchy := specReport.ContainerHierarchyTexts

			cleanLeafNode := getCleanLeafNodeText(specReport)
			containerHierarchy = append(containerHierarchy, cleanLeafNode)

			currentNestingLevel := spec
			for _, item := range containerHierarchy {
				// If the item is not yet in the map, create a new map for the underlying levels
				// and add the item to the map.
				content, found := currentNestingLevel.Contents[item]
				if !found {
					content = newNestedContents()
					currentNestingLevel.Contents[item] = content
				}
				currentNestingLevel = content
			}
		}
	}

	return spec, nil
}

// getCleanLeafNodeText returns the clean leaf node text for the given Ginkgo spec report.
//
// Example: An It block with the text "should return the correct value" will be converted
// to "it should return the correct value".
func getCleanLeafNodeText(specReport ginkgotypes.SpecReport) string {
	var leafNodeTypeStr string

	leafNodeType := specReport.LeafNodeType
	switch leafNodeType {
	case ginkgotypes.NodeTypeIt:
		leafNodeTypeStr = "it"
	default:
		panic(fmt.Sprintf("unknown leaf node type: %d", leafNodeType))
	}

	leafNodeText := specReport.LeafNodeText
	cleanLeafNode := leafNodeTypeStr + " " + leafNodeText

	return cleanLeafNode
}
