package serializer

import "go-store/model"

type Address struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	CreatedAt int64  `json:"created_at"`
}

func BuildAddress(address *model.Address) Address {
	return Address{
		ID:        int(address.ID),
		UserID:    address.UserID,
		Name:      address.Name,
		Phone:     address.Phone,
		Address:   address.Address,
		CreatedAt: address.CreatedAt.Unix(),
	}
}

func BuildAddresses(item []*model.Address) (address []Address) {
	for _, item := range item {
		address = append(address, BuildAddress(item))
	}
	return address
}
