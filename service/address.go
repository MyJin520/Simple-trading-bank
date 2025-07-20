package service

import (
	"context"
	"go-store/dao"
	"go-store/model"
	"go-store/pkg/e"
	"go-store/serializer"
	"net/http"
	"strconv"
)

type AddressService struct {
	Address string `json:"address" form:"address"`
	Name    string `json:"name" form:"name"`
	Phone   string `json:"phone" form:"phone"`
}

func (s *AddressService) CreateAddressService(context context.Context, uid int) interface{} {
	var address *model.Address
	code := http.StatusOK
	addressDao := dao.NewAddressDao(context)
	address = &model.Address{
		UserID:  uid,
		Address: s.Address,
		Name:    s.Name,
		Phone:   s.Phone,
	}
	err := addressDao.CreateAddress(address)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMSG(code),
	}
}

func (s *AddressService) GetAddressByID(ctx context.Context, aid string) serializer.Response {
	addressId, _ := strconv.Atoi(aid)
	code := http.StatusOK
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressByID(addressId)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    "查询地址不存在",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMSG(code),
		Data:   serializer.BuildAddress(address),
	}
}

func (s *AddressService) ShowAddress(ctx context.Context, id int) serializer.Response {
	code := http.StatusOK
	addressDao := dao.NewAddressDao(ctx)
	addressList, err := addressDao.ListAddressByID(id)
	if err != nil {
		code = http.StatusBadRequest
		return serializer.Response{
			Status: code,
			Msg:    e.GetMSG(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildAddresses(addressList),
		Msg:    e.GetMSG(code),
	}
}

func (s *AddressService) UpdateAddress(ctx context.Context, uid int, aID string) interface{} {
	var address *model.Address
	addressDao := dao.NewAddressDao(ctx)
	address = &model.Address{
		Address: s.Address,
		Name:    s.Name,
		Phone:   s.Phone,
		UserID:  uid,
	}
	addressID, _ := strconv.Atoi(aID)
	err := addressDao.UpdateAddressByID(uid, addressID, address)
	if err != nil {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    e.GetMSG(http.StatusBadRequest),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    e.GetMSG(http.StatusOK),
	}
}

func (s *AddressService) DeleteAddress(ctx context.Context, uid int, aid string) interface{} {
	addressDao := dao.NewAddressDao(ctx)
	addrID, _ := strconv.Atoi(aid)
	err := addressDao.DeleteAddressByID(addrID, uid)
	if err != nil {
		return serializer.Response{
			Status: http.StatusBadRequest,
			Msg:    e.GetMSG(http.StatusBadRequest),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: http.StatusOK,
		Msg:    e.GetMSG(http.StatusOK),
	}
}
