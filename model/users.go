package model

type (
	Users struct {
		Id            string  `json:"id"`
		Username      string  `json:"username"`
		Email         string  `json:"email"`
		Password      string  `json:"password"`
		DepositAmount float64 `json:"deposit_amount"`
		Role          string  `json:"role"`
	}

	UsersResponse struct {
		Id            string  `json:"id"`
		Username      string  `json:"username"`
		Email         string  `json:"email"`
		DepositAmount float64 `json:"deposit_amount"`
		Role          string  `json:"role"`
	}
)
