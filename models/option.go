package models

type Option struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	AttributeId int64  `json:"attribute_id"`
}

type OptionCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	AttributeId int64  `json:"-"`
}

type OptionUpdateRequest struct {
	Id          int64  `json:"-"`
	Name        string `json:"name" validate:"required"`
	AttributeId int64  `json:"-"`
}

type OptionFilter struct {
	AttributeId int64
}
