package main

import (
	proto "code.xxxxxx.com/micro/proto/scores"

	"code.xxxxxx.com/micro/qrcode-srv/handler"
	"github.com/micro/go-log"

	"code.xxxxxx.com/micro/common"
	"code.xxxxxx.com/micro/qrcode-srv/store"
)

func main() {
	service := common.NewService("latest", "scores", common.NilInit)

	store, err := store.New()
	if err != nil {
		log.Fatal(err)
	}

	proto.RegisterScoresHandler(service.Server(), handler.New(store))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
