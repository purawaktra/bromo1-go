package module

import (
	"fmt"
	"github.com/purawaktra/bromo1/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BromoRepo struct {
	db *gorm.DB
}

func CreateBromoRepo(gorm *gorm.DB) BromoRepo {
	return BromoRepo{
		db: gorm,
	}
}

type BromoRepoInterface interface {
	SelectAccountById(query entities.Accounts, offset int) ([]entities.Accounts, error)
	SelectAccountByFirstName(query entities.Accounts, offset int) ([]entities.Accounts, error)
	SelectAccountByLastName(query entities.Accounts, offset int) ([]entities.Accounts, error)
	SelectAccountByEmail(query entities.Accounts, offset int) ([]entities.Accounts, error)
	SelectAllAccount(offset int) ([]entities.Accounts, error)
	InsertSingleAccount(query entities.Accounts) (entities.Accounts, error)
	InsertMultipleAccount(queries []entities.Accounts) ([]entities.Accounts, []error)
	UpdateSingleAccountById(query entities.Accounts) (entities.Accounts, error)
	UpdateMultipleAccountById(queries []entities.Accounts) ([]entities.Accounts, []error)
	DeleteSingleAccountById(query entities.Accounts) (entities.Accounts, error)
	DeleteMultipleAccountById(queries []entities.Accounts) ([]entities.Accounts, []error)
	SelectCityById(query entities.Cities, offset int) ([]entities.Cities, error)
	SelectCityByProvince(query entities.Cities, offset int) ([]entities.Cities, error)
	SelectAllCity(offset int) ([]entities.Cities, error)
	SelectProvinceById(query entities.Provinces, offset int) ([]entities.Provinces, error)
	SelectAllProvince(offset int) ([]entities.Provinces, error)
	SelectCredentialByAccountId(query entities.Credentials, offset int) ([]entities.Credentials, error)
	InsertSingleCredential(query entities.Credentials) (entities.Credentials, error)
	InsertMultipleCredential(queries []entities.Credentials) ([]entities.Credentials, []error)
	UpdateSingleCredentialByAccountId(query entities.Credentials) (entities.Credentials, error)
	UpdateMultipleCredentialByAccountId(queries []entities.Credentials) ([]entities.Credentials, []error)
	DeleteSingleCredentialByAccountId(query entities.Credentials) (entities.Credentials, error)
	DeleteMultipleCredentialByAccountId(queries []entities.Credentials) ([]entities.Credentials, []error)
}

func (br BromoRepo) SelectAccountById(query entities.Accounts, offset int) ([]entities.Accounts, error) {
	var accounts = make([]entities.Accounts, 0)
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM customers WHERE account_id = ", query.AccountId, " LIMIT 10 OFFSET ", offset)).Scan(
		&accounts)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (br BromoRepo) SelectAccountByFirstName(query entities.Accounts, offset int) ([]entities.Accounts, error) {
	var accounts []entities.Accounts
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM customers WHERE first_name LIKE \"%", query.FirstName, "%\" LIMIT 10 OFFSET ", offset)).Scan(
		&accounts)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (br BromoRepo) SelectAccountByLastName(query entities.Accounts, offset int) ([]entities.Accounts, error) {
	var accounts []entities.Accounts
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM customers WHERE first_name LIKE \"%", query.LastName, "%\" LIMIT 10 OFFSET ", offset)).Scan(
		&accounts)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (br BromoRepo) SelectAccountByEmail(query entities.Accounts, offset int) ([]entities.Accounts, error) {
	var accounts []entities.Accounts
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM customers WHERE first_name LIKE \"%", query.LastName, "%\" LIMIT 10 OFFSET ", offset)).Scan(
		&accounts)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (br BromoRepo) SelectAllAccount(offset int) ([]entities.Accounts, error) {
	var accounts []entities.Accounts
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM customers LIMIT 10 OFFSET ", offset)).Scan(
		&accounts)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

func (br BromoRepo) InsertSingleAccount(query entities.Accounts) (entities.Accounts, error) {
	var account entities.Accounts
	tx := br.db.Model(&account).Clauses(clause.Returning{}).Create(query)
	err := tx.Error
	if err != nil {
		return entities.Accounts{}, err
	}
	return account, nil
}

func (br BromoRepo) InsertMultipleAccount(queries []entities.Accounts) ([]entities.Accounts, []error) {
	var account entities.Accounts
	var accounts []entities.Accounts
	var errors []error
	for _, query := range queries {
		tx := br.db.Model(&account).Clauses(clause.Returning{}).Create(query)
		accounts = append(accounts, account)
		err := tx.Error
		if err != nil {
			errors = append(errors, err)
		} else {
			errors = append(errors, nil)
		}
	}
	return accounts, errors
}

func (br BromoRepo) UpdateSingleAccountById(query entities.Accounts) (entities.Accounts, error) {
	var account entities.Accounts
	tx := br.db.Model(&account).Clauses(clause.Returning{}).Updates(query)
	err := tx.Error
	if err != nil {
		return entities.Accounts{}, err
	}
	return account, nil
}

func (br BromoRepo) UpdateMultipleAccountById(queries []entities.Accounts) ([]entities.Accounts, []error) {
	var account entities.Accounts
	var accounts []entities.Accounts
	var errors []error
	for _, query := range queries {
		tx := br.db.Model(&account).Clauses(clause.Returning{}).Updates(query)
		accounts = append(accounts, account)
		err := tx.Error
		if err != nil {
			errors = append(errors, err)
		} else {
			errors = append(errors, nil)
		}
	}
	return accounts, errors
}

func (br BromoRepo) DeleteSingleAccountById(query entities.Accounts) (entities.Accounts, error) {
	tx := br.db.Clauses(clause.Returning{}).Delete(query)
	account := query
	err := tx.Error
	if err != nil {
		return entities.Accounts{}, err
	}
	return account, nil
}

func (br BromoRepo) DeleteMultipleAccountById(queries []entities.Accounts) ([]entities.Accounts, []error) {
	var errors []error
	var accounts []entities.Accounts
	for _, query := range queries {
		tx := br.db.Clauses(clause.Returning{}).Delete(&query)
		accounts = append(accounts, query)
		err := tx.Error
		if err != nil {
			errors = append(errors, err)
		} else {
			errors = append(errors, nil)
		}
	}
	return accounts, errors
}

func (br BromoRepo) SelectCityById(query entities.Cities, offset int) ([]entities.Cities, error) {
	var cities = make([]entities.Cities, 0)
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM cities WHERE city_id = ", query.CityId, " LIMIT 10 OFFSET ", offset)).Scan(
		&cities)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (br BromoRepo) SelectCityByProvince(query entities.Cities, offset int) ([]entities.Cities, error) {
	var cities = make([]entities.Cities, 0)
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM cities WHERE province = ", query.Province, " LIMIT 10 OFFSET ", offset)).Scan(
		&cities)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (br BromoRepo) SelectAllCity(offset int) ([]entities.Cities, error) {
	var cities = make([]entities.Cities, 0)
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM cities LIMIT 10 OFFSET ", offset)).Scan(
		&cities)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return cities, nil
}

func (br BromoRepo) SelectProvinceById(query entities.Provinces, offset int) ([]entities.Provinces, error) {
	var provinces = make([]entities.Provinces, 0)
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM provinces WHERE province_id = ", query.ProvinceId, " LIMIT 10 OFFSET ", offset)).Scan(
		&provinces)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return provinces, nil
}

func (br BromoRepo) SelectAllProvince(offset int) ([]entities.Provinces, error) {
	var provinces = make([]entities.Provinces, 0)
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM provinces LIMIT 10 OFFSET ", offset)).Scan(
		&provinces)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return provinces, nil
}

func (br BromoRepo) SelectCredentialByAccountId(query entities.Credentials, offset int) ([]entities.Credentials, error) {
	var credentials = make([]entities.Credentials, 0)
	tx := br.db.Raw(
		fmt.Sprint("SELECT * FROM cities WHERE credentials = ", query.AccountId, " LIMIT 10 OFFSET ", offset)).Scan(
		&credentials)
	err := tx.Error
	if err != nil {
		return nil, err
	}
	return credentials, nil
}

func (br BromoRepo) InsertSingleCredential(query entities.Credentials) (entities.Credentials, error) {
	var credentials entities.Credentials
	tx := br.db.Model(&credentials).Clauses(clause.Returning{}).Create(query)
	err := tx.Error
	if err != nil {
		return entities.Credentials{}, err
	}
	return credentials, nil
}

func (br BromoRepo) InsertMultipleCredential(queries []entities.Credentials) ([]entities.Credentials, []error) {
	var credential entities.Credentials
	var credentials []entities.Credentials
	var errors []error
	for _, query := range queries {
		tx := br.db.Model(&credential).Clauses(clause.Returning{}).Create(query)
		credentials = append(credentials, credential)
		err := tx.Error
		if err != nil {
			errors = append(errors, err)
		} else {
			errors = append(errors, nil)
		}
	}
	return credentials, errors
}

func (br BromoRepo) UpdateSingleCredentialByAccountId(query entities.Credentials) (entities.Credentials, error) {
	var credential entities.Credentials
	tx := br.db.Model(&credential).Clauses(clause.Returning{}).Updates(query)
	err := tx.Error
	if err != nil {
		return entities.Credentials{}, err
	}
	return credential, nil
}

func (br BromoRepo) UpdateMultipleCredentialByAccountId(queries []entities.Credentials) ([]entities.Credentials, []error) {
	var credential entities.Credentials
	var credentials []entities.Credentials
	var errors []error
	for _, query := range queries {
		tx := br.db.Model(&credential).Clauses(clause.Returning{}).Updates(query)
		credentials = append(credentials, credential)
		err := tx.Error
		if err != nil {
			errors = append(errors, err)
		} else {
			errors = append(errors, nil)
		}
	}
	return credentials, errors
}

func (br BromoRepo) DeleteSingleCredentialByAccountId(query entities.Credentials) (entities.Credentials, error) {
	tx := br.db.Clauses(clause.Returning{}).Delete(query)
	credential := query
	err := tx.Error
	if err != nil {
		return entities.Credentials{}, err
	}
	return credential, nil
}

func (br BromoRepo) DeleteMultipleCredentialByAccountId(queries []entities.Credentials) ([]entities.Credentials, []error) {
	var errors []error
	var credentials []entities.Credentials
	for _, query := range queries {
		tx := br.db.Clauses(clause.Returning{}).Delete(&query)
		credentials = append(credentials, query)
		err := tx.Error
		if err != nil {
			errors = append(errors, err)
		} else {
			errors = append(errors, nil)
		}
	}
	return credentials, errors
}
