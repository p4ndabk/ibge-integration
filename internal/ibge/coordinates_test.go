package ibge

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCoordinatesIBGE(t *testing.T) {
	coordinates, err := GetCoordinatesIBGE("4205407")

	fmt.Println(coordinates)
	assert.Nil(t, err)
	assert.NotNil(t, coordinates)
}
