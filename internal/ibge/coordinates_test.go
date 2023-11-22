package ibge

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetCoordinatesIBGE(t *testing.T) {
	coordinates, err := GetCoordinatesIBGE("4205407")
	assert.Nil(t, err)
	assert.NotNil(t, coordinates)
}
