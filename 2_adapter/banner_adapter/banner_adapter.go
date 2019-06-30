package banner_adapter

import "github.com/saekis/go-design-pattern/2_adapter/banner"

type Adapter interface {
	GetWeak() string
	GetStrong() string
}

type bannerAdapter struct {
	banner banner.Banner
}

func NewBannerAdapter(str string) Adapter {
	return &bannerAdapter{banner: banner.Banner{Str: str}}
}

func (b *bannerAdapter) GetWeak() string {
	return b.banner.WithParen()
}

func (b *bannerAdapter) GetStrong() string {
	return b.banner.WithAster()
}
