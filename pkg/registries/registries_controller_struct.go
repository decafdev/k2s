package registries

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/techdecaf/k2s/v2/pkg/global"
)

// NewRegistryController function description
func NewRegistryController(app *gin.Engine, registry *RegistryService) *RegistryController {
	controller := &RegistryController{
		registry: registry,
	}

	// register routes
	router := app.Group("/registries")

	router.GET("", controller.ListRegistries)

	router.GET("/:name", controller.GetRegistry)

	// create or update registry data
	router.PUT("", controller.CreateRegistrySecret)
	// copy registry to new namespace
	router.PUT("/:name/copy-to/:namespace", controller.CopyRegistrySecret)

	return controller
}

// RegistryController struct
type RegistryController struct {
	registry *RegistryService
}

// @Summary list managed docker registry secrets
// @Description list managed docker registry secrets
// @Accept application/json
// @Produce json
// @Success 200 {object} []PrivateRegistryDTO
// @Router /registries [GET]
func (t *RegistryController) ListRegistries(context *gin.Context) {
	if body, err := t.registry.ListPrivateRegisties(); err != nil {
		global.GinerateError(context, global.KubeError(err))
	} else {
		context.JSON(http.StatusOK, body)
	}
}

// @Summary get managed docker registry secret by name
// @Description get managed docker registry secret by name
// @Accept application/json
// @Produce json
// @Success 200 {object} PrivateRegistryDTO
// @Router /registries/:name [GET]
func (t *RegistryController) GetRegistry(context *gin.Context) {
	name := context.Param("name")
	if body, err := t.registry.ListPrivateRegisties(); len(*body) > 0 {
		for _, v := range *body {
			if v.Name == name {
				context.JSON(http.StatusOK, v)
				return
			}
		}
	} else {
		global.GinerateError(context, global.KubeError(err))
	}
	err := fmt.Errorf("registry [name: %s], not found", name)
	global.GinerateError(context, global.NotFoundError(err))
}

// @Summary Create a name managed docker registry secret
// @Description Create a name managed docker registry secret
// @Accept application/json
// @Produce json
// @Param CreateRegistryDTO body CreateRegistryDTO true "create registry request body"
// @Success 200 {object} map[string]interface{}
// @Router /registries [PUT]
func (t *RegistryController) CreateRegistrySecret(context *gin.Context) {
	// var registry CreateRegistryDTO

	// if err := context.ShouldBind(&registry); err != nil {
	// 	global.GinerateError(context, global.BadRequestError(err))
	// 	return
	// }

	// body, err := t.registry.CreateRegistrySecret(&registry)
	// if err != nil {
	// 	global.GinerateError(context, global.KubeError(err))
	// 	return
	// }

	// context.JSON(http.StatusOK, body)
}

// @Summary copy a managed docker registry secret from the k2s namespace to another namespace
// @Description copy a managed docker registry secret from the k2s namespace to another namespace
// @Accept application/json
// @Produce json
// @Success 200 {object} PrivateRegistryDTO{}
// @Router /registries/:name/copy-to/:namespace [PUT]
func (t *RegistryController) CopyRegistrySecret(context *gin.Context) {

	// if body, err := t.registry.CopyRegistry(context.Param("name"), context.Param("namespace")); err != nil {
	// 	global.GinerateError(context, global.KubeError(err))
	// 	return
	// } else {
	// 	context.JSON(http.StatusOK, body)
	// }
}
