package {{.NameImport}}

import (
	"github.com/Kong/go-pdk"
)

type Config struct {
	Param string `json:"param"`
}

func New() interface{} {
	return &Config{}
}

var pluginName = "{{.Name}}"

func (conf Config) Access(kong *pdk.PDK) {
	log("Executing {{.Name}} plugin..")

	// Example code below
	//
	// Get Header values
	// header, err := kong.Request.GetHeader("Authorization")
	//
	// Exit on error
	// kong.Response.Exit(401, "", x)
	//
	// Access outgoing service request:
	// kong.ServiceRequest.
	//
	// Override Response
	// king.Response.
}

/*
	Logger function to format output for clarification on which plugin produced the log.
*/
func log(kong *pdk.PDK, message string) {
	kong.Log.Info(fmt.Sprintf("[%s] %s", PluginName, message))
}
