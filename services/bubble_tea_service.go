package services

import (
	"fmt"

	"github.com/onelzyugy/projects/golang-microservice/types"
)

// OrderBubbleTea allows users to order a bubble tea drink
func OrderBubbleTea(b types.BubbleTeaRequest) error {
	fmt.Println("bubble tea: ", b)
	//check if bubble tea request is valid or not here
	return nil
}
