package rewards

import (
	"math"
	"strings"
	"time"
	"unicode"

	"github.com/slimus/fetch-task/internal/model"
)

// TODO: this function should be refactored to be more readable and maintainable
// Maybe we should have slice of functions that calculate rewards and iterate over them
func CalculateRewards(reciept *model.Reciept) int {
	rewards := 0

	rewards += calculateRetailerNameRewards(reciept.Retailer)
	rewards += calculateTotalRewards(reciept.Total)
	rewards += calculateItemRewards(reciept.Items)
	rewards += calculatePurchaseDateRewards(reciept.PurchaseTime)
	return rewards
}

func calculateRetailerNameRewards(retailerName string) int {
	rewards := 0
	for _, char := range retailerName {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			rewards++
		}
	}
	return rewards
}

func calculateTotalRewards(total float64) int {
	rewards := 0
	if total == float64(int(total)) {
		rewards += 50
	}
	if int(total*100)%25 == 0 {
		rewards += 25
	}
	return rewards
}

func calculateItemRewards(items []model.Item) int {
	rewards := 0

	rewards += len(items) / 2 * 5

	for _, item := range items {
		trimmedLength := len(strings.Trim(item.ShortDescription, " "))
		if trimmedLength%3 == 0 {
			rewards += int(math.Ceil(item.Price * 0.2))
		}
	}

	return rewards
}

func calculatePurchaseDateRewards(purchaseTime time.Time) int {
	rewards := 0

	if purchaseTime.Day()%2 != 0 {
		rewards += 6
	}

	hour := purchaseTime.Hour()
	if hour >= 14 && hour < 16 {
		rewards += 10
	}

	return rewards
}
