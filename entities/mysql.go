package entities

type Accounts struct {
	AccountId    uint   `gorm:"primary_key"`
	FirstName    string `gorm:"column:first_name"`
	LastName     string `gorm:"column:last_name"`
	Address      string `gorm:"column:address"`
	City         uint   `gorm:"column:city"`
	Province     uint   `gorm:"column:province"`
	Zipcode      string `gorm:"column:zipcode"`
	EmailAddress string `gorm:"column:email_address"`
	PhoneNumber  uint   `gorm:"column:phone_number"`
}

type Cities struct {
	CityId   uint   `gorm:"primary_key"`
	Name     string `gorm:"column:name"`
	Province uint   `gorm:"column:province"`
}

type Provinces struct {
	ProvinceId uint   `gorm:"primary_key"`
	Name       string `gorm:"column:name"`
}

type Credentials struct {
	CredentialId uint   `gorm:"primary_key"`
	AccountId    uint   `gorm:"column:account_id"`
	PasswordHash string `gorm:"column:password_hash"`
	Salt         string `gorm:"column:salt"`
}
