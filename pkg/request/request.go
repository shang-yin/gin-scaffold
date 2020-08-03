package request

import (
	"encoding/json"
	"fmt"
	"gin-scaffold/pkg/comm"
	"gin-scaffold/pkg/e"
	"gin-scaffold/pkg/translation"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ParseRequest(c *gin.Context, request interface{}) (err error) {
	err = c.ShouldBind(request)
	if err != nil {
		errMap := make(map[string]interface{}, 0)
		errType := e.ErrError
		switch err.(type) {
		case validator.ValidationErrors:
			errMap = translation.Translate(err.(validator.ValidationErrors))
			errType = e.ErrParamIsInvalid
		case *json.UnmarshalTypeError:
			unmarshalTypeError := err.(*json.UnmarshalTypeError)
			errMap[unmarshalTypeError.Field] = fmt.Sprintf("%s类型错误，期望类型%s",
				unmarshalTypeError.Field, unmarshalTypeError.Type.String())
			errType = e.ErrParamTypeBindError
		default:
			errMap["err"] = err
		}
		comm.ReturnJSON(c, errType, errMap)
		return err
	}
	return nil
}
