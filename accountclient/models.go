package accountclient

// AccountData contains the data for an account
type AccountData struct {
    Data Account `json:"data"`
}

// Account represents a registered bank account
type Account struct {
    AccountType    string `json:"type"`
    ID             string `json:"id"`
    OrganisationID string `json:"organisation_id"`
    Attributes     AccountAttributes
}

// AccountAttributes are the attributes of an account
type AccountAttributes struct {
    Country       string `json:"country"`
    BaseCurrency  string `json:"base_currency"`
    AccountNumber string `json:"account_number"`
    BankID        string `json:"bank_id"`
    BankIDCode    string `json:"bank_id_code"`
    Bic           string `json:"bic"`
    Iban          string `json:"iban"`
    Title string `json:"title"`
    FirstName string `json:"first_name"`
    BankAccountName string `json:"bank_account_name"`
    AlternativeBankAccountNames []string `json:"alternative_bank_account_names"`
    AccountClassification string `json:"account_classification"`
    JointAccount bool `json:"joint_account"`
    AccountMatchingOptOut bool `json:"account_matching_opt_out"`
    SecondaryIdentification string `json:"secondary_identification"`
}