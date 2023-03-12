import (
	"context"
	"fmt"
	{{if .time}}"time"{{end}}

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)
