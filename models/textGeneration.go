package models

import (
	"context"
	"os"

	"github.com/hupe1980/go-huggingface"
)

func TextGeneration(message string) (*string, error) {
	hg := huggingface.NewInferenceClient(os.Getenv("HUGGING_FACE_TOKEN"))
	waitmodel := true
	res, err := hg.TextGeneration(context.Background(), &huggingface.TextGenerationRequest{
		Inputs: message,
		Options: huggingface.Options{
			WaitForModel: &waitmodel,
		},
		Model: "meta-llama/Meta-Llama-3-8B-Instruct",
	})
	if err != nil {
		return nil, err
	}
	return &res[0].GeneratedText, nil
}
