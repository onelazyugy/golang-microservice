package services

import (
	"fmt"

	"github.com/onelzyugy/projects/golang-microservice/types"
)

var birds []types.Bird

// RetrieveBirds add a bird
func RetrieveBirds(birds []types.Bird) []types.Bird {
	for _, b := range birds {
		fmt.Println(b)
	}
	return birds
}

// AddBird get all birds
func AddBird() {

}
