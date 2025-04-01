package uploadprovider

import (
	"context"
	"rest/common"
)

type UploadProvide interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*common.Image, error)
}