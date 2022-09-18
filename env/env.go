package env

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

var (
	LogLevel string

	SrvEnv              string
	SrvPort             string
	SrvShutdownDeadline time.Duration

	MongoDBUrl  string
	MongoDBname string
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Print("no .env file")
	}

	LogLevel = getEnv("LOG_LEVEL", "debug")

	SrvEnv = getEnv("SERVER_ENV", "dev")
	SrvPort = getEnv("SERVER_PORT", "3000")
	SrvShutdownDeadline = mustStrToDuration(getEnv("SERVER_SHUTDOWN_DEADLINE_SEC", "30"), time.Second)

	MongoDBUrl = getEnv("MONGO_DB_URL", "mongodb://localhost:27017")
	MongoDBname = getEnv("MONGO_DB_NAME", "dev")
}

func getEnv(envName, defaultVal string) string {
	envVal := os.Getenv(envName)
	if envVal == "" {
		envVal = defaultVal
	}
	return envVal
}

func mustAtoi(val string) int {
	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(err)
	}
	return intVal
}

func mustStrToDuration(val string, unit time.Duration /*time.Nanosecond,...*/) time.Duration {
	intVal := mustAtoi(val)
	return time.Duration(intVal) * unit
}

// Do not use any whitespace as a delimiter
func mustSplitStr(val, delimiter string) []string {
	noSpaceStr := strings.ReplaceAll(val, " ", "")
	return strings.Split(noSpaceStr, delimiter)
}
