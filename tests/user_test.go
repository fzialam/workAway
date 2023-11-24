package tests

import (
	"log"
	"testing"

	"github.com/fzialam/workAway/exception"
	"golang.org/x/crypto/bcrypt"
)

func TestHash(t *testing.T) {
	pasw := []string{"password", "pimpinan", "tuunesa", "pegawai2", "pegawai3", "pegawai4"}

	has := []string{
		"$2a$10$TmgbLdG/6qcHgE0vcsbs7ewbm5sOA4cWWN1LpeEByXV5Aex9mQspO",
		"$2a$10$5R3zGo9EDQJgsVSFvXPbhuOMXZlcztlLuUzTHz1F65nHk.p7bxM96",
		"$2a$10$T6iloKq5NszdJHHCBKi5eOnqcuYSIcDTl6XiLVxvXc.y55A1JXYfW",
		"$2a$10$pxPhPRfFYfVq7NWPOFQcBOYZCMeb4Zds3y8KHPU4ZTZuPEiNvAloS",
		"$2a$10$4vhNI3t25xIcpBLrzFEGy.h7IQOs.q12T9M.nnNniq70NQiz2es2a",
		"$2a$10$4wb8WVCgGYFCkw9NTw0E5OUexqRqd6VO2Wd5.kFx1fl1f8dqOhafS",
	}

	for i := 0; i < len(pasw); i++ {
		err := bcrypt.CompareHashAndPassword([]byte(has[i]), []byte(pasw[i]))
		if err != nil {
			log.Println(err)
			panic(exception.NewUnauthorized(err.Error()))
		}
		log.Println(err)

	}

}
