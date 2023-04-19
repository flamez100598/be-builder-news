package usecase

import (
	"lykafe/news/core/data/model"
	"lykafe/news/core/data/dto"
	"lykafe/news/core/data/codes"
)

type Subscription interface {
	Subscribe(email string) (*model.Subscription, codes.Code, error)
	SearchSubscription(req *dto.SearchSubscriptionReq) ([]*model.Subscription, error)
	SearchSubscriptionCount(req *dto.SearchSubscriptionReq) (int, error)
	DeleteSubscription(email string) (*model.Subscription, error)
	SendContact(req *dto.ContactReq) (*model.Contact, error)
	SearchContact(req *dto.ContactReq) ([]*model.Contact, error)
	SearchContactCount(req *dto.ContactReq) (int, error)
	GetPageContent(page string) (*model.PageContent, error)
	UpdatePageContent(content, updatedBy, page string) (*model.PageContent, error)
}