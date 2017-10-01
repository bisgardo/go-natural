package strcmp

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__sort_versions(t *testing.T) {
	ss := []string{
		"1.2.3",
		"0.2.3",
		"1.10.3",
		"1.2.10",
	}
	sort.Sort(Naturally(ss))

	expected := []string{
		"0.2.3",
		"1.2.3",
		"1.2.10",
		"1.10.3",
	}

	assert.Equal(t, expected, ss)
}
