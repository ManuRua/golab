module github.com/ManuRua/golab

go 1.14

// +heroku goVersion go1.14
// +heroku install ./telegram-bot/...

require (
	github.com/joho/godotenv v1.3.0
	github.com/json-iterator/go v1.1.9
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.3.0
	github.com/yanzay/tbot/v2 v2.1.0
)
