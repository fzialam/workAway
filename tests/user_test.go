package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/fzialam/workAway/helper"
	"golang.org/x/crypto/bcrypt"
)

func TestHash(t *testing.T) {
	pasw := []string{"ojan", "karim", "panji", "yusuf", "sulis", "marco", "ucup", "munir"}

	// has := []string{
	// 	"$2a$10$TmgbLdG/6qcHgE0vcsbs7ewbm5sOA4cWWN1LpeEByXV5Aex9mQspO",
	// 	"$2a$10$5R3zGo9EDQJgsVSFvXPbhuOMXZlcztlLuUzTHz1F65nHk.p7bxM96",
	// 	"$2a$10$T6iloKq5NszdJHHCBKi5eOnqcuYSIcDTl6XiLVxvXc.y55A1JXYfW",
	// 	"$2a$10$pxPhPRfFYfVq7NWPOFQcBOYZCMeb4Zds3y8KHPU4ZTZuPEiNvAloS",
	// 	"$2a$10$4vhNI3t25xIcpBLrzFEGy.h7IQOs.q12T9M.nnNniq70NQiz2es2a",
	// 	"$2a$10$4wb8WVCgGYFCkw9NTw0E5OUexqRqd6VO2Wd5.kFx1fl1f8dqOhafS",
	// }

	// for i := 0; i < len(pasw); i++ {
	// 	hash, err := bcrypt.GenerateFromPassword([]byte(pasw[i]), bcrypt.DefaultCost)
	// 	if err != nil {
	// 		tln(err)
	// 		panic(exception.NewUnauthorized(err.Error()))
	// 	}
	// 	tln(string(hash))

	// }

	hash := "$2a$10$ShndwhNULzQGf6CxP995fe6CGnt8iwGymXV55XXVytqc7khlH8Dx."

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pasw[0]))
	helper.PanicIfError(err)

}

func TestGeneratePassword(t *testing.T) {
	pasw := []string{"password", "pimpinan", "tuunesa", "pegawai2", "pegawai3", "keuangan"}

	result := []string{}

	for i := 0; i < len(pasw); i++ {
		hash, err := bcrypt.GenerateFromPassword([]byte(pasw[i]), bcrypt.DefaultCost)
		helper.PanicIfError(err)
		result = append(result, string(hash))
		fmt.Println(pasw[i], result[i])
	}

}

func TestImageBase64(t *testing.T) {
	bs := ImageToBase64()

	f, err := os.Create("base64.log")
	helper.PanicIfError(err)

	_, err = f.WriteString(bs)

	helper.PanicIfError(err)
}
