package stark

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFixMessage(t *testing.T) {
	fixed, _ := fixMessage("0x123456789abcdef")
	assert.Equal(t, "123456789abcdef", fixed)

	fixed, _ = fixMessage("0x074180eaec7e68712b5a0fbf5d63a70c33940c9b02e60565e36f84d705b669e")
	assert.Equal(t, "074180eaec7e68712b5a0fbf5d63a70c33940c9b02e60565e36f84d705b669e0", fixed)

	fixed, _ = fixMessage("074180eaec7e68712b5a0fbf5d63a70c33940c9b02e60565e36f84d705b669e")
	assert.Equal(t, "074180eaec7e68712b5a0fbf5d63a70c33940c9b02e60565e36f84d705b669e0", fixed)
}
