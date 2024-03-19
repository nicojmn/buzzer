package database 

import (
	"log"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Admin struct {
	ID       int    `gorm:"primaryKey;type:int;autoIncrement:true;not null;unique"`
	Username string `gorm:"unique;type:string;not null;unique"`
	Password string `gorm:"type:string;not null"`
}

type Team struct {
	ID    int    `gorm:"primaryKey;type:int;autoIncrement:true;not null;unique"`
	Name  string `gorm:"type:string;not null;unique"`
	Score int	 `gorm:"type:int;not null"`
}


func InitDB(path string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatal(err);
		return nil, err;
	}

	// create tables
	db.AutoMigrate(&Admin{}, &Team{})

	return db, nil
}

func ClearDB(db *gorm.DB) {
	db.Migrator().DropTable(&Admin{}, &Team{})
	db.AutoMigrate(&Admin{}, &Team{})
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(hash), nil
}

func AddAdmin(db *gorm.DB, username string, password string) (int, error) {
	// check if user exists with gorm
	query := db.Model(&Admin{}).Where("username = ?", username)
	if query.RowsAffected > 0 {
		log.Printf("User %s already exists", username)
		return -1, nil
	}

	// hash password
	password, err := hashPassword(password)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}

	// add user to db
	result := db.Create(&Admin{Username: username, Password: password})
	if result.Error != nil {
		log.Fatal(result.Error)
		return -1, result.Error
	}
	log.Printf("User %s added", username)
	return 0, nil
}

func AddTeam(db *gorm.DB, name string) (int, error) {
	// check if there are not more then 13 teams
	var teams []Team
	db.Find(&teams)
	// TODO : fix number of teams bug
	if len(teams) > 13 {
		log.Printf("DB :there are already 13 teams")
		return -1, errors.New("maxTeamsReached")
	}

	// check if team exists
	query := db.Model(&Team{}).Where("name = ?", name)
	if query.RowsAffected > 0 {
		log.Printf("Team %s already exists", name)
		return -1, errors.New("Team already exists")
	}

	// add team to db
	result := db.Create(&Team{Name: name, Score: 0})
	if result.Error != nil {
		log.Fatal(result.Error)
		return -1, result.Error
	}

	log.Printf("Team %s added", name)
	return 0, nil
}

func GetAdmin(db *gorm.DB, username string) (Admin, error) {
	var admin Admin
	result := db.Where("username = ?", username).First(&admin)
	if result.Error != nil {
		log.Fatal(result.Error)
		return admin, result.Error
	}
	return admin, nil
}

func GetTeam(db *gorm.DB, name string) (Team, error) {
	var team Team
	result := db.Where("name = ?", name).First(&team)
	if result.Error != nil {
		log.Println(result.Error)
		return team, result.Error
	}
	return team, nil
}

func GetTeamID(db *gorm.DB, id int ) (Team, error) {
	var team Team
	result := db.Where("id = ?", id).First(&team)
	if result.Error != nil {
		log.Fatal(result.Error)
		return team, result.Error
	}
	return team, nil
}

func UpdateScore(db *gorm.DB, team Team, score int) (int, error) {
	team.Score = score
	result := db.Model(&team).Where("name = ?", team.Name).Update("score", score)
	if result.Error != nil {
		log.Fatal(result.Error)
		return -1, result.Error
	}
	return 0, nil
}