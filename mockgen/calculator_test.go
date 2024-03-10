package mockgen_test

import (
	"testing"

	"github.com/artemxgod/code-gen-go/mockgen"
	"go.uber.org/mock/gomock"
)

func TestComputer_AddAndSubtract(t *testing.T) {
    // Create a new mock controller
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    // Create a mock calculator
    mockCalc := mockgen.NewMockCalculator(ctrl)

    // Set expectations
    mockCalc.EXPECT().Add(3, 2).Return(5)
    mockCalc.EXPECT().Subtract(3, 2).Return(1)

    // Create a computer with the mock calculator
    computer := mockgen.NewComputer(mockCalc)

    // Call the method being tested
    sum, difference := computer.AddAndSubtract(3, 2)

    // Check the results
    if sum != 5 {
        t.Errorf("Expected sum to be 5, got %d instead", sum)
    }
    if difference != 1 {
        t.Errorf("Expected difference to be 1, got %d instead", difference)
    }
}