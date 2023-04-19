package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"lykafe/news/core/data/model"
	"lykafe/news/core/data/dto"
	"time"
	"strings"
)

type NewsRepo struct {
	db *sql.DB
	
	getNewsByTag *sql.Stmt
	getNewsByTagCount *sql.Stmt
	getListFeaturedNews *sql.Stmt
	getNewsBySlug *sql.Stmt
	getNewsById *sql.Stmt
	getNewsByCategory *sql.Stmt
	getNewsByCategoryNotIncludeSubs  *sql.Stmt
	getNewsBySubCategory *sql.Stmt
	adminSearchNews *sql.Stmt
	adminSearchNewsCount *sql.Stmt
	adminSearchRelatedNews *sql.Stmt
	adminSearchRelatedNewsCount *sql.Stmt
	searchNews *sql.Stmt
	searchNewsCount *sql.Stmt
	getRelatedNews *sql.Stmt
	getNewsTags *sql.Stmt
	getTagsNews *sql.Stmt
	getTagsNewsCount *sql.Stmt
	reorderFeatured *sql.Stmt
	getFeaturedOrder *sql.Stmt
	setFeaturedOrder *sql.Stmt
	getNewsUrl *sql.Stmt
}

func NewNewsRepo() *NewsRepo {
	getNewsByTag, err := db.Prepare("SELECT DISTINCT HEX(n.`id`), `title`, `description`, `img_url`, `meta_kw`, `meta_desc`, `slug`, `category`, `sub_category`, `comment_num`, `vote_num`, `view_num`, `status`, HEX(`publish_by`), `ranking`, n.`created_at`, `updated_at`, `published_at`, `username` as `username_publish_by`, `name` as `name_publish_by`, avatar as `avatar_publish_by` FROM news as n INNER JOIN news_tag as t ON n.id=t.news_id LEFT JOIN  `user` as u ON n.publish_by=u.id WHERE t.tag_name = ? AND n.published_at <= NOW() AND n.status <> ? ORDER BY n.id DESC LIMIT ?, ?")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getListFeaturedNews ", err)
	}

	getNewsByTagCount, err := db.Prepare("SELECT COUNT(DISTINCT n.`id`) as c FROM news as n INNER JOIN news_tag as t ON n.id=t.news_id LEFT JOIN  `user` as u ON n.publish_by=u.id WHERE t.tag_name = ? AND n.published_at <= NOW() AND n.status <> ?")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getListFeaturedNews ", err)
	}

	getListFeaturedNews, err := db.Prepare("SELECT HEX(n.`id`), `title`, `description`, `img_url`, `meta_kw`, `meta_desc`, `slug`, `category`, `sub_category`, `comment_num`, `vote_num`, `view_num`, `status`, HEX(`publish_by`), `ranking`, n.`created_at`, `updated_at`, `published_at`, `order`, `username` as `username_publish_by`, `name` as `name_publish_by`, avatar as `avatar_publish_by` FROM featured as f INNER JOIN news as n ON f.news_id=n.id INNER JOIN user as u ON u.id=n.publish_by WHERE n.published_at <= NOW() AND n.status <> 3")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getListFeaturedNews ", err)
	}
	getNewsBySlug, err := db.Prepare("SELECT HEX(n.`id`), `title`, `content`, `description`, `img_url`, `meta_kw`, `meta_desc`, `slug`, `category`, `sub_category`, `comment_num`, `vote_num`, `view_num`, `status`, HEX(`publish_by`), `ranking`, n.`created_at`, `updated_at`, `published_at`, `username` as `username_publish_by`, `name` as `name_publish_by`, avatar as `avatar_publish_by` FROM news as n INNER JOIN user as u ON u.id=n.publish_by WHERE slug=? AND n.published_at <= NOW() AND n.status <> 3")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getNewsBySlug ", err)
	}
	getNewsById, err := db.Prepare("SELECT HEX(n.`id`), `title`, `content`, `description`, `img_url`, `meta_kw`, `meta_desc`, `slug`, `category`, `sub_category`, `comment_num`, `vote_num`, `view_num`, `status`, HEX(`publish_by`), `ranking`, n.`created_at`, `updated_at`, `published_at`, `username` as `username_publish_by`, `name` as `name_publish_by`, avatar as `avatar_publish_by` FROM news as n INNER JOIN user as u ON u.id=n.publish_by WHERE n.`id`=UNHEX(?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getNewsById ", err)
	}

	getNewsByCategory, err := db.Prepare("SELECT HEX(n.`id`), `title`, `description`, `img_url`, `meta_kw`, `meta_desc`,`slug`, `category`, `sub_category`, `comment_num`, `vote_num`, `view_num`, `status`, HEX(`publish_by`), `ranking`, n.`created_at`, `updated_at`, `published_at`, `username` as `username_publish_by`, `name` as `name_publish_by`, avatar as `avatar_publish_by` FROM news as n INNER JOIN user as u ON u.id=n.publish_by WHERE n.category=? AND n.published_at <= NOW() AND n.status <> 3 ORDER BY n.id DESC LIMIT ?, ?")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getNewsByCategory ", err)
	}

	getNewsByCategoryNotIncludeSubs, err := db.Prepare("SELECT HEX(n.`id`), `title`, `description`, `img_url`, `meta_kw`, `meta_desc`,`slug`, `category`, `sub_category`, `comment_num`, `vote_num`, `view_num`, `status`, HEX(`publish_by`), `ranking`, n.`created_at`, `updated_at`, `published_at`, `username` as `username_publish_by`, `name` as `name_publish_by`, avatar as `avatar_publish_by` FROM news as n INNER JOIN user as u ON u.id=n.publish_by WHERE n.category=? AND n.sub_category = 0 AND n.published_at <= NOW() AND n.status <> 3 ORDER BY n.id DESC LIMIT ?, ?")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getNewsByCategory ", err)
	}

	getNewsBySubCategory, err := db.Prepare("SELECT HEX(n.`id`), `title`, `description`, `img_url`, `meta_kw`, `meta_desc`, `slug`, `category`, `sub_category`, `comment_num`, `vote_num`, `view_num`, `status`, HEX(`publish_by`), `ranking`, n.`created_at`, `updated_at`, `published_at`, `username` as `username_publish_by`, `name` as `name_publish_by`, avatar as `avatar_publish_by` FROM news as n INNER JOIN user as u ON u.id=n.publish_by WHERE n.sub_category=? AND n.published_at <= NOW() AND n.status <> 3 ORDER BY n.id DESC LIMIT ?, ?")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getNewsBySubCategory ", err)
	}

	adminSearchNews, err := db.Prepare("call AdminSearchNews(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare adminSearchNews ", err)
	}
	adminSearchNewsCount, err := db.Prepare("call AdminSearchNewsCount(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare adminSearchNewsCount ", err)
	}

	adminSearchRelatedNews, err := db.Prepare("call AdminSearchRelatedNews(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare adminSearchRelatedNews ", err)
	}

	adminSearchRelatedNewsCount, err := db.Prepare("call AdminSearchRelatedNewsCount(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare adminSearchRelatedNewsCount ", err)
	}

	searchNews, err := db.Prepare("SELECT HEX(n.`id`), `title`, `description`, `img_url`, `meta_kw`, `meta_desc`, `slug`, `category`, `sub_category`, `comment_num`, `vote_num`, `view_num`, `status`, HEX(`publish_by`), `ranking`, n.`created_at`, `updated_at`, `published_at`, `username` as `username_publish_by`, `name` as `name_publish_by`, avatar as `avatar_publish_by` FROM news as n INNER JOIN user as u ON u.id=n.publish_by WHERE (n.title LIKE CONCAT('%', ?, '%') OR n.content LIKE CONCAT('%', ?, '%')) AND n.published_at <= NOW() AND n.status <> 3 ORDER BY n.id DESC LIMIT ?, ?")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare searchNews ", err)
	}

	searchNewsCount, err := db.Prepare("SELECT COUNT(n.`id`) AS c FROM news as n WHERE (n.title LIKE CONCAT('%', ?, '%') OR n.content LIKE CONCAT('%', ?, '%')) AND n.published_at <= NOW() AND n.status <> 3")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare searchNewsCount ", err)
	}

	getRelatedNews, err := db.Prepare("SELECT HEX(n.`id`), `title`, `description`, `img_url`, `meta_kw`, `meta_desc`, `slug`, `category`, `sub_category`, `comment_num`, `vote_num`, `view_num`, `status`, HEX(`publish_by`), `ranking`, n.`created_at`, `updated_at`, `published_at` FROM related as r INNER JOIN news as n ON r.related_id = n.id WHERE r.news_id=UNHEX(?) AND n.published_at <= NOW() AND n.status <> 3 ORDER BY n.id DESC")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getRelatedNews ", err)
	}

	getNewsTags, err := db.Prepare("SELECT `tag_name` FROM news_tag WHERE news_id=UNHEX(?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getNewsTags ", err)
	}

	getTagsNews, err := db.Prepare("SELECT `name`, `count` FROM tag WHERE name LIKE CONCAT(?, '%') ORDER BY `count` DESC LIMIT ?, ?")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getTagsNews ", err)
	}

	getTagsNewsCount, err := db.Prepare("SELECT COUNT(`name`) as c FROM tag WHERE name LIKE CONCAT(?, '%')")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getTagsNewsCount ", err)
	}

	reorderFeatured, err := db.Prepare("UPDATE featured SET `order`=? WHERE news_id=UNHEX(?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare reorderFeatured ", err)
	}

	getFeaturedOrder, err := db.Prepare("SELECT `order` FROM featured WHERE news_id=UNHEX(?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getFeaturedOrder ", err)
	}

	setFeaturedOrder, err := db.Prepare("UPDATE featured SET `order` = ? WHERE news_id=UNHEX(?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare setFeaturedOrder ", err)
	}

	getNewsUrl, err := db.Prepare("SELECT n.slug, c.slug FROM news AS n LEFT JOIN category AS c ON n.category=c.id")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getNewsUrl ", err)
	}

	return &NewsRepo {
		db: db,

		getNewsByTag: getNewsByTag,
		getNewsByTagCount: getNewsByTagCount,
		getListFeaturedNews: getListFeaturedNews,
		getNewsBySlug: getNewsBySlug,
		getNewsById: getNewsById,
		getNewsByCategory: getNewsByCategory,
		getNewsByCategoryNotIncludeSubs: getNewsByCategoryNotIncludeSubs,
		getNewsBySubCategory: getNewsBySubCategory,
		adminSearchNews: adminSearchNews,
		adminSearchNewsCount: adminSearchNewsCount,
		adminSearchRelatedNews: adminSearchRelatedNews,
		adminSearchRelatedNewsCount: adminSearchRelatedNewsCount,
		searchNews: searchNews,
		searchNewsCount: searchNewsCount,
		getRelatedNews: getRelatedNews,
		getNewsTags: getNewsTags,
		getTagsNews: getTagsNews,
		getTagsNewsCount: getTagsNewsCount,
		reorderFeatured: reorderFeatured,
		getFeaturedOrder: getFeaturedOrder,
		setFeaturedOrder: setFeaturedOrder,
		getNewsUrl: getNewsUrl,
	}
}

func (n *NewsRepo) PublishNews(req *dto.PublishNewsReq) (*model.News, error) {
	tx, err := n.db.Begin()
  if err != nil {
		log.Println("[DEBUG] PublishNews begin tx err: ", err)
    return nil, err
  }
	txn := NewTxn(tx)
	id, err := txn.PublishNews(req)
	if err != nil {
		return nil, err
	}
  return &model.News{
		Id: id,
		Title: req.Title,
		Content: req.Content,
		PublishBy: req.PublishBy,
		Slug: req.Slug,
	}, nil
}

func (n *NewsRepo) EditNewsPost(req *dto.PublishNewsReq) (*model.News, error) {
	news, err := n.GetNewsById(req.Id)
	newsTags, err := n.GetNewsTags(req.Id)
	newsFeaturedOrder, err := n.GetFeaturedOrder(req.Id)
	tx, err := n.db.Begin()
  if err != nil {
		log.Println("[DEBUG] EditNewsPost begin tx err: ", err)
    return nil, err
  }
	txn := NewTxn(tx)
	err = txn.EditNews(news, newsTags, newsFeaturedOrder, req)
	if err != nil {
		return nil, err
	}
  return &model.News{
		Id: req.Id,
	}, nil
	return nil, nil
}

func (n *NewsRepo) DeleteNewsPost(id string) (*model.News, error) {
	newsFeaturedOrder, err := n.GetFeaturedOrder(id)
	tx, err := n.db.Begin()
  if err != nil {
		log.Println("[DEBUG] DeleteNewsPost begin tx err: ", err)
    return nil, err
  }
	txn := NewTxn(tx)
	err = txn.DeleteNews(id, newsFeaturedOrder)
	if err != nil {
		return nil, err
	}
  return &model.News{
		Id: id,
	}, nil
}

func (n *NewsRepo) GetFeaturedNews()  ([]*model.NewsView, error) {
	rows, err := n.getListFeaturedNews.Query()
	if err != nil {
		log.Println("[DEBUG] getListFeaturedNews err: ", err)
		return nil, err
	}
	listNewsView := []*model.NewsView{}
	defer rows.Close()
	for rows.Next() {
		nv := model.NewsView{}
		var createdAt []uint8
		var updatedAt []uint8
		var publishedAt []uint8
		err = rows.Scan(&nv.Id, &nv.Title, &nv.Description, &nv.ImgUrl, &nv.MetaKw, &nv.MetaDesc, &nv.Slug, &nv.Category, &nv.SubCategory, &nv.CommentNum, &nv.VoteNum, &nv.ViewNum, &nv.Status, &nv.PublishBy, &nv.Ranking, &createdAt, &updatedAt, &publishedAt, &nv.Order, &nv.UsernamePublishBy, &nv.NamePublishBy, &nv.AvatarPublishBy)
		if err != nil {
			return nil, err
		}
		nv.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", string(createdAt))
		nv.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		nv.PublishedAt, _ = time.Parse("2006-01-02 15:04:05", string(publishedAt))
		listNewsView = append(listNewsView, &nv)
	}
	return listNewsView, err
}

func (n *NewsRepo) GetNewsBySlug(slug string) (*model.NewsView, error) {
	var createdAt []uint8
	var updatedAt []uint8
	var publishedAt []uint8
	nv := model.NewsView{}
	err := n.getNewsBySlug.QueryRow(slug).Scan(&nv.Id, &nv.Title, &nv.Content, &nv.Description, &nv.ImgUrl, &nv.MetaKw, &nv.MetaDesc, &nv.Slug, &nv.Category, &nv.SubCategory, &nv.CommentNum, &nv.VoteNum, &nv.ViewNum, &nv.Status, &nv.PublishBy, &nv.Ranking, &createdAt, &updatedAt, &publishedAt, &nv.UsernamePublishBy, &nv.NamePublishBy, &nv.AvatarPublishBy)
	if err != nil {
		log.Println("[DEBUG] GetNewsBySlug err: ", err)
		return nil, err
	}
	nv.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", string(createdAt))
	nv.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))
	nv.PublishedAt, _ = time.Parse("2006-01-02 15:04:05", string(publishedAt))
	return &nv, err
}

func (n *NewsRepo) GetNewsById(id string) (*model.NewsView, error) {
	var createdAt []uint8
	var updatedAt []uint8
	var publishedAt []uint8
	nv := model.NewsView{}
	err := n.getNewsById.QueryRow(id).Scan(&nv.Id, &nv.Title, &nv.Content, &nv.Description, &nv.ImgUrl, &nv.MetaKw, &nv.MetaDesc, &nv.Slug, &nv.Category, &nv.SubCategory, &nv.CommentNum, &nv.VoteNum, &nv.ViewNum, &nv.Status, &nv.PublishBy, &nv.Ranking, &createdAt, &updatedAt, &publishedAt, &nv.UsernamePublishBy, &nv.NamePublishBy, &nv.AvatarPublishBy)
	if err != nil {
		log.Println("[DEBUG] GetNewsById err: ", err)
		return nil, err
	}
	nv.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", string(createdAt))
	nv.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))
	nv.PublishedAt, _ = time.Parse("2006-01-02 15:04:05", string(publishedAt))
	return &nv, err
}

func (n *NewsRepo) GetNewsCategory(includeSubs bool, id, offset, limit int) ([]*model.NewsView, error) {
	stmt := n.getNewsByCategory
	if !includeSubs {
		stmt = n.getNewsByCategoryNotIncludeSubs
	}
	rows, err :=stmt.Query(id, offset, limit)
	if err != nil {
		log.Println("[DEBUG] GetNewsCategory err: ", err)
		return nil, err
	}
	listNewsView := []*model.NewsView{}
	defer rows.Close()
	for rows.Next() {
		nv := model.NewsView{}
		var createdAt []uint8
		var updatedAt []uint8
		var publishedAt []uint8
		err = rows.Scan(&nv.Id, &nv.Title, &nv.Description, &nv.ImgUrl, &nv.MetaKw, &nv.MetaDesc, &nv.Slug, &nv.Category, &nv.SubCategory, &nv.CommentNum, &nv.VoteNum, &nv.ViewNum, &nv.Status, &nv.PublishBy, &nv.Ranking, &createdAt, &updatedAt, &publishedAt, &nv.UsernamePublishBy, &nv.NamePublishBy, &nv.AvatarPublishBy)
		if err != nil {
			return nil, err
		}
		nv.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", string(createdAt))
		nv.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		nv.PublishedAt, _ = time.Parse("2006-01-02 15:04:05", string(publishedAt))
		listNewsView = append(listNewsView, &nv)
	}
	return listNewsView, err
}

func (n *NewsRepo) GetNewsSubCategory(id, offset, limit int) ([]*model.NewsView, error) {
	rows, err := n.getNewsBySubCategory.Query(id, offset, limit)
	if err != nil {
		log.Println("[DEBUG] GetNewsSubCategory err: ", err)
		return nil, err
	}
	listNewsView := []*model.NewsView{}
	defer rows.Close()
	for rows.Next() {
		nv := model.NewsView{}
		var createdAt []uint8
		var updatedAt []uint8
		var publishedAt []uint8
		err = rows.Scan(&nv.Id, &nv.Title, &nv.Description, &nv.ImgUrl, &nv.MetaKw, &nv.MetaDesc, &nv.Slug, &nv.Category, &nv.SubCategory, &nv.CommentNum, &nv.VoteNum, &nv.ViewNum, &nv.Status, &nv.PublishBy, &nv.Ranking, &createdAt, &updatedAt, &publishedAt, &nv.UsernamePublishBy, &nv.NamePublishBy, &nv.AvatarPublishBy)
		if err != nil {
			return nil, err
		}
		nv.CreatedAt, _ = time.Parse("2006-01-02 15:04:05", string(createdAt))
		nv.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		nv.PublishedAt, _ = time.Parse("2006-01-02 15:04:05", string(publishedAt))
		listNewsView = append(listNewsView, &nv)
	}
	return listNewsView, err
}

func (n *NewsRepo) AdminSearchNews(req *dto.AdminSearchNewsReq)  ([]*model.NewsView, error) {
	var from *time.Time = nil
	var to *time.Time = nil
	if !req.From.IsZero() {
		from = &req.From
	}
	if !req.To.IsZero() {
		to = &req.To
	}
	log.Println("[DEBUG] adminSearchNews from: ", from)
	log.Println("[DEBUG] adminSearchNews to: ", to)
	rows, err := n.adminSearchNews.Query(req.Title, req.Content, req.Description, req.Category, req.SubCategory, req.PublishBy, req.UpdateBy, from, to, req.Status, req.Offset, req.Limit)
	if err != nil {
		log.Println("[DEBUG] adminSearchNews err: ", err)
		return nil, err
	}
	listNewsView := []*model.NewsView{}
	defer rows.Close()
	for rows.Next() {
		nv := model.NewsView{}
		var createdAt []uint8
		var updatedAt []uint8
		var publishedAt []uint8
		var publishByName sql.NullString
		var publishByUsername sql.NullString
		var publishByAvatar sql.NullString
		var updatedByName sql.NullString
		var updatedByUsername sql.NullString
		var updatedByAvatar sql.NullString
		err = rows.Scan(&nv.Id, &nv.Title, &nv.Content, &nv.Description, &nv.ImgUrl, &nv.MetaKw, &nv.MetaDesc, &nv.Slug, &nv.Category, &nv.SubCategory, &nv.CommentNum, &nv.VoteNum, &nv.ViewNum, &nv.Status, &nv.PublishBy, &nv.Ranking, &createdAt, &updatedAt, &publishedAt, &publishByUsername, &publishByName, &publishByAvatar, &updatedByUsername, &updatedByName, &updatedByAvatar)
		if err != nil {
			return nil, err
		}
		nv.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		nv.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		nv.PublishedAt, _ = time.Parse("2006-01-02 15:04:05", string(publishedAt))
		nv.UsernamePublishBy = publishByUsername.String
		nv.NamePublishBy = publishByName.String
		nv.AvatarPublishBy = publishByAvatar.String
		nv.UsernameUpdateBy = updatedByUsername.String
		nv.AvatarUpdateBy = updatedByAvatar.String
		nv.NameUpdateBy = updatedByName.String
		if err != nil {
			log.Println("	time.Parse err", 	err)
		}
		listNewsView = append(listNewsView, &nv)
	}
	return listNewsView, err
}

func (n *NewsRepo) AdminSearchNewsCount(req *dto.AdminSearchNewsReq) (int, error) {
	var from *time.Time = nil
	var to *time.Time = nil
	if !req.From.IsZero() {
		from = &req.From
	}
	if !req.To.IsZero() {
		to = &req.To
	}
	var c int
	err := n.adminSearchNewsCount.QueryRow(req.Title, req.Content, req.Description, req.Category, req.SubCategory, req.PublishBy, req.UpdateBy, from, to, req.Status).Scan(&c)
	return c, err
}

func (n *NewsRepo) AdminSearchRelatedNews(req *dto.AdminSearchRelatedNewsReq) ([]*model.NewsView, error) {
	var from *time.Time = nil
	var to *time.Time = nil
	if !req.From.IsZero() {
		from = &req.From
	}
	if !req.To.IsZero() {
		to = &req.To
	}
	tags := strings.Join(req.Tags, "','")
	tags = "'" + tags + "'"
	rows, err := n.adminSearchRelatedNews.Query(tags, req.Title, req.Content, req.Category, req.SubCategory, req.PublishBy, from, to, req.Offset, req.Limit)
	if err != nil {
		log.Println("[DEBUG] adminSearchRelatedNews err: ", err)
		return nil, err
	}
	listNewsView := []*model.NewsView{}
	defer rows.Close()
	for rows.Next() {
		nv := model.NewsView{}
		var createdAt []uint8
		var updatedAt []uint8
		var publishedAt []uint8
		err = rows.Scan(&nv.Id, &nv.Title, &nv.Content, &nv.Description, &nv.ImgUrl, &nv.MetaKw, &nv.MetaDesc, &nv.Slug, &nv.Category, &nv.SubCategory, &nv.CommentNum, &nv.VoteNum, &nv.ViewNum, &nv.Status, &nv.PublishBy, &nv.Ranking, &createdAt, &updatedAt, &publishedAt, &nv.UsernamePublishBy, &nv.NamePublishBy, &nv.AvatarPublishBy)
		if err != nil {
			return nil, err
		}
		nv.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		nv.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		nv.PublishedAt, _ = time.Parse("2006-01-02 15:04:05", string(publishedAt))
		if err != nil {
			log.Println("	time.Parse err", 	err)
		}
		listNewsView = append(listNewsView, &nv)
	}
	return listNewsView, err
}

func (n *NewsRepo) AdminSearchRelatedNewsCount(req *dto.AdminSearchRelatedNewsReq) (int, error) {
	var from *time.Time = nil
	var to *time.Time = nil
	if !req.From.IsZero() {
		from = &req.From
	}
	if !req.To.IsZero() {
		to = &req.To
	}
	var c int
	tags := strings.Join(req.Tags, "','")
	tags = "'" + tags + "'"
	err := n.adminSearchRelatedNewsCount.QueryRow(tags, req.Title, req.Content, req.Category, req.SubCategory, req.PublishBy, from, to).Scan(&c)
	return c, err
}

func (n *NewsRepo) SearchNews(req *dto.SearchNewsReq)  ([]*model.NewsView, error) {
	rows, err := n.searchNews.Query(req.Keywords, req.Keywords, req.Offset, req.Limit)
	if err != nil {
		log.Println("[DEBUG] searchNews err: ", err)
		return nil, err
	}
	listNewsView := []*model.NewsView{}
	defer rows.Close()
	for rows.Next() {
		nv := model.NewsView{}
		var createdAt []uint8
		var updatedAt []uint8
		var publishedAt []uint8
		err = rows.Scan(&nv.Id, &nv.Title, &nv.Description, &nv.ImgUrl, &nv.MetaKw, &nv.MetaDesc, &nv.Slug, &nv.Category, &nv.SubCategory, &nv.CommentNum, &nv.VoteNum, &nv.ViewNum, &nv.Status, &nv.PublishBy, &nv.Ranking, &createdAt, &updatedAt, &publishedAt, &nv.UsernamePublishBy, &nv.NamePublishBy, &nv.AvatarPublishBy)
		if err != nil {
			return nil, err
		}
		nv.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		nv.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		nv.PublishedAt, _ = time.Parse("2006-01-02 15:04:05", string(publishedAt))
		if err != nil {
			log.Println("	time.Parse err", 	err)
		}
		listNewsView = append(listNewsView, &nv)
	}
	return listNewsView, err
}

func (n *NewsRepo) SearchNewsCount(req *dto.SearchNewsReq) (int, error) {
	var c int
	err := n.searchNewsCount.QueryRow(req.Keywords, req.Keywords).Scan(&c)
	return c, err
}

func (n *NewsRepo) GetNewsByTag(tag string, offset, limit int, includeDeleted bool) ([]*model.NewsView, error) {
	statusDeleted := 3
	if includeDeleted == true {
		statusDeleted = 100
	}
	rows, err := n.getNewsByTag.Query(tag, statusDeleted, offset, limit)
	if err != nil {
		log.Println("[DEBUG] GetNewsByTags err: ", err)
		return nil, err
	}
	listNewsView := []*model.NewsView{}
	defer rows.Close()
	for rows.Next() {
		nv := model.NewsView{}
		var createdAt []uint8
		var updatedAt []uint8
		var publishedAt []uint8
		err = rows.Scan(&nv.Id, &nv.Title, &nv.Description, &nv.ImgUrl, &nv.MetaKw, &nv.MetaDesc, &nv.Slug, &nv.Category, &nv.SubCategory, &nv.CommentNum, &nv.VoteNum, &nv.ViewNum, &nv.Status, &nv.PublishBy, &nv.Ranking, &createdAt, &updatedAt, &publishedAt, &nv.UsernamePublishBy, &nv.NamePublishBy, &nv.AvatarPublishBy)
		if err != nil {
			return nil, err
		}
		nv.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		nv.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		nv.PublishedAt, _ = time.Parse("2006-01-02 15:04:05", string(publishedAt))
		if err != nil {
			log.Println("	time.Parse err", 	err)
		}
		listNewsView = append(listNewsView, &nv)
	}
	return listNewsView, err
}

func (n *NewsRepo) GetNewsByTagCount(tag string, includeDeleted bool) (int, error) {
	statusDeleted := 3
	if includeDeleted == true {
		statusDeleted = 100
	}
	var c int
	err := n.getNewsByTagCount.QueryRow(tag, statusDeleted).Scan(&c)
	return c, err
}

func (n *NewsRepo) GetRelatedNews(id string) ([]*model.NewsView, error) {
	rows, err := n.getRelatedNews.Query(id)
	if err != nil {
		log.Println("[DEBUG] GetRelatedNews err: ", err)
		return nil, err
	}
	listNewsView := []*model.NewsView{}
	defer rows.Close()
	for rows.Next() {
		nv := model.NewsView{}
		var createdAt []uint8
		var updatedAt []uint8
		var publishedAt []uint8
		err = rows.Scan(&nv.Id, &nv.Title, &nv.Description, &nv.ImgUrl, &nv.MetaKw, &nv.MetaDesc, &nv.Slug, &nv.Category, &nv.SubCategory, &nv.CommentNum, &nv.VoteNum, &nv.ViewNum, &nv.Status, &nv.PublishBy, &nv.Ranking, &createdAt, &updatedAt, &publishedAt)
		if err != nil {
			return nil, err
		}
		nv.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAt))
		nv.UpdatedAt, _ = time.Parse("2006-01-02 15:04:05", string(updatedAt))
		nv.PublishedAt, _ = time.Parse("2006-01-02 15:04:05", string(publishedAt))
		if err != nil {
			log.Println("	time.Parse err", 	err)
		}
		listNewsView = append(listNewsView, &nv)
	}
	return listNewsView, err
}

func (n *NewsRepo) GetNewsTags(id string) ([]string, error) {
	rows, err := n.getNewsTags.Query(id)
	if err != nil {
		log.Println("[DEBUG] GetNewsTags err: ", err)
		return nil, err
	}
	tags := []string{}
	defer rows.Close()
	for rows.Next() {
		var tag string
		err = rows.Scan(&tag)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, err
}

// get list tags
func (n *NewsRepo) GetTagsNews(tag string, offset, limit int) ([]*model.Tag, error) {
	rows, err := n.getTagsNews.Query(tag, offset, limit)
	if err != nil {
		log.Println("[DEBUG] GetTagsNews err: ", err)
		return nil, err
	}
	tags := []*model.Tag{}
	defer rows.Close()
	for rows.Next() {
		t := model.Tag{}
		err = rows.Scan(&t.Name, &t.Count)
		if err != nil {
			return nil, err
		}
		tags = append(tags, &t)
	}
	return tags, err
}

// get list tags count
func (n *NewsRepo) GetTagsNewsCount(tag string) (int, error) {
	var c int
	err := n.getTagsNewsCount.QueryRow(tag).Scan(&c)
	return c, err
}

func (n *NewsRepo) ReorderFeatured(req []*dto.NewsOrder) ([]*dto.NewsOrder, error) {
	for _, o := range req {
		_, err := n.reorderFeatured.Exec(o.Order, o.Id ) 
  	if err != nil {
			log.Println("[DEBUG] ReorderFeatured err: ", err)
			return nil, err
		}
  }
	return req, nil
}

func (n *NewsRepo) GetFeaturedOrder(newsId string) (int, error) {
	var o int
	err := n.getFeaturedOrder.QueryRow(newsId).Scan(&o)
	if err == sql.ErrNoRows {
		return MaxNumberFeaturedPosts, nil
	}
	return o, err
}

func (n *NewsRepo) SetFeaturedOrder(newsId string, order int) error {
	_, err := n.setFeaturedOrder.Exec(order, newsId ) 
	if err != nil {
		log.Println("[DEBUG] SetFeaturedOrder err: ", err)
	}
	return err
}

func (n *NewsRepo) GetAllNewsUrls() ([]*dto.NewsUrl, error) {
	rows, err := n.getNewsUrl.Query()
	if err != nil {
		log.Println("[DEBUG] getNewsUrl err: ", err)
		return nil, err
	}
	urls := []*dto.NewsUrl{}
	defer rows.Close()
	for rows.Next() {
		url := dto.NewsUrl{}
		err = rows.Scan(&url.NewsSlug, &url.CategorySlug)
		if err != nil {
			return nil, err
		}
		urls = append(urls, &url)
	}
	return urls, err
}