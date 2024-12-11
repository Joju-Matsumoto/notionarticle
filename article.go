package notionarticle

import (
	"github.com/jomei/notionapi"
)

type Article struct {
	Page   notionapi.Page   `json:"page"`
	Blocks notionapi.Blocks `json:"blocks"`
}
