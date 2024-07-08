package admin

import "github.com/ukane-philemon/labtracka-api/db"

// AdminLabs returns a list of labs added to the db for only super admin. The
// provided adminID must match a super admin.
func (m *MongoDB) AdminLabs(adminID string) ([]*db.LaboratoryInfo, error) {
	return nil, nil
}
