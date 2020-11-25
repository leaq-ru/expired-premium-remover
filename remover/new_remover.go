package remover

import "github.com/nnqq/scr-proto/codegen/go/parser"

func NewRemover(companyClient parser.CompanyClient) Remover {
	return Remover{
		companyClient: companyClient,
	}
}
