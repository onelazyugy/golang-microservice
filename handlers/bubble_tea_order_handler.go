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
	bubbleTeaResponse := &types.BubbleTeaResponse{}
	bubbleTeaRequest, unMarshallError := unMarshallToBubbleTeaRequestType(r)
	if unMarshallError != nil {
		error := fmt.Errorf("Error: %v", unMarshallError)
		fmt.Println(error)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}
	serviceError := services.OrderBubbleTea(bubbleTeaRequest)
	if serviceError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(serviceError.Error()))
		return
	}
	createBubbleTeaResponse(bubbleTeaResponse, "success", true, http.StatusOK)
	bubbleTeaResponseBytes, marshallBubbleTeaResponseError := marshallBubbleTeaResponseType(bubbleTeaResponse)
	if marshallBubbleTeaResponseError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(marshallBubbleTeaResponseError.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bubbleTeaResponseBytes)
}

func createBubbleTeaResponse(bubbleTeaResponse *types.BubbleTeaResponse, message string, success bool, statusCd int) {
	bubbleTeaResponse.Message = message
	bubbleTeaResponse.Success = success
	bubbleTeaResponse.StatusCd = statusCd
}

func unMarshallToBubbleTeaRequestType(r *http.Request) (types.BubbleTeaRequest, error) {
	var bubbleTea types.BubbleTeaRequest
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return bubbleTea, fmt.Errorf("Unable to read request body: %w", err)
	}
	if err := r.Body.Close(); err != nil {
		return bubbleTea, fmt.Errorf("Unable to close request body: %w", err)
	}
	if err := json.Unmarshal(body, &bubbleTea); err != nil {
		return bubbleTea, fmt.Errorf("Unable to convert request to BubbleTea: %w", err)
	}
	return bubbleTea, nil
}

func marshallBubbleTeaResponseType(bubbleTeaResponse *types.BubbleTeaResponse) ([]byte, error) {
	bubbleTeaResponseBytes, err := json.Marshal(bubbleTeaResponse)
	if err != nil {
		return bubbleTeaResponseBytes, errors.New("Unable to convert response to bytes")
	}
	return bubbleTeaResponseBytes, nil
}
