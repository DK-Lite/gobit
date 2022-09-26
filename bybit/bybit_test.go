package bybit

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/DK-Lite/gobit/config"
	"github.com/stretchr/testify/assert"
)

var envLocal *config.Config

func init() {
	fmt.Println("Test Start")
}

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	teardown()
	os.Exit(exitCode)
}

func setup() {
	envLocal, _ = config.LoadConfig("../config/config.yaml")
}

func teardown() {
	fmt.Println("teardown...")
}

func TestGetServerTime(t *testing.T) {
	c := NewClient(Options{
		AccessKey: envLocal.Bybit.AccessKey,
		SecretKey: envLocal.Bybit.SecretKey,
	})
	curTime := c.GetServerTime()

	if curTime < 1655529694 {
		t.Errorf("Could not get Exact Time")
	}
	log.Println(curTime)
	assert.NotNil(t, curTime, "expecting non-nil re")
}

func TestGetCandles(t *testing.T) {
	c := NewClient(Options{
		AccessKey: envLocal.Bybit.AccessKey,
		SecretKey: envLocal.Bybit.SecretKey,
	})
	res, err := c.GetCandle("BTCUSDT", 1, 2)
	if err != nil {
		t.Errorf("Expected person, received %v", err)
	}

	log.Printf("%+v", res[0].Close)
	assert.NotNil(t, res, "expecting non-nil res")
	assert.Nil(t, err, "expecting nil err")
	assert.Equal(t, 2, len(res), "expection 2 candle found")
}

func TestGetBalance(t *testing.T) {
	c := NewClient(Options{
		AccessKey: envLocal.Bybit.AccessKey,
		SecretKey: envLocal.Bybit.SecretKey,
	})
	res, err := c.GetBalance("USDT")
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", res)
}

func TestGetPosition(t *testing.T) {
	c := NewClient(Options{
		AccessKey: envLocal.Bybit.AccessKey,
		SecretKey: envLocal.Bybit.SecretKey,
	})
	res, err := c.GetPosition("BTCUSDT")
	if err != nil {
		t.Error(err)
	}
	log.Printf("%+v", res)
}

// func TestPlace(t *testing.T) {
// 	c := NewClient(envLocal.Bybit.AccessKey, envLocal.Bybit.SecretKey)
// 	res, err := c.PlaceBetting("BTCUSDT", BETTING_SHORT, 30000, 0.005)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	tUUID = res
// 	log.Printf("%+v", res)
// 	time.Sleep(time.Second)
// }

// var tUUID string

// func TestReplaceOrder(t *testing.T) {
// 	c := NewClient(envLocal.Bybit.AccessKey, envLocal.Bybit.SecretKey)
// 	res, err := c.Replace("BTCUSDT", tUUID, 31000, 0.005)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	log.Printf("%+v", res)
// 	time.Sleep(time.Second)
// }

// func TestSearchOrder(t *testing.T) {
// 	c := NewClient(envLocal.Bybit.AccessKey, envLocal.Bybit.SecretKey)
// 	res, err := c.Search("BTCUSDT")
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	log.Printf("%+v", res)
// }

// func TestCancelOrder(t *testing.T) {
// 	c := NewClient(envLocal.Bybit.AccessKey, envLocal.Bybit.SecretKey)
// 	res, err := c.Cancel("BTCUSDT", tUUID)
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	log.Printf("%+v", res)
// 	time.Sleep(time.Second)
// }
