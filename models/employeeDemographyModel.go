package models

type EmployeeDemography struct {
	DateKey            string
	TenantUuid         string
	DemographyCategory string
	DemographyName     string
	DemographyValue    float64
}

type Tabler interface {
	TableName() string
}

func (EmployeeDemography) TableName() string {
	return "dm_employee_demography"
}
