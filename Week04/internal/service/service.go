package service

import (
	pb "geekGo/Week04/api/blog/v1"
	"geekGo/Week04/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewBlogService)

type BlogService struct {
	pb.UnimplementedBlogServiceServer

	log *log.Helper

	article *biz.ArticleUsecase
}
