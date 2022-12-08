package where

import "strconv"

type ConditionInterface interface {
	Build(n int) (string, []interface{})
}

// ////////////////////////////////////////////////////////
// top level builder struct
type ConditionBuilder struct {
	SubCondition ConditionInterface
}

func (cb *ConditionBuilder) Where(arg ConditionInterface) *ConditionBuilder {
	cb.SubCondition = arg
	return cb
}

func (cb *ConditionBuilder) And(args ...ConditionInterface) *ConditionBuilder {
	and_condition := AndCondition{}
	and_condition.Add(cb.SubCondition)
	for _, a := range args {
		and_condition.Add(a)
	}
	cb.SubCondition = &and_condition
	return cb
}

func (condition *ConditionBuilder) Build(n int) (
	str string,
	val []interface{},
) {

	if condition.SubCondition == nil {
		return "", nil
	} else {
		str, val = condition.SubCondition.Build(n)
		return " WHERE " + str, val
	}
}

// ////////////////////////////////////////////////////////
// simpla condition
type Condition struct {
	Target    string
	Condition string
	Value     interface{}
}

func (condition *Condition) Build(n int) (
	str string,
	val []interface{},
) {
	str = condition.Target + " " + condition.Condition + " $" + strconv.Itoa(n)
	val = append(val, condition.Value)
	return str, val
}

// ////////////////////////////////////////////////////////
// and condition

func And(args ...ConditionInterface) ConditionInterface {
	and_condition := AndCondition{}
	for _, a := range args {
		and_condition.Add(a)
	}
	return &and_condition
}

type AndCondition struct {
	SubConditions []ConditionInterface
}

func (condition *AndCondition) Add(arg ConditionInterface) {
	condition.SubConditions = append(condition.SubConditions, arg)
}

func (condition *AndCondition) Build(n int) (
	str string,
	val []interface{},
) {

	if len(condition.SubConditions) == 0 {
		return "", val
	} else if len(condition.SubConditions) == 1 {
		return condition.SubConditions[0].Build(n)

	} else {
		s, v := condition.SubConditions[0].Build(n)
		val = append(val, v...)
		str = "( " + s
		n++
		for _, c := range condition.SubConditions[1:] {
			s, v = c.Build(n)
			val = append(val, v...)
			str = str + " AND " + s
			n++
		}
		str = str + " )"
		return str, val
	}
}

// ////////////////////////////////////////////////////////
// or condition

func Or(args ...ConditionInterface) ConditionInterface {
	or_condition := OrCondition{}
	for _, a := range args {
		or_condition.Add(a)
	}
	return &or_condition
}

type OrCondition struct {
	SubConditions []ConditionInterface
}

func (condition *OrCondition) Add(arg ConditionInterface) {
	condition.SubConditions = append(condition.SubConditions, arg)
}

func (condition *OrCondition) Build(n int) (
	str string,
	val []interface{},
) {
	if len(condition.SubConditions) == 0 {
		return "", val
	} else if len(condition.SubConditions) == 1 {
		return condition.SubConditions[0].Build(n)

	} else {
		s, v := condition.SubConditions[0].Build(n)
		val = append(val, v...)
		str = "( " + s
		n++
		for _, c := range condition.SubConditions[1:] {
			s, v = c.Build(n)
			val = append(val, v...)
			str = str + " OR " + s
			n++
		}
		str = str + " )"
		return str, val
	}
}
