package interfaces

import (
	"fmt"
	"time"
	"untitled/internal/domain"
	"untitled/internal/infrastructure"
)

type AccountSQLRepository struct {
	DB infrastructure.DBConnector
}

func (repo *AccountSQLRepository) SelectAccountByID(id int) (domain.AccountAggregate, error) {
	var err error
	sqlStatement := `
		SELECT id, email, email_activation_token, email_verified_at, password, role,
				first_name, last_name, phone, company, about_me, avatar,
				city, country, post_index
		  FROM users as u
		 WHERE u.id = $1
	;`
	account := domain.AccountAggregate{
		Identity: domain.IdentityEntity{
			ID:        0,
			Password:  "",
			Email:     domain.EmailValueObject{
				Value:           "",
				VerifiedAt:      time.Time{},
				ActivationToken: "",
			},
			Role:      "",
			Status:    "",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: time.Time{},
		},
		Profile: domain.ProfileValueObject{
			FullName: domain.FullNameValueObject{
				FirstName: "",
				LastName:  "",
			},
			Phone: domain.PhoneValueObject{
				FullPhone: "",
				Country: "",
				Code:    "",
				Number:  "",
			},
			AboutMe:  "",
			Avatar:   "",
			Company:  "",
		},
		Address:  domain.AddressValueObject{
			Country: "",
			City:    "",
			Post:    "",
		},
	}
	err = repo.DB.QueryRow(sqlStatement, id).Scan(
		&account.Identity.ID,
		&account.Identity.Email.Value,
		&account.Identity.Email.ActivationToken,
		&account.Identity.Email.VerifiedAt,
		&account.Identity.Password,
		&account.Identity.Role,

		&account.Profile.FullName.FirstName,
		&account.Profile.FullName.LastName,
		&account.Profile.Phone.FullPhone,
		&account.Profile.Company,
		&account.Profile.AboutMe,
		&account.Profile.Avatar,

		&account.Address.City,
		&account.Address.Country,
		&account.Address.Post)
	if err != nil {
		return account, err
	}

	return account, nil
}

// GetAccountByEmail used for authenticate purposes, will load only required fields
func (repo *AccountSQLRepository) SelectAccountByCredentials(aggregate domain.AccountAggregate) (domain.AccountAggregate, error) {
	var err error
	sqlStatement := `
		SELECT id, first_name, last_name, email, password, role
		  FROM users AS u
		 WHERE u.email = $1 AND u.password = $2
	;`
	a := domain.AccountAggregate{}
	err = repo.DB.QueryRow(sqlStatement,
		aggregate.Identity.Email.Value,
		aggregate.Identity.Password).
		Scan(
		&a.Identity.ID,
		&a.Profile.FullName.FirstName,
		&a.Profile.FullName.LastName,
		&a.Identity.Email.Value,
		&a.Identity.Password,
		&a.Identity.Role)
	if err != nil {
		return a, err
	}

	return a, nil
}

// InsertAccount used once for User registration action, so can omny many fields
func (repo *AccountSQLRepository) InsertAccount(a domain.AccountAggregate) error  {
	// transaction
	sqlStatement := `
		INSERT INTO users (first_name, last_name, phone, email, password, role)
 			 VALUES ($1, $2, $3, $4, $5, $6)
		  RETURNING id
	;`
	err := repo.DB.QueryRow(sqlStatement,
		a.Profile.FullName.FirstName,
		a.Profile.FullName.LastName,
		a.Profile.Phone.FullPhone,
		a.Identity.Email.Value,
		a.Identity.Password,
		a.Identity.Role,
	).Scan(&a.Identity.ID)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

/*
func (repo *AccountSQLRepository) UpdateAccount(user domain.AccountAggregate) error  {
	fmt.Print(user)
	sqlStatement := `
   UPDATE users
      SET first_name=$2, last_name=$3, email=$4, phone=$5, role=$6, company=$7, address=$8, city=$9, 
			country=$10, post_index=$11, about_me=$12, updated_at=$13
    WHERE id = $1;`
	err := repo.DB. QueryRow(sqlStatement,
		u.ID,
		u.FirstName,
		u.LastName,
		u.Value,
		u.Phone,
		u.Role,
		u.Company,
		u.Address,
		u.City,
		u.Country,
		u.PostIndex,
		u.AboutMe,
		u.UpdatedAt,
	).Scan(&u.ID)
	if err != nil {
		return err
	}
	return nil
}

*/