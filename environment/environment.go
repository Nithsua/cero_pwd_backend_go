package environment

import "github.com/nitsuan/cero_pwd_backend_go/psqldatabase"

//
func LoadEnvironment() {
	psqldatabase.GetDatabaseEnv()
}
