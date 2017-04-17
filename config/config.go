package config

var SmartPOS = &configStruct{}

type configStruct struct {
	App struct {
		    HTTPAddr string
	    }
}
