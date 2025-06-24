package init

import (
	"github.com/matheushermes/IAResumeScanner/configs"
	unipdfLicense "github.com/unidoc/unipdf/v4/common/license"
	uniofficeLicense "github.com/unidoc/unioffice/v2/common/license"
)

func init() {
	configs.LoadingEnvironmentVariables()

	//Ativar licença do UniPDF
	err := unipdfLicense.SetMeteredKey(configs.UNIDOC_LICENSE_API_KEY)
	if err != nil {
		panic("Falha ao ativar licença UniPDF: " + err.Error())
	}

	//Ativar licença do Unioffice
	err = uniofficeLicense.SetMeteredKey(configs.UNIDOC_LICENSE_API_KEY)
	if err != nil {
		panic("Falha ao ativar licença Unioffice: " + err.Error())
	}
}
