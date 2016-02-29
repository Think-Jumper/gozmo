package gozmo

type EngineSingleton struct {
    registeredComponents map[string]*RegisteredComponent
}

var Engine EngineSingleton

func RegisterComponent(name string, generator func([]interface{}) Component) {
    if Engine.registeredComponents == nil {
        Engine.registeredComponents = make(map[string]*RegisteredComponent)
    }

    rc := RegisteredComponent{Name: name, Init: generator}

    Engine.registeredComponents[name] = &rc
}