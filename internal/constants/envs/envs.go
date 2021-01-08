package envs

import "errors"

// Env represents the available environments
type Env string

// ErrNotValidEnv is used when env is not any of the available env modes
var ErrNotValidEnv = errors.New("Should be one of Env.(DEBUG|TEST|LOCAL|DEV|STAGE|PROD)")

const (
	// DEBUG env
	DEBUG Env = "debug"
	// TEST env
	TEST Env = "test"
	// LOCAL env
	LOCAL Env = "local"
	// DEV env
	DEV Env = "dev"
	// STAGE env
	STAGE Env = "stage"
	// PROD env
	PROD Env = "prod"
)

func (tm Env) String() string {
	if err := tm.Validate(); err != nil {
		return ""
	}

	return string(tm)
}

// Validate the env mode
func (tm Env) Validate() (err error) {
	envModes := [...]Env{DEBUG, TEST, LOCAL, DEV, STAGE, PROD}

	for _, v := range envModes {
		if v == tm {
			return
		}
	}

	err = ErrNotValidEnv

	return
}
