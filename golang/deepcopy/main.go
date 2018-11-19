package main

import (
	"fmt"
	"reflect"
	"time"
)

type (
	Player struct {
		Id     int
		Level  int
		Heroes map[int]*Hero
		Equips []*Equip
	}

	Hero struct {
		Id     int
		Level  int
		Skills []*Skill
	}

	Equip struct {
		Id    int
		Level int
	}

	Skill struct {
		Id    int
		Level int
	}
)

func NewHero() *Hero {
	return &Hero{
		Id:     1,
		Level:  1,
		Skills: append([]*Skill{NewSkill()}, NewSkill(), NewSkill()),
	}
}

func NewSkill() *Skill {
	return &Skill{1, 1}
}

func NewEquip() *Equip {
	return &Equip{1, 1}
}

func NewPlayer() *Player {
	return &Player{
		Id:     1,
		Level:  1,
		Heroes: map[int]*Hero{1: NewHero(), 2: NewHero(), 3: NewHero()},
		Equips: append([]*Equip{NewEquip()}, NewEquip(), NewEquip()),
	}
}

func (self *Hero) Print() {
	fmt.Printf("Id=%d, Level=%d\n", self.Id, self.Level)
	for _, v := range self.Skills {
		fmt.Printf("%v\n", *v)
	}
}

func (self *Player) Print() {
	fmt.Printf("Id=%d, Level=%d\n", self.Id, self.Level)
	for _, v := range self.Heroes {
		v.Print()
	}

	for _, v := range self.Equips {
		fmt.Printf("%+v\n", *v)
	}
}

type Interface interface {
	DeepCopy() interface{}
}

func Copy(src interface{}) interface{} {
	if src == nil {
		return nil
	}
	original := reflect.ValueOf(src)
	cpy := reflect.New(original.Type()).Elem()
	copyRecursive(original, cpy)

	return cpy.Interface()
}

func copyRecursive(src, dst reflect.Value) {
	if src.CanInterface() {
		if copier, ok := src.Interface().(Interface); ok {
			dst.Set(reflect.ValueOf(copier.DeepCopy()))
			return
		}
	}

	switch src.Kind() {
	case reflect.Ptr:
		originalValue := src.Elem()

		if !originalValue.IsValid() {
			return
		}
		dst.Set(reflect.New(originalValue.Type()))
		copyRecursive(originalValue, dst.Elem())

	case reflect.Interface:
		if src.IsNil() {
			return
		}
		originalValue := src.Elem()
		copyValue := reflect.New(originalValue.Type()).Elem()
		copyRecursive(originalValue, copyValue)
		dst.Set(copyValue)

	case reflect.Struct:
		t, ok := src.Interface().(time.Time)
		if ok {
			dst.Set(reflect.ValueOf(t))
			return
		}
		for i := 0; i < src.NumField(); i++ {
			if src.Type().Field(i).PkgPath != "" {
				continue
			}
			copyRecursive(src.Field(i), dst.Field(i))
		}

	case reflect.Slice:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeSlice(src.Type(), src.Len(), src.Cap()))
		for i := 0; i < src.Len(); i++ {
			copyRecursive(src.Index(i), dst.Index(i))
		}

	case reflect.Map:
		if src.IsNil() {
			return
		}
		dst.Set(reflect.MakeMap(src.Type()))
		for _, key := range src.MapKeys() {
			originalValue := src.MapIndex(key)
			copyValue := reflect.New(originalValue.Type()).Elem()
			copyRecursive(originalValue, copyValue)
			copyKey := Copy(key.Interface())
			dst.SetMapIndex(reflect.ValueOf(copyKey), copyValue)
		}

	default:
		dst.Set(src)
	}
}

func main() {
	p1 := NewPlayer()
	p2 := Copy(p1).(*Player)
	fmt.Println(reflect.DeepEqual(p1, p2))
}
