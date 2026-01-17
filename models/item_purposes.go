package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Item_purposes struct {
	ItemPurposeId int64     `orm:"auto"`
	Item          *Items    `orm:"rel(fk)"`
	Purpose       *Purposes `orm:"rel(fk)"`
	Active        int
	DateCreated   time.Time `orm:"type(datetime)"`
	DateModified  time.Time `orm:"type(datetime)"`
	CreatedBy     int
	ModifiedBy    int
}

func init() {
	orm.RegisterModel(new(Item_purposes))
}

// AddItem_purposes insert a new Item_purposes into database and returns
// last inserted Id on success.
func AddItem_purposes(m *Item_purposes) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetItem_purposesById retrieves Item_purposes by Id. Returns error if
// Id doesn't exist
func GetItem_purposesById(id int64) (v *Item_purposes, err error) {
	o := orm.NewOrm()
	v = &Item_purposes{ItemPurposeId: id}
	if err = o.QueryTable(new(Item_purposes)).Filter("ItemPurposeId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetItem_purposesByItemId retrieves Item_purposes by item Id. Returns error if
// Id doesn't exist
func GetItem_purposesByItemId(id int64) (v *[]Item_purposes, err error) {
	o := orm.NewOrm()
	var l []Item_purposes
	if _, err = o.QueryTable(new(Item_purposes)).Filter("Item", id).RelatedSel().All(&l); err == nil {
		return &l, nil
	}
	return nil, err
}

// GetItem_purposesByItemId retrieves Item_purposes by purpose Id. Returns error if
// Id doesn't exist
func GetItem_purposesByPurposeId(id int64) (v *[]Item_purposes, err error) {
	o := orm.NewOrm()
	var l []Item_purposes
	if _, err = o.QueryTable(new(Item_purposes)).Filter("Purpose", id).RelatedSel().All(&l); err == nil {
		return &l, nil
	}
	return nil, err
}

// GetAllItem_purposes retrieves all Item_purposes matches certain condition. Returns empty list if
// no records exist
func GetAllItem_purposes(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Item_purposes))
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

	var l []Item_purposes
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
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

// UpdateItem_purposes updates Item_purposes by Id and returns error if
// the record to be updated doesn't exist
func UpdateItem_purposesById(m *Item_purposes) (err error) {
	o := orm.NewOrm()
	v := Item_purposes{ItemPurposeId: m.ItemPurposeId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteItem_purposes deletes Item_purposes by Id and returns error if
// the record to be deleted doesn't exist
func DeleteItem_purposes(id int64) (err error) {
	o := orm.NewOrm()
	v := Item_purposes{ItemPurposeId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Item_purposes{ItemPurposeId: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func DeleteItem_purposesByItem(itemId int64) (err error) {
	o := orm.NewOrm()
	var num int64
	if num, err = o.QueryTable(new(Item_purposes)).Filter("Item", itemId).Delete(); err == nil {
		fmt.Println("Number of records deleted in database:", num)
	}
	return
}
