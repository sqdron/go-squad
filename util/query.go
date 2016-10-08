package util

type Grouped map[interface{}][]interface{}

func GroupBy(query []interface{}, expression func(v interface{}) interface{}) Grouped {
	result := make(Grouped)
	for _, v := range query {
		eval := expression(v)
		if result[eval] == nil {
			result[eval] = []interface{}{}
		}
		result[eval] = append(result[eval], v)
	}
	return result
}
