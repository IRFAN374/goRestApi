package service

import service "github.com/IRFAN374/goRestApi/grocery"

type Middleware func(service.Service) service.Service
