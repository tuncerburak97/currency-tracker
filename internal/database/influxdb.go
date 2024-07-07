package database

import (
	"context"
	"currency-tracker/internal/config"
	"currency-tracker/pkg/logger"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
	"time"
)

type InfluxDB interface {
	WriteData(org, bucket, measurement string, tags map[string]string, fields map[string]interface{}) error
	QueryData(org, bucket, query string) *api.QueryTableResult
}

type InfluxDBClient struct {
	client influxdb2.Client
}

func GetInfluxDBClient() InfluxDB {
	return &InfluxDBClient{client: client}
}

var client influxdb2.Client

func InitInfluxDatabase() {
	influxConfig := config.GetConfig().Database.InfluxDB
	client = influxdb2.NewClient(influxConfig.URL, influxConfig.Token)
	logger.Logger.Info("Connected to InfluxDB")
}

func (s *InfluxDBClient) WriteData(org, bucket, measurement string, tags map[string]string, fields map[string]interface{}) error {
	writeAPI := s.client.WriteAPIBlocking(org, bucket)
	point := write.NewPoint(measurement, tags, fields, time.Now())
	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		return err
	}
	return nil
}

func (s *InfluxDBClient) QueryData(org, bucket, query string) *api.QueryTableResult {
	queryAPI := s.client.QueryAPI(org)
	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	return results
}

func Close() {
	client.Close()
}
