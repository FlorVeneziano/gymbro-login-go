package envs

func GetEnvs() *env {
	once.Do(initializeEnvs)
	return envs
}
