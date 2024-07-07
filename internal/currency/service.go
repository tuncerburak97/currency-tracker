package currency

import (
	"currency-tracker/internal/config"
	"currency-tracker/internal/database"
	"fmt"
)

type Service struct {
	influxClient database.InfluxDB
	Config       *config.Config
}

func NewService() *Service {
	return &Service{
		influxClient: database.GetInfluxDBClient(),
		Config:       config.GetConfig(),
	}
}

func (s *Service) GetGoldRateByName(goldName string) RateResponse {
	databaseConfig := s.Config.Database.InfluxDB
	bucket := databaseConfig.Bucket

	query := fmt.Sprintf(`
      from(bucket: "%s")
        |> range(start: -1h)
        |> filter(fn: (r) => r["_measurement"] == "%s")
        |> filter(fn: (r) => r["_field"] == "buy_price")
        |> filter(fn: (r) => r["code"] == "%s")
        |> yield(name: "price_changes")
    `, bucket, databaseConfig.Measurement.Gold, goldName)

	queryResponse := s.influxClient.QueryData(databaseConfig.Org, bucket, query)

	var rates []Rate
	for queryResponse.Next() {
		record := queryResponse.Record()
		rates = append(rates, Rate{
			Price:     record.Value().(float64),
			Timestamp: record.Time(),
		})
	}
	if queryResponse.Err() != nil {
		return RateResponse{}
	}

	return RateResponse{
		Name:  goldName,
		Rates: rates,
	}

}

func (s *Service) GetCurrencyRateByName(currencyName string) RateResponse {

	databaseConfig := s.Config.Database.InfluxDB
	bucket := databaseConfig.Bucket

	query := fmt.Sprintf(`
	  from(bucket: "%s")
		|> range(start: -1h)
		|> filter(fn: (r) => r["_measurement"] == "%s")
		|> filter(fn: (r) => r["_field"] == "buy_price")
		|> filter(fn: (r) => r["code"] == "%s")
		|> yield(name: "price_changes")
	`, bucket, databaseConfig.Measurement.Currency, currencyName)

	queryResponse := s.influxClient.QueryData(databaseConfig.Org, bucket, query)

	var rates []Rate
	for queryResponse.Next() {
		record := queryResponse.Record()
		rates = append(rates, Rate{
			Price:     record.Value().(float64),
			Timestamp: record.Time(),
		})
	}
	if queryResponse.Err() != nil {
		return RateResponse{}
	}

	return RateResponse{
		Name:  currencyName,
		Rates: rates,
	}

}
