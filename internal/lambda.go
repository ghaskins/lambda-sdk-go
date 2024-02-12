// Generated by `wit-bindgen` 0.17.0. DO NOT EDIT!
package internal

// #include "lambda.h"
import "C"
import "unsafe"

// Import functions from manetu:internal/sparql@0.0.1
func ManetuLambda0_0_1_SparqlQuery(expr string) string {
	var lower_expr C.lambda_string_t

	// use unsafe.Pointer to avoid copy
	lower_expr.ptr = (*uint8)(unsafe.Pointer(C.CString(expr)))
	lower_expr.len = C.size_t(len(expr))
	var ret C.lambda_string_t
	C.manetu_lambda_sparql_query(&lower_expr, &ret)
	var lift_ret string
	lift_ret = C.GoStringN((*C.char)(unsafe.Pointer(ret.ptr)), C.int(ret.len))
	return lift_ret
}

// Export functions from lambda
var lambda Lambda = nil

// `SetLambda` sets the `Lambda` interface implementation.
// This function will need to be called by the init() function from the guest application.
// It is expected to pass a guest implementation of the `Lambda` interface.
func SetLambda(i Lambda) {
	lambda = i
}

type Lambda interface {
	Handler(request string) string
}

//export lambda_handler
func lambdaHandler(request *C.lambda_string_t, ret *C.lambda_string_t) {
	var lift_request string
	lift_request = C.GoStringN((*C.char)(unsafe.Pointer(request.ptr)), C.int(request.len))
	result := lambda.Handler(lift_request)
	var lower_result C.lambda_string_t

	// use unsafe.Pointer to avoid copy
	lower_result.ptr = (*uint8)(unsafe.Pointer(C.CString(result)))
	lower_result.len = C.size_t(len(result))
	*ret = lower_result

}
