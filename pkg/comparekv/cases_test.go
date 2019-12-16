package comparekv

type args struct {
	struct1 map[string]interface{}
	struct2 map[string]interface{}
	fields  []string
}

var map1 = map[string]interface{}{
	"test_field_1": "value_1",
	"test_field_2": "value_2",
	"test_field_3": "value_3",
}

var map2 = map[string]interface{}{
	"test_field_1": "value_1",
	"test_field_2": "value_2",
}

var map3 = map[string]interface{}{
	"test_field_1": "value_1",
	"test_field_2": "value_2",
	"test_field_3": "value_n",
}

var map4 = map[string]interface{}{
	"test_field_1": "value_1",
	"test_field_2": "value_nn",
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "Проверка выборочных полей - все совпадают",
		args: args{map1, map2, []string{"test_field_1", "test_field_2"}},
		want: true,
	},
	{
		name: "Проверка выборочных полей - не совпадают значения",
		args: args{map1, map4, []string{"test_field_1", "test_field_2"}},
		want: false,
	},
	{
		name: "Проверка по всем полям - все совпадают",
		args: args{map1, map1, []string{"test_field_1", "test_field_2", "test_field_3"}},
		want: true,
	},
	{
		name: "Проверка по всем полям - нет нужного поля",
		args: args{map1, map2, []string{"test_field_1", "test_field_2", "test_field_3"}},
		want: false,
	},
	{
		name: "Проверка по всем полям - не совпадают значения",
		args: args{map1, map3, []string{"test_field_1", "test_field_2", "test_field_3"}},
		want: false,
	},
}
