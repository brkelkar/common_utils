package utils

import "reflect"

func GetBatchSize(i interface{}) int {

	v := reflect.ValueOf(i)
	var batchCount int
	batchCount = (2100 / v.NumField()) - 1

	return batchCount
}
