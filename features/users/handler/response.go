package handler

import "Absensi-App/features/users"

type LoginResponse struct {
	ID    uint   `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Role  string `json:"role,omitempty"`
	Token string `json:"token,omitempty"`
}

type UserResponse struct {
	ID           uint               `json:"id,omitempty"`
	Name         string             `json:"name,omitempty"`
	Email        string             `json:"email,omitempty"`
	PhoneNumber  string             `json:"phone_number,omitempty"`
	Password     string             `json:"password,omitempty"`
	Address      string             `json:"address,omitempty"`
	ProfilePhoto string             `json:"profil_photo,omitempty"`
	UploadKTP    string             `json:"image,omitempty"`
	MembershipID uint               `json:"membership_id,omitempty"`
	Membership   MembershipResponse `json:"membership,omitempty"`
}

type MembershipResponse struct {
	JenisMembership string `json:"jenis_membership"`
	Status          string `json:"status"`
}

func UserCoreToResponse(input users.UserCore) UserResponse {
	return UserResponse{
		ID:           input.ID,
		Name:         input.Name,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		Password:     input.Password,
		Address:      input.Address,
		ProfilePhoto: input.ProfilePhoto,
		UploadKTP:    input.UploadKTP,
		MembershipID: input.MembershipID,
		Membership:   MembershipToResponse(input.Membership),
	}
}

func MembershipToResponse(input users.MembershipCore) MembershipResponse {
	return MembershipResponse{
		JenisMembership: input.JenisMembership,
		Status:          input.Status,
	}
}
