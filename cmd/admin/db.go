package admin

import "github.com/ukane-philemon/labtracka-api/db"

type Database interface {
	/**** AUTHENTICATION METHODS ****/

	// LoginAdmin logs an admin into their account. Returns an
	// ErrorInvalidRequest is user email or password is invalid/not correct or
	// does not exist or an ErrorOTPRequired if otp validation is required for
	// this account.
	LoginAdmin(req *db.LoginRequest) (*db.Admin, error)
	// ResetPassword reset the password of an existing admin. Returns an
	// ErrorInvalidRequest if the adminEmail is not tied to an existing admin.
	ResetPassword(adminEmail, password string) error
	// ChangePassword updates the password for an existing admin. Returns an
	// ErrorInvalidRequest if adminID is not tied to an existing admin or
	// current password is incorrect.
	ChangePassword(adminID, currentPassword, newPassword string) error
	// UpdateAdminInfo updates the admin with the provided adminID if the
	// parameters are provided.
	// UpdateAdminInfo(adminID string, profileImageURL string, firstName, lastName, phoneNumber string) error
	// WithdrawalRequests returns withdrawal request information for the
	// provided statuses. branchID is compulsory if adminID is not a server
	// admin but optional for server admin.
	// WithdrawalRequests(adminID, branchID string, status ...db.WithdrawalRequestStatus) (map[db.WithdrawalRequestStatus][]*db.WithdrawalRequest, error)

	/**** AUTHENTICATION METHODS END ****/

	/**** UNAUTHENTICATED METHODS ****/

	// SaveCreateLabRequests saves a new create lab request that would be
	// reviewed and approved by a Server Admin. Returns ErrorInvalidRequest if a
	// super admin with the manager email exists or if the lab RC number exits.
	// SaveCreateLabRequests(req *db.CreateLabRequestInfo) error

	/****  UNAUTHENTICATED METHODS END ****/

	/**** SERVER ADMIN METHODS ****/

	/**** SERVER ADMIN END ****/

	// Notifications returns all the notifications for patient sorted by unread
	// first.
	Notifications(email string) ([]*db.Notification, error)
	// MarkNotificationsAsRead marks the notifications with the provided noteIDs
	// as read.
	MarkNotificationsAsRead(email string, noteIDs ...string) error
	// Faqs returns information about frequently asked questions and help links.
	Faqs() ([]*db.Faq, error)
	Shutdown()
}

type ServerAdminDatabase interface {
	// PendinCreateLabRequests retrieves all the create lab requests that have
	// not been approved.
	PendinCreateLabRequests() ([]*db.CreateLabRequest, error)
	// UpdateLabRequest approves or rejects a request to  register a laboratory.
	// Rejected requests are deleted from the db.
	UpdateLabRequest(requestID string, decision *db.Decision) (*db.CreateLabRequest, error)
	// CreateLab adds a new laboratory to the database. Returns
	// ErrorInvalidRequest if a super admin with the manager email exists or if
	// the lab RC number exits.
	CreateLab(request *db.CreateLabRequestInfo) error
	// AdminLabs returns a list of labs added to the db for only super admin.
	// The provided adminID must match a super admin.
	AdminLabs() ([]*db.LaboratoryInfo, error)
	// ToggleLabFeaturedStatus sets a lab as featured or not.
	ToggleLabFeaturedStatus(labID string, feature bool) error
	AddFaq(faq *db.FaqInfo) ([]*db.Faq, error)
	HideFaq(faqID string, hide bool) ([]*db.Faq, error)
	DeleteFaq(faqID string) ([]*db.Faq, error)
	Admins() []*db.Admin
	ToggleLabStatus(labID string, disable bool) error
	ToggleAdminStatus(adminID string, disable bool) error
	TogglePatientStatus(patientID string, disable bool) error
	AddTest()
	ToggleTestStatus(testID string, disable bool) error
	DeleteTest(testID string) ([]*db.LabTest, error)
}

// CompleteInfoReader combines information from the admin and the patient db to
// give complete information.
type CompleteInfoReader interface {
	// AdminPatientInfo returns data for all patients if id is not provided.
	AdminPatientInfo(adminID string, id ...string) (*db.AdminPatientInfo, error)
	// WalletBalance returns the wallet balance for the provided labIDs. This is
	// paid total orders - approved and paid withdrawal requests. Only super
	// admin is allowed to read total wallet balance. labIDs cannot be empty if
	// not server admin.
	WalletBalance(adminID string, labIDs ...string) (float64, error)
}

type PatientDatabase interface {
	// Orders returns all the orders pending orders in the patient db for the
	// provided labIDs.
	Orders(labIDs ...string) (map[string]map[string][]*db.Order, error)
	// AdminStats returns some admin stats for display. If no lab id is
	// returned, all current stats will be returned.
	AdminStats(isBranch bool, labIDs ...string) (map[string]db.AdminStats, error)
}
