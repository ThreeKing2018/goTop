package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdd(t *testing.T) {
	a := 10
	b := 20
	ret := Add(a, b)
	assert.Equal(t, ret, 30, "等于30") //判断ret与30是否相等, 如果不相等,则会报错

	a = 1
	b = 1
	ret = Add(a, b)
	assert.NotEqual(t, ret, 3) //判断ret与3是否不相等, 如果成立则OK, 反之
}

func TestDiv(t *testing.T) {
	a := 10
	b := 5
	ret, err := Div(a, b)
	assert.Nil(t, err)  //判断是否为空, 如果成立则OK, 反之
	assert.Equal(t, ret, 2, "相等")

	a = 10
	b = 0
	ret, err = Div(a, b)
	assert.NotNil(t, err) //判断是否不为空, 如果成立则OK, 反之
	assert.Zero(t, ret)
}

func TestSubset(t *testing.T) {
	assert.Subset(t, []int{1, 2, 3}, []int{1})
	assert.NotSubset(t,  []string{"hello", "world"}, []string{"hello", "testify"})
	assert.FileExists(t, "/var/aa")
}

func TestRequire(t *testing.T) {
	require.Equal(t, 1, 1)
	require.NotEqual(t, 1, 20)
	require.Contains(t, "Hello World", "Hello")
	require.NotContains(t, []string{"aa", "bb"}, []string{"bb", "aa"})
}
