package rewards

import (
	"testing"
	"time"

	"github.com/slimus/fetch-task/internal/model"
)

func TestCalculateRewardsExample1(t *testing.T) {
	reciept := model.Reciept{
		Retailer:     "Target",
		PurchaseTime: time.Date(2022, 01, 01, 13, 01, 0, 0, time.UTC),
		Items: []model.Item{
			{
				ShortDescription: "Mountain Dew 12PK",
				Price:            6.49,
			},
			{
				ShortDescription: "Emils Cheese Pizza",
				Price:            12.25,
			},
			{
				ShortDescription: "Knorr Creamy Chicken",
				Price:            1.26,
			},
			{
				ShortDescription: "Doritos Nacho Cheese",
				Price:            3.35,
			},
			{
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            12.00,
			},
		},
		Total: 35.35,
	}

	expected := 28
	actual := CalculateRewards(&reciept)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestCalculateRewardsExample2(t *testing.T) {
	reciept := model.Reciept{
		Retailer:     "M&M Corner Market",
		PurchaseTime: time.Date(2022, 03, 20, 14, 33, 0, 0, time.UTC),
		Items: []model.Item{
			{
				ShortDescription: "Gatorade",
				Price:            2.25,
			},
			{
				ShortDescription: "Gatorade",
				Price:            2.25,
			},
			{
				ShortDescription: "Gatorade",
				Price:            2.25,
			},
			{
				ShortDescription: "Gatorade",
				Price:            2.25,
			},
		},
		Total: 9.00,
	}

	expected := 109
	actual := CalculateRewards(&reciept)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
