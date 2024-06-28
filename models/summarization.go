package models

import (
	"context"
	"os"

	"github.com/hupe1980/go-huggingface"
)

func Summary(input string) (*string, error) {
	hg := huggingface.NewInferenceClient(os.Getenv("HUGGING_FACE_TOKEN"))
	waitmodel := true
	res, err := hg.Summarization(context.Background(), &huggingface.SummarizationRequest{
		Inputs: []string{input},
		Options: huggingface.Options{
			WaitForModel: &waitmodel,
		},
		Model: "Falconsai/text_summarization",
	})
	if err != nil {
		return nil, err
	}
	return &res[0].SummaryText, nil
}
