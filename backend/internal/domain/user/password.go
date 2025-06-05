package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	OK int = iota
	SHORTPWD
	LONGPWD
	FORBIDDENCHARS
)

func CreatePasswordHash(pwd string) (pwdHash string, err error) {
	err = pwdValidation(pwd)
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	PasswordHash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return "", fmt.Errorf("failed to create user hash to error: %v", err)
	}
	return string(PasswordHash), nil
}

func isValid(pwd string) int {
	len := len(pwd)
	code := OK
	if len < 5 {
		code = SHORTPWD
	}
	if len > 20 {
		code = LONGPWD
	} else {
		runesArr := []rune(pwd)
		for i := 0; i < len; i++ {
			if runesArr[i] < '!' || runesArr[i] > '~' {
				code = FORBIDDENCHARS
			}
		}
	}
	return code
}

func pwdValidation(pwd string) error {
	val := isValid(pwd)
	var err error
	switch val {
	case SHORTPWD:
		// h.logger.Infoln("pwd is short")
		// h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Password is too short")
		err = fmt.Errorf("pwd is short")
	case LONGPWD:
		// h.logger.Infoln("pwd is long")
		// h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Password is too long")
		err = fmt.Errorf("pwd is long")
	case FORBIDDENCHARS:
		// h.logger.Infoln("pwd contains forbidden chars")
		// h.respMsg.SendMsgJson(w, http.StatusBadRequest, "Bad Request", "Password contains forbidden chars")
		err = fmt.Errorf("pwd contains forbiddent chars")
	default:
		err = nil
	}
	return err
}
