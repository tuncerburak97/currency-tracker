package currency

import (
	"currency-tracker/internal/config"
	"currency-tracker/internal/database"
	"currency-tracker/pkg/logger"
	"github.com/robfig/cron/v3"
	"strconv"
	"strings"
)

type Scheduler struct {
	currencyClient *Client
	cnf            *config.Config
	influxClient   database.InfluxDB
	cron           *cron.Cron
}

var schedulerInstance *Scheduler

func GetCurrencyScheduler() *Scheduler {
	if schedulerInstance == nil {

		schedulerInstance = &Scheduler{
			currencyClient: GetCurrencyClient(),
			cnf:            config.GetConfig(),
			influxClient:   database.GetInfluxDBClient(),
			cron:           cron.New(),
		}
	}
	return schedulerInstance
}

func InitScheduler() {
	schedulerInstance = GetCurrencyScheduler()
	schedulerInstance.StartScheduler()
}

func (s *Scheduler) StartScheduler() {
	schedule := config.GetConfig().Scheduler.Currency.Interval
	logger.Logger.Infof("Starting scheduler with schedule: %s", schedule)
	_, err := s.cron.AddFunc(schedule, s.fetchAndProcessData)
	if err != nil {
		logger.Logger.Errorf("Could not add function to scheduler: %v", err)
		return
	}

	s.cron.Start()
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
}

func (s *Scheduler) fetchAndProcessData() {
	go s.fetchAndWriteGoldPrices()
	go s.fetchAndWriteCurrencyPrices()
}

func (s *Scheduler) fetchAndWriteGoldPrices() {
	goldPrices, err := s.currencyClient.FetchGoldPrices()
	if err != nil {
		logger.Logger.Errorf("Could not fetch gold prices: %v", err)
		return
	}

	org := s.cnf.Database.InfluxDB.Org
	bucket := s.cnf.Database.InfluxDB.Bucket

	for _, price := range goldPrices {

		buyParsedValue, err := parseGoldValue(price.BuyPrice)
		if err != nil {
			logger.Logger.Errorf("Could not parse gold buy price: %v", err)
			continue
		}

		sellParsedValue, err := parseGoldValue(price.SellPrice)
		if err != nil {
			logger.Logger.Errorf("Could not parse gold sell price: %v", err)
			continue
		}

		tags := map[string]string{
			"code": price.MobileDescription,
		}
		fields := map[string]interface{}{
			"buy_price":  buyParsedValue,
			"sell_price": sellParsedValue,
		}

		goldMeasurement := s.cnf.Database.InfluxDB.Measurement.Gold

		if err := s.influxClient.WriteData(org, bucket, goldMeasurement, tags, fields); err != nil {
			logger.Logger.Errorf("Could not write gold price to InfluxDB: %v", err)
		}

	}
}

func (s *Scheduler) fetchAndWriteCurrencyPrices() {
	currencyPrices, err := s.currencyClient.FetchCurrencyPrices()
	if err != nil {
		logger.Logger.Errorf("Could not fetch currency prices: %v", err)
		return
	}

	for _, price := range currencyPrices {

		parsedBuyValue, err := parseCurrencyValue(price.BuyPrice)
		if err != nil {
			continue
		}

		parsedSellValue, err := parseCurrencyValue(price.SellPrice)
		if err != nil {
			continue
		}

		tags := map[string]string{
			"code": price.MobileDescription,
		}
		fields := map[string]interface{}{
			"buy_price":  parsedBuyValue,
			"sell_price": parsedSellValue,
		}

		org := s.cnf.Database.InfluxDB.Org
		bucket := s.cnf.Database.InfluxDB.Bucket
		currencyMeasurement := s.cnf.Database.InfluxDB.Measurement.Currency

		if err := s.influxClient.WriteData(org, bucket, currencyMeasurement, tags, fields); err != nil {
			logger.Logger.Errorf("Could not write currency price to InfluxDB: %v", err)
		}
	}
}

func parseGoldValue(value string) (float64, error) {
	value = strings.Replace(value, ".", "", -1)
	value = strings.Replace(value, ",", ".", 1)
	return strconv.ParseFloat(value, 64)
}

func parseCurrencyValue(value string) (float64, error) {
	value = strings.Replace(value, ",", ".", 1)
	return strconv.ParseFloat(value, 64)
}
