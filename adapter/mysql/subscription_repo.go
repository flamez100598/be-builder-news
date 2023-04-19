package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"lykafe/news/core/data/apperr"
	"lykafe/news/core/data/model"
	"lykafe/news/core/data/dto"
	"lykafe/news/lib"
	"time"
)

type SubscriptionRepo struct {
	db *sql.DB
	getSubscriptionByEmail *sql.Stmt
	insertSubscription *sql.Stmt
	searchSubscription *sql.Stmt
	searchSubscriptionCount *sql.Stmt
	deleteSubscription *sql.Stmt
	insertContact *sql.Stmt
	searchContact *sql.Stmt
	searchContactCount  *sql.Stmt
	getAboutUs  *sql.Stmt
	getCareer  *sql.Stmt
	getNewsletter  *sql.Stmt
	updateAboutUs  *sql.Stmt
	updateCareer  *sql.Stmt
	updateNewsletter  *sql.Stmt
}

func NewSubscriptionRepo() *SubscriptionRepo {
	getSubscriptionByEmail, err := db.Prepare(`SELECT email, name, status, created_at FROM subscription WHERE email = ?`)
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getSubscriptionByEmail ", err)
	}

	insertSubscription, err := db.Prepare(`INSERT INTO subscription(email) VALUES ( ?)`)
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare insertSubscription ", err)
	}

	searchSubscription, err := db.Prepare("call SearchSubscription(?, ?, ?, ?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare searchSubscription ", err)
	}
	
	searchSubscriptionCount, err := db.Prepare("call SearchSubscriptionCount(?, ?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare searchSubscriptionCount ", err)
	}

	deleteSubscription, err := db.Prepare("UPDATE `subscription` SET `status`=? WHERE `email`=?")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare deleteSubscription ", err)
	}

	insertContact, err := db.Prepare(`INSERT INTO contact(id, email, phone, name, subject, msg) VALUES (UNHEX(?), ?, ?, ?, ?, ?)`)
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare insertSubscription ", err)
	}

	searchContact, err := db.Prepare("call SearchContact(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare searchContact ", err)
	}
	searchContactCount, err := db.Prepare("call SearchContactCount(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare searchContactCount ", err)
	}

	getAboutUs, err := db.Prepare("SELECT p.id, p.title, p.content, p.updated_at, u.username, u.name, u.email, u.avatar FROM p_about_us p INNER JOIN `user` u ON p.updated_by=u.id")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getAboutUs ", err)
	}

	getCareer, err := db.Prepare("SELECT p.id, p.title, p.content, p.updated_at, u.username, u.name, u.email, u.avatar FROM p_career p INNER JOIN `user` u ON p.updated_by=u.id")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getCareer ", err)
	}

	getNewsletter, err := db.Prepare("SELECT p.id, p.title, p.content, p.updated_at, u.username, u.name, u.email, u.avatar FROM p_newsletter p INNER JOIN `user` u ON p.updated_by=u.id")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getNewsletter ", err)
	}

	updateAboutUs, err := db.Prepare("UPDATE p_about_us SET content=?, updated_at=NOW(), updated_by=UNHEX(?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare updateAboutUs ", err)
	}

	updateCareer, err := db.Prepare("UPDATE p_career SET content=?, updated_at=NOW(), updated_by=UNHEX(?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare updateCareer ", err)
	}

	updateNewsletter, err := db.Prepare("UPDATE p_newsletter SET content=?, updated_at=NOW(), updated_by=UNHEX(?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare updateNewsletter ", err)
	}

	return &SubscriptionRepo {
		db: db,
		getSubscriptionByEmail: getSubscriptionByEmail,
		insertSubscription: insertSubscription,
		searchSubscription: searchSubscription,
		searchSubscriptionCount: searchSubscriptionCount,
		deleteSubscription: deleteSubscription,
		insertContact: insertContact,
		searchContact: searchContact,
		searchContactCount: searchContactCount,
		getAboutUs: getAboutUs,
		getCareer: getCareer,
		getNewsletter: getNewsletter,
		updateAboutUs: updateAboutUs,
		updateCareer: updateCareer,
		updateNewsletter: updateNewsletter,
	}
}

func (s *SubscriptionRepo) InsertSubscription(email string) (*model.Subscription, error) {
	_, err := s.insertSubscription.Exec(email)
	sub := model.Subscription{
		Email: email,
	}
	if err != nil {
		log.Println("[DEBUG] InsertSubscription err: ", err)
		return nil, err
	}
	return &sub, err
}

func (s *SubscriptionRepo) GetSubscriptionByEmail(email string) (*model.Subscription, error) {
	sub := model.Subscription{}
	var createdAt []uint8
	err := s.getSubscriptionByEmail.QueryRow(email).Scan(&sub.Email,&sub.Name, &sub.Status, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, apperror.ErrNoRows
		}
		log.Println("[DEBUG] GetSubscriptionByEmail err: ", err)
	}
	sub.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", string(createdAt))
	return &sub, err
}

func (s *SubscriptionRepo) SearchSubscription(req *dto.SearchSubscriptionReq) ([]*model.Subscription, error) {
	rows, err := s.searchSubscription.Query(req.Email, req.Status, req.Offset, req.Limit)
	if err != nil {
		log.Println("[DEBUG] SearchSubscription err: ", err)
		return nil, err
	}
	subscriptions := []*model.Subscription{}
	defer rows.Close()
	for rows.Next() {
		sub := model.Subscription{}
		var createdAt []uint8
		err = rows.Scan(&sub.Email, &sub.Name, &sub.Status, &createdAt)
		if err != nil {
			return nil, err
		}
		sub.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		log.Println("SearchSubscription result", sub)
		if err != nil {
			log.Println("	time.Parse err", 	err)
		}
		subscriptions = append(subscriptions, &sub)
	}
	return subscriptions, err
}

func (s *SubscriptionRepo)	SearchSubscriptionCount(req *dto.SearchSubscriptionReq) (int, error) {
	var c int
	err := s.searchSubscriptionCount.QueryRow(req.Email, req.Status).Scan(&c)
	return c, err
}

func (s *SubscriptionRepo)	DeleteSubscription(email string) (*model.Subscription, error) {
	const DELETED_STATUS = 2
	_, err := s.deleteSubscription.Exec(DELETED_STATUS, email)
	return &model.Subscription {
		Email: email,
	}, err
}

func (s *SubscriptionRepo)	SendContact(req *dto.ContactReq) (*model.Contact, error) {
	id, err := lib.OrderedUuidV1()
	if err != nil {
		log.Println("[DEBUG] SendContact cannot OrderedUuidV1: ", err)
		return nil, err
	}
	_, err = s.insertContact.Exec(id, req.Email, req.Phone, req.Name, req.Subject, req.Msg)
	contact := model.Contact{
		Id: id,
		Email: req.Email,
		Phone: req.Phone,
		Name: req.Name,
		Subject: req.Subject,
		Msg: req.Msg,
	}
	if err != nil {
		log.Println("[DEBUG] insertContact err: ", err)
		return nil, err
	}
	return &contact, err
}

func (s *SubscriptionRepo)	SearchContact(req *dto.ContactReq) ([]*model.Contact, error) {
	rows, err := s.searchContact.Query(req.Email, req.Phone, req.Name, req.Subject, req.Msg, req.Offset, req.Limit)
	if err != nil {
		log.Println("[DEBUG] searchContact err: ", err)
		return nil, err
	}
	contacts := []*model.Contact{}
	defer rows.Close()
	for rows.Next() {
		con := model.Contact{}
		var createdAt []uint8
		err = rows.Scan(&con.Id, &con.Email, &con.Phone, &con.Name, &con.Subject, &con.Msg, &con.Status, &createdAt)
		if err != nil {
			return nil, err
		}
		con.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		log.Println("SearchContact result", con)
		if err != nil {
			log.Println("	time.Parse err", 	err)
		}
		contacts = append(contacts, &con)
	}
	return contacts, err
}

func (s *SubscriptionRepo)	SearchContactCount(req *dto.ContactReq) (int, error) {
	var c int
	err := s.searchContactCount.QueryRow(req.Email, req.Phone, req.Name, req.Subject, req.Msg).Scan(&c)
	return c, err
}

func (s *SubscriptionRepo)	GetPageContent(page string) (*model.PageContent, error) {
	var stmt *sql.Stmt
	switch page {
	case "about_us":
		stmt = s.getAboutUs
	case "career":
		stmt = s.getCareer
	case "newsletter":
		stmt = s.getNewsletter
	default:
		stmt = nil
	}
	p := model.PageContent{}
	var updatedAt []uint8
	err := stmt.QueryRow().Scan(&p.Id, &p.Title, &p.Content, &updatedAt, &p.UsernameUpdatedBy, &p.NameUpdatedBy, &p.EmailUpdatedBy, &p.AvatarUpdatedBy)
	p.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", string(updatedAt))
	return &p, err
}

func (s *SubscriptionRepo)	UpdatePageContent(content, updatedBy, page string) (*model.PageContent, error) {
	log.Println("updatedBy", updatedBy)
	log.Println("content", content)
	log.Println("page", page)
	var stmt *sql.Stmt
	switch page {
	case "about_us":
		stmt = s.updateAboutUs
	case "career":
		stmt = s.updateCareer
	case "newsletter":
		stmt = s.updateNewsletter
	default:
		stmt = nil
	}
	_, err := stmt.Exec(content, updatedBy)
	p := model.PageContent{
		Id: 1,
		Content: content,
	}
	return &p, err
}