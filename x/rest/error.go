package rest

import (
	"net/http"
)

var (
	ErrUserExisted         = BadRequestError("Tài khoản đã tồn tại. Vui lòng kiểm tra lại thông tin tài khoản")
	ErrInvalidCredentials  = BadRequestError("Thông tin nhập không chính xác")
	ErrInvalidUser         = BadRequestError("Người dùng không hợp lệ")
	ErrInternalServerError = InternalServerError("Internal Server Error")
	ErrConflict            = BadRequestError("Your Item already exist")
	ErrBadParamInput       = BadRequestError("Given Param is not valid")
	ErrValidation          = BadRequestError("Validation Error")
	ErrDuplicatedKey       = BadRequestError("Unique Violation")
	ErrInvalidImage        = BadRequestError("Invalid Image")
	// ErrDataNotFound        = BadRequestError("Data Not Found")
)

type IHttpError interface {
	StatusCode() int
}

type BadRequestError string

func (e BadRequestError) Error() string {
	return string(e)
}

func (e BadRequestError) StatusCode() int {
	return http.StatusBadRequest
}

func WrapBadRequest(err error, message string) error {
	if err != nil {
		return BadRequestError(message + ": " + err.Error())
	}
	return nil
}

type UnauthorizedError string

func (e UnauthorizedError) Error() string {
	return string(e)
}

func (e UnauthorizedError) StatusCode() int {
	return http.StatusUnauthorized
}

type InternalServerError string

func (e InternalServerError) Error() string {
	return string(e)
}

func (e InternalServerError) StatusCode() int {
	return http.StatusInternalServerError
}

func AssertNil(errs ...error) {
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
}
