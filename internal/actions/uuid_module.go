package actions

import (
	"context"

	"github.com/dop251/goja"
	"github.com/gofrs/uuid"
	"github.com/zitadel/logging"
)

type uid struct {
	value string
}

func WithUuid(ctx context.Context) Option {
	logging.Debug("WITH UUID")
	return func(c *runConfig) {
		c.modules["zitadel/uuid"] = func(runtime *goja.Runtime, module *goja.Object) {
			m := module.Get("exports").(*goja.Object)
			m.Set("v4", func(call goja.FunctionCall) goja.Value {
				uid, _ := uuid.NewV4()
				return runtime.ToValue(uid.String())
				// m := module.Get("exports").(*goja.Object)
				// logging.OnError(m.Set("v4", v4(runtime))).Warn("unable to set modules")
			})
		}
	}
}

/*func v4(runtime *goja.Runtime) func(call goja.FunctionCall) goja.Value {
	logging.Info("V4() CALLED")
	return func(call goja.FunctionCall) goja.Value {
		uid, err := uuid.NewV4()
		if err != nil {
			logging.WithError(err).Debug("call failed")
			panic(err)
		}
		return goja.New().ToValue(uid.String())
	}
}
*/
