package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"lykafe/news/core/data/model"
)

type CategoryRepo struct {
	db *sql.DB
	getAllCategories *sql.Stmt
	updateStatusCategory *sql.Stmt
	editCategory *sql.Stmt
	newCategory *sql.Stmt
}

func NewCategoryRepo() *CategoryRepo {
	getAllCategories, err := db.Prepare("SELECT `id`, `parent_id`, `name`, `description`, `slug`, `news_count`, `status`, `order` FROM category")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare getAllCategories ", err)
	}
	updateStatusCategory, err := db.Prepare("UPDATE category SET `status`=? WHERE id=?")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare updateStatusCategory ", err)
	}
	editCategory, err := db.Prepare("UPDATE category SET `name`=?, `slug`=? WHERE id=?")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare editCategory ", err)
	}
	newCategory, err := db.Prepare("INSERT INTO category(parent_id,`name`,`slug`,`order`) VALUES(?,?,?,?)")
	if err != nil {
		log.Fatal("[DEBUG] Cannot prepare newCategory ", err)
	}

	return &CategoryRepo {
		db: db,
		getAllCategories: getAllCategories,
		updateStatusCategory: updateStatusCategory,
		editCategory: editCategory,
		newCategory: newCategory,
	}
}

func (c *CategoryRepo) GetAllCategories() ([]*model.Category, error) {
	rows, err := c.getAllCategories.Query()
	if err != nil {
		log.Println("[DEBUG] GetAllCategories err: ", err)
		return nil, err
	}
	categories := []*model.Category{}
	defer rows.Close()
	for rows.Next() {
		cat := model.Category{}
		err = rows.Scan(&cat.Id, &cat.ParentId, &cat.Name, &cat.Description, &cat.Slug, &cat.NewsCount, &cat.Status, &cat.Order)
		if err != nil {
			return nil, err
		}
		categories = append(categories, &cat)
	}
	return categories, err
}

func (this *CategoryRepo) DeleteCategory(id int) (*model.Category, error) {
	const DELETED_STATUS = 3
	_, err := this.updateStatusCategory.Exec(DELETED_STATUS, id ) 
	if err != nil {
		log.Println("[DEBUG] DeleteCategory err: ", err)
	}
	return &model.Category{
		Id: id,
	}, err
}

func (this *CategoryRepo) UndeleteCategory(id int) (*model.Category, error) {
	const UNDELETED_STATUS = 0
	_, err := this.updateStatusCategory.Exec(UNDELETED_STATUS, id ) 
	if err != nil {
		log.Println("[DEBUG] UndeleteCategory err: ", err)
	}
	return &model.Category{
		Id: id,
	}, err
}

func (this *CategoryRepo) EditCategory(id int, name, slug string) (*model.Category, error) {
	_, err := this.editCategory.Exec(name, slug, id ) 
	if err != nil {
		log.Println("[DEBUG] editCategory err: ", err)
	}
	return &model.Category{
		Id: id,
		Name: name,
		Slug: slug,
	}, err
}

func (this *CategoryRepo) NewCategory(cat *model.Category) (*model.Category, error) {
	res, err := this.newCategory.Exec(cat.ParentId, cat.Name, cat.Slug, cat.Order ) 
	if err != nil {
		log.Println("[DEBUG] newCategory err: ", err)
	}
	cid, err := res.LastInsertId()
	if err != nil {
		log.Println("[DEBUG] newCategory LastInsertId err: ", err)
	}
	return &model.Category{
		Id: int(cid),
		ParentId: cat.ParentId,
		Name: cat.Name,
		Slug: cat.Slug,
		Order: cat.Order,
	}, err
}

func (this *CategoryRepo) ReorderCategories(categories []*model.Category) ([]*model.Category, error) {
	tx, err := this.db.Begin()
  if err != nil {
		log.Println("[DEBUG] ReorderCategories begin tx err: ", err)
    return nil, err
  }
	txn := NewTxn(tx)
	err = txn.ReorderCategories(categories)
	if err != nil {
		return nil, err
	}
  return categories, nil
}