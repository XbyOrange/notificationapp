package config

const(
SUCCESS = "SUCCESS"
FAILED = "FAILED"
)
var(
	SrvPort = ":8080"
	SrvName = "NotificationApp"
	ComponentName="NotificationApp"
	Api = ComponentName+"-api"
	Database = ComponentName+"-db"
	Channel = ComponentName + "-channel"
	Template = ComponentName + "-template"
)

