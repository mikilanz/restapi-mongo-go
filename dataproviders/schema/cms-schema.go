package schema

import "go.mongodb.org/mongo-driver/bson/primitive"

type CmsConfigs struct {
	Id              primitive.ObjectID `json:"id,omitempty"`
	CompanyCode     string             `json:"companyCode,omitempty" validate:"required"`
	PrimaryColor    string             `json:"primaryColor,omitempty"`
	BackgroundColor string             `json:"backgroundColor,omitempty"`
}
