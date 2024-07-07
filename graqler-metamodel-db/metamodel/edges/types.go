package edges

import "graqler-metamodel-db/metamodel/nodes"

type OutFrom struct {
	FromNode *nodes.NodeType `json:"fromNode"`
	ViaEdge  *nodes.EdgeType `json:"viaEdge"`
}

type InTo struct {
	ViaEdge *nodes.EdgeType `json:"viaEdge"`
	ToNode  *nodes.NodeType `json:"toNode"`
}
