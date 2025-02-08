package receipts

import (
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/slimus/fetch-task/internal/model"
)

type StorageInterface interface {
	SaveReciept(reciept *model.Reciept) (uuid.UUID, error)
	GetById(id string) (*model.Reciept, error)
}

type App struct {
	db StorageInterface
}

func NewApp(storage StorageInterface) *App {
	return &App{
		db: storage,
	}
}

func convertDateAndTimeToTime(dateString, timeString string) (*time.Time, error) {
	parsedTime, err := time.Parse("2006-01-02 15:04", fmt.Sprintf("%s %s", dateString, timeString))
	if err != nil {
		return nil, fmt.Errorf("can't convert date + time string to time: %w", err)
	}

	return &parsedTime, nil
}

func ConvertRequestRecieptToDB(requestReciept RecieptProcessRequest) (*model.Reciept, error) {
	items := make([]model.Item, 0)
	for _, requestItem := range requestReciept.Items {
		floatPrice, err := strconv.ParseFloat(requestItem.Price, 64)
		if err != nil {
			return nil, fmt.Errorf("Items.Price: Can't convert price from string to float: %w", err)
		}

		items = append(items, model.Item{
			ShortDescription: requestItem.ShortDescription,
			Price:            floatPrice,
		})
	}

	totalPrice, err := strconv.ParseFloat(requestReciept.Total, 64)
	if err != nil {
		return nil, fmt.Errorf("TotalPrice: Can't convert total_price from string to float: %w", err)
	}

	parsedTime, err := convertDateAndTimeToTime(requestReciept.PurchaseDate, requestReciept.PurchaseTime)
	if err != nil {
		return nil, err
	}

	return &model.Reciept{
		Retailer:     requestReciept.Retailer,
		PurchaseTime: *parsedTime,
		Items:        items,
		Total:        totalPrice,
	}, nil
}
