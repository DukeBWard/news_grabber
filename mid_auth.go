package main

import (
	"net/http"

	"github.com/dukebward/news_grabber/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg)