package member

import "time"

type MemberFormatter struct {
	ID       int `json:"id"`
	RegionID int `json:"region_id"`
	// Region    RegionFormatter `json:"region"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Gender    string    `json:"gender"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RegionFormatter struct {
	Name string `json:"name"`
}

func FormatMember(member Member) MemberFormatter {
	// member = region.Member
	// regionFormatter := RegionFormatter{}
	formatter := MemberFormatter{
		ID:        member.ID,
		RegionID:  member.RegionID,
		Name:      member.Name,
		Phone:     member.Phone,
		Address:   member.Address,
		Gender:    member.Gender,
		CreatedAt: member.CreatedAt,
		UpdatedAt: member.UpdatedAt,
		// region:    region.Name,
	}

	return formatter
}

func FormatMembers(member []Member) []MemberFormatter {
	membersFormatter := []MemberFormatter{}

	for _, member := range member {
		memberFormatter := FormatMember(member)
		membersFormatter = append(membersFormatter, memberFormatter)
	}

	return membersFormatter
}
