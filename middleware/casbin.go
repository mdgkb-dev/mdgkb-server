package middleware

import (
	"github.com/gin-gonic/gin"
)

//
//// Middleware returns a CasbinAuth middleware.
////
//// For valid credentials it calls the next handler.
//// For missing or invalid credentials, it sends "401 - Unauthorized" response.
//func Middleware(ce *casbin.Enforcer) echo.MiddlewareFunc {
//	ce.AddFunction("mask", maskMatchFunc)
//	c := DefaultConfig
//	c.Enforcer = ce
//	return MiddlewareWithConfig(c)
//}
//
//// MiddlewareWithConfig returns a CasbinAuth middleware with config.
//// See `Middleware()`.
//func MiddlewareWithConfig(config Config) echo.MiddlewareFunc {
//	// Defaults
//	if config.Skipper == nil {
//		config.Skipper = DefaultConfig.Skipper
//	}
//
//	return func(next echo.HandlerFunc) echo.HandlerFunc {
//		return func(c echo.Context) error {
//			c.Set("ce", config.Enforcer)
//			if config.Skipper(c) || config.CheckPermission(c) {
//				return next(c)
//			}
//			return echo.ErrForbidden
//		}
//	}
//}
//
//func maskMatch(mask1 uint, mask2 uint) bool {
//	return mask1&mask2 != 0
//}
//
//func maskMatchFunc(args ...interface{}) (interface{}, error) {
//	mask1 := uint(args[0].(float64))
//	name2 := args[1].(string)
//	mask2, err := strconv.ParseUint(name2, 10, 32)
//	if err != nil {
//		return false, err
//	}
//	return maskMatch(mask1, uint(mask2)), nil
//}
//

func (m *Middleware) checkPermission(c *gin.Context) bool {
	role := "ADMIN"
	method := "GET"
	path := c.FullPath()
	///api/v1/doctors
	re, err := m.enforcer.Enforce(role, path, method)
	if err != nil || !re {
		return false
	}
	return true
}
