package models

import (
	"encoding/json"
	"os"
)

//Configuration settings
type Configuration struct {
	DatabaseConnectionString    string                      `json:"connectionString"`
	IdentityServerConfiguration IdentityServerConfiguration `json:"identity_server_config"`
	ServiceUrl                  string                      `json:"serviceUrl"`

}

//IdentityServerConfiguration configuration
type IdentityServerConfiguration struct {
	URL         string `json:"url"`
	Certificate string `json:"certificate"`
}

var CurrentConfiguration Configuration

func init()  {
	configFile, err := os.Open("appsettings.json")
	defer configFile.Close()
	if err != nil {
		panic(err)
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&CurrentConfiguration)
	
}


