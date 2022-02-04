package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClock_NowString_ShouldFormatAsExpected(t *testing.T) {
	// Arrange
	expectedDate := "13/01/2012"
	now := func() time.Time {
		t, err := time.Parse("02/01/2006", expectedDate)
		if err != nil {
			panic(err)
		}
		return t
	}
	clock := NewClockWithNowFunc(now)

	// Act
	date := clock.NowString()

	// Assert
	require.Equal(t, expectedDate, date)
}
