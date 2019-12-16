package comparekv

// CompareFields возвращает признак сравнения двух map[string]interface{}
// fields - массив строк
func CompareFields(struct1, struct2 map[string]interface{}, fields ...string) bool {
	for _, fieldname := range fields {
		// намерено нет проверки на существование поля в struct1
		// так как подразумевается что fields гарантировано содержит поля из struct1
		if w, ok := struct2[fieldname]; !ok || struct1[fieldname] != w {
			return false
		}
	}
	return true
}
