package schema

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CmsConfigs struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id, omitempty"`
	CompanyCode     string             `json:"companyCode,omitempty" validate:"required"`
	Ips             []string           `json:"ips,omitempty"`
	HeadLink        []interface{}      `json:"headLink,omitempty"`
	Config          interface{}        `json:"config,omitempty"`
	Keywords        []string           `json:"keywords,omitempty"`
	PrimaryColor    string             `json:"primaryColor,omitempty"`
	BackgroundColor string             `json:"backgroundColor,omitempty"`
	Display         string             `json:"display,omitempty"`
	ThemeColor      string             `json:"themeColor,omitempty"`
	CompanyLogoUrl  string             `json:"companyLogoUrl,omitempty"`
	Scope           string             `json:"scope,omitempty"`
	CreatedAt       time.Time          `json:"createdAt,omitempty"`
}
