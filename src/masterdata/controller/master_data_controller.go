package controller

import (
	"apigouniv/src/domain"
	"apigouniv/src/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MasterDataController struct {
	mdService domain.MasterDataService
}

func CreateMasterDataController(r *gin.Engine, mdService domain.MasterDataService) {
	MasterDataController := MasterDataController{mdService: mdService}
	v1 := r.Group("/master")
	{
		v1.GET("/university", MasterDataController.GetListUniversity)
		v1.GET("/api", MasterDataController.GetListAPIByAuth)
	}
}

// Get List University godoc
// @Tags Master Data
// @Summary List University
// @Description This is API to get list university
// @Accept json
// @Produce json
// @Param name query string false "Params 1"
// @Param country query string false "Params 2"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /master/university [get]
func (md *MasterDataController) GetListUniversity(c *gin.Context) {
	// Set params
	name := c.Query("name")
	ctr := c.Query("country")

	// Get Country
	result, err := md.mdService.FindUniversity(name, ctr)
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in FindUniversity")
		return
	}

	util.HandleSuccess(c, http.StatusOK, "000", "Data found", gin.H{
		"count": len(result),
		"data":  result,
	}, "TotalData: "+strconv.Itoa(len(result)), "Success")
}

// Get List API godoc
// @Tags Master Data
// @Summary List Public API
// @Description This is API to get list of public API by auth or category
// @Accept json
// @Produce json
// @Param auth query boolean false "Params 1"
// @Param category query string false "Params 2"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /master/api [get]
func (md *MasterDataController) GetListAPIByAuth(c *gin.Context) {
	// Set params
	auth := c.Query("auth")
	cat := c.Query("category")

	// Check isAuth and add kind of auth
	isAuth, _ := strconv.ParseBool(auth)
	var auths []string
	if !isAuth && auth != "" {
		auths = append(auths, "null")
	} else if isAuth && auth != "" {
		auths = append(auths, "apikey", "oauth")
	}

	// Get API
	result, err := md.mdService.FindAPI(auths, cat)
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in FindAPI")
		return
	}

	resp, _ := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Data found", resp, "TotalData: "+strconv.Itoa(result.Count), "Success")
}
