package go_stringable

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringable_StrPad(t *testing.T) {
	assert.Equal(t, "0001", StrPadLeft("1", rune('0'), 4))
	assert.Equal(t, "1000", StrPadRight("1", rune('0'), 4))

	assert.Equal(t, "0001", StrPadLeft("1", '0', 4))
	assert.Equal(t, "1000", StrPadRight("1", '0', 4))
}

func TestStringable_Md5(t *testing.T) {
	assert.Equal(t, "ebcd082fe9479bc278628297985ca2ba", Md5("Hyperf"))
}

func TestStringable_Random(t *testing.T) {
	res, _ := Random(6)
	assert.NotEmpty(t, res)
}
