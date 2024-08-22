package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CountriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CountriesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CountriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CountriesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CountriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CountriesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CountriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CountriesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CountriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CountriesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CurrenciesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CurrenciesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CurrenciesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CurrenciesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CurrenciesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CurrenciesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"],
        beego.ControllerComments{
            Method: "ChangeVisibility",
            Router: `/change-visibility/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"],
        beego.ControllerComments{
            Method: "GetItemFeaturesByFeature",
            Router: `/features/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"],
        beego.ControllerComments{
            Method: "GetItemPurposesByPurpose",
            Router: `/purposes/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetItemFeatures",
            Router: `/features/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetItemPurposes",
            Router: `/purposes/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetItemQuantity",
            Router: `/quantity/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ObjectController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ObjectController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ObjectController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ObjectController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ObjectController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:PurposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:PurposesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:PurposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:PurposesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:PurposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:PurposesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:PurposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:PurposesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:PurposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:PurposesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:PurposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:PurposesController"],
        beego.ControllerComments{
            Method: "ChangeVisibility",
            Router: `/change-visibility/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:UserController"] = append(beego.GlobalControllerRouter["item_processor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:UserController"] = append(beego.GlobalControllerRouter["item_processor/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:UserController"] = append(beego.GlobalControllerRouter["item_processor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:UserController"] = append(beego.GlobalControllerRouter["item_processor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:UserController"] = append(beego.GlobalControllerRouter["item_processor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:UserController"] = append(beego.GlobalControllerRouter["item_processor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:UserController"] = append(beego.GlobalControllerRouter["item_processor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
