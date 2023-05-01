package db

import (
	"database/sql"
	"fmt"
	"github.com/DamirLuketic/vehicle/config"
	"github.com/DamirLuketic/vehicle/db/data"
	mysqld "github.com/go-sql-driver/mysql"
	"gopkg.in/retry.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type MariaDBDataStore struct {
	db *gorm.DB
}

func NewMariaDBDataStore(c *config.ServerConfig) *MariaDBDataStore {
	ds := MariaDBDataStore{}
	var err error
	dsn := getDSN(c)
	createDB()
	for a := retry.Start(connectRetryStrategy(), nil); a.Next(); {
		ds.db, err = gorm.Open(mysql.Open(dsn))
		if err == nil {
			break
		} else {
			log.Println("DB connect fail")
		}
	}
	if err != nil {
		log.Fatalf("DB connect fail. Error: %s", err.Error())
	}
	ds.db.Migrator().DropTable(&Vehicle{})
	ds.db.AutoMigrate(&Vehicle{})
	ds.setData()
	return &ds
}

func (ds *MariaDBDataStore) CreateVehicle(vehicle Vehicle) (Vehicle, error) {
	re := ds.db.Create(&vehicle)
	if re.Error != nil {
		return vehicle, re.Error
	}
	return vehicle, nil
}

func (ds *MariaDBDataStore) DeleteVehicle(vehicleOID string) error {
	re := ds.db.Where("oid = ?", vehicleOID).Delete(&Vehicle{})
	if re.Error != nil {
		return re.Error
	}
	return nil
}

func (ds *MariaDBDataStore) GetVehicles() (v []Vehicle, err error) {
	re := ds.db.Find(&v)
	if re.Error != nil {
		return nil, re.Error
	}
	return v, nil
}

func (ds *MariaDBDataStore) setData() {
	data := getData()
	log.Println("TEST2 DATA", data)
	for _, d := range data {
		_, err := ds.CreateVehicle(d)
		if err != nil {
			log.Fatalf("Error on inserting data. Error: %s", err.Error())
		}
		log.Printf("Data for vehicle with oid: %s created in DB", d.OID)
	}
}

func getData() []Vehicle {
	data := data.GetDBData()
	var paredData []Vehicle
	for _, d := range data {
		paredData = append(paredData, Vehicle{
			Make:  d.Make,
			Model: d.Model,
			Year:  d.Year,
			OID:   d.Id.Oid,
		})
	}
	return paredData
}

func getDSN(c *config.ServerConfig) string {
	cfg := mysqld.NewConfig()
	cfg.DBName = c.MySQLDatabase
	cfg.ParseTime = true
	cfg.User = c.MySQLUser
	cfg.Passwd = c.MySQLPassword
	cfg.Net = "tcp"
	cfg.Params = map[string]string{
		"charset": "utf8mb4",
		"loc":     "Local",
	}
	cfg.Addr = fmt.Sprintf("%v:%v", c.MySQLHost, c.MySQLPort)
	dsn := cfg.FormatDSN()
	return dsn
}

func connectRetryStrategy() retry.Strategy {
	return retry.LimitTime(30*time.Second,
		retry.Exponential{
			Initial: 1 * time.Second,
			Factor:  1.5,
		},
	)
}

func createDB() {
	db, err := sql.Open("mysql", "vehicle:vehicle@tcp(db:3306)/")
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Printf("MySQL connected successfully")
	}
	for a := retry.Start(connectRetryStrategy(), nil); a.Next(); {
		_, err = db.Exec("DROP DATABASE IF EXISTS vehicle")
		if err == nil {
			break
		} else {
			log.Println("Unsuccessful database drop, will retry...")
		}
	}
	if err != nil {
		log.Fatalf("DB connect fail. Error: %s", err.Error())
	}
	_, err = db.Exec("CREATE DATABASE vehicle")
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		log.Println("Successfully created database..")
	}
}
