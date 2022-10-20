package render

import (
	_ "embed"
)

//go:embed tmpl/map.html
var MapTmpl []byte

//go:generate genopts --params --function Map latitude:float64:40.7701286 longitude:float64:-73.9829762 zoom:int:14 sleep:int:5000
