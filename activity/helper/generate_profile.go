package helper

import (
	"gitlab.com/stream/buffers/common"
	pb "gitlab.com/stream/buffers/profile"
)

func GenerateSampleProfile() *pb.ProfileInfo {
	return &pb.ProfileInfo{
		Type: "user",
		Id:   "4870",
		Name: &pb.Name{
			FirstName: "Axel",
			LastName:  "Rose",
		},
		Gender: "Male",
		Address: &common.Address{
			Type: "address",
			Division: &common.CodeSystem{
				Code:    "123",
				Display: "Dhaka",
				Ref:     "Division",
			},
			District: &common.CodeSystem{
				Code:    "1234",
				Display: "Dhaka",
				Ref:     "District",
			},
			Area: &common.CodeSystem{
				Code:    "2345",
				Display: "Banasree",
				Ref:     "Area",
			},
			Street: "Road#08 , Block#E",
			Zip:    1212,
		},
		Hobbies:    []string{"Guitar", "Books"},
		Dob:        "4-12-1993",
		Registered: "Registered",
		Picture: &pb.Picture{
			Profile: "Profile Picture",
			Cover:   "Cover Picture",
		},
		WorkEducation: []*pb.WorkEducation{
			&pb.WorkEducation{
				Type:       "Work",
				Place:      "Shimahin Ltd .",
				Department: "Software Developer",
				Location: &common.Address{
					Type: "address",
					Division: &common.CodeSystem{
						Code:    "123",
						Display: "Dhaka",
						Ref:     "Division",
					},
					District: &common.CodeSystem{
						Code:    "1234",
						Display: "Dhaka",
						Ref:     "District",
					},
					Area: &common.CodeSystem{
						Code:    "2345",
						Display: "Banani",
						Ref:     "Area",
					},
					Street: "Road#05 , Block#G",
					Zip:    1212,
				},
			},
		},
		TagLine:  "Happy",
		Verified: true,
	}
}
