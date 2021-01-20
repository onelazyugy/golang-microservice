package services

import (
	"errors"
	"fmt"

	"github.com/onelzyugy/projects/golang-microservice/types"
)

// OrderBubbleTea allows users to order a bubble tea drink
func OrderBubbleTea(b types.BubbleTeaRequest) error {
	fmt.Println("bubble tea: ", b)
	//check if bubble tea request is valid or not here
	return errors.New("unable to order bubble tea")
}
