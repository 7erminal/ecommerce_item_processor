// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"item_processor/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.SetStaticPath("/uploads", "uploads")

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/categories",
			beego.NSInclude(
				&controllers.CategoriesController{},
			),
		),
		beego.NSNamespace("/features",
			beego.NSInclude(
				&controllers.FeaturesController{},
			),
		),
		beego.NSNamespace("/purposes",
			beego.NSInclude(
				&controllers.PurposesController{},
			),
		),
		beego.NSNamespace("/items",
			beego.NSInclude(
				&controllers.ItemsController{},
			),
		),
		beego.NSNamespace("/item-images",
			beego.NSInclude(
				&controllers.Item_imagesController{},
			),
		),
		beego.NSNamespace("/item-purposes",
			beego.NSInclude(
				&controllers.Item_purposesController{},
			),
		),
		beego.NSNamespace("/item-features",
			beego.NSInclude(
				&controllers.Item_featuresController{},
			),
		),
		beego.NSNamespace("/item-reviews",
			beego.NSInclude(
				&controllers.Item_reviewsController{},
			),
		),
		beego.NSNamespace("/item-types",
			beego.NSInclude(
				&controllers.Item_typesController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
