package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/onelzyugy/projects/golang-microservice/services"
	"github.com/onelzyugy/projects/golang-microservice/types"
)

// OrderBubbleTeaHandler allows you to order bubble tea drink
func OrderBubbleTeaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// var bubbleTeaResponse types.BubbleTeaResponse
	bubbleTeaRequest, unMarshallError := unMarshallToBubbleTeaRequestType(r)
	if unMarshallError != nil {
		//form error response
	}
	//make call to service
	serviceError := services.OrderBubbleTea(bubbleTeaRequest)
	if serviceError != nil {
		//form error response
	}
	//return proper response to api call
	bubbleTeaResponse := types.BubbleTeaResponse{
		Success:  true,
		StatusCd: 200,
		Message:  "success",
	}
	bubbleTeaResponseBytes, marshallBubbleTeaResponseError := marshallBubbleTeaResponseType(bubbleTeaResponse)
	if marshallBubbleTeaResponseError != nil {
		//form error response
	}
	w.Write(bubbleTeaResponseBytes)
}

func unMarshallToBubbleTeaRequestType(r *http.Request) (types.BubbleTeaRequest, error) {
	var bubbleTeam types.BubbleTeaRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return bubbleTeam, fmt.Errorf("Unable to read request body: %w", err)
	}
	if err := r.Body.Close(); err != nil {
		return bubbleTeam, fmt.Errorf("Unable to close request body: %w", err)
	}
	if err := json.Unmarshal(body, &bubbleTeam); err != nil {
		return bubbleTeam, fmt.Errorf("Unable to convert request to BubbleTea: %w", err)
	}
	return bubbleTeam, nil
}

func marshallBubbleTeaResponseType(bubbleTeaResponse types.BubbleTeaResponse) ([]byte, error) {
	bubbleTeaResponseBytes, err := json.Marshal(bubbleTeaResponse)
	if err != nil {
		return bubbleTeaResponseBytes, errors.New("Unable to convert response to bytes")
	}
	return bubbleTeaResponseBytes, nil
}
