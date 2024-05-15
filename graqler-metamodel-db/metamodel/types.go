package metamodel

type NodeType struct {
	Name       string `json:"name"`
	SampleName string `json:"sampleName"`
}

type EdgeType struct {
	Name string    `json:"name"`
	From *NodeType `json:"from"`
	To   *NodeType `json:"to"`
}
