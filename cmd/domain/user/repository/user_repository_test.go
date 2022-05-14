package repository

import (
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"regexp"
	"testing"
	"time"
	"toko/cmd/domain/user/entity"
)

//var (
//	sqlMockDB, sqlMock, _ = sqlmock.New()
//	gormDB, _             = gorm.Open(mysql.New(mysql.Config{
//		Conn:                      sqlMockDB,
//		SkipInitializeWithVersion: true,
//	}))
//
//	userRepo = UserRepositoryImpl{
//		Db: gormDB,
//	}
//)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func InitConnection() (sqlmock.Sqlmock, UserRepositoryImpl) {
	sqlMockDB, sqlMock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlMockDB,
		SkipInitializeWithVersion: true,
	}))

	userRepo := UserRepositoryImpl{
		Db: gormDB,
	}
	return sqlMock, userRepo
}

func TestUserRepository_FindAllFail(t *testing.T) {
	sqlMock, userRepo := InitConnection()
	sqlMock.ExpectQuery("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL").
		WillReturnError(errors.New("can't fetch to mock db"))

	users, err := userRepo.FindAll()
	assert.NotNil(t, err)
	assert.Nil(t, users)
}

func TestUserRepository_FindAllSuccess(t *testing.T) {
	sqlMock, userRepo := InitConnection()
	sqlMock.ExpectQuery("SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}))
	users, err := userRepo.FindAll()

	assert.Nil(t, err)
	assert.NotNil(t, users)
}

func TestUserRepository_FindFail(t *testing.T) {
	sqlMock, userRepo := InitConnection()
	sqlMock.ExpectQuery(regexp.QuoteMeta(
		"SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1")).
		WillReturnError(errors.New("can't fetch to mock db"))
	users, err := userRepo.FindAll()
	assert.NotNil(t, err)
	assert.Nil(t, users)
}

func TestUserRepository_FindSuccess(t *testing.T) {
	sqlMock, userRepo := InitConnection()
	var (
		id        uint      = 1
		name      string    = "Abdul Kholiq"
		email     string    = "kholiqdev@icloud.com"
		password  string    = "bismillah"
		createdAt time.Time = time.Now()
		updatedAt time.Time = time.Now()
	)
	sqlMock.ExpectQuery("SELECT * FROM `users` WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1").WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).AddRow(id, name, email, password, createdAt, updatedAt))
	users, err := userRepo.Find(id)

	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.IsType(t, &entity.User{}, users)
}

func TestUserRepository_FindByEmailFail(t *testing.T) {
	sqlMock, userRepo := InitConnection()
	sqlMock.ExpectQuery("SELECT * FROM `users` WHERE email = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1").
		WillReturnError(errors.New("can't fetch to mock db"))

	users, err := userRepo.FindAll()
	assert.NotNil(t, err)
	assert.Nil(t, users)
}

func TestUserRepository_FindByEmailSuccess(t *testing.T) {
	sqlMock, userRepo := InitConnection()
	var (
		id        uint      = 1
		name      string    = "Abdul Kholiq"
		email     string    = "kholiqdev@icloud.com"
		password  string    = "bismillah"
		createdAt time.Time = time.Now()
		updatedAt time.Time = time.Now()
	)

	sqlMock.ExpectQuery("SELECT * FROM `users` WHERE email = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1").WithArgs(email).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).AddRow(id, name, email, password, createdAt, updatedAt))
	users, err := userRepo.FindByEmail(email)

	assert.Nil(t, err)
	assert.NotNil(t, users)
	assert.IsType(t, &entity.User{}, users)
}

func TestUserRepository_InsertFail(t *testing.T) {
	sqlMock, userRepo := InitConnection()
	var (
		//id        uint      = 1
		name      string  = "Abdul Kholiq"
		email     string  = "kholiqdev@icloud.com"
		password  string  = "bismillah"
		createdAt AnyTime = AnyTime{}
		updatedAt AnyTime = AnyTime{}
	)

	sqlMock.ExpectExec("INSERT INTO `users` (`name`,`email`,`password`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?)").WithArgs(name, email, password, createdAt, updatedAt, nil).
		WillReturnError(errors.New("can't fetch to mock db"))

	userEntity := entity.User{
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "bismillah",
	}

	users, err := userRepo.Insert(&userEntity)
	assert.NotNil(t, err)
	assert.Nil(t, users)
}

func TestUserRepository_InsertSuccess(t *testing.T) {
	sqlMock, userRepo := InitConnection()
	var (
		//id        uint      = 1
		name      string  = "Abdul Kholiq"
		email     string  = "kholiqdev@icloud.com"
		password  string  = "bismillah"
		createdAt AnyTime = AnyTime{}
		updatedAt AnyTime = AnyTime{}
	)

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec("INSERT INTO `users` (`name`,`email`,`password`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?)").
		WithArgs(name, email, password, createdAt, updatedAt, nil).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	userEntity := entity.User{
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "bismillah",
	}
	user, err := userRepo.Insert(&userEntity)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.IsType(t, &entity.User{}, user)
	assert.Equal(t, email, user.Email)
}
