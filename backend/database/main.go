package database

import (
	"buzzer/config"
	"buzzer/hashing"
	"buzzer/observer"
	"errors"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Admin struct {
	ID       int    `gorm:"primaryKey;type:int;autoIncrement:true;not null;unique"`
	Username string `gorm:"unique;type:string;not null;"`
	Password string `gorm:"type:string;not null"`
}

type Team struct {
	ID    int    `gorm:"primaryKey;type:int;autoIncrement:true;not null;unique"`
	Name  string `gorm:"type:string;not null;unique"`
	Score int	 `gorm:"type:int;not null;default:0"`
	PressedAt int `gorm:"type:int;not null; default:0"`
	Locked bool `gorm:"type:bool;not null;default:false"`
}



type DBconn struct {
	DB *gorm.DB
}

var dbInstance *DBconn



func (db *DBconn) InitDB(path string) error {
	var err error
	db.DB, err = gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Println(err);
		return err;
	}
	
	// create tables
	db.DB.AutoMigrate(&Admin{}, &Team{})

	log.Println("Database initialized")
	
	return nil
}

func GetInstance(path string) *DBconn {
	if dbInstance == nil {
		dbInstance = &DBconn{}
		err := dbInstance.InitDB(path)
		if err != nil {
			log.Fatalf("Error initializing database: %s", err)
		}
	}
	return dbInstance
}
func Clear() {
	DB := GetInstance("db.sqlite").DB
	DB.Migrator().DropTable(&Admin{}, &Team{})
	DB.AutoMigrate(&Admin{}, &Team{})
}


func AddAdmin(username string, password string) (error) {
	DB := GetInstance("db.sqlite").DB
	// check if user exists with gorm
	query := DB.Model(&Admin{}).Where("username = ?", username)
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
	result := DB.Create(&Admin{Username: username, Password: password})
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	log.Printf("User %s added", username)
	return nil
}

func AddTeam(name string) (error) {
	DB := GetInstance("db.sqlite").DB
	// check if there are not more then 13 teams
	var teams []Team
	config_data, err := config.LoadConfig("config.yaml")

	if err != nil {
		log.Printf("Database : error loading config file (%s)", err);
		return err;
	}

	DB.Find(&teams)
	if len(teams) >= config_data.MaxTeams {
		log.Printf("DB :too many teams, maximum number allowed %d", config_data.MaxTeams)
		return errors.New("maxTeamsReached")
	}

	// check if team exists
	query := DB.Model(&Team{}).Where("name = ?", name)
	if query.RowsAffected > 0 {
		log.Printf("Team %s already exists", name)
		return errors.New("Team already exists")
	}

	// add team to db
	result := DB.Create(&Team{Name: name, Score: 0, PressedAt: 0, Locked: false})
	if result.Error != nil {
		log.Printf("Database : error in creation of user, %s", result.Error)
		return result.Error
	}

	log.Printf("Team %s added", name)
	return nil
}

func GetAdmin(username string) (Admin, error) {
	DB := GetInstance("db.sqlite").DB
	var admin Admin
	result := DB.Where("username = ?", username).First(&admin)
	if result.Error != nil {
		log.Println(result.Error)
		return admin, result.Error
	}
	return admin, nil
}

func GetTeam(name string) (Team, error) {
	DB := GetInstance("db.sqlite").DB
	var team Team
	result := DB.Where("name = ?", name).First(&team)
	if result.Error != nil {
		log.Println(result.Error)
		return team, result.Error
	}
	return team, nil
}

func GetTeamID(id int) (Team, error) {
	DB := GetInstance("db.sqlite").DB
	var team Team
	result := DB.Where("id = ?", id).First(&team)
	if result.Error != nil {
		log.Println(result.Error)
		return team, result.Error
	}
	return team, nil
}

func GetTeams() ([]Team, error) {
	DB := GetInstance("db.sqlite").DB
	var teams []Team
	result := DB.Find(&teams)
	if result.Error != nil {
		log.Println(result.Error)
		return teams, result.Error
	}
	return teams, nil
}

func UpdateScore(team Team, score int) error {
	DB := GetInstance("db.sqlite").DB
	result := DB.Model(&team).Where("name = ?", team.Name).Update("score", team.Score + score)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	// notify observers
	observer.SubjectInstance.Notify(map[string]interface{}{
		"type": "scoreUpdate",
		"data": map[string]interface{}{
			"team": team.Name,
			"score": team.Score + score,
			"team_id": team.ID,
		},
	})
	return nil
}

func UpdatePressedAt(team Team) error {
	DB := GetInstance("db.sqlite").DB
	result := DB.Model(&team).Where("name = ?", team.Name).Update("pressed_at", time.Now().UnixMilli())
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	
	log.Printf("Team %s pressed at %d", team.Name, team.PressedAt)
	return nil
}

func LockTeam(team Team) error {
	DB := GetInstance("db.sqlite").DB
	result := DB.Model(&team).Where("name = ?", team.Name).Update("locked", true)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	
	log.Printf("Team %s locked", team.Name)
	return nil
}

func LockAllTeams() error {
	DB := GetInstance("db.sqlite").DB
	result := DB.Model(&Team{}).Where("locked = ?", false).Update("locked", true)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	//notify observers
	observer.SubjectInstance.Notify(map[string]interface{}{
		"type": "lockAll",
	})
	
	log.Printf("All teams locked")
	return nil
}

func UnlockAllTeams() error {
	DB := GetInstance("db.sqlite").DB
	result := DB.Model(&Team{}).Where("locked = ?", true).Update("locked", false)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	//notify observers
	observer.SubjectInstance.Notify(map[string]interface{}{
		"type": "unlockAll",
	})
	
	log.Printf("All teams unlocked")
	return nil
}

func LockState() (bool, error) {
	// return true if all teams are locked, false otherwise
	DB := GetInstance("db.sqlite").DB
	var teams []Team

	result := DB.Find(&teams)
	if result.Error != nil {
		log.Println(result.Error)
		return false, result.Error
	}

	for _, team := range teams {
		if !team.Locked {
			return false, nil
		}
	}
	return true, nil
}

func GetLockedTeams() ([]Team, error) {
	DB := GetInstance("db.sqlite").DB
	var teams []Team
	result := DB.Where("locked = ?", true).Find(&teams)
	if result.Error != nil {
		log.Println(result.Error)
		return teams, result.Error
	}
	return teams, nil
}
