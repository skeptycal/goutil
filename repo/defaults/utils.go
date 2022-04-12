package defaults

import "reflect"

func GetKind(any Any) string             { return reflect.ValueOf(any).Kind().String() }
func CheckType(any Any, typ string) bool { return GetType(any) == typ }

func GetType(any Any) string {
	if any == nil {
		return "nil"
	}
	return reflect.ValueOf(any).Type().String()
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
