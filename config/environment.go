package config

import (
	"os"
	"github.com/alknopfler/notificationapp/logger"
	"strconv"
)

var(
	LogLevel 			=   logger.GetLevel(os.Getenv("SERVICE_LOG_LEVEL"))

	EmailSMTPHost		= 	os.Getenv("SESHOST")
	EmailSMTPPort,_		= 	strconv.Atoi(os.Getenv("SESPORT"))
	EmailSMTPUser		= 	os.Getenv("SESUSER")
	EmailSMTPPass		= 	os.Getenv("SESPASSWORD")
	EmailSMTPFrom		=   os.Getenv("SESFROM")
	EmailSMTPSender		= 	os.Getenv("SESSENDER")
	EmailSMTPTLS		= 	os.Getenv("SESTLS")

	SMSAccessKey		=   os.Getenv("SNSACCESSKEY")
	SMSSecretKey		=   os.Getenv("SNSSECRETKEY")
	SMSRegion			=   os.Getenv("SNSREGION")

	MongoUsername		= 	os.Getenv("MONGODBUSER")
	MongoPassword		= 	os.Getenv("MONGODBPASSWORD")
	MongoDB				= 	os.Getenv("MONGODB")
	MongoHost			=   os.Getenv("MONGODBHOST")
	MongoCollection		=   os.Getenv("MONGODBCOLLECTION")

)

