package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type ItemsBranchCountDTO struct {
	Branch    string `orm:"column(branch)"`
	ItemCount int64  `orm:"column(itemCount)"`
}

type ItemsCategoryCountDTO struct {
	Category  string `orm:"column(category_name)"`
	ItemCount int64  `orm:"column(itemCount)"`
}

type Items struct {
	ItemId          int64        `orm:"auto"`
	ItemName        string       `orm:"size(80)"`
	Description     string       `orm:"size(250);omitempty"`
	Weight          string       `orm:"size(20);omitempty"`
	Category        *Categories  `orm:"rel(fk)"`
	ItemPrice       *Item_prices `orm:"rel(fk);omitempty"`
	AvailableSizes  string       `orm:"size(250);omitempty"`
	AvailableColors string       `orm:"size(250);omitempty"`
	Material        string       `orm:"size(400);omitempty"`
	ImagePath       string       `orm:"size(250);omitempty"`
	Quantity        int          `orm:"omitempty"`
	QuantityAlert   int          `orm:"omitempty"`
	Active          int          `orm:"omitempty"`
	DateCreated     time.Time    `orm:"type(datetime);omitempty"`
	DateModified    time.Time    `orm:"type(datetime);omitempty"`
	CreatedBy       int          `orm:"omitempty"`
	ModifiedBy      int          `orm:"omitempty"`
	Country         *Countries   `orm:"rel(fk);column(country)"`
	Branch          *Branches    `orm:"rel(fk);column(branch)"`
}

func init() {
	orm.RegisterModel(new(Items))
}

// AddItems insert a new Items into database and returns
// last inserted Id on success.
func AddItems(m *Items) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetItemsStatsByBranch retrieves Items stats by branch. Returns error if
// Id doesn't exist
func GetItemsCategoryStatsByBranch(id int64) (count_ *[]ItemsCategoryCountDTO, err error) {
	o := orm.NewOrm()

	sql := `
        SELECT c.category_name, COUNT(i.item_id) itemCount
        FROM categories c LEFT JOIN items i ON c.category_id = i.category_id
		WHERE i.branch = 3
        GROUP BY c.category_name
    `
	var results []ItemsCategoryCountDTO
	if _, err := o.Raw(sql).QueryRows(&results); err == nil {
		logs.Info("Results:: ", results)
		// for _, result := range results {
		// 	logs.Info("Data fetched is ", result.Branch, " :: ", result.ItemCount)
		// }
		return &results, nil
	}
	return nil, err
}

// GetItemsStatsByCategory retrieves Items stats by category. Returns error if
// Id doesn't exist
func GetItemsStatsByCategory() (count_ *[]ItemsCategoryCountDTO, err error) {
	o := orm.NewOrm()

	sql := `
        SELECT c.category_name, COUNT(i.item_id) itemCount
        FROM categories c LEFT JOIN items i ON c.category_id = i.category_id
        GROUP BY c.category_name
    `
	var results []ItemsCategoryCountDTO
	if _, err := o.Raw(sql).QueryRows(&results); err == nil {
		logs.Info("Results:: ", results)
		// for _, result := range results {
		// 	logs.Info("Data fetched is ", result.Branch, " :: ", result.ItemCount)
		// }
		return &results, nil
	}
	// if _, err := o.QueryTable(new(Categories)).RelatedSel().All(results); err == nil {
	// 	logs.Info("Results:: ", results)
	// 	// for _, result := range results {
	// 	// 	logs.Info("Data fetched is ", result.Branch, " :: ", result.ItemCount)
	// 	// }
	// 	return &results, nil
	// }
	return nil, err
}

// GetItemsStatsByBranch retrieves Items stats by branch. Returns error if
// Id doesn't exist
// func GetItemsStatsByCategory() (v *Items, err error) {
// 	o := orm.NewOrm()
// 	v = &Items{}
// 	if err = o.QueryTable(new(Items)).Filter("ItemId", id).RelatedSel().One(v); err == nil {
// 		return v, nil
// 	}
// 	return nil, err
// }

// GetItemsById retrieves Items by Id. Returns error if
// Id doesn't exist
func GetItemsById(id int64) (v *Items, err error) {
	o := orm.NewOrm()
	v = &Items{ItemId: id}
	if err = o.QueryTable(new(Items)).Filter("ItemId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetItemsByPurposeId retrieves Items by purpose Id. Returns error if
// purpose Id doesn't exist
func GetItemsByPurposeId(id int64) (v *[]Items, err error) {
	o := orm.NewOrm()
	v = &[]Items{}

	if ft, err := GetPurposesById(id); err == nil {
		if _, err := o.QueryTable(new(Items)).RelatedSel().Filter("Item_purposes__Purpose", ft).All(v); err == nil {
			return v, nil
		}
	}
	return nil, err
}

// GetItemsByFeatureId retrieves Items by Feature Id. Returns error if
// Feature Id doesn't exist
func GetItemsByFeatureId(id int64) (v *[]Items, err error) {
	o := orm.NewOrm()
	v = &[]Items{}
	if ft, err := GetFeaturesById(id); err == nil {
		if _, err = o.QueryTable(new(Items)).RelatedSel().Filter("Item_features__Feature", ft).All(v); err == nil {
			return v, nil
		}
	}
	return nil, err
}

// GetItemsByFeatureId retrieves Items by Feature Id. Returns error if
// Feature Id doesn't exist
func GetItemsByCategoryId(id int64) (v *[]Items, err error) {
	o := orm.NewOrm()
	v = &[]Items{}
	if ft, err := GetCategoriesById(id); err == nil {
		if _, err = o.QueryTable(new(Items)).RelatedSel().Filter("Category", ft).All(v); err == nil {
			return v, nil
		}
	}
	return nil, err
}

// GetItemsByItemId retrieves Items by Item Id. Returns error if
// Item Id doesn't exist
func GetItemsFromFeaturesWithItem(id int64) (v *[]Items, err error) {
	o := orm.NewOrm()
	v = &[]Items{}
	if ft, err := GetItemsById(id); err == nil {
		if _, err = o.QueryTable(new(Items)).RelatedSel().Filter("Item_features__Item", ft).All(v); err == nil {
			return v, nil
		}
	}
	return nil, err
}

// GetAllItems retrieves all Items matches certain condition. Returns empty list if
// no records exist
func GetAllItems(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Items))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Items
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// GetAllItems retrieves all Items matches certain condition. Returns empty list if
// no records exist
func GetAllItemsByBranch(branch *Branches, query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Items))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Items
	qs = qs.Filter("Branch", branch).OrderBy(sortFields...).RelatedSel()
	if _, err = qs.All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateItems updates Items by Id and returns error if
// the record to be updated doesn't exist
func UpdateItemsById(m *Items) (err error) {
	o := orm.NewOrm()
	v := Items{ItemId: m.ItemId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteItems deletes Items by Id and returns error if
// the record to be deleted doesn't exist
func DeleteItems(id int64) (err error) {
	o := orm.NewOrm()
	v := Items{ItemId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Items{ItemId: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
