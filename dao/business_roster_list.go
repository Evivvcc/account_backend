package dao

type BusinessRosterList struct {
}

func (BusinessRosterList) TableName() string {
	return ""
}

func (b *BusinessRosterList) GetOne() error {
	return nil
}
