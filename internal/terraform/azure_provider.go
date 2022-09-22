package terraform

import (
	"github.com/hashicorp/hcl/v2/hclwrite"
	"strings"
	"text/template"
)

var (
	AzureTerraformProviderTemplateString string = `
 "terraform" {
	required_providers {
		azurerm = {
			source = "hashicorp/azurerm"
			version = ">={{.Version}}"
		}
	}
}
`
	AzureTerraformProviderTemplate = template.Must(template.New("azure-provider").Funcs(templateFunctionMap).Parse(AzureTerraformProviderTemplateString))
)

type AzureTerraformProvider struct {
	Version string
}

func (x AzureTerraformProvider) TerraformConfig() string {
	var buf strings.Builder
	err := AzureTerraformProviderTemplate.Execute(&buf, x)
	if err != nil {
		panic(err)
	}
	return string(hclwrite.Format([]byte(buf.String())))
}
