package db

type Vehicle struct {
	Make  string
	Model string
	Year  int64
	OID   string `gorm:"unique;column:oid;"`
}
