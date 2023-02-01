`gorm:"column:{{.field}}{{if eq .field "delete_time"}};not null{{else if eq .field "create_time"}};autoCreateTime{{else if eq .field "update_time"}};autoUpdateTime{{end}}"`
