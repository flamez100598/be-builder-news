package service

import (
	"lykafe/news/core/data/model"
	"lykafe/news/core/data/dto"
	"lykafe/news/core/data/codes"
	"lykafe/news/core/port/outbound"
	"lykafe/news/core/data/apperr"
	"errors"
)

type SubscriptionService struct {
	subscriptionRepo outbound.SubscriptionRepo
}

func NewSubscriptionService(subscriptionRepo outbound.SubscriptionRepo) *SubscriptionService {
	return &SubscriptionService{
		subscriptionRepo: subscriptionRepo,
	}
}

func (s *SubscriptionService) Subscribe(email string) (*model.Subscription, codes.Code, error) {
	sub, err := s.subscriptionRepo.GetSubscriptionByEmail(email)
	if err == nil && sub != nil {
		return sub, codes.Existed, nil
	}
	if errors.Is(err, apperror.ErrNoRows) {
		sub, err = s.subscriptionRepo.InsertSubscription(email)	
	}
	return sub, codes.Ok, err
}

func (s *SubscriptionService) SearchSubscription(req *dto.SearchSubscriptionReq) ([]*model.Subscription, error) {
	return s.subscriptionRepo.SearchSubscription(req)
}

func (s *SubscriptionService) SearchSubscriptionCount(req *dto.SearchSubscriptionReq) (int, error) {
	return s.subscriptionRepo.SearchSubscriptionCount(req)
}

func (s *SubscriptionService) DeleteSubscription(email string) (*model.Subscription, error) {
	return s.subscriptionRepo.DeleteSubscription(email)
}

func (s *SubscriptionService) SendContact(req *dto.ContactReq) (*model.Contact, error) {
	return s.subscriptionRepo.SendContact(req)
}

func (s *SubscriptionService)	SearchContact(req *dto.ContactReq) ([]*model.Contact, error) {
	return s.subscriptionRepo.SearchContact(req)
}

func (s *SubscriptionService)	SearchContactCount(req *dto.ContactReq) (int, error) {
	return s.subscriptionRepo.SearchContactCount(req)
}

func (s *SubscriptionService)	GetPageContent(page string) (*model.PageContent, error) {
	return s.subscriptionRepo.GetPageContent(page)
}

func (s *SubscriptionService)	UpdatePageContent(content, updatedBy, page string) (*model.PageContent, error) {
	return s.subscriptionRepo.UpdatePageContent(content, updatedBy, page)
}