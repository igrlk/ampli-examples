package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"ampli-example/ampli"
)

const userID = "ampli-go-user-id"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}
	apiKey := os.Getenv("AMPLITUDE_API_KEY")

	// Initialize the Ampli instance with LoadOptions and LoadClientOptions
	ampli.Instance.Load(ampli.LoadOptions{
		Client: ampli.LoadClientOptions{
			Configuration: ampli.NewClientConfig(apiKey),
		},
	})
	defer ampli.Instance.Shutdown()

	// Identify using IdentifyProperties in tracking plan
	ampli.Instance.Identify(userID, ampli.Identify.Builder().
		RequiredNumber(16.6).
		OptionalArray([]string{"abc", "test"}).
		Build())

	// Group Identify using GroupProperties in tracking plan
	ampli.Instance.GroupIdentify("Org", "Engineer", ampli.Group.Builder().
		RequiredBoolean(true).
		OptionalString("optional-string").
		Build())

	// Set groups for user
	ampli.Instance.SetGroup(userID, "Org", []string{"Engineer", "DevOp"})

	ampli.Instance.EventNoProperties(userID)

	ampli.Instance.Track(userID, ampli.EventWithAllProperties.Builder().
		RequiredArray([]string{"I'm", "required"}).
		RequiredBoolean(false).
		RequiredEnum(ampli.EventWithAllProperties.RequiredEnum.Enum1).
		RequiredInteger(42).
		RequiredNumber(1.23).
		RequiredString("Hi!").
		OptionalString("optional-string").
		Build())

	ampli.Instance.Track("", ampli.EventWithOptionalProperties.Builder().
		OptionalBoolean(true).
		Build(),
		ampli.EventOptions{DeviceID: "12345"})

	ampli.Instance.EventWithOptionalProperties(userID, ampli.EventWithOptionalProperties.Builder().
		OptionalBoolean(true).
		Build())

	ampli.Instance.EventMaxIntForTest(userID, ampli.EventMaxIntForTest.Builder().
		IntMax10(5).
		Build())

	ampli.Instance.EventWithConstTypes(userID)

	ampli.Instance.EventObjectTypes(userID, ampli.EventObjectTypes.Builder().
		RequiredObject(map[string]interface{}{"key-1": "value-1"}).
		RequiredObjectArray([]interface{}{
			map[string]interface{}{"key-1": "value-1"},
			map[string]interface{}{"key-2": "value-2"},
		}).
		Build())

	ampli.Instance.EventWithArrayTypes(userID, ampli.EventWithArrayTypes.Builder().
		RequiredBooleanArray([]bool{true, false}).
		RequiredNumberArray([]float64{1.2, 3, 4.56}).
		RequiredObjectArray([]interface{}{
			map[string]interface{}{"key-1": "value-1"},
			map[string]interface{}{"key-2": "value-2"},
		}).
		RequiredStringArray([]string{"string-1", "string-2", "string-3"}).
		Build())

	ampli.Instance.EventWithEnumTypes(userID, ampli.EventWithEnumTypes.Builder().
		RequiredEnum(ampli.EventWithEnumTypes.RequiredEnum.RequiredEnum2).
		Build())

	ampli.Instance.EventWithOptionalArrayTypes(userID, ampli.EventWithOptionalArrayTypes.Builder().
		OptionalBooleanArray([]bool{true, false}).
		Build())

	ampli.Instance.EventWithTemplateProperties(userID, ampli.EventWithTemplateProperties.Builder().
		RequiredEventProperty("event property").
		RequiredTemplateProperty("template property").
		OptionalTemplateProperty(1.23).
		Build())

	ampli.Instance.Track(userID, ampli.EventWithDifferentCasingTypes.Builder().
		EnumCamelCase(ampli.EventWithDifferentCasingTypes.EnumCamelCase.EnumCamelCase).
		EnumPascalCase(ampli.EventWithDifferentCasingTypes.EnumPascalCase.EnumPascalCase).
		EnumSnakeCase(ampli.EventWithDifferentCasingTypes.EnumSnakeCase.EnumSnakeCase).
		EnumWithSpace(ampli.EventWithDifferentCasingTypes.EnumWithSpace.EnumWithSpace).
		PropertyWithCamelCase("property with camel case").
		PropertyWithPascalCase("property with pascal case").
		PropertyWithSnakeCase("property with snake case").
		PropertyWithSpace("property with space").
		Build())
}
