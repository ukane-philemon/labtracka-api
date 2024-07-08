package admin

import (
	"time"

	"github.com/ukane-philemon/labtracka-api/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/**** Patient ****/

func (m *MongoDB) PatientLabStats(patientID string) (db.PatientStats, error) {
	return db.PatientStats{
		TotalNumberOfLabsVisited:     5,
		TotalNumberOfCompletedOrders: 20,
	}, nil
}

// Results returns all results for the patient with the specified email
// address.
func (m *MongoDB) Results(patientID string) ([]*db.LabResult, error) {
	nowUnix := time.Now().Unix()
	return []*db.LabResult{
		{
			ID:                  patientID,
			TestName:            "Malaria",
			LabName:             "Test Lab",
			Status:              db.ResultStatusPending,
			Data:                []string{"random base64 file or file url"},
			TurnAroundInSeconds: 60 * 60 * 4,
			CreatedAt:           nowUnix,
			LastUpdatedAt:       nowUnix,
		},
		{
			ID:                  patientID,
			TestName:            "Typhoid",
			LabName:             "Test Lab",
			Status:              db.ResultStatusInProgress,
			Data:                []string{"random base64 file or file url"},
			TurnAroundInSeconds: 60 * 60 * 4,
			CreatedAt:           nowUnix,
			LastUpdatedAt:       nowUnix,
		},
		{
			ID:                  patientID,
			TestName:            "Pregnancy",
			LabName:             "Test Lab",
			Status:              db.ResultStatusCompleted,
			Data:                []string{"random base64 file or file url"},
			TurnAroundInSeconds: 60 * 60 * 4,
			CreatedAt:           nowUnix,
			LastUpdatedAt:       nowUnix,
		},
	}, nil
}

/**** Labs ****/

// Labs returns a list of available labs.
func (m *MongoDB) Labs() ([]*db.BasicLabInfo, error) {
	return []*db.BasicLabInfo{
		{
			ID:      primitive.NewObjectID().Hex(),
			Name:    "Test Lab",
			LogoURL: "full path to logo url",
			Address: db.Address{
				Coordinates: "",
				HouseNumber: "29",
				StreetName:  "Test street, musa close",
				City:        "Port Harcourt",
				Country:     "Nigeria",
			},
			Featured: true,
		},
		{
			ID:      primitive.NewObjectID().Hex(),
			Name:    "Zion Test Lab",
			LogoURL: "full path to logo url",
			Address: db.Address{
				Coordinates: "",
				HouseNumber: "12",
				StreetName:  "Dynwell street",
				City:        "Port Harcourt",
				Country:     "Nigeria",
			},
			Featured: false,
		},
	}, nil
}

// LabTests returns a list of supported single lab tests and test packages
// for the lab with the provided labID. Returns an ErrorInvalidRequest if
// labID does not exist.
func (m *MongoDB) LabTests(labID string) (*db.LabTests, error) {
	nowUnix := time.Now().Unix()
	c1 := primitive.NewObjectID().Hex()
	c2 := primitive.NewObjectID().Hex()
	c3 := primitive.NewObjectID().Hex()
	return &db.LabTests{
		Categories: []*db.TestCategory{
			{
				ID:            c1,
				Name:          "Sexual Health",
				IsActive:      false,
				CreatedAt:     nowUnix,
				LastUpdatedAt: nowUnix,
			},
			{
				ID:            c2,
				Name:          "Female Health",
				IsActive:      false,
				CreatedAt:     nowUnix,
				LastUpdatedAt: nowUnix,
			},
			{
				ID:            c3,
				Name:          "Male Health",
				IsActive:      false,
				CreatedAt:     nowUnix,
				LastUpdatedAt: nowUnix,
			},
		},
		Tests: []*db.LabTest{
			{
				ID:                   primitive.NewObjectID().Hex(),
				Name:                 "Golden Package",
				LabID:                labID,
				LabName:              "Test Lab",
				Price:                45000,
				OldPrice:             0,
				Description:          "Golden Package is the perfect package you need to check your wellbeing",
				Gender:               "All",
				Categories:           []string{c1, c2, c3},
				IsDisabled:           false,
				SampleCollectionMode: []string{"Walk-In", "Home"},
				Tests:                []string{"test" + c1, "test" + c2, "test" + c3},
				CreatedAt:            nowUnix,
				LastUpdatedAt:        nowUnix,
			},
			{
				ID:                   "test" + c1,
				Name:                 "Malaria",
				LabID:                labID,
				LabName:              "Test Lab",
				Price:                3700,
				OldPrice:             0,
				Description:          "Malaria Test",
				Gender:               "All",
				SampleCollectionMode: []string{"Walk-In", "Home"},
				CreatedAt:            nowUnix,
				LastUpdatedAt:        nowUnix,
			},
		},
	}, nil
}

// LabTest returns the information of a lab test from the admin db.
func (m *MongoDB) LabTest(testID string) (*db.LabTest, error) {
	nowUnix := time.Now().Unix()
	return &db.LabTest{
		ID:                   testID,
		Name:                 "Malaria",
		LabID:                primitive.NewObjectID().Hex(),
		LabName:              "Test Lab",
		Price:                3700,
		OldPrice:             0,
		Description:          "Malaria Test",
		Gender:               "All",
		SampleCollectionMode: []string{"Walk-In", "Home"},
		CreatedAt:            nowUnix,
		LastUpdatedAt:        nowUnix,
	}, nil
}

/**** Server Info ****/

// Faqs returns information about frequently asked questions and help links.
func (m *MongoDB) Faqs() ([]*db.Faq, error) {
	return []*db.Faq{
		{
			ID: "01",
			FaqInfo: db.FaqInfo{
				Title:  "How does it work?",
				Answer: "Book a test and pay fot it, you'll have a phlebotomist at your door to collect your sample",
			},
			Hidden:    false,
			Timestamp: "12345678",
		},
	}, nil
}
