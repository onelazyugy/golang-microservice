package docs

import "github.com/onelzyugy/projects/golang-microservice/types"

// swagger:route POST /order order bubbleTeaRequest
// oder bubble tea drinks.
// responses:
//   200: bubbleTeaResponse

// swagger:parameters bubbleTeaRequest
type bubbleTeaRequestWrapper struct {
	// in:body
	Body types.BubbleTeaRequest
}

// swagger:response bubbleTeaResponse
type bubbleTeaResponseWrapper struct {
	// in:body
	Body types.BubbleTeaResponse
}

// swagger:route GET /health-check health check
// health check of the app
// responses:
//	 200: healthResponse

// swagger:response healthResponse
type healthResponseWrapper struct {
	// in:body
	Body types.Health
}
