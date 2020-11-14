package model

import (
	"crypto/rand"
	"fmt"
	"github.com/rs/zerolog/log"
)

func TokenGenerator() string{
	b := make([]byte,8)
	_, err := rand.Read(b)
	if err != nil{
		log.Info().Msg("Generate Token Error")
	}
	return fmt.Sprintf("%x", b)
}