package models

import (
	"context"
	"os"

	"github.com/hupe1980/go-huggingface"
)

func TableQA(query string, table map[string][]string) (*huggingface.TableQuestionAnsweringResponse, error) {
	hf := huggingface.NewInferenceClient(os.Getenv("HUGGING_FACE_TOKEN"))
	res, err := hf.TableQuestionAnswering(context.Background(), &huggingface.TableQuestionAnsweringRequest{
		Inputs: huggingface.TableQuestionAnsweringInputs{
			Query: query,
			Table: table,
		},
	})
	if err != nil {
		return nil, err
	}
	return res, nil
}
