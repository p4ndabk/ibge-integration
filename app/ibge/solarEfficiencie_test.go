package ibge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolarEfficiencieByCode(t *testing.T) {
	efficiencie, err := EfficiencieByIBGECode(4205407)
	assert.Nil(t, err)
	assert.NotNil(t, efficiencie)
	assert.Equal(t, "Brasil", efficiencie.Country)
	assert.Equal(t, 4262, efficiencie.Annual)
}