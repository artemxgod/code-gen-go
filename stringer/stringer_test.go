package stringer_test

import (
	"testing"

	"github.com/artemxgod/code-gen-go/stringer"
	"github.com/stretchr/testify/assert"
)

func TestStringerWindows(t *testing.T) {
	testCases := []struct {
		name               string
		opsystem           stringer.OS
		expectedSystemName string
	}{
		{
			name:               "windows",
			opsystem:           0,
			expectedSystemName: "Windows",
		},
		{
			name:               "darwin",
			opsystem:           1,
			expectedSystemName: "Darwin",
		},
		{
			name:               "linux",
			opsystem:           2,
			expectedSystemName: "Linux",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got_name := tc.opsystem.String()
			assert.Equal(t, tc.expectedSystemName, got_name)
		})
	}

}
