package services

import (
	"fmt"

	"github.com/onelazyugy/golang-microservice/types"
)

var orderedBubbleTea []types.BubbleTeaRequest

// OrderBubbleTea allows users to order a bubble tea drink
func OrderBubbleTea(b types.BubbleTeaRequest) (*types.BubbleTeaResponse, error) { //returning *types.BubbleTeaResponse which is a pointer to that type
	fmt.Println("bubble tea: ", b)
	orderedBubbleTea = append(orderedBubbleTea, b)
	response := &types.BubbleTeaResponse{ //when creating this, a memory slot is allocated for it
		StatusCd: 200,
		Success:  true,
		Message:  "successfully process bubble tea order",
	}
	return response, nil
}

// RetrieveOrderedBubbleTea allows user to retrieve all of the ordered bubble teas
func RetrieveOrderedBubbleTea() (types.RetrieveBubbleTeaResponse, error) {
	status := types.BubbleTeaResponse{
		StatusCd: 200,
		Message:  "successfully retreived all bubble teas",
		Success:  true,
	}
	res := types.RetrieveBubbleTeaResponse{
		OrderedBubbleTeas: orderedBubbleTea,
		BubbleTeaResponse: status,
	}
	return res, nil
}
