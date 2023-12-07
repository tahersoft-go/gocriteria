package gocriteria

import (
	"github.com/tahersoft-go/gocriteria/validation"
	"time"
)

type FuncSchema struct {
	Fn func(string, ...interface{}) (bool, error)
}

type Validator struct {
	Field        string
	Value        interface{}
	IsOptional   bool
	DefaultValue string
	Errors       []error
	Validations  []FuncSchema
}

func New(label string) *Validator {
	return &Validator{
		Field:       label,
		Validations: []FuncSchema{},
	}
}

func (v *Validator) Email() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsEmail,
	})
	return v
}

func (v *Validator) Required() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsRequired,
	})
	return v
}

func (v *Validator) Number() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsNumber,
	})
	return v
}

func (v *Validator) Url() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsUrl,
	})
	return v
}

func (v *Validator) Alpha(locale string) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsAlpha(f, v.Value, locale)
		},
	})
	return v
}

func (v *Validator) Filepath() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsFilepath,
	})
	return v
}

func (v *Validator) LowerCase() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsLowerCase,
	})
	return v
}

func (v *Validator) UpperCase() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsUpperCase,
	})
	return v
}

func (v *Validator) Int() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsInt,
	})
	return v
}

func (v *Validator) IntSlice() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsIntSlice,
	})
	return v
}

func (v *Validator) Float() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsFloat,
	})
	return v
}

func (v *Validator) Json() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsJson,
	})
	return v
}

func (v *Validator) Ip() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsIp,
	})
	return v
}

func (v *Validator) IpV4() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsIpV4,
	})
	return v
}

func (v *Validator) IpV6() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsIpV6,
	})
	return v
}

func (v *Validator) Port() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsPort,
	})
	return v
}

func (v *Validator) IsDNSName() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsDNSName,
	})
	return v
}

func (v *Validator) Host() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsHost,
	})
	return v
}

func (v *Validator) Latitude() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsLatitude,
	})
	return v
}

func (v *Validator) Logitude() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsLogitude,
	})
	return v
}

func (v *Validator) AlphaNum(locale string) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsAlphaNum(f, v.Value, locale)
		},
	})
	return v
}

func (v *Validator) InRange(from, to int) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsInRange(f, v.Value, from, to)
		},
	})
	return v
}

func (v *Validator) MinMaxLength(min, max int) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsMinMaxLength(f, v.Value, min, max)
		},
	})
	return v
}

func (v *Validator) MinLength(min int) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsMinLength(f, v.Value, min)
		},
	})
	return v
}

func (v *Validator) MaxLength(max int) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsMaxLength(f, v.Value, max)
		},
	})
	return v
}

func (v *Validator) In(in []string) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsIn(f, v.Value, in)
		},
	})
	return v
}

func (v *Validator) FilterOperators(operators ...string) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.FilterOperators(f, v.Value, operators)
		},
	})
	return v
}

func (v *Validator) CustomValidator(fn func(string, ...interface{}) (bool, error)) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			isValid, err := fn(v.Field, v.Value)
			label, value := validation.GetFieldLabelAndValue(v.Field, []interface{}{v.Value})
			if err != nil {
				err = validation.GetErrorMessageByFieldValue(err.Error(), label, value)
			}
			return isValid, err
		},
	})
	return v
}

func (v *Validator) Default(value string) *Validator {
	v.DefaultValue = value
	return v
}

func (v *Validator) Optional() *Validator {
	v.IsOptional = true
	return v
}

func (v *Validator) Min(min int) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsMin(f, v.Value, min)
		},
	})
	return v
}

func (v *Validator) Max(max int) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsMax(f, v.Value, max)
		},
	})
	return v
}

func (v *Validator) MinDate(min time.Time) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsMinDate(f, v.Value, min)
		},
	})
	return v
}

func (v *Validator) MaxDate(max time.Time) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsMaxDate(f, v.Value, max)
		},
	})
	return v
}

func (v *Validator) MinTime(min time.Time) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsMinTime(f, v.Value, min)
		},
	})
	return v
}

func (v *Validator) MaxTime(max time.Time) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsMaxTime(f, v.Value, max)
		},
	})
	return v
}

func (v *Validator) MinDateTime(min time.Time) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsMinDateTime(f, v.Value, min)
		},
	})
	return v
}

func (v *Validator) MaxDateTime(max time.Time) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsMaxDateTime(f, v.Value, max)
		},
	})
	return v
}

func (v *Validator) MinLengthIfPresent(min int) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsMinLengthIfPresent(f, v.Value, min)
		},
	})
	return v
}

func (v *Validator) MaxLengthIfPresent(max int) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsMaxLengthIfPresent(f, v.Value, max)
		},
	})
	return v
}

func (v *Validator) MinIfPresent(min int) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsMinIfPresent(f, v.Value, min)
		},
	})
	return v
}

func (v *Validator) StartWith(str string) *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: func(f string, i ...interface{}) (bool, error) {
			return validation.IsStartWith(f, v.Value, str)
		},
	})
	return v
}

func (v *Validator) HexColor() *Validator {
	v.Validations = append(v.Validations, FuncSchema{
		Fn: validation.IsHexColor,
	})
	return v
}
