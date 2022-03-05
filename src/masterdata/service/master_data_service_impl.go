package service

import (
	"apigouniv/src/domain"
	"errors"
)

type MasterDataServiceImpl struct {
	mdRepo domain.MasterDataRepo
}

func CreateMasterDataServiceImpl(mdRepo domain.MasterDataRepo) domain.MasterDataService {
	return &MasterDataServiceImpl{
		mdRepo: mdRepo,
	}
}

func (md *MasterDataServiceImpl) FindUniversity(name, ctr string) ([]domain.University, error) {
	var listResult []domain.University

	listData, err := md.mdRepo.GetAllUniversity("name", name, "country", ctr)
	if err != nil {
		return nil, err
	}

	for _, v := range listData {
		result := domain.University{
			Name:    v.Name,
			Country: v.Country,
			URL:     v.WebPages,
		}
		listResult = append(listResult, result)
	}
	return listResult, nil
}

func (md *MasterDataServiceImpl) FindAPI(auths []string, cat string) (*domain.PublicAPIs, error) {
	// get all public api
	if len(auths) == 0 && cat == "" {
		listData, err := md.mdRepo.GetAllAPI("auth", "", "category", "")
		if err != nil {
			return nil, err
		}

		return listData, nil
	}

	// get all public api by auth and category
	if len(auths) != 0 && cat != "" {
		listData, err := md.FindAPIByAuthCategory(auths, cat)
		if err != nil {
			return nil, err
		}

		return listData, nil
	}

	// Get public api by auth
	if len(auths) != 0 {
		listData, err := md.FindAPIByAuth(auths)
		if err != nil {
			return nil, err
		}
		return listData, nil
	}

	// Get public api by category
	if cat != "" {
		listData, err := md.FindAPIByCategory(cat)
		if err != nil {
			return nil, err
		}
		return listData, nil
	}

	return nil, errors.New("data not found")
}

func (md *MasterDataServiceImpl) FindAPIByAuth(auths []string) (*domain.PublicAPIs, error) {
	// Get public api by auth, its needed to be loop because if isAuth true, there's 2 kind of auth and we have to look up for it twice
	var listData domain.PublicAPIs
	for _, v := range auths {
		listAPI, err := md.mdRepo.GetAllAPI("auth", v, "category", "") // because we set QueryParams, we need to send another param even its value empty
		if err != nil {
			return nil, err
		}
		listData.TotalAuth = len(auths)
		listData.Auths = auths
		listData.Count += listAPI.Count
		listData.Entries = append(listData.Entries, listAPI.Entries...)
	}

	// Reassign data that match
	for _, v := range listData.Entries {
		if v.Auth == auths[listData.TotalAuth-1] {
			listData.Entries = append(listData.Entries, listData.Entries...)
		}
	}

	return &listData, nil
}

func (md *MasterDataServiceImpl) FindAPIByCategory(cat string) (*domain.PublicAPIs, error) {
	// Get public api by category
	listData, err := md.mdRepo.GetAllAPI("category", cat, "auth", "") // because we set QueryParams, we need to send another param even its value empty
	if err != nil {
		return nil, err
	}

	return listData, nil
}

func (md *MasterDataServiceImpl) FindAPIByAuthCategory(auths []string, cat string) (*domain.PublicAPIs, error) {
	// Get public api by auth, its needed to be loop because if isAuth true, there's 2 kind of auth and we have to look up for it twice
	var listData domain.PublicAPIs
	for _, v := range auths {
		listAPI, err := md.mdRepo.GetAllAPI("auth", v, "category", cat)
		if err != nil {
			return nil, err
		}
		listData.TotalAuth = len(auths)
		listData.Auths = auths
		listData.Count += listAPI.Count
		listData.Entries = append(listData.Entries, listAPI.Entries...)
	}

	// Reassign data that match
	for _, v := range listData.Entries {
		if v.Auth == auths[listData.TotalAuth-1] {
			listData.Entries = append(listData.Entries, listData.Entries...)
		}
	}

	return &listData, nil
}
