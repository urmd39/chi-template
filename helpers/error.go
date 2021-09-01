package helpers

import "go.mongodb.org/mongo-driver/mongo"

type ErrResponse struct {
	Success bool        `json:"success"`
	Message interface{} `json:"data"`
}

func CatchErrResponse(err error) (code int, errResponse ErrResponse) {
	switch err {
	case mongo.ErrNoDocuments:
		code = 404
		errResponse.Message = "Not Found"
	default:
		code = 406
		errResponse.Message = "Not Acceptable"
	}
	return
}
