package envs

import (
	"os"

	"github.com/joho/godotenv"
)

func initializeEnvs() {
	godotenv.Load()

	envs = &env{
		PORT: getEnv("PORT", "3030"),

		LOCAL: getEnvBool("LOCAL", false),

		PEPPER: getEnv("PEPPER", "secret"),
		ENV:    getEnv("ENV", "development"),

		MONGO_HOST:     getEnv("MONGO_HOST", "mongodb://localhost:27017"),
		MONGO_DATABASE: getEnv("MONGO_DATABASE", "example"),
	}

}

// * getEnv - Get env value (string)
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// * getEnvBool - Get env value (bool)
func getEnvBool(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		return parseBool(value)
	}

	return fallback
}

func parseBool(value string) bool {
	if value == "true" || value == "TRUE" || value == "1" {
		return true
	}

	return false
}
