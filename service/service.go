package service

import "github.com/alinz/goissue"

type Access struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	Message      int
}

func Service() {
	conf := issue.Conf.OAuth2Google

	access := &Access{
		ClientID: conf.ClientID,
	}

	if access == nil {
		//do something
	}

	//uncomment the following code to compile the code in go1.6
	//fmt.Println("")
}
