package repo

import (
	"apigouniv/src/domain"
	"apigouniv/src/util"
	"errors"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

type MasterDataRepoImpl struct {
}

func CreateMasterDataRepoImpl() domain.MasterDataRepo {
	return &MasterDataRepoImpl{}
}

func (md *MasterDataRepoImpl) GetAllUniversity(param1, value1, param2, value2 string) ([]domain.Universities, error) {
	// Consume third API
	client := resty.New()          // Create client
	var resp []domain.Universities // Initialize new variable to catch response from 3rd party
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			param1: value1,
			param2: value2,
		}).
		SetResult(&resp).
		Get(viper.GetString("url.universities"))
	util.Log3rdParty(res.Request.Method, res.Request.URL, "", string(res.Body()))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New("data not found")
	} else if res.IsSuccess() {
		return resp, nil
	}

	return nil, nil
}

func (md *MasterDataRepoImpl) GetAllAPI(param1, value1, param2, value2 string) (*domain.PublicAPIs, error) {
	// Consume third API
	client := resty.New()      // Create client
	var resp domain.PublicAPIs // Initialize new variable to catch response from 3rd party
	res, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(map[string]string{
			param1: value1,
			param2: value2,
		}).
		SetResult(&resp).
		Get(viper.GetString("url.public-apis"))
	util.Log3rdParty(res.Request.Method, res.Request.URL, "", "TotalData: "+strconv.Itoa(resp.Count))
	if err != nil {
		return nil, err
	}

	if res.IsError() {
		return nil, errors.New("data not found")
	} else if res.IsSuccess() {
		return &resp, nil
	}

	return nil, nil
}
