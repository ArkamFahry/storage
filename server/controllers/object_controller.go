package controllers

import (
	"github.com/ArkamFahry/storage/server/models"
	"github.com/ArkamFahry/storage/server/services"
	"github.com/gofiber/fiber/v2"
)

type ObjectController struct {
	objectService *services.ObjectService
}

func NewObjectController(objectService *services.ObjectService) *ObjectController {
	return &ObjectController{
		objectService: objectService,
	}
}

func (oc *ObjectController) RegisterObjectRoutes(app *fiber.App) {
	routes := app.Group("/api")

	routesV1 := routes.Group("/v1")

	routesV1.Post("/objects/:bucket_name/pre-signed/upload", oc.CreatePreSignedUploadObject)
	routesV1.Post("/objects/:bucket_name/pre-signed/upload/:object_id/complete", oc.CompletePreSignedObjectUpload)
	routesV1.Get("/objects/:bucket_name/pre-signed/download/:object_id", oc.CreatePreSignedDownloadObject)
	routesV1.Delete("/objects/:bucket_name/:object_id", oc.DeleteObject)
	routesV1.Get("/objects/:bucket_name/:object_id", oc.GetObject)
	routesV1.Get("/objects/:bucket_name/:object_path", oc.SearchObjects)
}

// CreatePreSignedUploadObject is used to create a pre signed upload object
// @Summary Create a pre signed upload object
// @Description Create a pre signed upload object
// @Tags objects
// @Accept json
// @Produce json
// @Param bucket_name path string true "Bucket Name"
// @Param bucket body models.PreSignedUploadObjectCreate true "Pre Signed Upload Object Create"
// @Success 201 {object} models.PreSignedUploadObject
// @Failure 400 {object} middleware.HttpError
// @Failure 500 {object} middleware.HttpError
// @Router /api/v1/objects/{bucket_name}/pre-signed/upload [post]
func (oc *ObjectController) CreatePreSignedUploadObject(ctx *fiber.Ctx) error {
	var preSignedUploadObjectCreate models.PreSignedUploadObjectCreate

	bucketName := ctx.Params("bucket_name")

	err := ctx.BodyParser(&preSignedUploadObjectCreate)
	if err != nil {
		return err
	}

	preSignedUploadObject, err := oc.objectService.CreatePreSignedUploadObject(ctx.Context(), bucketName, &preSignedUploadObjectCreate)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(preSignedUploadObject)
}

// CompletePreSignedObjectUpload is used to complete a pre signed upload session
// @Summary Complete a pre signed upload session
// @Description Complete a pre signed upload session
// @Tags objects
// @Accept json
// @Produce json
// @Param bucket_name path string true "Bucket Name"
// @Param object_id path string true "Object ID"
// @Success 202
// @Failure 400 {object} middleware.HttpError
// @Failure 500 {object} middleware.HttpError
// @Router /api/v1//objects/{bucket_name}/pre-signed/upload/{object_id}/complete [post]
func (oc *ObjectController) CompletePreSignedObjectUpload(ctx *fiber.Ctx) error {
	bucketName := ctx.Params("bucket_name")
	objectId := ctx.Params("object_id")

	err := oc.objectService.CompletePreSignedObjectUpload(ctx.Context(), bucketName, objectId)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusAccepted)
}

// CreatePreSignedDownloadObject is used to create a pre signed download object
// @Summary Create a pre signed download object
// @Description Create a pre signed download object
// @Tags objects
// @Accept json
// @Produce json
// @Param bucket_name path string true "Bucket Name"
// @Param object_id path string true "Object ID"
// @Param expires_in query int true "Expires In"
// @Success 201 {object} models.PreSignedDownloadObject
// @Failure 400 {object} middleware.HttpError
// @Failure 500 {object} middleware.HttpError
// @Router /api/v1/objects/{bucket_name}/pre-signed/download/{object_id} [post]
func (oc *ObjectController) CreatePreSignedDownloadObject(ctx *fiber.Ctx) error {
	bucketName := ctx.Params("bucket_name")
	objectId := ctx.Params("object_id")

	expiresIn := ctx.QueryInt("expires_in")

	preSignedDownloadObject, err := oc.objectService.CreatePreSignedDownloadObject(ctx.Context(), bucketName, objectId, int64(expiresIn))
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(preSignedDownloadObject)
}

// DeleteObject is used to delete an object
// @Summary Delete an object
// @Description Delete an object
// @Tags objects
// @Accept json
// @Produce json
// @Param bucket_name path string true "Bucket Name"
// @Param object_id path string true "Object ID"
// @Success 204
// @Failure 400 {object} middleware.HttpError
// @Failure 500 {object} middleware.HttpError
// @Router /api/v1/objects/{bucket_name}/{object_id} [delete]
func (oc *ObjectController) DeleteObject(ctx *fiber.Ctx) error {
	bucketName := ctx.Params("bucket_name")
	objectId := ctx.Params("object_id")

	err := oc.objectService.DeleteObject(ctx.Context(), bucketName, objectId)
	if err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// GetObject is used to get an object
// @Summary Get an object
// @Description Get an object
// @Tags objects
// @Accept json
// @Produce json
// @Param bucket_name path string true "Bucket Name"
// @Param object_id path string true "Object ID"
// @Success 200 {object} models.Object
// @Failure 400 {object} middleware.HttpError
// @Failure 500 {object} middleware.HttpError
// @Router /api/v1/objects/{bucket_name}/{object_id} [get]
func (oc *ObjectController) GetObject(ctx *fiber.Ctx) error {
	bucketName := ctx.Params("bucket_name")
	objectId := ctx.Params("object_id")

	object, err := oc.objectService.GetObject(ctx.Context(), bucketName, objectId)
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(object)
}

// SearchObjects is used to search objects
// @Summary Search objects by path
// @Description Search objects by path
// @Tags objects
// @Accept json
// @Produce json
// @Param bucket_name path string true "Bucket Name"
// @Param object_path path string true "Object Path"
// @Param levels query int true "Levels"
// @Param limit query int true "Limit"
// @Param offset query int true "Offset"
// @Success 200 {array} models.Object
// @Failure 400 {object} middleware.HttpError
// @Failure 500 {object} middleware.HttpError
// @Router /api/v1/objects/{bucket_name}/{object_path} [get]
func (oc *ObjectController) SearchObjects(ctx *fiber.Ctx) error {
	bucketName := ctx.Params("bucket_name")
	objectPath := ctx.Params("object_path")

	levels := ctx.QueryInt("levels")
	limit := ctx.QueryInt("limit")
	offset := ctx.QueryInt("offset")

	objects, err := oc.objectService.SearchObjects(ctx.Context(), bucketName, objectPath, int32(levels), int32(limit), int32(offset))
	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusOK).JSON(objects)
}