package dto

type CreateAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type PatchAccountRequest struct {
	Name string `json:"name"`
}

type ChangeAccountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type DeleteAccountRequest struct {
	Name string `json:"name"`
}

type ChangeAccountNameRequest struct {
	Name    string `json:"name"`
	NameNew string `json:"nameNew"`
}

type ChangeAccountBalanceRequest struct {
	Name      string `json:"name"`
	AmountNew int    `json:"amountNew"`
}
