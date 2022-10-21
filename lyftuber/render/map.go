package render

import (
	_ "embed"
)

//go:embed tmpl/map.html
var MapTmpl []byte

//go:generate genopts --params --function Map --required "lyftToken string, uberCSID string, uberSID string" latitude:float64:40.7701286 longitude:float64:-73.9829762 zoom:int:14 sleepSecs:int:0
