package embed

import _ "embed"

//go:embed dev/global.yml
var ConfigDev string

//go:embed prod/global.yml
var ConfigProd string

//go:embed dev/auth.yml
var AuthDev string

//go:embed prod/auth.yml
var AuthProd string

//go:embed dev/proxy.yml
var ProxyDev string

//go:embed prod/proxy.yml
var ProxyProd string
