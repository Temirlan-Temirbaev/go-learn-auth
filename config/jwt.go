package config

import "time"

var JWTSecret = []byte("foofie213")
var TokenExpiration = time.Hour * 24
