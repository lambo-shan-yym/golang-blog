package tools

import "golang-blog/src/common/error"

func CheckParamNotBlank(str string, key string) {
	if str == "" {
		panic(code.FormatException(code.RequiredParam, key))
	}
}

func CheckParamGreaterThanZero(value int64, key string) {
	if value <= 0 {
		panic(code.FormatException(code.ParamInvalid, key))
	}
}

func CheckParams4Paginator(limit *int, offset int, limitKey string, offsetKey string) {
	if offset < 0 {
		panic(code.FormatException(code.ParamInvalid, limitKey))
	}
	if *limit < 0 {
		panic(code.FormatException(code.ParamInvalid, limitKey))
	}

	if *limit > 200 {
		*limit = 200
	}
}
