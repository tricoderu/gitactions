// main_test.go
// go get -u github.com/stretchr/testify
package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxInt(t *testing.T) {
	a, b := 2, 7
	res := MaxInt(a, b)

	// Используя assert
	assert.Equal(t, b, res, "Expected %d, but got %d", b, res)

	// Или использовать require для завершения теста при ошибке
	// require.Equal(t, b, res, "Expected %d, but got %d", b, res)
}

// функция не содержит никакой логики, удаляю
/*func TestMain(m *testing.M) {
	main()
}*/
