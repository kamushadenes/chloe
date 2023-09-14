package cost

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCategoryCost(t *testing.T) {
	// Test getting a category for an existing category
	existingCategory := "completion"
	expectedCost := 0.0

	actualCost := GetCategoryCost(existingCategory)
	assert.Equal(t, expectedCost, actualCost)

	// Test getting a category for a non-existing category
	nonExistingCategory := "non-existing"
	nonExistentExpectedCost := 0.0

	actualCost = GetCategoryCost(nonExistingCategory)
	assert.Equal(t, nonExistentExpectedCost, actualCost)
}

func TestGetCategories(t *testing.T) {
	cat := GetCategories()
	assert.NotNil(t, cat)
}

func TestGetTotalSessionCost(t *testing.T) {
	// Initialize all  to 0
	ResetSessionCost()

	// Set  for a few categories
	AddCategoryCost("completion", 10)
	AddCategoryCost("transcription", 5)
	AddCategoryCost("moderation", 12)

	expectedTotal := 27.0
	actualTotal := GetTotalSessionCost()

	assert.Equal(t, expectedTotal, actualTotal)
}

func TestAddCategoryCost(t *testing.T) {
	ResetSessionCost()

	// Add for an existing category
	AddCategoryCost("completion", 5)

	expectedCost := 5.0
	actualCost := GetCategoryCost("completion")
	assert.Equal(t, expectedCost, actualCost)

	// Add for a non-existing category
	AddCategoryCost("non-existing", 10)

	nonExistentExpectedCost := 10.0
	actualCost = GetCategoryCost("non-existing")
	assert.Equal(t, nonExistentExpectedCost, actualCost)
}

func TestResetSessionCost(t *testing.T) {
	// Add some  to the sessionCost map
	AddCategoryCost("completion", 10)
	AddCategoryCost("transcription", 5)
	AddCategoryCost("moderation", 12)

	// Verify the map is not empty
	assert.NotEmpty(t, GetCategories())

	// Reset the map
	ResetSessionCost()

	// Verify the map is empty
	assert.Empty(t, GetCategories())
}

func TestSubtractCategoryCost(t *testing.T) {
	ResetSessionCost()

	// Set  for a few categories
	AddCategoryCost("completion", 10)
	AddCategoryCost("transcription", 5)
	AddCategoryCost("moderation", 12)

	// Subtract for an existing category
	SubtractCategoryCost("completion", 5)

	expectedCost := 5.0
	actualCost := GetCategoryCost("completion")
	assert.Equal(t, expectedCost, actualCost)

	// Subtract for a non-existing category (should not go negative)
	SubtractCategoryCost("transcription", 7)

	nonExistentExpectedCost := -2.0
	actualCost = GetCategoryCost("transcription")
	assert.Equal(t, nonExistentExpectedCost, actualCost)
}
