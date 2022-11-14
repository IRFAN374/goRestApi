package service

import  service "github.com/IRFAN374/goRestApi/repository/mygrocery"

type Middleware func(service.Repository) service.Repository