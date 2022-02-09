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

func (conf Config) Access(kong *pdk.PDK) {
	kong.Log.Info("Executing {{.Name}} plugin..")

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
