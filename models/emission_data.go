package models

type EmissionData struct {
	ID             uint    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	ActivityType   string  `gorm:"TYPE:VARCHAR(50)" json:"activity_type"`
	InputValue     float64 `json:"input_value"`
	EmissionFactor float64 `json:"emission_factor"`
	TotalEmission  float64 `json:"total_emission"`
}


// selectQuery := `
//         SELECT 
//             emission_datas.id,
//             emission_data.activity_type,
//             emission_data.input_value,
//             emission_data.emission_factor,
//             emission_data.input_value * emission_data.emission_factor AS total_emission
//         FROM emission_data
// `

// updateQuery := `
//             UPDATE emission_data SET total_emission = input_value * emission_datas.emission_factor
// `