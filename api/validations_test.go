package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidQueryParams(t *testing.T) {
	assert := assert.New(t)
	queryParmasBad := "bad query parmams!"
	queryParmasGood := "technology"

	assert.False(IsValidQueryParams(queryParmasBad))
	assert.False(IsValidQueryParams(queryParmasBad, queryParmasGood))
	assert.True(IsValidQueryParams(queryParmasGood))
}
