package osvscanner

import (
	"github.com/dbtedman/tosarif/internal/domain/sarif"
)

type ConvertUseCase struct{}

func (my *ConvertUseCase) Convert(source *Source) (*sarif.Sarif, error) {
	if source == nil {
		return nil, ErrorNilSource
	}

	if err := source.Validate(); err != nil {
		return nil, err
	}

	return &sarif.Sarif{}, nil
}
