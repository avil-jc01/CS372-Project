package models

    import (
    )

    type TemplatePayload struct {

    	Authorized bool
    	Session Session
    	Account Account
    	Category Category
    	QuickCategorys []Category
    }
