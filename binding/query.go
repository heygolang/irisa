package binding

import (
	"github.com/gin-gonic/gin"
)

// Query
func Query(ctx *gin.Context, ptr interface{}) error {
	err := shouldBindQuery(ctx, ptr)
	if err != nil {
		return err
	}
	return Validate(ptr)
}

// Param
func Param(ctx *gin.Context, ptr interface{}) error {
	err := shouldBindUri(ctx, ptr)
	if err != nil {
		return err
	}
	err = shouldBindQuery(ctx, ptr)
	if err != nil {
		return err
	}
	return Validate(ptr)
}

// PathVariable
func PathVariable(ctx *gin.Context, ptr interface{}) error {
	err := shouldBindUri(ctx, ptr)
	if err != nil {
		return err
	}
	return Validate(ptr)
}

// shouldBindUri
func shouldBindUri(ctx *gin.Context, ptr interface{}) error {
	m := make(map[string][]string)
	for _, v := range ctx.Params {
		m[v.Key] = []string{v.Value}
	}
	if err := mappingByPtr(ptr, m, "json"); err != nil {
		return err
	}
	return nil
}

// shouldBindQuery
func shouldBindQuery(ctx *gin.Context, ptr interface{}) error {
	values := ctx.Request.URL.Query()
	if err := mappingByPtr(ptr, values, "json"); err != nil {
		return err
	}
	return nil
}
