package tests

import (
	"logger"
	"dao"
	"model"
	"config"
)

func setUpTest() {
	env := "test"
	logger.Init(env)
	config.Init(env)
	dao.Init(env)
	model.Migrate()
}

func destroy() {
	// TODO drop all tables
	// delete files
}
