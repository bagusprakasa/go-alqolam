package region

import "time"

type RegionFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatRegion(region Region) RegionFormatter {
	formatter := RegionFormatter{
		ID:        region.ID,
		Name:      region.Name,
		CreatedAt: region.CreatedAt,
		UpdatedAt: region.UpdatedAt,
	}

	return formatter
}

func FormatRegions(region []Region) []RegionFormatter {
	regionsFormatter := []RegionFormatter{}

	for _, region := range region {
		regionFormatter := RegionFormatter(region)
		regionsFormatter = append(regionsFormatter, regionFormatter)
	}

	return regionsFormatter
}
