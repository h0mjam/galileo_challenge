package main

import (
	"galileo/handler"
	"galileo/server"
	"galileo/store"
	"github.com/sirupsen/logrus"
)

func main() {
	srv := server.NewServer(":3003")
	storeUser := store.NewUserStore()
	storeDevice := store.NewDeviceStore()

	srv.Post("/user", handler.UserRegistry(storeUser))
	srv.Post("/device", handler.DeviceRegistry(storeDevice, storeUser))
	srv.Get("/user/devices", handler.UserDevices(storeUser))
	srv.Post("/metrics", handler.MetricsAppend(storeDevice))
	srv.Get("/device/stat/", handler.DeviceStat(storeUser))

	err := srv.Listen()

	if err != nil {
		logrus.WithError(err).Fatal("exit")
	}
}
