// ampli.go
//
// Ampli - A strong typed wrapper for your Analytics
//
// This file is generated by Amplitude.
// To update run 'ampli pull go-ampli'
//
// Required dependencies: github.com/amplitude/analytics-go/amplitude@v0.0.5
// Tracking Plan Version: 0
// Build: 1.0.0
// Runtime: go-ampli
//
// View Tracking Plan: https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest
//
// Full Setup Instructions: https://data.amplitude.com/test-codegen/Test%20Codegen/implementation/main/latest/getting-started/go-ampli
//

package ampli

import (
	"log"
	"sync"

	"github.com/amplitude/analytics-go/amplitude"
)

type (
	EventOptions  = amplitude.EventOptions
	ExecuteResult = amplitude.ExecuteResult
)

const (
	IdentifyEventType      = amplitude.IdentifyEventType
	GroupIdentifyEventType = amplitude.GroupIdentifyEventType

	ServerZoneUS = amplitude.ServerZoneUS
	ServerZoneEU = amplitude.ServerZoneEU
)

var (
	NewClientConfig = amplitude.NewConfig
	NewClient       = amplitude.NewClient
)

var Instance = Ampli{}

type Environment string

const (
	EnvironmentDevelopment Environment = `development`

	EnvironmentProduction Environment = `production`
)

var APIKey = map[Environment]string{
	EnvironmentDevelopment: ``,

	EnvironmentProduction: ``,
}

// LoadClientOptions is Client options setting to initialize Ampli client.
//
// Params:
//   - APIKey: the API key of Amplitude project
//   - Instance: the core SDK instance used by Ampli client
//   - Configuration: the core SDK client configuration instance
type LoadClientOptions struct {
	APIKey        string
	Instance      amplitude.Client
	Configuration amplitude.Config
}

// LoadOptions is options setting to initialize Ampli client.
//
// Params:
//   - Environment: the environment of Amplitude Data project
//   - Disabled: the flag of disabled Ampli client
//   - Client: the LoadClientOptions struct
type LoadOptions struct {
	Environment Environment
	Disabled    bool
	Client      LoadClientOptions
}

type baseEvent struct {
	eventType  string
	properties map[string]interface{}
}

type Event interface {
	ToAmplitudeEvent() amplitude.Event
}

func newBaseEvent(eventType string, properties map[string]interface{}) baseEvent {
	return baseEvent{
		eventType:  eventType,
		properties: properties,
	}
}

func (event baseEvent) ToAmplitudeEvent() amplitude.Event {
	return amplitude.Event{
		EventType:       event.eventType,
		EventProperties: event.properties,
	}
}

var Identify = struct {
	Builder func() interface {
		RequiredNumber(requiredNumber float64) IdentifyBuilder
	}
}{
	Builder: func() interface {
		RequiredNumber(requiredNumber float64) IdentifyBuilder
	} {
		return &identifyBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type IdentifyEvent interface {
	Event
	identify()
}

type identifyEvent struct {
	baseEvent
}

func (e identifyEvent) identify() {
}

type IdentifyBuilder interface {
	Build() IdentifyEvent
	OptionalArray(optionalArray []string) IdentifyBuilder
}

type identifyBuilder struct {
	properties map[string]interface{}
}

func (b *identifyBuilder) RequiredNumber(requiredNumber float64) IdentifyBuilder {
	b.properties[`requiredNumber`] = requiredNumber

	return b
}

func (b *identifyBuilder) OptionalArray(optionalArray []string) IdentifyBuilder {
	b.properties[`optionalArray`] = optionalArray

	return b
}

func (b *identifyBuilder) Build() IdentifyEvent {
	return &identifyEvent{
		newBaseEvent(`Identify`, b.properties),
	}
}

func (e identifyEvent) ToAmplitudeEvent() amplitude.Event {
	identify := amplitude.Identify{}
	for name, value := range e.properties {
		identify.Set(name, value)
	}

	return amplitude.Event{
		EventType:      IdentifyEventType,
		UserProperties: identify.Properties,
	}
}

var Group = struct {
	Builder func() interface {
		RequiredBoolean(requiredBoolean bool) GroupBuilder
	}
}{
	Builder: func() interface {
		RequiredBoolean(requiredBoolean bool) GroupBuilder
	} {
		return &groupBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type GroupEvent interface {
	Event
	group()
}

type groupEvent struct {
	baseEvent
}

func (e groupEvent) group() {
}

type GroupBuilder interface {
	Build() GroupEvent
	OptionalString(optionalString string) GroupBuilder
}

type groupBuilder struct {
	properties map[string]interface{}
}

func (b *groupBuilder) RequiredBoolean(requiredBoolean bool) GroupBuilder {
	b.properties[`requiredBoolean`] = requiredBoolean

	return b
}

func (b *groupBuilder) OptionalString(optionalString string) GroupBuilder {
	b.properties[`optionalString`] = optionalString

	return b
}

func (b *groupBuilder) Build() GroupEvent {
	return &groupEvent{
		newBaseEvent(`Group`, b.properties),
	}
}

func (e groupEvent) ToAmplitudeEvent() amplitude.Event {
	identify := amplitude.Identify{}
	for name, value := range e.properties {
		identify.Set(name, value)
	}

	return amplitude.Event{
		EventType:       GroupIdentifyEventType,
		GroupProperties: identify.Properties,
	}
}

var EventMaxIntForTest = struct {
	Builder func() interface {
		IntMax10(intMax10 int) EventMaxIntForTestBuilder
	}
}{
	Builder: func() interface {
		IntMax10(intMax10 int) EventMaxIntForTestBuilder
	} {
		return &eventMaxIntForTestBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type EventMaxIntForTestEvent interface {
	Event
	eventMaxIntForTest()
}

type eventMaxIntForTestEvent struct {
	baseEvent
}

func (e eventMaxIntForTestEvent) eventMaxIntForTest() {
}

type EventMaxIntForTestBuilder interface {
	Build() EventMaxIntForTestEvent
}

type eventMaxIntForTestBuilder struct {
	properties map[string]interface{}
}

func (b *eventMaxIntForTestBuilder) IntMax10(intMax10 int) EventMaxIntForTestBuilder {
	b.properties[`intMax10`] = intMax10

	return b
}

func (b *eventMaxIntForTestBuilder) Build() EventMaxIntForTestEvent {
	return &eventMaxIntForTestEvent{
		newBaseEvent(`EventMaxIntForTest`, b.properties),
	}
}

var EventNoProperties = struct {
	Builder func() EventNoPropertiesBuilder
}{
	Builder: func() EventNoPropertiesBuilder {
		return &eventNoPropertiesBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type EventNoPropertiesEvent interface {
	Event
	eventNoProperties()
}

type eventNoPropertiesEvent struct {
	baseEvent
}

func (e eventNoPropertiesEvent) eventNoProperties() {
}

type EventNoPropertiesBuilder interface {
	Build() EventNoPropertiesEvent
}

type eventNoPropertiesBuilder struct {
	properties map[string]interface{}
}

func (b *eventNoPropertiesBuilder) Build() EventNoPropertiesEvent {
	return &eventNoPropertiesEvent{
		newBaseEvent(`Event No Properties`, b.properties),
	}
}

var EventObjectTypes = struct {
	Builder func() interface {
		RequiredObject(requiredObject interface{}) interface {
			RequiredObjectArray(requiredObjectArray []interface{}) EventObjectTypesBuilder
		}
	}
}{
	Builder: func() interface {
		RequiredObject(requiredObject interface{}) interface {
			RequiredObjectArray(requiredObjectArray []interface{}) EventObjectTypesBuilder
		}
	} {
		return &eventObjectTypesBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type EventObjectTypesEvent interface {
	Event
	eventObjectTypes()
}

type eventObjectTypesEvent struct {
	baseEvent
}

func (e eventObjectTypesEvent) eventObjectTypes() {
}

type EventObjectTypesBuilder interface {
	Build() EventObjectTypesEvent
}

type eventObjectTypesBuilder struct {
	properties map[string]interface{}
}

func (b *eventObjectTypesBuilder) RequiredObject(requiredObject interface{}) interface {
	RequiredObjectArray(requiredObjectArray []interface{}) EventObjectTypesBuilder
} {
	b.properties[`requiredObject`] = requiredObject

	return b
}

func (b *eventObjectTypesBuilder) RequiredObjectArray(requiredObjectArray []interface{}) EventObjectTypesBuilder {
	b.properties[`requiredObjectArray`] = requiredObjectArray

	return b
}

func (b *eventObjectTypesBuilder) Build() EventObjectTypesEvent {
	return &eventObjectTypesEvent{
		newBaseEvent(`Event Object Types`, b.properties),
	}
}

type EventWithAllPropertiesRequiredEnum string

var EventWithAllProperties = struct {
	RequiredEnum struct {
		Enum1 EventWithAllPropertiesRequiredEnum

		Enum2 EventWithAllPropertiesRequiredEnum
	}
	Builder func() interface {
		RequiredArray(requiredArray []string) interface {
			RequiredBoolean(requiredBoolean bool) interface {
				RequiredEnum(requiredEnum EventWithAllPropertiesRequiredEnum) interface {
					RequiredInteger(requiredInteger int) interface {
						RequiredNumber(requiredNumber float64) interface {
							RequiredString(requiredString string) EventWithAllPropertiesBuilder
						}
					}
				}
			}
		}
	}
}{
	RequiredEnum: struct {
		Enum1 EventWithAllPropertiesRequiredEnum

		Enum2 EventWithAllPropertiesRequiredEnum
	}{
		Enum1: `Enum1`,

		Enum2: `Enum2`,
	},
	Builder: func() interface {
		RequiredArray(requiredArray []string) interface {
			RequiredBoolean(requiredBoolean bool) interface {
				RequiredEnum(requiredEnum EventWithAllPropertiesRequiredEnum) interface {
					RequiredInteger(requiredInteger int) interface {
						RequiredNumber(requiredNumber float64) interface {
							RequiredString(requiredString string) EventWithAllPropertiesBuilder
						}
					}
				}
			}
		}
	} {
		return &eventWithAllPropertiesBuilder{
			properties: map[string]interface{}{
				`requiredConst`: `some-const-value`,
			},
		}
	},
}

type EventWithAllPropertiesEvent interface {
	Event
	eventWithAllProperties()
}

type eventWithAllPropertiesEvent struct {
	baseEvent
}

func (e eventWithAllPropertiesEvent) eventWithAllProperties() {
}

type EventWithAllPropertiesBuilder interface {
	Build() EventWithAllPropertiesEvent
	OptionalString(optionalString string) EventWithAllPropertiesBuilder
}

type eventWithAllPropertiesBuilder struct {
	properties map[string]interface{}
}

func (b *eventWithAllPropertiesBuilder) RequiredArray(requiredArray []string) interface {
	RequiredBoolean(requiredBoolean bool) interface {
		RequiredEnum(requiredEnum EventWithAllPropertiesRequiredEnum) interface {
			RequiredInteger(requiredInteger int) interface {
				RequiredNumber(requiredNumber float64) interface {
					RequiredString(requiredString string) EventWithAllPropertiesBuilder
				}
			}
		}
	}
} {
	b.properties[`requiredArray`] = requiredArray

	return b
}

func (b *eventWithAllPropertiesBuilder) RequiredBoolean(requiredBoolean bool) interface {
	RequiredEnum(requiredEnum EventWithAllPropertiesRequiredEnum) interface {
		RequiredInteger(requiredInteger int) interface {
			RequiredNumber(requiredNumber float64) interface {
				RequiredString(requiredString string) EventWithAllPropertiesBuilder
			}
		}
	}
} {
	b.properties[`requiredBoolean`] = requiredBoolean

	return b
}

func (b *eventWithAllPropertiesBuilder) RequiredEnum(requiredEnum EventWithAllPropertiesRequiredEnum) interface {
	RequiredInteger(requiredInteger int) interface {
		RequiredNumber(requiredNumber float64) interface {
			RequiredString(requiredString string) EventWithAllPropertiesBuilder
		}
	}
} {
	b.properties[`requiredEnum`] = requiredEnum

	return b
}

func (b *eventWithAllPropertiesBuilder) RequiredInteger(requiredInteger int) interface {
	RequiredNumber(requiredNumber float64) interface {
		RequiredString(requiredString string) EventWithAllPropertiesBuilder
	}
} {
	b.properties[`requiredInteger`] = requiredInteger

	return b
}

func (b *eventWithAllPropertiesBuilder) RequiredNumber(requiredNumber float64) interface {
	RequiredString(requiredString string) EventWithAllPropertiesBuilder
} {
	b.properties[`requiredNumber`] = requiredNumber

	return b
}

func (b *eventWithAllPropertiesBuilder) RequiredString(requiredString string) EventWithAllPropertiesBuilder {
	b.properties[`requiredString`] = requiredString

	return b
}

func (b *eventWithAllPropertiesBuilder) OptionalString(optionalString string) EventWithAllPropertiesBuilder {
	b.properties[`optionalString`] = optionalString

	return b
}

func (b *eventWithAllPropertiesBuilder) Build() EventWithAllPropertiesEvent {
	return &eventWithAllPropertiesEvent{
		newBaseEvent(`Event With All Properties`, b.properties),
	}
}

var EventWithArrayTypes = struct {
	Builder func() interface {
		RequiredBooleanArray(requiredBooleanArray []bool) interface {
			RequiredNumberArray(requiredNumberArray []float64) interface {
				RequiredObjectArray(requiredObjectArray []interface{}) interface {
					RequiredStringArray(requiredStringArray []string) EventWithArrayTypesBuilder
				}
			}
		}
	}
}{
	Builder: func() interface {
		RequiredBooleanArray(requiredBooleanArray []bool) interface {
			RequiredNumberArray(requiredNumberArray []float64) interface {
				RequiredObjectArray(requiredObjectArray []interface{}) interface {
					RequiredStringArray(requiredStringArray []string) EventWithArrayTypesBuilder
				}
			}
		}
	} {
		return &eventWithArrayTypesBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type EventWithArrayTypesEvent interface {
	Event
	eventWithArrayTypes()
}

type eventWithArrayTypesEvent struct {
	baseEvent
}

func (e eventWithArrayTypesEvent) eventWithArrayTypes() {
}

type EventWithArrayTypesBuilder interface {
	Build() EventWithArrayTypesEvent
}

type eventWithArrayTypesBuilder struct {
	properties map[string]interface{}
}

func (b *eventWithArrayTypesBuilder) RequiredBooleanArray(requiredBooleanArray []bool) interface {
	RequiredNumberArray(requiredNumberArray []float64) interface {
		RequiredObjectArray(requiredObjectArray []interface{}) interface {
			RequiredStringArray(requiredStringArray []string) EventWithArrayTypesBuilder
		}
	}
} {
	b.properties[`requiredBooleanArray`] = requiredBooleanArray

	return b
}

func (b *eventWithArrayTypesBuilder) RequiredNumberArray(requiredNumberArray []float64) interface {
	RequiredObjectArray(requiredObjectArray []interface{}) interface {
		RequiredStringArray(requiredStringArray []string) EventWithArrayTypesBuilder
	}
} {
	b.properties[`requiredNumberArray`] = requiredNumberArray

	return b
}

func (b *eventWithArrayTypesBuilder) RequiredObjectArray(requiredObjectArray []interface{}) interface {
	RequiredStringArray(requiredStringArray []string) EventWithArrayTypesBuilder
} {
	b.properties[`requiredObjectArray`] = requiredObjectArray

	return b
}

func (b *eventWithArrayTypesBuilder) RequiredStringArray(requiredStringArray []string) EventWithArrayTypesBuilder {
	b.properties[`requiredStringArray`] = requiredStringArray

	return b
}

func (b *eventWithArrayTypesBuilder) Build() EventWithArrayTypesEvent {
	return &eventWithArrayTypesEvent{
		newBaseEvent(`Event With Array Types`, b.properties),
	}
}

var EventWithConstTypes = struct {
	Builder func() EventWithConstTypesBuilder
}{
	Builder: func() EventWithConstTypesBuilder {
		return &eventWithConstTypesBuilder{
			properties: map[string]interface{}{
				`Boolean Const`: true,

				`Integer Const`: 10,

				`Number Const`: 2.2,

				`String Const`: `String-Constant`,

				`String Const WIth Quotes`: `"String "Const With" Quotes"`,

				`String Int Const`: 0,
			},
		}
	},
}

type EventWithConstTypesEvent interface {
	Event
	eventWithConstTypes()
}

type eventWithConstTypesEvent struct {
	baseEvent
}

func (e eventWithConstTypesEvent) eventWithConstTypes() {
}

type EventWithConstTypesBuilder interface {
	Build() EventWithConstTypesEvent
}

type eventWithConstTypesBuilder struct {
	properties map[string]interface{}
}

func (b *eventWithConstTypesBuilder) Build() EventWithConstTypesEvent {
	return &eventWithConstTypesEvent{
		newBaseEvent(`Event With Const Types`, b.properties),
	}
}

type EventWithDifferentCasingTypesEnumCamelCase string

type EventWithDifferentCasingTypesEnumPascalCase string

type EventWithDifferentCasingTypesEnumSnakeCase string

type EventWithDifferentCasingTypesEnumWithSpace string

var EventWithDifferentCasingTypes = struct {
	EnumCamelCase struct {
		EnumCamelCase EventWithDifferentCasingTypesEnumCamelCase
	}
	EnumPascalCase struct {
		EnumPascalCase EventWithDifferentCasingTypesEnumPascalCase
	}
	EnumSnakeCase struct {
		EnumSnakeCase EventWithDifferentCasingTypesEnumSnakeCase
	}
	EnumWithSpace struct {
		EnumWithSpace EventWithDifferentCasingTypesEnumWithSpace
	}
	Builder func() interface {
		EnumCamelCase(enumCamelCase EventWithDifferentCasingTypesEnumCamelCase) interface {
			EnumPascalCase(enumPascalCase EventWithDifferentCasingTypesEnumPascalCase) interface {
				EnumSnakeCase(enumSnakeCase EventWithDifferentCasingTypesEnumSnakeCase) interface {
					EnumWithSpace(enumWithSpace EventWithDifferentCasingTypesEnumWithSpace) interface {
						PropertyWithCamelCase(propertyWithCamelCase string) interface {
							PropertyWithPascalCase(propertyWithPascalCase string) interface {
								PropertyWithSnakeCase(propertyWithSnakeCase string) interface {
									PropertyWithSpace(propertyWithSpace string) EventWithDifferentCasingTypesBuilder
								}
							}
						}
					}
				}
			}
		}
	}
}{
	EnumCamelCase: struct {
		EnumCamelCase EventWithDifferentCasingTypesEnumCamelCase
	}{
		EnumCamelCase: `enumCamelCase`,
	},
	EnumPascalCase: struct {
		EnumPascalCase EventWithDifferentCasingTypesEnumPascalCase
	}{
		EnumPascalCase: `EnumPascalCase`,
	},
	EnumSnakeCase: struct {
		EnumSnakeCase EventWithDifferentCasingTypesEnumSnakeCase
	}{
		EnumSnakeCase: `enum_snake_case`,
	},
	EnumWithSpace: struct {
		EnumWithSpace EventWithDifferentCasingTypesEnumWithSpace
	}{
		EnumWithSpace: `enum with space`,
	},
	Builder: func() interface {
		EnumCamelCase(enumCamelCase EventWithDifferentCasingTypesEnumCamelCase) interface {
			EnumPascalCase(enumPascalCase EventWithDifferentCasingTypesEnumPascalCase) interface {
				EnumSnakeCase(enumSnakeCase EventWithDifferentCasingTypesEnumSnakeCase) interface {
					EnumWithSpace(enumWithSpace EventWithDifferentCasingTypesEnumWithSpace) interface {
						PropertyWithCamelCase(propertyWithCamelCase string) interface {
							PropertyWithPascalCase(propertyWithPascalCase string) interface {
								PropertyWithSnakeCase(propertyWithSnakeCase string) interface {
									PropertyWithSpace(propertyWithSpace string) EventWithDifferentCasingTypesBuilder
								}
							}
						}
					}
				}
			}
		}
	} {
		return &eventWithDifferentCasingTypesBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type EventWithDifferentCasingTypesEvent interface {
	Event
	eventWithDifferentCasingTypes()
}

type eventWithDifferentCasingTypesEvent struct {
	baseEvent
}

func (e eventWithDifferentCasingTypesEvent) eventWithDifferentCasingTypes() {
}

type EventWithDifferentCasingTypesBuilder interface {
	Build() EventWithDifferentCasingTypesEvent
}

type eventWithDifferentCasingTypesBuilder struct {
	properties map[string]interface{}
}

func (b *eventWithDifferentCasingTypesBuilder) EnumCamelCase(enumCamelCase EventWithDifferentCasingTypesEnumCamelCase) interface {
	EnumPascalCase(enumPascalCase EventWithDifferentCasingTypesEnumPascalCase) interface {
		EnumSnakeCase(enumSnakeCase EventWithDifferentCasingTypesEnumSnakeCase) interface {
			EnumWithSpace(enumWithSpace EventWithDifferentCasingTypesEnumWithSpace) interface {
				PropertyWithCamelCase(propertyWithCamelCase string) interface {
					PropertyWithPascalCase(propertyWithPascalCase string) interface {
						PropertyWithSnakeCase(propertyWithSnakeCase string) interface {
							PropertyWithSpace(propertyWithSpace string) EventWithDifferentCasingTypesBuilder
						}
					}
				}
			}
		}
	}
} {
	b.properties[`enumCamelCase`] = enumCamelCase

	return b
}

func (b *eventWithDifferentCasingTypesBuilder) EnumPascalCase(enumPascalCase EventWithDifferentCasingTypesEnumPascalCase) interface {
	EnumSnakeCase(enumSnakeCase EventWithDifferentCasingTypesEnumSnakeCase) interface {
		EnumWithSpace(enumWithSpace EventWithDifferentCasingTypesEnumWithSpace) interface {
			PropertyWithCamelCase(propertyWithCamelCase string) interface {
				PropertyWithPascalCase(propertyWithPascalCase string) interface {
					PropertyWithSnakeCase(propertyWithSnakeCase string) interface {
						PropertyWithSpace(propertyWithSpace string) EventWithDifferentCasingTypesBuilder
					}
				}
			}
		}
	}
} {
	b.properties[`EnumPascalCase`] = enumPascalCase

	return b
}

func (b *eventWithDifferentCasingTypesBuilder) EnumSnakeCase(enumSnakeCase EventWithDifferentCasingTypesEnumSnakeCase) interface {
	EnumWithSpace(enumWithSpace EventWithDifferentCasingTypesEnumWithSpace) interface {
		PropertyWithCamelCase(propertyWithCamelCase string) interface {
			PropertyWithPascalCase(propertyWithPascalCase string) interface {
				PropertyWithSnakeCase(propertyWithSnakeCase string) interface {
					PropertyWithSpace(propertyWithSpace string) EventWithDifferentCasingTypesBuilder
				}
			}
		}
	}
} {
	b.properties[`enum_snake_case`] = enumSnakeCase

	return b
}

func (b *eventWithDifferentCasingTypesBuilder) EnumWithSpace(enumWithSpace EventWithDifferentCasingTypesEnumWithSpace) interface {
	PropertyWithCamelCase(propertyWithCamelCase string) interface {
		PropertyWithPascalCase(propertyWithPascalCase string) interface {
			PropertyWithSnakeCase(propertyWithSnakeCase string) interface {
				PropertyWithSpace(propertyWithSpace string) EventWithDifferentCasingTypesBuilder
			}
		}
	}
} {
	b.properties[`enum with space`] = enumWithSpace

	return b
}

func (b *eventWithDifferentCasingTypesBuilder) PropertyWithCamelCase(propertyWithCamelCase string) interface {
	PropertyWithPascalCase(propertyWithPascalCase string) interface {
		PropertyWithSnakeCase(propertyWithSnakeCase string) interface {
			PropertyWithSpace(propertyWithSpace string) EventWithDifferentCasingTypesBuilder
		}
	}
} {
	b.properties[`propertyWithCamelCase`] = propertyWithCamelCase

	return b
}

func (b *eventWithDifferentCasingTypesBuilder) PropertyWithPascalCase(propertyWithPascalCase string) interface {
	PropertyWithSnakeCase(propertyWithSnakeCase string) interface {
		PropertyWithSpace(propertyWithSpace string) EventWithDifferentCasingTypesBuilder
	}
} {
	b.properties[`PropertyWithPascalCase`] = propertyWithPascalCase

	return b
}

func (b *eventWithDifferentCasingTypesBuilder) PropertyWithSnakeCase(propertyWithSnakeCase string) interface {
	PropertyWithSpace(propertyWithSpace string) EventWithDifferentCasingTypesBuilder
} {
	b.properties[`property_with_snake_case`] = propertyWithSnakeCase

	return b
}

func (b *eventWithDifferentCasingTypesBuilder) PropertyWithSpace(propertyWithSpace string) EventWithDifferentCasingTypesBuilder {
	b.properties[`property with space`] = propertyWithSpace

	return b
}

func (b *eventWithDifferentCasingTypesBuilder) Build() EventWithDifferentCasingTypesEvent {
	return &eventWithDifferentCasingTypesEvent{
		newBaseEvent(`event withDifferent_CasingTypes`, b.properties),
	}
}

type EventWithEnumTypesOptionalEnum string

type EventWithEnumTypesRequiredEnum string

var EventWithEnumTypes = struct {
	OptionalEnum struct {
		OptionalEnum1 EventWithEnumTypesOptionalEnum

		OptionalEnum2 EventWithEnumTypesOptionalEnum
	}
	RequiredEnum struct {
		RequiredEnum1 EventWithEnumTypesRequiredEnum

		RequiredEnum2 EventWithEnumTypesRequiredEnum
	}
	Builder func() interface {
		RequiredEnum(requiredEnum EventWithEnumTypesRequiredEnum) EventWithEnumTypesBuilder
	}
}{
	OptionalEnum: struct {
		OptionalEnum1 EventWithEnumTypesOptionalEnum

		OptionalEnum2 EventWithEnumTypesOptionalEnum
	}{
		OptionalEnum1: `optional enum 1`,

		OptionalEnum2: `optional enum 2`,
	},
	RequiredEnum: struct {
		RequiredEnum1 EventWithEnumTypesRequiredEnum

		RequiredEnum2 EventWithEnumTypesRequiredEnum
	}{
		RequiredEnum1: `required enum 1`,

		RequiredEnum2: `required enum 2`,
	},
	Builder: func() interface {
		RequiredEnum(requiredEnum EventWithEnumTypesRequiredEnum) EventWithEnumTypesBuilder
	} {
		return &eventWithEnumTypesBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type EventWithEnumTypesEvent interface {
	Event
	eventWithEnumTypes()
}

type eventWithEnumTypesEvent struct {
	baseEvent
}

func (e eventWithEnumTypesEvent) eventWithEnumTypes() {
}

type EventWithEnumTypesBuilder interface {
	Build() EventWithEnumTypesEvent
	OptionalEnum(optionalEnum EventWithEnumTypesOptionalEnum) EventWithEnumTypesBuilder
}

type eventWithEnumTypesBuilder struct {
	properties map[string]interface{}
}

func (b *eventWithEnumTypesBuilder) RequiredEnum(requiredEnum EventWithEnumTypesRequiredEnum) EventWithEnumTypesBuilder {
	b.properties[`required enum`] = requiredEnum

	return b
}

func (b *eventWithEnumTypesBuilder) OptionalEnum(optionalEnum EventWithEnumTypesOptionalEnum) EventWithEnumTypesBuilder {
	b.properties[`optional enum`] = optionalEnum

	return b
}

func (b *eventWithEnumTypesBuilder) Build() EventWithEnumTypesEvent {
	return &eventWithEnumTypesEvent{
		newBaseEvent(`Event With Enum Types`, b.properties),
	}
}

var EventWithOptionalArrayTypes = struct {
	Builder func() EventWithOptionalArrayTypesBuilder
}{
	Builder: func() EventWithOptionalArrayTypesBuilder {
		return &eventWithOptionalArrayTypesBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type EventWithOptionalArrayTypesEvent interface {
	Event
	eventWithOptionalArrayTypes()
}

type eventWithOptionalArrayTypesEvent struct {
	baseEvent
}

func (e eventWithOptionalArrayTypesEvent) eventWithOptionalArrayTypes() {
}

type EventWithOptionalArrayTypesBuilder interface {
	Build() EventWithOptionalArrayTypesEvent
	OptionalBooleanArray(optionalBooleanArray []bool) EventWithOptionalArrayTypesBuilder
	OptionalJsonArray(optionalJsonArray []interface{}) EventWithOptionalArrayTypesBuilder
	OptionalNumberArray(optionalNumberArray []float64) EventWithOptionalArrayTypesBuilder
	OptionalStringArray(optionalStringArray []string) EventWithOptionalArrayTypesBuilder
}

type eventWithOptionalArrayTypesBuilder struct {
	properties map[string]interface{}
}

func (b *eventWithOptionalArrayTypesBuilder) OptionalBooleanArray(optionalBooleanArray []bool) EventWithOptionalArrayTypesBuilder {
	b.properties[`optionalBooleanArray`] = optionalBooleanArray

	return b
}

func (b *eventWithOptionalArrayTypesBuilder) OptionalJsonArray(optionalJsonArray []interface{}) EventWithOptionalArrayTypesBuilder {
	b.properties[`optionalJSONArray`] = optionalJsonArray

	return b
}

func (b *eventWithOptionalArrayTypesBuilder) OptionalNumberArray(optionalNumberArray []float64) EventWithOptionalArrayTypesBuilder {
	b.properties[`optionalNumberArray`] = optionalNumberArray

	return b
}

func (b *eventWithOptionalArrayTypesBuilder) OptionalStringArray(optionalStringArray []string) EventWithOptionalArrayTypesBuilder {
	b.properties[`optionalStringArray`] = optionalStringArray

	return b
}

func (b *eventWithOptionalArrayTypesBuilder) Build() EventWithOptionalArrayTypesEvent {
	return &eventWithOptionalArrayTypesEvent{
		newBaseEvent(`Event With Optional Array Types`, b.properties),
	}
}

var EventWithOptionalProperties = struct {
	Builder func() EventWithOptionalPropertiesBuilder
}{
	Builder: func() EventWithOptionalPropertiesBuilder {
		return &eventWithOptionalPropertiesBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type EventWithOptionalPropertiesEvent interface {
	Event
	eventWithOptionalProperties()
}

type eventWithOptionalPropertiesEvent struct {
	baseEvent
}

func (e eventWithOptionalPropertiesEvent) eventWithOptionalProperties() {
}

type EventWithOptionalPropertiesBuilder interface {
	Build() EventWithOptionalPropertiesEvent
	OptionalArrayNumber(optionalArrayNumber []float64) EventWithOptionalPropertiesBuilder
	OptionalArrayString(optionalArrayString []string) EventWithOptionalPropertiesBuilder
	OptionalBoolean(optionalBoolean bool) EventWithOptionalPropertiesBuilder
	OptionalNumber(optionalNumber float64) EventWithOptionalPropertiesBuilder
	OptionalString(optionalString string) EventWithOptionalPropertiesBuilder
}

type eventWithOptionalPropertiesBuilder struct {
	properties map[string]interface{}
}

func (b *eventWithOptionalPropertiesBuilder) OptionalArrayNumber(optionalArrayNumber []float64) EventWithOptionalPropertiesBuilder {
	b.properties[`optionalArrayNumber`] = optionalArrayNumber

	return b
}

func (b *eventWithOptionalPropertiesBuilder) OptionalArrayString(optionalArrayString []string) EventWithOptionalPropertiesBuilder {
	b.properties[`optionalArrayString`] = optionalArrayString

	return b
}

func (b *eventWithOptionalPropertiesBuilder) OptionalBoolean(optionalBoolean bool) EventWithOptionalPropertiesBuilder {
	b.properties[`optionalBoolean`] = optionalBoolean

	return b
}

func (b *eventWithOptionalPropertiesBuilder) OptionalNumber(optionalNumber float64) EventWithOptionalPropertiesBuilder {
	b.properties[`optionalNumber`] = optionalNumber

	return b
}

func (b *eventWithOptionalPropertiesBuilder) OptionalString(optionalString string) EventWithOptionalPropertiesBuilder {
	b.properties[`optionalString`] = optionalString

	return b
}

func (b *eventWithOptionalPropertiesBuilder) Build() EventWithOptionalPropertiesEvent {
	return &eventWithOptionalPropertiesEvent{
		newBaseEvent(`Event With Optional Properties`, b.properties),
	}
}

var EventWithTemplateProperties = struct {
	Builder func() interface {
		RequiredEventProperty(requiredEventProperty string) interface {
			RequiredTemplateProperty(requiredTemplateProperty string) EventWithTemplatePropertiesBuilder
		}
	}
}{
	Builder: func() interface {
		RequiredEventProperty(requiredEventProperty string) interface {
			RequiredTemplateProperty(requiredTemplateProperty string) EventWithTemplatePropertiesBuilder
		}
	} {
		return &eventWithTemplatePropertiesBuilder{
			properties: map[string]interface{}{},
		}
	},
}

type EventWithTemplatePropertiesEvent interface {
	Event
	eventWithTemplateProperties()
}

type eventWithTemplatePropertiesEvent struct {
	baseEvent
}

func (e eventWithTemplatePropertiesEvent) eventWithTemplateProperties() {
}

type EventWithTemplatePropertiesBuilder interface {
	Build() EventWithTemplatePropertiesEvent
	OptionalEventProperty(optionalEventProperty float64) EventWithTemplatePropertiesBuilder
	OptionalTemplateProperty(optionalTemplateProperty float64) EventWithTemplatePropertiesBuilder
}

type eventWithTemplatePropertiesBuilder struct {
	properties map[string]interface{}
}

func (b *eventWithTemplatePropertiesBuilder) RequiredEventProperty(requiredEventProperty string) interface {
	RequiredTemplateProperty(requiredTemplateProperty string) EventWithTemplatePropertiesBuilder
} {
	b.properties[`required_event_property`] = requiredEventProperty

	return b
}

func (b *eventWithTemplatePropertiesBuilder) RequiredTemplateProperty(requiredTemplateProperty string) EventWithTemplatePropertiesBuilder {
	b.properties[`required_template_property`] = requiredTemplateProperty

	return b
}

func (b *eventWithTemplatePropertiesBuilder) OptionalEventProperty(optionalEventProperty float64) EventWithTemplatePropertiesBuilder {
	b.properties[`optional_event_property`] = optionalEventProperty

	return b
}

func (b *eventWithTemplatePropertiesBuilder) OptionalTemplateProperty(optionalTemplateProperty float64) EventWithTemplatePropertiesBuilder {
	b.properties[`optional_template_property`] = optionalTemplateProperty

	return b
}

func (b *eventWithTemplatePropertiesBuilder) Build() EventWithTemplatePropertiesEvent {
	return &eventWithTemplatePropertiesEvent{
		newBaseEvent(`Event With Template Properties`, b.properties),
	}
}

type Ampli struct {
	Disabled bool
	Client   amplitude.Client
	mutex    sync.RWMutex
}

// Load initializes the Ampli wrapper.
// Call once when your application starts.
func (a *Ampli) Load(options LoadOptions) {
	if a.Client != nil {
		log.Print("Warn: Ampli is already initialized. Ampli.Load() should be called once at application start up.")

		return
	}

	var apiKey string
	switch {
	case options.Client.APIKey != "":
		apiKey = options.Client.APIKey
	case options.Environment != "":
		apiKey = APIKey[options.Environment]
	default:
		apiKey = options.Client.Configuration.APIKey
	}

	if apiKey == "" && options.Client.Instance == nil {
		log.Print("Error: Ampli.Load() requires option.Environment, " +
			"and apiKey from either options.Instance.APIKey or APIKey[options.Environment], " +
			"or options.Instance.Instance")
	}

	clientConfig := options.Client.Configuration

	if clientConfig.Plan == nil {
		clientConfig.Plan = &amplitude.Plan{
			Branch:    `main`,
			Source:    `go-ampli`,
			Version:   `0`,
			VersionID: `79154a50-f057-4db5-9755-775e4e9f05e6`,
		}
	}

	if clientConfig.IngestionMetadata == nil {
		clientConfig.IngestionMetadata = &amplitude.IngestionMetadata{
			SourceName:    `go-go-ampli`,
			SourceVersion: `2.0.0`,
		}
	}

	if options.Client.Instance != nil {
		a.Client = options.Client.Instance
	} else {
		clientConfig.APIKey = apiKey
		a.Client = amplitude.NewClient(clientConfig)
	}

	a.mutex.Lock()
	a.Disabled = options.Disabled
	a.mutex.Unlock()
}

// InitializedAndEnabled checks if Ampli is initialized and enabled.
func (a *Ampli) InitializedAndEnabled() bool {
	if a.Client == nil {
		log.Print("Error: Ampli is not yet initialized. Have you called Ampli.Load() on app start?")

		return false
	}

	a.mutex.RLock()
	defer a.mutex.RUnlock()

	return !a.Disabled
}

func (a *Ampli) setUserID(userID string, eventOptions *EventOptions) {
	if userID != "" {
		eventOptions.UserID = userID
	}
}

// Track tracks an event.
func (a *Ampli) Track(userID string, event Event, eventOptions ...EventOptions) {
	if !a.InitializedAndEnabled() {
		return
	}

	var options EventOptions
	if len(eventOptions) > 0 {
		options = eventOptions[0]
	}

	a.setUserID(userID, &options)

	baseEvent := event.ToAmplitudeEvent()
	baseEvent.EventOptions = options

	a.Client.Track(baseEvent)
}

// Identify identifies a user and set user properties.
func (a *Ampli) Identify(userID string, identify IdentifyEvent, eventOptions ...EventOptions) {
	a.Track(userID, identify, eventOptions...)
}

// GroupIdentify identifies a group and set group properties.
func (a *Ampli) GroupIdentify(groupType string, groupName string, group GroupEvent, eventOptions ...EventOptions) {
	event := group.ToAmplitudeEvent()
	event.Groups = map[string][]string{groupType: {groupName}}
	if len(eventOptions) > 0 {
		event.EventOptions = eventOptions[0]
	}

	a.Client.Track(event)
}

// SetGroup sets group for the current user.
func (a *Ampli) SetGroup(userID string, groupType string, groupName []string, eventOptions ...EventOptions) {
	var options EventOptions
	if len(eventOptions) > 0 {
		options = eventOptions[0]
	}
	a.setUserID(userID, &options)

	a.Client.SetGroup(groupType, groupName, options)
}

// Flush flushes events waiting in buffer.
func (a *Ampli) Flush() {
	if !a.InitializedAndEnabled() {
		return
	}

	a.Client.Flush()
}

// Shutdown disables and shutdowns Ampli Instance.
func (a *Ampli) Shutdown() {
	if !a.InitializedAndEnabled() {
		return
	}

	a.mutex.Lock()
	a.Disabled = true
	a.mutex.Unlock()

	a.Client.Shutdown()
}

func (a *Ampli) EventMaxIntForTest(userID string, event EventMaxIntForTestEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) EventNoProperties(userID string, eventOptions ...EventOptions) {
	event := EventNoProperties.Builder().Build()
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) EventObjectTypes(userID string, event EventObjectTypesEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) EventWithAllProperties(userID string, event EventWithAllPropertiesEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) EventWithArrayTypes(userID string, event EventWithArrayTypesEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) EventWithConstTypes(userID string, eventOptions ...EventOptions) {
	event := EventWithConstTypes.Builder().Build()
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) EventWithDifferentCasingTypes(userID string, event EventWithDifferentCasingTypesEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) EventWithEnumTypes(userID string, event EventWithEnumTypesEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) EventWithOptionalArrayTypes(userID string, event EventWithOptionalArrayTypesEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) EventWithOptionalProperties(userID string, event EventWithOptionalPropertiesEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}

func (a *Ampli) EventWithTemplateProperties(userID string, event EventWithTemplatePropertiesEvent, eventOptions ...EventOptions) {
	a.Track(userID, event, eventOptions...)
}
