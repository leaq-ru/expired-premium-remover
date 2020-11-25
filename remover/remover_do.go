package remover

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

func (r Remover) Do(ctx context.Context) (err error) {
	_, err = r.companyClient.UnsetExpiredPremium(ctx, &empty.Empty{})
	return
}
