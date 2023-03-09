package gh

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type ErrHandler func(c *gin.Context, code int, err error)

type argParser func(c *gin.Context) (reflect.Value, error)

// MustBind bind given function into a gin.HandlerFunc,
// which can decode arguments from request and compose the invocation
// of the given function.
func MustBind(fn interface{}) *Binder {
	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		panic("the given function must be a function object")
	}

	argParsers := make([]argParser, 0, fnType.NumIn())
	for i := 0; i < fnType.NumIn(); i++ {
		in := fnType.In(i).Elem()
		switch in.String() {
		case "gin.Context":
			argParsers = append(argParsers, func(c *gin.Context) (reflect.Value, error) {
				return reflect.ValueOf(c), nil
			})
		default:
			uriBuilding := containsURIBinding(fnType.In(i))
			argParsers = append(argParsers, func(c *gin.Context) (reflect.Value, error) {
				inObj := reflect.New(in)
				if uriBuilding {
					if err := c.BindUri(inObj.Interface()); err != nil {
						return reflect.Value{}, err
					}
				}
				if err := c.Bind(inObj.Interface()); err != nil {
					return reflect.Value{}, err
				}
				return inObj, nil
			})
		}
	}
	return &Binder{
		fn:         fn,
		argParsers: argParsers,
		fnType:     fnType,
	}
}

func containsURIBinding(t reflect.Type) bool {
	tEl := t.Elem()
	for i := 0; i < tEl.NumField(); i++ {
		f := tEl.Field(i)
		if _, ok := f.Tag.Lookup("uri"); ok {
			return true
		}
	}
	return false
}

// Binder to customize.
type Binder struct {
	fn         interface{}
	argParsers []argParser
	fnType     reflect.Type
	respondErr ErrHandler
	respondOk  func(c *gin.Context, r interface{})
}

// SetRespondErr define how to respond error
func (b *Binder) SetRespondErr(f ErrHandler) *Binder {
	b.respondErr = f
	return b
}

// SetRespondOk define how to respond ok.
func (b *Binder) SetRespondOk(f func(c *gin.Context, r interface{})) *Binder {
	b.respondOk = f
	return b
}

// HandlerFunc returns composed handler function of gin.
func (b *Binder) HandlerFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		args := make([]reflect.Value, 0, b.fnType.NumIn())
		for i := 0; i < b.fnType.NumIn(); i++ {
			arg, err := b.argParsers[i](c)
			if err != nil {
				b.handleErr(c, http.StatusBadRequest, err)
				return
			}
			args = append(args, arg)
		}

		outs := reflect.ValueOf(b.fn).Call(args)
		for _, out := range outs {
			if out.IsNil() {
				continue
			}

			v := out.Interface()

			switch v := v.(type) {
			case error:
				b.handleErr(c, http.StatusInternalServerError, v)
				return
			default:
				if b.respondOk != nil {
					b.respondOk(c, v)
				} else {
					c.JSON(http.StatusOK, v)
				}
			}
		}
	}
}

func (b *Binder) handleErr(c *gin.Context, status int, err error) {
	if b.respondErr != nil {
		b.respondErr(c, status, err)
		return
	}
	c.Status(status)
	c.Writer.Write([]byte(err.Error()))
}
