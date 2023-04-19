package service

import (
	"lykafe/news/core/port/outbound"
	"lykafe/news/core/data/model"
	"context"
	"fmt"
	"encoding/json"
)

type ConsumerService struct {
	mq outbound.MessageQueue
	userRepo outbound.UserRepo
}

func NewConsumerService(mq outbound.MessageQueue, userRepo outbound.UserRepo) *ConsumerService{
	return &ConsumerService{
		mq: mq,
		userRepo: userRepo,
	}
}

func (c *ConsumerService) Start() {
	go func() {
		for {
			_, v, err := c.mq.ReadMessage(context.Background())
			if err != nil {
        fmt.Println("Consume error: ", err)
    	}
			if v != nil && len(v) > 0 {
				u := model.User{}
				err = json.Unmarshal(v, &u)
				if err != nil {
					fmt.Println("[DEBUG] Unmarshal user ", err)
				} else {
					c.userRepo.UpdateUser(u.Id,  u.Username, u.Email, u.Name, u.Avatar)
				}
			}
		}
	}()
}