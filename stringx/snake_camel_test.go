package stringx_test

import (
	"github.com/go-leo/gox/stringx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCamel2Snake(t *testing.T) {
	assert.Equal(t, "", stringx.Camel2Snake(""))
	assert.Equal(t, "a", stringx.Camel2Snake("a"))
	assert.Equal(t, "ab", stringx.Camel2Snake("ab"))
	assert.Equal(t, "ab_f", stringx.Camel2Snake("abF"))
	assert.Equal(t, "ab_f", stringx.Camel2Snake("ab_F"))
	assert.Equal(t, "ab_f-_f", stringx.Camel2Snake("ab_F-F"))
	assert.Equal(t, "ab_f_f", stringx.Camel2Snake("abFF"))
	assert.Equal(t, "ab_ff_ffe_d_d", stringx.Camel2Snake("abFfFfeDD"))
	assert.Equal(t, "ab+_ff-_ffe_d_d", stringx.Camel2Snake("ab+Ff-Ffe_DD"))
	assert.Equal(t, "ab+_ff-_ffe_d_d", stringx.Camel2Snake("Ab+Ff-Ffe_DD"))
}

func TestSnake2Camel(t *testing.T) {
	assert.Equal(t, "", stringx.Snake2Camel("", true))
	assert.Equal(t, "A", stringx.Snake2Camel("a", true))
	assert.Equal(t, "a", stringx.Snake2Camel("a", false))
	assert.Equal(t, "Ab", stringx.Snake2Camel("ab", true))
	assert.Equal(t, "ab", stringx.Snake2Camel("ab", false))
	assert.Equal(t, "AbF", stringx.Snake2Camel("ab_f", true))
	assert.Equal(t, "abF", stringx.Snake2Camel("ab_f", false))
	assert.Equal(t, "AbF", stringx.Snake2Camel("ab_F", true))
	assert.Equal(t, "abF", stringx.Snake2Camel("ab_F", false))
	assert.Equal(t, "AbFF", stringx.Snake2Camel("ab_f_f", true))
	assert.Equal(t, "abFF", stringx.Snake2Camel("ab_f_f", false))
	assert.Equal(t, "AbFf", stringx.Snake2Camel("ab_ff", true))
	assert.Equal(t, "abFf", stringx.Snake2Camel("ab_ff", false))
	assert.Equal(t, "AbFfFfeDd", stringx.Snake2Camel("ab_ff_ffe_dd", true))
	assert.Equal(t, "abFfFfeDd", stringx.Snake2Camel("ab_ff_ffe_dd", false))
	assert.Equal(t, "Ab+Ff-FfeDD", stringx.Snake2Camel("ab+_ff-_ffe_d_d", true))
	assert.Equal(t, "ab+Ff-FfeDD", stringx.Snake2Camel("ab+_ff-_ffe_d_d", false))
	assert.Equal(t, "Ab+Ff-FfeDD", stringx.Snake2Camel("_ab+_ff-_ffe_d_d", true))
	assert.Equal(t, "ab+Ff-FfeDD", stringx.Snake2Camel("_ab+_ff-_ffe_d_d", false))
	assert.Equal(t, "Ab+Ff-FfeDD", stringx.Snake2Camel("_ab+_ff-_ffe_d_d_", true))
	assert.Equal(t, "ab+Ff-FfeDD", stringx.Snake2Camel("_ab+_ff-_ffe_d_d_", false))
}
