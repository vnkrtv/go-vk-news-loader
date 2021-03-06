package main

import (
	"log"
	"time"

	"github.com/vnkrtv/go-vk-news-loader/pkg/service"
)

const (
	groupsPath = "config/groups.json"
)

func main() {
	cfg, err := service.GetConfig()
	if err != nil {
		log.Fatal(err)
	}
	newsService, err := service.NewNewsService(
		cfg.VKToken, cfg.PGUser, cfg.PGPass, cfg.PGHost, cfg.PGPort, cfg.PGName)
	if err != nil {
		log.Fatal(err)
	}

	groupsScreenNames,err := service.GetGroupsScreenNames(groupsPath)
	if err != nil {
		log.Fatal(err)
	}
	if err := newsService.InitDB(); err != nil {
		log.Println(err)
	}
	if err := newsService.AddNewsGroups(groupsScreenNames); err != nil {
		log.Fatal(err)
	}
	for {
		if err := newsService.AddNewsGroups(groupsScreenNames); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
		if err := newsService.LoadNews(100); err != nil {
			log.Println(err)
		} else {
			log.Println()
		}
		time.Sleep(time.Duration(cfg.Interval) * time.Second)
	}
}
