package response

import (
	"net/http"
)

func Pagination(data, limit, offset, total interface{}) (int, interface{}) {
	return http.StatusOK, map[string]interface{}{
		"data":   data,
		"limit":  limit,
		"offset": offset,
		"total":  total,
	}
}
func Data(code int, data interface{}) (int, interface{}) {
	return code, map[string]interface{}{
		"data": data,
	}
}

func NewResponse(code int, data interface{}) (int, interface{}) {
	return code, data
}

func NewOKResponse(data interface{}) (int, interface{}) {
	return http.StatusOK, map[string]interface{}{
		"data":    data,
		"code":    http.StatusOK,
		"content": "successfully",
	}
}

func OK(data interface{}) (int, interface{}) {
	return http.StatusOK, data
}

func NewCreatedResponse(data map[string]interface{}) (int, interface{}) {
	result := map[string]interface{}{
		"code":    http.StatusCreated,
		"content": "successfully",
	}
	for key, value := range data {
		result[key] = value
	}
	return http.StatusCreated, result
}

func NewErrorResponse(code int, msg interface{}) (int, interface{}) {
	return code, map[string]interface{}{
		"error":   http.StatusText(code),
		"code":    code,
		"content": msg,
	}
}

func ServiceUnavailable() (int, interface{}) {
	return http.StatusServiceUnavailable, map[string]interface{}{
		"error":   http.StatusText(http.StatusServiceUnavailable),
		"code":    http.StatusBadRequest,
		"content": http.StatusText(http.StatusServiceUnavailable),
	}
}

func ServiceUnavailableMsg(msg interface{}) (int, interface{}) {
	return http.StatusServiceUnavailable, map[string]interface{}{
		"error":   http.StatusText(http.StatusServiceUnavailable),
		"code":    http.StatusBadRequest,
		"content": msg,
	}
}

func BadRequest() (int, interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"error":   http.StatusText(http.StatusBadRequest),
		"code":    http.StatusBadRequest,
		"content": http.StatusText(http.StatusBadRequest),
	}
}

func BadRequestMsg(msg interface{}) (int, interface{}) {
	return http.StatusBadRequest, map[string]interface{}{
		"error":   http.StatusText(http.StatusBadRequest),
		"code":    http.StatusBadRequest,
		"content": msg,
	}
}

func NotFound() (int, interface{}) {
	return http.StatusNotFound, map[string]interface{}{
		"error":   http.StatusText(http.StatusNotFound),
		"code":    http.StatusNotFound,
		"content": http.StatusText(http.StatusNotFound),
	}
}

func NotFoundMsg(msg interface{}) (int, interface{}) {
	return http.StatusNotFound, map[string]interface{}{
		"error":   http.StatusText(http.StatusNotFound),
		"code":    http.StatusNotFound,
		"content": msg,
	}
}

func Forbidden() (int, interface{}) {
	return http.StatusForbidden, map[string]interface{}{
		"error":   "Do not have permission for the request.",
		"code":    http.StatusForbidden,
		"content": http.StatusText(http.StatusForbidden),
	}
}

func Unauthorized() (int, interface{}) {
	return http.StatusUnauthorized, map[string]interface{}{
		"error":   http.StatusText(http.StatusUnauthorized),
		"code":    http.StatusUnauthorized,
		"content": http.StatusText(http.StatusUnauthorized),
	}
}
