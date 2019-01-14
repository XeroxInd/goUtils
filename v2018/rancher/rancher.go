package rancher

import (
	"io/ioutil"
	"net/http"

	"git.libmed.fr/LibMed/goUtil/v2018/v2018/env"
	"git.libmed.fr/LibMed/goUtil/v2018/v2018/logs"
)

// Stack :
// 1 : QA
// 2 : PREPROD
// 3 : PROD
// 4 : UNKNOWN, which is the case when you test outside of Rancher (I.E. your laptop)
type Stack int

const (
	QA Stack = 1 + iota
	PREPROD
	EXTERNAL
	UNKNOWN
)

// GetStack return the current Rancher Stack :
// 1 : QA
// 2 : PREPROD
// 3 : PROD
// 4 : UNKNOWN, which is the case when you test outside of Rancher (I.E. your laptop)
func GetStack() Stack {
	//Reflect Rancher environment
	resp, err := http.Get("http://rancher-metadata/latest/self/stack/name")
	if err != nil {
		logs.Error(err, "Rancher reflection")
		return UNKNOWN
	} else {
		stackName, _ := ioutil.ReadAll(resp.Body)
		stackNameString := string(stackName)
		var ret Stack
		switch stackNameString {
		case "libmedqa":
			ret = QA
		case "libmedpreprod":
			ret = PREPROD
		case "libmedexternal":
			ret = EXTERNAL
		default:
			ret = UNKNOWN
		}
		return ret
	}
}

// GetStackLogo return the logo URL for the current stack.
// In case of unknown stack, the prod logo is returned.
func GetStackLogo(c Stack) string {
	qaLogo := env.GetEnvOrElse("QA_LOGO_URL", "https://qa.libertymedical.fr/assets/mail_assets/logo_couleur_rvb_720_135_qa.png")
	prodLogo := env.GetEnvOrElse("PROD_LOGO_URL", "https://www.libertymedical.fr/assets/mail_assets/logo_couleur_rvb_720_135.png")
	preprodLogo := env.GetEnvOrElse("PREPROD_LOGO_URL", "https://preprod.libertymedical.fr/assets/mail_assets/logo_couleur_rvb_720_135_pp.png")

	var logoUrl string
	switch c {
	case UNKNOWN:
		logoUrl = prodLogo
	case EXTERNAL:
		logoUrl = prodLogo
	case PREPROD:
		logoUrl = preprodLogo
	case QA:
		logoUrl = qaLogo
	default:
		logoUrl = prodLogo
	}
	return logoUrl
}
