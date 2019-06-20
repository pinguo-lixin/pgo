package pgo

import (
    "reflect"
)

type (
    // HookContainerBind hook attached to container bind function.
    // Return a name to bind to container, then can access the instance by the specify name.
    // Must be configured before controllers and commands bound.
    HookContainerBind func(i interface{}, name string) (bindName string)

    // HookAfterCreateController hook attached to controller & action created
    HookAfterCreateController func(ctx *Context, c, a reflect.Value, route string, params []string) (reflect.Value, reflect.Value)
)

// Hook lifecycle hooks
type Hook struct {
    ContainerBind         HookContainerBind
    AfterCreateController HookAfterCreateController
}
