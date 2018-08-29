package db

import "os"

type Database interface {
	Insert(TemplateDB) error
	FindAll() ([]TemplateDB, error)
	Find(ID string) (TemplateDB, error)
	Remove(ID string) bool
}

var TypeDb = GetDriver()

func GetDriver() Database {
	switch os.Getenv("MOCK")  {
	case "MOCK":
		return &Fake{}
	default:
		return &MongoDB{}
	}
}
