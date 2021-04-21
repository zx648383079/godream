package batch

type BatchHandleFunc = func(data interface{}) interface{}

type BatchMap = map[string]BatchHandleFunc

func InvokeBatch(params map[string]interface{}, items BatchMap) (map[string]interface{}, error) {
	data := make(map[string]interface{})
	for path, v := range items {
		if val, ok := params[path]; ok {
			data[path] = v(val)
		}
	}
	return data, nil
}
