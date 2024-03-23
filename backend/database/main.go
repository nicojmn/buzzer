package database 

import (
	"log"
	"errors"
	"buzzer/config"
	"buzzer/hashing"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

type Admin struct {
	ID       int    `gorm:"primaryKey;type:int;autoIncrement:true;not null;unique"`
	Username string `gorm:"unique;type:string;not null;"`
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

func Clear(db *gorm.DB) {
	db.Migrator().DropTable(&Admin{}, &Team{})
	db.AutoMigrate(&Admin{}, &Team{})
}


func AddAdmin(db *gorm.DB, username string, password string) (error) {
	// check if user exists with gorm
	query := db.Model(&Admin{}).Where("username = ?", username)
	if query.RowsAffected > 0 {
		log.Printf("User %s already exists", username)
		return errors.New("user already exists")
	}

	// hash password
	password, err := hashing.HashPassword(password)
	if err != nil {
		log.Println(err)
		return err
	}

	// add user to db
	result := db.Create(&Admin{Username: username, Password: password})
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	log.Printf("User %s added", username)
	return nil
}

func AddTeam(db *gorm.DB, name string) (error) {
	// check if there are not more then 13 teams
	var teams []Team
	config_data, err := config.LoadConfig("config.yaml")

	if err != nil {
		log.Printf("Database : error loading config file (%s)", err);
		return err;
	}

	db.Find(&teams)
	if len(teams) >= config_data.MaxTeams {
		log.Printf("DB :too many teams, maximum number allowed %d", config_data.MaxTeams)
		return errors.New("maxTeamsReached")
	}

	// check if team exists
	query := db.Model(&Team{}).Where("name = ?", name)
	if query.RowsAffected > 0 {
		log.Printf("Team %s already exists", name)
		return errors.New("Team already exists")
	}

	// add team to db
	result := db.Create(&Team{Name: name, Score: 0})
	if result.Error != nil {
		log.Printf("Database : error in creation of user, %s", result.Error)
		return result.Error
	}

	log.Printf("Team %s added", name)
	return nil
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
	return score, nil
}