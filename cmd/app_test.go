package main

import (
	"context"
	"flag"
	modles "github.com/paraggarg37/order_app/src/domain/models"
	"github.com/paraggarg37/order_app/src/interfaces/web/api"
	"log"
	"os"
	"sync"
	"testing"
)

var integration = flag.Bool("integration", false, "run database integration tests")

var LogisticsApi *api.API

func TestMain(m *testing.M) {
	flag.Parse()

	if *integration {
		log.Println("Running Integration Test")
		setup()
	}
	code := m.Run()

	if *integration {
		log.Println("Tearing down")
		tearDown()
	}

	os.Exit(code)
}

func setup() {
	LogisticsApi = Setup()
}

func TestCreateOrder(t *testing.T) {
	if !*integration {
		t.SkipNow()
	}

	var err error

	ctx := context.Background()
	_, err = LogisticsApi.Interactor.OrderInteractor.CreateOrder(ctx, &modles.OrderReq{
		Destination: []string{},
		Origin:      []string{},
	})

	if err == nil {
		t.Error("must through error for empty args")
	}

	_, err = LogisticsApi.Interactor.OrderInteractor.CreateOrder(ctx, &modles.OrderReq{
		Destination: []string{"wronglat", "wronglng"},
		Origin:      []string{"wronglat", "wronglng"},
	})

	if err == nil {
		t.Error("must through error for wrong args")
	}
	order, err := LogisticsApi.Interactor.OrderInteractor.CreateOrder(ctx, &modles.OrderReq{
		Destination: []string{"40.43206", "-80.38992"},
		Origin:      []string{"41.43206", "-80.388"},
	})

	if err != nil {
		t.Error(err)
		return
	}

	if order.Id == 0 {
		t.Error("id must be greater than zero")
		return
	} else {
		t.Logf("created test order successfully %+v", order)
	}

}

func TestTakeOrderParallel(t *testing.T) {

	if !*integration {
		t.SkipNow()
	}

	var err error

	ctx := context.Background()
	order, err := LogisticsApi.Interactor.OrderInteractor.CreateOrder(ctx, &modles.OrderReq{
		Destination: []string{"40.43206", "-80.38992"},
		Origin:      []string{"41.43206", "-80.388"},
	})

	if err != nil {
		t.Error(err)
		return
	}

	var err1, err2 error
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		_, err1 = LogisticsApi.Interactor.OrderInteractor.TakeOrder(ctx, &modles.TakeOrderReq{
			Status: "TAKEN",
		}, order.Id)

		wg.Done()

	}()

	go func() {
		_, err2 = LogisticsApi.Interactor.OrderInteractor.TakeOrder(ctx, &modles.TakeOrderReq{
			Status: "TAKEN",
		}, order.Id)
		wg.Done()
	}()

	wg.Wait()

	if err1 == nil && err2 == nil {
		t.Error("atleast one should have failed")
	} else if err1 != nil && err2 != nil {
		t.Error("atleast one should be successfull")
	} else {
		t.Log(err1, err2)
	}

}

func TestTakeOrder(t *testing.T) {
	if !*integration {
		t.SkipNow()
	}

	var err error

	ctx := context.Background()
	order, err := LogisticsApi.Interactor.OrderInteractor.CreateOrder(ctx, &modles.OrderReq{
		Destination: []string{"40.43206", "-80.38992"},
		Origin:      []string{"41.43206", "-80.388"},
	})

	if err != nil {
		t.Error(err)
		return
	}

	resp, err := LogisticsApi.Interactor.OrderInteractor.TakeOrder(ctx, &modles.TakeOrderReq{
		Status: "INVALID_ORDER",
	}, order.Id)

	if err == nil {
		t.Error("take order must throw error for invalid state")
		return
	}

	resp, err = LogisticsApi.Interactor.OrderInteractor.TakeOrder(ctx, &modles.TakeOrderReq{
		Status: "TAKEN",
	}, order.Id)

	if err != nil {
		t.Error(err)
	} else if resp.Status != "SUCCESS" {
		t.Error("response should be success")
		return
	} else {
		t.Log("Successfully taken order")
	}

	resp, err = LogisticsApi.Interactor.OrderInteractor.TakeOrder(ctx, &modles.TakeOrderReq{
		Status: "TAKEN",
	}, order.Id)

	if err == nil {
		t.Error("must throw error, order is already taken")
		return
	}

	resp, err = LogisticsApi.Interactor.OrderInteractor.TakeOrder(ctx, &modles.TakeOrderReq{
		Status: "TAKEN",
	}, 0)

	if err == nil {
		t.Error("must through error for invalid order id")
		return
	}
}

func TestListOrder(t *testing.T) {
	if !*integration {
		t.SkipNow()
	}

	ctx := context.Background()
	order, err := LogisticsApi.Interactor.OrderInteractor.CreateOrder(ctx, &modles.OrderReq{
		Destination: []string{"40.43206", "-80.38992"},
		Origin:      []string{"41.43206", "-80.388"},
	})

	if err != nil {
		t.Error(err)
		return
	}

	resp, err := LogisticsApi.Interactor.OrderInteractor.ListOrders(ctx, 1, 1)

	if err != nil {
		t.Error(err)
		return
	}

	if len(resp) != 1 {
		t.Error("lenght of response must be 1")
		return
	} else {
		t.Logf("order list = %+v", order)
	}

}

func tearDown() {
	LogisticsApi.DB.Close()
}
