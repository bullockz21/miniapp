package db

import (
	"miniapp/internal/infrastructure/database/storage"
	"miniapp/pkg/logger"
	"miniapp/pkg/postgresql"
	"strings"
)

type repository struct {
	client postgresql.Client
	logger *logger.Logger
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func NewRepository(client postgresql.Client, logger *logger.Logger) storage.Storage {
	return &repository{
		client: client,
		logger: logger,
	}
}
