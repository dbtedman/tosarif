package osvscanner

import (
	"github.com/dbtedman/tosarif/internal/domain/sarif"
)

type ConvertUseCase struct{}

const OSVScanner = "osv-scanner"

func (my *ConvertUseCase) Convert(source *Source) (*sarif.Root, error) {
	if source == nil {
		return nil, ErrorNilSource
	}

	if err := source.Validate(); err != nil {
		return nil, err
	}

	return &sarif.Root{
		Runs: []sarif.Run{
			{
				Tool: sarif.Tool{
					Driver: sarif.ToolComponent{
						Name: OSVScanner,
					},
				},
			},
		},
		Schema:  sarif.CurrentSchema,
		Version: sarif.CurrentVersion,
	}, nil
}
