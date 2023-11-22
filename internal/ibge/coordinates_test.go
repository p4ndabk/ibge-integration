package ibge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCoordinatesIBGE(t *testing.T) {
	coordinates, err := GetCoordinatesIBGE("4205407")
	assert.Nil(t, err)
	assert.NotNil(t, coordinates)
}
