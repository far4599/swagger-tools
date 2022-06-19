package merger

import (
	"testing"

	"github.com/go-openapi/spec"
	"github.com/stretchr/testify/require"
)

func Test_uniqueTags(t *testing.T) {
	testCases := []struct {
		input, expect []string
	}{
		{input: []string{"1", "1", "1"}, expect: []string{"1"}},
		{input: []string{"1", "2", "1"}, expect: []string{"1", "2"}},
		{input: []string{"1", "2", "3"}, expect: []string{"1", "2", "3"}},
		{input: []string{}, expect: []string{}},
		{input: nil, expect: nil},
	}

	t.Parallel()

	for i := range testCases {
		tc := testCases[i]
		t.Run("", func(t *testing.T) {
			input := _convertStringsToTags(tc.input)
			expect := _convertStringsToTags(tc.expect)

			result := uniqueTags(input)

			require.ElementsMatch(t, expect, result)

			if len(expect) == 0 {
				require.Equal(t, expect, result)
			}
		})
	}
}

func Test_uniqueStrings(t *testing.T) {
	testCases := []struct {
		input, expect []string
	}{
		{input: []string{"1", "1", "1"}, expect: []string{"1"}},
		{input: []string{"1", "2", "1"}, expect: []string{"1", "2"}},
		{input: []string{"1", "2", "3"}, expect: []string{"1", "2", "3"}},
		{input: []string{}, expect: []string{}},
		{input: nil, expect: nil},
	}

	t.Parallel()

	for i := range testCases {
		tc := testCases[i]
		t.Run("", func(t *testing.T) {
			result := uniqueStrings(tc.input)

			require.ElementsMatch(t, tc.expect, result)

			if len(tc.expect) == 0 {
				require.Equal(t, tc.expect, result)
			}
		})
	}
}

func _convertStringsToTags(input []string) []spec.Tag {
	if input == nil {
		return nil
	}

	result := make([]spec.Tag, 0, len(input))
	for i := range input {
		tag := spec.Tag{}
		tag.Name = input[i]
		result = append(result, tag)
	}

	return result
}
