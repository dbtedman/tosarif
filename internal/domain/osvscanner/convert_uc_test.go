package osvscanner_test

import (
	"testing"

	"github.com/dbtedman/tosarif/internal/domain/osvscanner"
	"github.com/stretchr/testify/assert"
)

func TestOSVScannerUCConvertAcceptsValidSource(t *testing.T) {
	source := osvscanner.Source{}
	uc := osvscanner.ConvertUseCase{}

	result, resultError := uc.Convert(&source)

	assert.NotNil(t, result)
	assert.Nil(t, resultError)
}

func TestOSVScannerUCConvertRejectsNilSource(t *testing.T) {
	uc := osvscanner.ConvertUseCase{}

	result, resultError := uc.Convert(nil)

	assert.Nil(t, result)
	assert.ErrorIs(t, resultError, osvscanner.ErrorNilSource)
}
