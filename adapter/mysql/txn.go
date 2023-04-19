package mysql

import (
	"database/sql"
	"log"
	"lykafe/news/core/data/dto"
	"lykafe/news/core/data/model"
	"lykafe/news/lib"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Txn struct {
	tx *sql.Tx
}

func NewTxn(tx *sql.Tx) *Txn {
	return &Txn{
		tx: tx,
	}
}

const MaxNumberFeaturedPosts = 16
const MaxNumberHotPosts = 31
const MaxNumberTrendingPosts = 31

func (this *Txn) PublishNews(req *dto.PublishNewsReq) (string, error) {
	commit := false
	defer this.close(&commit)
	id, err := this.insertNews(req.Title, req.Content, req.ContentJp, req.Description, req.DescriptionJp, req.ImgUrl, req.MetaKw, req.MetaDesc, req.Slug, req.Category, req.SubCategory, req.AuthorId, req.PublishAt)
	if err != nil {
		return id, err
	}
	err = this.upsertTagsCount(req.Tags)
	if err != nil {
		return id, err
	}
	err = this.insertNewsTags(id, req.Tags)
	if err != nil {
		return id, err
	}
	err = this.increaseCategoryCount(req.Category)
	if err != nil {
		return id, err
	}
	err = this.increaseCategoryCount(req.SubCategory)
	if err != nil {
		return id, err
	}
	err = this.insertRelated(id, req.Related)
	if err != nil {
		return id, err
	}

	if req.IsFeatured {
		err = this.increaseAllFeaturedOrder(MaxNumberFeaturedPosts)
		if err != nil {
			return id, err
		}
		err = this.deleteBiggerFeatured()
		if err != nil {
			return id, err
		}
		err = this.insertFeatured(id)
		if err != nil {
			return id, err
		}
	}

	if req.IsTrending {
		err = this.increaseAllTrendingOrder(MaxNumberTrendingPosts)
		if err != nil {
			return id, err
		}
		err = this.deleteBiggerTrending()
		if err != nil {
			return id, err
		}
		err = this.insertTrending(id)
		if err != nil {
			return id, err
		}
	}

	if req.IsHot {
		err = this.increaseAllHotOrder(MaxNumberHotPosts)
		if err != nil {
			return id, err
		}
		err = this.deleteBiggerHot()
		if err != nil {
			return id, err
		}
		err = this.insertHot(id)
		if err != nil {
			return id, err
		}
	}
	commit = true
	return id, err
}

func (this *Txn) EditNews(
	news *model.NewsView,
	newsTags []string,
	newsFeaturedOrder int,
	req *dto.PublishNewsReq) error {
	commit := false
	defer this.close(&commit)
	err := this.updateNews(req.Id, req.Title, req.Content, req.ContentJp, req.Description, req.DescriptionJp, req.ImgUrl, req.MetaKw, req.MetaDesc, req.Slug, req.Category, req.SubCategory, req.AuthorId, req.PublishBy, req.PublishAt)
	if err != nil {
		return err
	}
	err = this.decreaseTagsCount(newsTags)
	if err != nil {
		return err
	}
	err = this.upsertTagsCount(req.Tags)
	if err != nil {
		return err
	}
	err = this.deleteNewsTags(news.Id)
	if err != nil {
		return err
	}
	err = this.insertNewsTags(req.Id, req.Tags)
	if err != nil {
		return err
	}
	err = this.decreaseCategoryCount(news.Category)
	if err != nil {
		return err
	}
	err = this.decreaseCategoryCount(news.SubCategory)
	if err != nil {
		return err
	}
	err = this.increaseCategoryCount(req.Category)
	if err != nil {
		return err
	}
	err = this.increaseCategoryCount(req.SubCategory)
	if err != nil {
		return err
	}
	err = this.deleteRelated(req.Id)
	if err != nil {
		return err
	}
	err = this.insertRelated(req.Id, req.Related)
	if err != nil {
		return err
	}
	if req.IsFeatured {
		if newsFeaturedOrder >= MaxNumberFeaturedPosts {
			err = this.deleteBiggerFeatured()
			if err != nil {
				return err
			}
			err = this.insertFeatured(req.Id)
			if err != nil {
				return err
			}
		}
	} else {
		if newsFeaturedOrder < MaxNumberFeaturedPosts {
			err = this.deleteFeaturedNews(req.Id)
			if err != nil {
				return err
			}
			err = this.decreaseAllFeaturedOrder(newsFeaturedOrder)
			if err != nil {
				return err
			}
		}
	}

	// if req.IsFeatured {
	// 	err = this.increaseAllFeaturedOrder(newsFeaturedOrder)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if newsFeaturedOrder < MaxNumberFeaturedPosts {
	// 		err = this.updateFeaturedOrder(req.Id, 1)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	} else {
	// 		err = this.deleteBiggerFeatured()
	// 		if err != nil {
	// 			return err
	// 		}
	// 		err = this.insertFeatured(req.Id)
	// 		if err != nil {
	// 			return err
	// 		}
	// 	}
	// }
	commit = true
	return nil
}

func (this *Txn) DeleteNews(newsId string, newsFeaturedOrder int) error {
	commit := false
	defer this.close(&commit)
	err := this.updateNewsStatus(newsId)
	if err != nil {
		return err
	}
	err = this.deleteFeaturedNews(newsId)
	if err != nil {
		return err
	}
	if newsFeaturedOrder < MaxNumberFeaturedPosts {
		err = this.decreaseAllFeaturedOrder(newsFeaturedOrder)
		if err != nil {
			return err
		}
	}
	commit = true
	return err
}

func (this *Txn) ReorderCategories(categories []*model.Category) error {
	commit := false
	defer this.close(&commit)
	reorderCategorySql := "UPDATE category set `order`=? WHERE `id`=?"
	stmt, err := this.tx.Prepare(reorderCategorySql)
	if err != nil {
		log.Println("[DEBUG] tx Prepare reorderCategorySql err: ", err)
		return err
	}
	defer stmt.Close()
	for _, category := range categories {
		_, err = stmt.Exec(category.Order, category.Id)
		if err != nil {
			log.Println("[DEBUG] tx Exec reorderCategorySql err: ", err)
			return err
		}
	}
	commit = true
	return nil
}

func (this *Txn) close(commit *bool) {
	if *commit {
		if err := this.tx.Commit(); err != nil {
			log.Println("Commit sql transaction err: ", err)
		}
	} else {
		if err := this.tx.Rollback(); err != nil {
			log.Println("Rollback sql transcation err: ", err)
		}
	}
}

func (this *Txn) exec(cmdSql string, params ...any) error {
	stmt, err := this.tx.Prepare(cmdSql)
	if err != nil {
		log.Println("[DEBUG] txn Prepare: ", cmdSql, " err: ", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(params...)
	if err != nil {
		log.Println("[DEBUG] txn Exec: ", cmdSql, " err: ", err)
		return err
	}
	return nil
}

func (this *Txn) insertNews(title, content, contentJp, description, descriptionJp, imgUrl, metaKw, metaDesc, slug string,
	category, subCategory int,
	publishBy string,
	publishAt time.Time) (string, error) {
	insertNewsSql := `INSERT INTO news(id, title, content, content_jp, description, description_jp, img_url, meta_kw, meta_desc, slug,
		category, sub_category, publish_by, published_at) VALUES(UNHEX(?), ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, UNHEX(?),?)`
	id, err := lib.OrderedUuidV1()
	if err != nil {
		log.Println("[DEBUG] PublishNews cannot OrderedUuidV1: ", err)
		return "", err
	}
	err = this.exec(insertNewsSql, id, title, content, contentJp, description, descriptionJp, imgUrl, metaKw, metaDesc, slug, category, subCategory, publishBy, publishAt)
	return id, err
}

func (this *Txn) updateNews(id, title, content, contentJp, description, descriptionJp, imgUrl, metaKw, metaDesc, slug string,
	category, subCategory int,
	authorId, publishBy string,
	publishAt time.Time) error {
	updateNewsSql := `UPDATE news SET title = ?, content = ?, content_jp = ?, description = ?, description_jp = ?, img_url = ?, meta_kw = ?, meta_desc = ?, slug = ?,
		category = ?, sub_category = ?, published_at = ?, publish_by=UNHEX(?), updated_by=UNHEX(?), updated_at=NOW(), status = 0 WHERE id = UNHEX(?)`
	return this.exec(updateNewsSql, title, content, contentJp, description, descriptionJp, imgUrl, metaKw, metaDesc, slug, category, subCategory, publishAt, authorId, publishBy, id)
}

func (this *Txn) upsertTagsCount(tags []string) error {
	upsertTagCountSql := `INSERT INTO tag(name, count) VALUES(?, 1) ON DUPLICATE KEY UPDATE count = count + 1`
	stmt, err := this.tx.Prepare(upsertTagCountSql)
	if err != nil {
		log.Println("[DEBUG] tx Prepare upsertTagCountSql err: ", err)
		return err
	}
	defer stmt.Close()
	for _, tag := range tags {
		_, err = stmt.Exec(tag)
		if err != nil {
			log.Println("[DEBUG] tx Exec upsertTagCountSql err: ", err)
			return err
		}
	}
	return nil
}

func (this *Txn) decreaseTagsCount(tags []string) error {
	decreaseTagCountSql := "UPDATE tag SET `count` = `count` - 1 WHERE `name` = ?"
	stmt, err := this.tx.Prepare(decreaseTagCountSql)
	if err != nil {
		log.Println("[DEBUG] tx Prepare decreaseTagCountSql err: ", err)
		return err
	}
	defer stmt.Close()
	for _, tag := range tags {
		_, err = stmt.Exec(tag)
		if err != nil {
			log.Println("[DEBUG] tx Exec decreaseTagCountSql err: ", err)
			return err
		}
	}
	return nil
}

func (this *Txn) deleteNewsTags(newsId string) error {
	deleteNewsTagsSql := "DELETE FROM news_tag WHERE news_id=UNHEX(?)"
	return this.exec(deleteNewsTagsSql, newsId)
}

func (this *Txn) insertNewsTags(newsId string, tags []string) error {
	insertNewsTagSql := `INSERT INTO news_tag(id, news_id, tag_name) VALUES(UNHEX(?), UNHEX(?), ?) `
	stmt, err := this.tx.Prepare(insertNewsTagSql)
	if err != nil {
		log.Println("[DEBUG] PublishNews tx Prepare insertNewsTagSql err: ", err)
		return err
	}
	defer stmt.Close()
	for _, tag := range tags {
		id, _ := lib.OrderedUuidV1()
		_, err = stmt.Exec(id, newsId, tag)
		if err != nil {
			log.Println("[DEBUG] PublishNews tx Exec insertNewsTagSql err: ", err)
			return err
		}
	}
	return nil
}

func (this *Txn) increaseCategoryCount(catId int) error {
	if catId == 0 {
		return nil
	}
	increaseCategoryCountSql := "UPDATE category SET news_count = news_count + 1 WHERE id = ?"
	return this.exec(increaseCategoryCountSql, catId)
}

func (this *Txn) decreaseCategoryCount(catId int) error {
	if catId == 0 {
		return nil
	}
	decreaseCategoryCountSql := "UPDATE category SET news_count = news_count -1 WHERE id = ?"
	return this.exec(decreaseCategoryCountSql, catId)
}

func (this *Txn) deleteRelated(newsId string) error {
	deleteRelatedSql := "DELETE FROM related WHERE news_id = UNHEX(?)"
	return this.exec(deleteRelatedSql, newsId)
}

func (this *Txn) insertRelated(newsId string, related []string) error {
	insertRelatedSql := "INSERT INTO related(id, news_id, related_id) VALUES(UNHEX(?), UNHEX(?), UNHEX(?))"
	stmt, err := this.tx.Prepare(insertRelatedSql)
	if err != nil {
		log.Println("[DEBUG] PublishNews tx Prepare insertRelatedSql err: ", err)
		return err
	}
	defer stmt.Close()
	for _, rel := range related {
		relId, err := lib.OrderedUuidV1()
		if err != nil {
			log.Println("[DEBUG] PublishNews cannot OrderedUuidV1: ", err)
			return err
		}
		_, err = stmt.Exec(relId, newsId, rel)
		if err != nil {
			log.Println("[DEBUG] PublishNews tx Exec insertRelatedSql err: ", err)
			return err
		}
	}
	return nil
}

func (this *Txn) increaseAllFeaturedOrder(topThreshold int) error {
	increaseAllOrderFeaturedSql := "UPDATE featured SET `order` = `order` + 1 WHERE `order` < ?"
	return this.exec(increaseAllOrderFeaturedSql, topThreshold)
}

func (this *Txn) decreaseAllFeaturedOrder(bottomThreshold int) error {
	decreaseAllOrderFeaturedSql := "UPDATE featured SET `order` = `order` - 1 WHERE `order` > ?"
	return this.exec(decreaseAllOrderFeaturedSql, bottomThreshold)
}

func (this *Txn) deleteBiggerFeatured() error {
	deleteBiggerFeaturedSql := "DELETE FROM featured WHERE `order` > 15"
	return this.exec(deleteBiggerFeaturedSql)
}

func (this *Txn) insertFeatured(newsId string) error {
	insertFeaturedSql := "INSERT INTO featured(news_id, `order`) VALUES(UNHEX(?), 1)"
	return this.exec(insertFeaturedSql, newsId)
}

func (this *Txn) increaseAllTrendingOrder(limit int) error {
	increaseAllOrderTrendingSql := "UPDATE trending SET `order` = `order` + 1 WHERE `order` < ?"
	return this.exec(increaseAllOrderTrendingSql, limit)
}

func (this *Txn) deleteBiggerTrending() error {
	deleteBiggerTrendingSql := "DELETE FROM trending WHERE `order` > 30"
	return this.exec(deleteBiggerTrendingSql)
}

func (this *Txn) insertTrending(newsId string) error {
	insertTrendingSql := "INSERT INTO trending(news_id, `order`) VALUES(UNHEX(?), 1)"
	return this.exec(insertTrendingSql, newsId)
}

func (this *Txn) increaseAllHotOrder(limit int) error {
	increaseAllOrderHotSql := "UPDATE hot SET `order` = `order` + 1 WHERE `order` < ?"
	return this.exec(increaseAllOrderHotSql, limit)
}

func (this *Txn) deleteBiggerHot() error {
	deleteBiggerHotSql := "DELETE FROM hot WHERE `order` > 30"
	return this.exec(deleteBiggerHotSql)
}

func (this *Txn) insertHot(newsId string) error {
	insertHotSql := "INSERT INTO hot(news_id, `order`) VALUES(UNHEX(?), 1)"
	return this.exec(insertHotSql, newsId)
}

func (this *Txn) updateNewsStatus(newsId string) error {
	updateNewsStatusSql := "UPDATE news SET `status`=?  WHERE `id`=UNHEX(?)"
	const DELETED_STATUS = 3
	return this.exec(updateNewsStatusSql, DELETED_STATUS, newsId)
}

func (this *Txn) deleteFeaturedNews(newsId string) error {
	deleteFeaturedNewsSql := "DELETE FROM featured WHERE news_id=UNHEX(?)"
	return this.exec(deleteFeaturedNewsSql, newsId)
}

func (this *Txn) updateFeaturedOrder(newsId string, order int) error {
	updateFeaturedOrderSql := "UPDATE featured SET `order` = ? WHERE news_id=UNHEX(?)"
	return this.exec(updateFeaturedOrderSql, order, newsId)
}
