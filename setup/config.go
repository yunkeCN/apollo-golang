package setup

import (
	"fmt"
	"github.com/goinggo/mapstructure"
	"github.com/philchia/agollo/v4"
	"reflect"
	"strconv"
)

var WatchConfigFields = make(map[string]*Field)

// Field ...
type Field struct {
	Name          string
	ApolloKeyName string
	Value         reflect.Value
	Type          reflect.Type
}

// SetNewValue ...
func (f *Field) SetNewValue(newValue string) error {
	switch f.Type.Kind() {
	case reflect.Int:
		againValue, err := strconv.Atoi(newValue)
		if err != nil {
			return err
		}
		f.Value.Set(reflect.ValueOf(againValue))
	case reflect.Int8:
		againValue, err := strconv.Atoi(newValue)
		if err != nil {
			return err
		}
		f.Value.Set(reflect.ValueOf(int8(againValue)))
	case reflect.Int16:
		againValue, err := strconv.Atoi(newValue)
		if err != nil {
			return err
		}
		f.Value.Set(reflect.ValueOf(int16(againValue)))
	case reflect.Int32:
		againValue, err := strconv.Atoi(newValue)
		if err != nil {
			return err
		}
		f.Value.Set(reflect.ValueOf(int32(againValue)))
	case reflect.Int64:
		againValue, err := strconv.Atoi(newValue)
		if err != nil {
			return err
		}
		f.Value.Set(reflect.ValueOf(int64(againValue)))
	case reflect.Uint:
		againValue, err := strconv.Atoi(newValue)
		if err != nil {
			return err
		}
		f.Value.Set(reflect.ValueOf(uint(againValue)))
	case reflect.Uint8:
		againValue, err := strconv.Atoi(newValue)
		if err != nil {
			return err
		}
		f.Value.Set(reflect.ValueOf(uint8(againValue)))
	case reflect.Uint16:
		againValue, err := strconv.Atoi(newValue)
		if err != nil {
			return err
		}
		f.Value.Set(reflect.ValueOf(uint16(againValue)))
	case reflect.Uint32:
		againValue, err := strconv.Atoi(newValue)
		if err != nil {
			return err
		}
		f.Value.Set(reflect.ValueOf(uint32(againValue)))
	case reflect.Uint64:
		againValue, err := strconv.Atoi(newValue)
		if err != nil {
			return err
		}
		f.Value.Set(reflect.ValueOf(uint64(againValue)))
	case reflect.Float32:
		againValue, err := strconv.ParseFloat(newValue, 32)
		if err != nil {
			return err
		}
		f.Value.Set(reflect.ValueOf(againValue))
	case reflect.Float64:
		againValue, err := strconv.ParseFloat(newValue, 64)
		if err != nil {
			return err
		}
		f.Value.Set(reflect.ValueOf(againValue))
	case reflect.String:
		f.Value.Set(reflect.ValueOf(newValue))
	case reflect.Bool:
		againValue, err := strconv.ParseBool(newValue)
		if err != nil {
			return err
		}
		f.Value.Set(reflect.ValueOf(againValue))
	default:
		return fmt.Errorf("Unkonwn field type")
	}

	return nil
}

// GetReflectFields gets reflect fields from v.
func GetReflectFields(section string, v interface{}) (map[string]*Field, error) {
	typeOf := reflect.TypeOf(v)
	valueOf := reflect.ValueOf(v)
	if typeOf.Kind() == reflect.Ptr {
		typeOf = typeOf.Elem()
	}
	if valueOf.Kind() == reflect.Ptr {
		valueOf = valueOf.Elem()
	}
	if typeOf.Kind() != reflect.Struct {
		return nil, fmt.Errorf("Type must be a struct")
	}

	result := make(map[string]*Field)
	fieldCnt := typeOf.NumField()
	for i := 0; i < fieldCnt; i++ {
		name := typeOf.Field(i).Name
		apolloKeyName := fmt.Sprintf("%s.%s", section, name)
		result[apolloKeyName] = &Field{
			Name:          name,
			ApolloKeyName: apolloKeyName,
			Type:          typeOf.Field(i).Type,
			Value:         valueOf.Field(i),
		}
	}

	return result, nil
}

// SaveWatchConfigField save uses reflection fields to get value from apollo, and convert
// it into the given native structure v.
func SaveWatchConfigField(v interface{}, fields map[string]*Field, client agollo.Client, namespace string) error {
	c := &apolloClient{
		client:    client,
		namespace: namespace,
	}

	configValues := make(map[string]interface{})
	for apolloKeyName, field := range fields {
		var value interface{}
		switch field.Type.Kind() {
		case reflect.Int,
			reflect.Int8,
			reflect.Int16,
			reflect.Int32,
			reflect.Int64,
			reflect.Uint,
			reflect.Uint8,
			reflect.Uint16,
			reflect.Uint32,
			reflect.Uint64:
			value = c.GetIntValue(apolloKeyName, 0)
		case reflect.String:
			value = c.GetStringValue(apolloKeyName, "")
		case reflect.Bool:
			value = c.GetBoolValue(apolloKeyName, false)
		case reflect.Float32,
			reflect.Float64:
			value = c.GetFloatValue(apolloKeyName, 0)
		default:
			return fmt.Errorf("current field type is not be supported")
		}
		configValues[field.Name] = value
	}
	if len(configValues) == 0 {
		return nil
	}

	if err := mapstructure.Decode(configValues, v); err != nil {
		return err
	}

	return nil
}

type apolloClient struct {
	client    agollo.Client
	namespace string
}

func (c *apolloClient) GetIntValue(key string, defaultValue int) int {
	v := c.client.GetString(key, agollo.WithNamespace(c.namespace))
	if v == "" {
		return defaultValue
	}

	vv, _ := strconv.Atoi(v)
	return vv
}

func (c *apolloClient) GetStringValue(key string, defaultValue string) string {
	v := c.client.GetString(key, agollo.WithNamespace(c.namespace))
	if v == "" {
		return defaultValue
	}

	return v
}

func (c *apolloClient) GetBoolValue(key string, defaultValue bool) bool {
	v := c.client.GetString(key, agollo.WithNamespace(c.namespace))
	if v == "" {
		return defaultValue
	}

	vv, _ := strconv.ParseBool(v)
	return vv
}

func (c *apolloClient) GetFloatValue(key string, defaultValue float64) float64 {
	v := c.client.GetString(key, agollo.WithNamespace(c.namespace))
	if v == "" {
		return defaultValue
	}

	vv, _ := strconv.ParseFloat(v, 64)
	return vv
}
