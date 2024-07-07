package metamodel

import (
	"github.com/gin-gonic/gin"
	"graqler-metamodel-db/metamodel/nodes"
	"net/http"
)

var allNodeTypes = []nodes.NodeType{
	{
		Name:       "Employee",
		SampleName: "employee",
	},
	{
		Name:       "Organization",
		SampleName: "organization",
	},
}

var allEdgeTypes = []nodes.EdgeType{
	{
		Name: "IS_PART_OF",
	},
	{
		Name: "WORKS_FOR",
	},
}

func findAllNodeTypes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, allNodeTypes)
}

func findAllEdgeTypes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, allEdgeTypes)
}

func AddRoutes(router *gin.Engine) {
	router.POST("/queries/findAllNodeTypes", findAllNodeTypes)
	router.POST("/queries/findAllEdgeTypes", findAllEdgeTypes)
}
