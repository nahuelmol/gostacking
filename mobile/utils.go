package mobile

import (
    "reflect"
    "errors"
)

func CleanXY(body map[string]interface{}) (int32, int32, error){

    xlocation, isx := body["xlocation"]
    ylocation, isy := body["ylocation"]
    
    var ix int32
    var iy int32
    if isx && isy {
        xtype := reflect.TypeOf(xlocation)
        ytype := reflect.TypeOf(ylocation)
        if xtype.Kind() == reflect.Float64 {
            f64x := xlocation.(float64)
            ix = int32(f64x)
        }
        if ytype.Kind() == reflect.Float64 {
            f64y := ylocation.(float64)
            iy = int32(f64y)
        }
    } else {
        return 0,0, errors.New("error extracting data")
    }

    return ix, iy, nil
}
