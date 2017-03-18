package main

var Config = make(map[string]string)

func initializeConfig() error {
	Config["mongoDatasource"] = "mongodb://data:27017/betterRecipe"

	return nil
}
func getMongoDatasource() string {
  return Config["mongoDatasource"]
}
