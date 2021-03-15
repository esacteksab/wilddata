package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Asset represents the asset model stored in our database
// This is used for assets, storing Name of asset, tags related
// to assets and the OrgID (OWner) of the asset.
type Asset struct {
	gorm.Model
	OrgID uint           `gorm:"not_null;index" json:"org_id"`
	Name  string         `gorm:"not null" json:"name"`
	Tags  datatypes.JSON `gorm: "-" json:"tags"`
}

// AssetService is a set of methods used to manipulate and work with
// the asset model
type AssetService interface {
	// Probably need a method here to look up if an asset by the same Name 
	// already exists if it doesn't, create it. If it does, append to the 
	// existing asset ID.
	AssetDB
}

// AssetDB is used to interface with the assets database.
// For pretty much all single asset queries:
// If asset is found, we will return a nil error
// If the asset is not found, we will return ErrNotFound
// If there is another error, we will return an error with
// more informationa bout what went wrong. 
//
// For single asset queries, any error but ErrNotFound should
// probably resturn in a 500 error.
type AssetDB interface {
	ByID(id uint) (*Asset, error)
	ByOrgID(orgid uint) ([]Asset, error)
	Create(asset *Asset) error
	Update(asset *Asset) error
	Delete(id uint) error
}

func NewAssetService(db *gorm.DB) AssetService {
	return &assetService{
		AssetDB: &AssetValidator{&AssetGorm{db}},
	}
}

type assetService struct {
	AssetDB
}

func NewAssetService(db *gorm.DB) AssetService {
	ag := &AssetGorm{db}
	return &assetService{
		AssetDB: ag,
	}
}

var _ AssetService = &assetService{}

type assetValidator struct {
	AssetDB
}

func (av *assetValidator) Create(asset *Asset) error {
	err := runAssetValFuncs(asset,
		av.orgIDRequired,
		av.nameRequired)
	if err != nil {
		return err
	}
	return av.AssetDB.Create(asset)
}

func (av *assetValidator) Update(asset *Asset) error {
	err := runAssetValFuncs(asset,
		av.orgIDRequired,
		av.nameRequired)
	if err != nil {
		return err
	}
	return av.AssetDB.Update(asset)
}

func (av *assetValidator) orgIDRequired(a *Asset) error {
	if av.OrgID <= 0 {
		return ErrOrgIDRequired
	}
	return nil
}

func (av *assetValidator) nameRequired(a *Asset) error {
	if av.Name == "" {
		return ErrNameRequired
	}
	return nil
}

type assetValFunc func(*Asset) error

func runAssetValFuncs(asset *Asset, fns ...assetValFunc) err {
	for _, fn :range fns {
		if err := fn(asset); err != nil {
			return err
		}
	}
	return nil
}