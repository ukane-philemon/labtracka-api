package db

import "errors"

// CreateLabRequest is like CreateLabRequestInfo but with unique ID.
type CreateLabRequest struct {
	ID string `json:"id"`
	CreateLabRequestInfo
	Timestamp int64 `json:"timestamp"`
}

// CreateLabRequestInfo is information required to register a laboratory.
type CreateLabRequestInfo struct {
	ManagerName           string   `json:"manager_name" bson:"manager_name"`
	ManagerEmail          string   `json:"manager_email" bson:"manager_email"`
	ManagerPhoneNumber    string   `json:"manager_phone_number" bson:"manager_phone_number"`
	LabName               string   `json:"lab_name" bson:"lab_name"`
	LabRegistrationNumber string   `json:"lab_registration_number" bson:"lab_registration_number"`
	LabAddress            *Address `json:"lab_address" bson:"lab_address"`
}

// LaboratoryInfo is information about a lab view.
type LaboratoryInfo struct {
	ID        string   `json:"id"`
	DisplayID string   `json:"display_id"`
	LogoURL   string   `json:"logo_url" bson:"logo_url"`
	Name      string   `json:"name"`
	Address   *Address `json:"address"`
	Disabled  bool     `json:"disabled"`
	Featured  bool     `json:"featured"`

	RegistrationNumber string `json:"registration_number"`
	ManagerName        string `json:"manager_name" bson:"manager_name"`
	ManagerEmail       string `json:"manager_email" bson:"manager_email"`
	ManagerPhoneNumber string `json:"manager_phone_number" bson:"manager_phone_number"`

	CreatedAt     int64 `json:"created_at"`
	LastUpdatedAt int64 `json:"last_updated_at"`
}

// Decision is a mandate to approve or reject and action.
type Decision struct {
	Approve bool   `json:"approve"`
	Reject  bool   `json:"reject"`
	Reason  string `json:"reason"`
}

// Validate checks that this Decision is valid.
func (d *Decision) Validate() error {
	if d.Approve == d.Reject {
		return errors.New("approve or reject action, not both")
	}

	if d.Reject && d.Reason == "" {
		return errors.New("reason is compulsory to reject action")
	}

	return nil
}

type AdminStats struct {
	TotalUsers         int64 `json:"total_users"`
	TotalRevenue       int64 `json:"total_revenue"`
	TotalTestOrders    int64 `json:"total_test_orders"`
	TotalPendingOrders int64 `json:"total_pending_orders"`
}

type AdminPatientInfo struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	Address       *Address `json:"address"`
	Disabled      bool     `json:"disabled"`
	CreatedAt     int64    `json:"created_at" bson:"created_at"`
	LastUpdatedAt int64    `json:"last_updated_at" bson:"last_updated_at"`
}

type WithdrawalRequestStatus string

const (
	WithdrawalRequestStatusPending  WithdrawalRequestStatus = "Pending"
	WithdrawalRequestStatusApproved WithdrawalRequestStatus = "Approved"
	WithdrawalRequestStatusDeclined WithdrawalRequestStatus = "Declined"
)

type WithdrawalRequest struct {
	ID          string                  `json:"id"`
	InvoiceID   string                  `json:"invoice_id"`
	Amount      float64                 `json:"amount"`
	Status      WithdrawalRequestStatus `json:"status"`
	LabID       string                  `json:"lab_id"`
	LabName     string                  `json:"lab_name"`
	LabBranchID string                  `json:"lab_branch_id"`
	AdminID     string                  `json:"admin_id"`
	AdminName   string                  `json:"admin_name"`
	Receipt     *WithdrawalReceipt      `json:"receipt"`
	Timestamp   int64                   `json:"timestamp"`
}

type TransactionStatus string

const (
	TransactionStatusSuccessful TransactionStatus = "Successful"
	TransactionStatusFailed     TransactionStatus = "Failed"
	TransactionStatusPending    TransactionStatus = "Pending"
)

type WithdrawalReceipt struct {
	Status                TransactionStatus `json:"status"`
	AmountPaid            float64           `json:"amount_paid"`
	Fees                  float64           `json:"fees"`
	TransactionsReference string            `json:"transaction_reference"`
	Timestamp             int64             `json:"timestamp"`
}
