package utils

type Config struct {
	Title    string
	Gameplay Gameplay
	Debug    Debug
}

type Gameplay struct {
	SizeField     [2]int
	CountResorces int
}

type Debug struct {
	LogLevel string
}
