package metamodel

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var allNodeTypes = []NodeType{
	{
		Name:       "Employee",
		SampleName: "employee",
	},
	{
		Name:       "Organization",
		SampleName: "organization",
	},
}

var allEdgeTypes = []EdgeType{
	{
		Name: "IS-PART-OF",
		From: &allNodeTypes[1],
		To:   &allNodeTypes[1],
	},
	{
		Name: "WORKS-FOR",
		From: &allNodeTypes[0],
		To:   &allNodeTypes[1],
	},
}

func findAllNodeTypes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, allNodeTypes)
}

func findAllEdgeTypes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, allEdgeTypes)
}

type JsonRpcRequest struct {
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
	Id      string `json:"id"`
}

func HandleRpc(c *gin.Context) {
	var rpcRequest JsonRpcRequest

	// Bind the incoming request
	if err := c.BindJSON(&rpcRequest); err != nil {
		errInvalidJson(c, err.Error())
		return
	}

	switch rpcRequest.Method {
	case "findAllNodeTypes":
		findAllNodeTypes(c)
	case "findAllEdgeTypes":
		findAllEdgeTypes(c)
	default:
		errUnknownMethod(c, rpcRequest.Method, rpcRequest.Id)
	}

}

func errUnknownMethod(c *gin.Context, method string, id string) {
	var msg = "Unknown method: " + method
	_ = c.Error(errors.New(msg))
	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"jsonrpc": "2.0",
		"error": gin.H{
			"code":    2,
			"message": msg,
		},
		"id": id,
	})
}

func errInvalidJson(c *gin.Context, errorMsg string) {
	var msg = "Invalid JSON: " + errorMsg
	_ = c.Error(errors.New(msg))
	c.IndentedJSON(http.StatusBadRequest, gin.H{
		"jsonrpc": "2.0",
		"error": gin.H{
			"code":    1,
			"message": msg,
		},
		"id": nil,
	})
}
