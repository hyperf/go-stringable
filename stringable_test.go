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
	assert.Equal(t, "ebcd082fe9479bc278628297985ca2ba", Md5ByteArray([]byte("Hyperf")))
}

func TestStringable_Random(t *testing.T) {
	res, _ := Random(6)
	assert.NotEmpty(t, res)

	res, _ = RandomNumber(6)
	assert.NotEmpty(t, res)
}

func TestStringable_Repeat(t *testing.T) {
	assert.Equal(t, "****", Repeat("*", 4))
	assert.Equal(t, "****", Repeat("**", 2))
}

func TestStringable_Mask(t *testing.T) {
	assert.Equal(t, "***erf", Mask("Hyperf", 0, 3, "*"))
	assert.Equal(t, "H***rf", Mask("Hyperf", 1, 3, "*"))
	assert.Equal(t, "***erf", MaskByStar("Hyperf", 0, 3))
	assert.Equal(t, "Hyp***", MaskLastByStar("Hyperf", 3))
	assert.Equal(t, "Hy***f", Mask("Hyperf", -1, 3, "*"))
	assert.Equal(t, "***erf", Mask("Hyperf", -3, 0, "*"))
	assert.Equal(t, "******erf", Mask("Hyperf", -3, 6, "*"))
	assert.Equal(t, "你好**", Mask("你好世界", 2, 2, "*"))
}

func TestSnake(t *testing.T) {
	assert.Equal(t, "hyperf", Snake("Hyperf"))
	assert.Equal(t, "hello_world", Snake("HelloWorld"))
	assert.Equal(t, "hello_world", Snake("helloWorld"))
	assert.Equal(t, "hello_world", Snake("hello_world"))
}

func TestStudly(t *testing.T) {
	assert.Equal(t, "Hyperf", Studly("hyperf"))
	assert.Equal(t, "HelloWorld", Studly("hello_world"))
	assert.Equal(t, "HelloWoRld", Studly("hello_wo-rld"))
}
