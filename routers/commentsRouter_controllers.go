package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:CategoriesController"],
        beego.ControllerComments{
            Method: "GetCategoryByName",
            Router: "/name/:name",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"],
        beego.ControllerComments{
            Method: "ChangeVisibility",
            Router: "/change-visibility/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"],
        beego.ControllerComments{
            Method: "GetAllFeaturesAndTheirItems",
            Router: "/items",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:FeaturesController"],
        beego.ControllerComments{
            Method: "GetOneByName",
            Router: "/name",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"],
        beego.ControllerComments{
            Method: "GetItemFeaturesByFeature",
            Router: "/features/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_featuresController"],
        beego.ControllerComments{
            Method: "GetItemFeaturesByItem",
            Router: "/features/item/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_imagesController"],
        beego.ControllerComments{
            Method: "UploadPictures",
            Router: "/upload-pictures",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_pricesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"],
        beego.ControllerComments{
            Method: "GetItemPurposesByPurpose",
            Router: "/purposes/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_purposesController"],
        beego.ControllerComments{
            Method: "GetItemPurposesByItem",
            Router: "/purposes/item/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_quantityController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:Item_reviewsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetAllByBranch",
            Router: "/branch/:branch_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetItemsByCategory",
            Router: "/categories/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetItemCount",
            Router: "/count/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetItemCountByTypeAndBranch",
            Router: "/count/type",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetItemCountByType",
            Router: "/count/type/:name",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetItemFeatures",
            Router: "/features/feature/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetItemStats",
            Router: "/get-item-stats/:branch_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetItemPurposes",
            Router: "/purposes/purpose/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "GetItemQuantity",
            Router: "/quantity/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ItemsController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ItemsController"],
        beego.ControllerComments{
            Method: "UpdateItemImage",
            Router: "/update-item-image/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ObjectController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ObjectController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ObjectController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ObjectController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:ObjectController"] = append(beego.GlobalControllerRouter["item_processor/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:PurposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:PurposesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:PurposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:PurposesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:PurposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:PurposesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:PurposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:PurposesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:PurposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:PurposesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:PurposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:PurposesController"],
        beego.ControllerComments{
            Method: "ChangeVisibility",
            Router: "/change-visibility/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:PurposesController"] = append(beego.GlobalControllerRouter["item_processor/controllers:PurposesController"],
        beego.ControllerComments{
            Method: "GetAllPurposesAndTheirItems",
            Router: "/items",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:UserController"] = append(beego.GlobalControllerRouter["item_processor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:UserController"] = append(beego.GlobalControllerRouter["item_processor/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:UserController"] = append(beego.GlobalControllerRouter["item_processor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:uid",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:UserController"] = append(beego.GlobalControllerRouter["item_processor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:uid",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:UserController"] = append(beego.GlobalControllerRouter["item_processor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:uid",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:UserController"] = append(beego.GlobalControllerRouter["item_processor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["item_processor/controllers:UserController"] = append(beego.GlobalControllerRouter["item_processor/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
