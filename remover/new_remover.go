package remover

import "github.com/leaq-ru/proto/codegen/go/parser"

func NewRemover(companyClient parser.CompanyClient) Remover {
	return Remover{
		companyClient: companyClient,
	}
}
