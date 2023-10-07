package emailn

import (
	"emailn/internal/domain/campaign"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	contacts := []campaign.Contacts{campaign.Contacts{Email: "emailqualquer@teste"}}
	campaign := campaign.Campaign{Contacts: contacts}
	validate := validator.New()

	err := validate.Struct(campaign)
	if err == nil {
		fmt.Println("Nenhum erro")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {
			// fmt.Println(v.Error())

			switch v.Tag() {
			case "required":
				fmt.Println(v.StructField() + " is required")
			case "min":
				fmt.Println(v.StructField() + " is required with min " + v.Param())
			case "max":
				fmt.Println(v.StructField() + " is required with max " + v.Param())
			case "email":
				fmt.Println(v.StructField() + " is invalid")
			}

		}
	}
}
