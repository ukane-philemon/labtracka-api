package admin

import "github.com/ukane-philemon/labtracka-api/db"

// UpdateFaqs updates the faqs in the database. This is a super admin only
// feature.
func (m *MongoDB) UpdateFaqs(faq []*db.Faq) ([]*db.Faq, error) {
	return nil, nil
}
