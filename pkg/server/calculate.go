package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	apiv1 "github.com/tmeshorer/volume/pkg/api/v1"
	"net/http"
)

func (s *Server) Calculate(c *gin.Context) {
	var req *apiv1.CalculateRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, apiv1.ErrResp("could not parse calculate request"))
		return
	}

	// validate the flight format, assure that each flight has source and dest
	for _, v := range req.Flights {
		if len(v) != 2 {
			c.Error(errors.Errorf("invalid flight format, flight must be of [source,destination] %q", v))
			c.JSON(http.StatusBadRequest, apiv1.ErrResp("invalid request format"))
		}

	}

	start, end, err := calculateRoute(req.Flights)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, apiv1.ErrResp("failed to calculate routes"))
		return
	}

	res := &apiv1.CalculateReply{[]string{start, end}}

	c.JSON(http.StatusOK, res)
}

// calculateRoute the actual route calculation. this is the real alg
func calculateRoute(legs [][]string) (string, string, error) {
	// represent the flight legs as graph
	vertices := map[string][]string{}
	// populate the vertices
	for _, v := range legs {
		_, ok := vertices[v[0]]
		// if the vertice does not exist,add it
		if !ok {
			var edges []string
			edges = append(edges, v[1])
			vertices[v[0]] = edges
			// if the vertices exist
		} else {
			edges := vertices[v[0]]
			edges = append(edges, v[1])
			vertices[v[0]] = edges
		}
		// also create an entry for the To node.
		_, ok = vertices[v[1]]
		if !ok {
			vertices[v[1]] = make([]string, 0)
		}

	}

	// the end node is a node in the graph that has no outgoing nodes
	endNode := ""
	for key, _ := range vertices {
		// if there are no exit nodes.
		if len(vertices[key]) == 0 { // no outgoing legs
			endNode = key
			break
		}
	}

	// to find the start node, we will try to visit all the nodes in the graph, the node that was not visited,
	// is the first node.
	visited := make(map[string]bool)

	// first fill the visited array
	for key, _ := range vertices {
		visited[key] = false
	}

	for _, nodes := range vertices {
		// if there are no exit nodes.
		for _, v := range nodes {
			visited[v] = true
		}
	}

	// the start node is the first node that is not visited
	startNode := ""
	for key, v := range visited {
		if !v {
			startNode = key
			break
		}
	}
	return startNode, endNode, nil
}
