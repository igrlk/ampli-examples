/* tslint:disable */
/* eslint-disable */
/**
 * Ampli - A strong typed wrapper for your Analytics
 *
 * This file is generated by Amplitude.
 * To update run 'ampli pull node-ts-ampli'
 *
 * Required dependencies: @amplitude/node@^1.9.0
 * Tracking Plan Version: 0
 * Build: 1.0.0
 *
 * [View Tracking Plan](https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest)
 *
 * [Full Setup Instructions](https://data.amplitude.com/test-codegen/Test%20Codegen/implementation/node-ts-ampli)
 */

import { Identify as AmplitudeIdentify } from '@amplitude/identify';
import { init as initNodeClient, NodeClient, Response, Status } from '@amplitude/node';
import { BaseEvent, Event, EventOptions, IdentifyOptions, MiddlewareExtra, Options, } from '@amplitude/types';

export enum Environment {
  development = 'development',
  production = 'production'
}

export const ApiKey: Record<Environment, string> = {
  development: '',
  production: ''
};

/**
* Default NodeClient Options. Contains tracking plan information.
*/
export const DefaultOptions: Partial<Options> = {
  plan: {
    version: '0',
    branch: 'main',
    source: 'node-ts-ampli'
  }
};

export interface EventProperties {
    Context?:                       ContextProperties;
    EventMaxIntForTest?:            EventMaxIntForTestProperties;
    EventNoProperties?:             EventNoPropertiesProperties;
    EventObjectTypes?:              EventObjectTypesProperties;
    EventWithAllProperties?:        EventWithAllPropertiesProperties;
    EventWithArrayTypes?:           EventWithArrayTypesProperties;
    EventWithConstTypes?:           EventWithConstTypesProperties;
    EventWithDifferentCasingTypes?: EventWithDifferentCasingTypesProperties;
    EventWithEnumTypes?:            EventWithEnumTypesProperties;
    EventWithOptionalArrayTypes?:   EventWithOptionalArrayTypesProperties;
    EventWithOptionalProperties?:   EventWithOptionalPropertiesProperties;
    Group?:                         GroupProperties;
    Identify?:                      IdentifyProperties;
}

export interface ContextProperties {
}

/**
 * Event to test schema validation
 */
export interface EventMaxIntForTestProperties {
    /**
     * property to test schema validation
     */
    intMax10: number;
}

/**
 * Event w no properties description
 */
export interface EventNoPropertiesProperties {
}

/**
 * Event with Object and Object Array
 */
export interface EventObjectTypesProperties {
    /**
     * Property Object Type
     */
    requiredObject: { [key: string]: any };
    /**
     * Property Object Array Type
     */
    requiredObjectArray: { [key: string]: any }[];
}

/**
 * Event w all properties description
 */
export interface EventWithAllPropertiesProperties {
    /**
     * Event 2 Property - Optional String    *     * Examples:    * Some string, or another
     */
    optionalString?: string;
    /**
     * Event 2 Property - Array
     */
    requiredArray: string[];
    /**
     * Event 2 Property - Boolean
     */
    requiredBoolean: boolean;
    /**
     * Event 2 Property - Enum
     */
    requiredEnum: RequiredEnum;
    /**
     * Event 2 Property - Integer    *     * Examples:    * 5, 4, 3
     */
    requiredInteger: number;
    /**
     * Event 2 Property - Number
     */
    requiredNumber: number;
    /**
     * Event 2 Property - String
     */
    requiredString: string;
}

/**
 * Event 2 Property - Enum
 */
export enum RequiredEnum {
    Enum1 = "Enum1",
    Enum2 = "Enum2",
}

/**
 * Description for event with Array Types
 */
export interface EventWithArrayTypesProperties {
    /**
     * description for required boolean array
     */
    requiredBooleanArray: boolean[];
    /**
     * Description for required number array
     */
    requiredNumberArray: number[];
    /**
     * Description for required object array
     */
    requiredObjectArray: { [key: string]: any }[];
    /**
     * description for required string array
     */
    requiredStringArray: string[];
}

/**
 * Description for event with const types
 */
export interface EventWithConstTypesProperties {
}

/**
 * Description for case with space
 */
export interface EventWithDifferentCasingTypesProperties {
    /**
     * Description for enum with space
     */
    "enum with space": EnumWithSpace;
    /**
     * description_for_enum_snake_case
     */
    enum_snake_case: EnumSnakeCase;
    /**
     * descriptionForEnumCamelCase
     */
    enumCamelCase: EnumCamelCase;
    /**
     * DescirptionForEnumPascalCase
     */
    EnumPascalCase: EnumPascalCase;
    /**
     * Description for case with space
     */
    "property with space": string;
    /**
     * Description_for_snake_case
     */
    property_with_snake_case: string;
    /**
     * descriptionForCamelCase
     */
    propertyWithCamelCase: string;
    /**
     * DescriptionForPascalCase
     */
    PropertyWithPascalCase: string;
}

/**
 * DescirptionForEnumPascalCase
 */
export enum EnumPascalCase {
    EnumPascalCase = "EnumPascalCase",
}

/**
 * Description for enum with space
 */
export enum EnumWithSpace {
    EnumWithSpace = "enum with space",
}

/**
 * descriptionForEnumCamelCase
 */
export enum EnumCamelCase {
    EnumCamelCase = "enumCamelCase",
}

/**
 * description_for_enum_snake_case
 */
export enum EnumSnakeCase {
    EnumSnakeCase = "enum_snake_case",
}

/**
 * Description for event with enum types
 */
export interface EventWithEnumTypesProperties {
    /**
     * Description for required enum
     */
    "optional enum"?: OptionalEnum;
    /**
     * Description for optional enum
     */
    "required enum": RequiredEnumEnum;
}

/**
 * Description for required enum
 */
export enum OptionalEnum {
    OptionalEnum1 = "optional enum 1",
    OptionalEnum2 = "optional enum 2",
}

/**
 * Description for optional enum
 */
export enum RequiredEnumEnum {
    RequiredEnum1 = "required enum 1",
    RequiredEnum2 = "required enum 2",
}

/**
 * Description for event with optional array types
 */
export interface EventWithOptionalArrayTypesProperties {
    /**
     * Description for optional boolean array
     */
    optionalBooleanArray?: boolean[];
    /**
     * Description for optional object array
     */
    optionalJSONArray?: { [key: string]: any }[];
    /**
     * Description for optional number array
     */
    optionalNumberArray?: number[];
    /**
     * Description for optional string array
     */
    optionalStringArray?: string[];
}

/**
 * Event w optional properties description
 */
export interface EventWithOptionalPropertiesProperties {
    optionalArrayNumber?: number[];
    optionalArrayString?: string[];
    optionalBoolean?:     boolean;
    optionalNumber?:      number;
    /**
     * Optional String property description
     */
    optionalString?: string;
}

export interface GroupProperties {
    /**
     * Description for group optionalString
     */
    optionalString?: string;
    /**
     * Description for group requiredBoolean
     */
    requiredBoolean: boolean;
}

export interface IdentifyProperties {
    /**
     * Description for identify optionalArray
     */
    optionalArray?: string[];
    /**
     * Description for identify requiredNumber
     */
    requiredNumber: number;
}


export class Context implements BaseEvent {
  event_type = 'Context';
  constructor() {}
}
export class Identify implements BaseEvent {
  event_type = 'Identify';
  event_properties: IdentifyProperties;

  constructor(event_properties: IdentifyProperties) {
    this.event_properties = event_properties;
  }
}
export class Group implements BaseEvent {
  event_type = 'Group';
  event_properties: GroupProperties;

  constructor(event_properties: GroupProperties) {
    this.event_properties = event_properties;
  }
}

export class EventMaxIntForTest implements BaseEvent {
  event_type = 'EventMaxIntForTest';

  constructor(
    public event_properties: EventMaxIntForTestProperties,
  ) {}
}

export class EventNoProperties implements BaseEvent {
  event_type = 'Event No Properties';
}

export class EventObjectTypes implements BaseEvent {
  event_type = 'Event Object Types';

  constructor(
    public event_properties: EventObjectTypesProperties,
  ) {}
}

export class EventWithAllProperties implements BaseEvent {
  event_type = 'Event With All Properties';
  event_properties: EventWithAllPropertiesProperties & {
    'requiredConst': "some-const-value";
  };

  constructor(
    event_properties: EventWithAllPropertiesProperties,
  ) {
    this.event_properties = {
        ...event_properties,
        'requiredConst': "some-const-value",
      };
  }
}

export class EventWithArrayTypes implements BaseEvent {
  event_type = 'Event With Array Types';

  constructor(
    public event_properties: EventWithArrayTypesProperties,
  ) {}
}

export class EventWithConstTypes implements BaseEvent {
  event_type = 'Event With Const Types';
  event_properties = {
    'Integer Const': 10,
    'Boolean Const': true,
    'String Int Const': 0,
    'Number Const': 2.2,
    'String Const WIth Quotes': "\"String \"Const With\" Quotes\"",
    'String Const': "String-Constant",
  };
}

export class EventWithDifferentCasingTypes implements BaseEvent {
  event_type = 'event withDifferent_CasingTypes';

  constructor(
    public event_properties: EventWithDifferentCasingTypesProperties,
  ) {}
}

export class EventWithEnumTypes implements BaseEvent {
  event_type = 'Event With Enum Types';

  constructor(
    public event_properties: EventWithEnumTypesProperties,
  ) {}
}

export class EventWithOptionalArrayTypes implements BaseEvent {
  event_type = 'Event With Optional Array Types';

  constructor(
    public event_properties?: EventWithOptionalArrayTypesProperties,
  ) {}
}

export class EventWithOptionalProperties implements BaseEvent {
  event_type = 'Event With Optional Properties';

  constructor(
    public event_properties?: EventWithOptionalPropertiesProperties,
  ) {}
}

// TODO: Should we keep LoadOptions flat?
// export interface LoadOptions {
//   environment?: Environment;
//   disabled?: boolean;
//
//   apiKey?: string;
//   clientOptions?: Partial<Options>;
//   client?: NodeClient;
// }

// export const LoadOptionsDefault: LoadOptions = {
//   environment: Environment.development,
//   clientOptions: DefaultOptions,
//   disabled: false,
// };

export interface LoadOptions {
  environment?: Environment;
  disabled?: boolean;
  client?: {
    apiKey?: string;
    options?: Partial<Options>;
    instance?: NodeClient;
  }
}

function getDefaultPromiseResponse(): Promise<Response> {
  return Promise.resolve<Response>({
    status: Status.Skipped,
    statusCode: 200,
  });
}

// prettier-ignore
export class Ampli {
  private disabled: boolean;
  private amplitude: NodeClient | undefined;
  
  get client() {
    return this.amplitude;
  }

  private isInitializedAndEnabled(): boolean {
    if (!this.amplitude) {
      throw new Error('Itly is not yet initialized. Have you called `itly.load()` on app start?');
    }
    return !this.disabled;
  }

  load(options: LoadOptions): void {
    this.disabled = options.disabled || false;
    const apiKey = options.client.apiKey || ApiKey[options.environment];
    if (!apiKey) {
      throw new Error(`No 'environment' or 'apiKey' provided to ampli.load()`);
    }
    this.amplitude = initNodeClient(apiKey, { ...DefaultOptions, ...options.client.options });
  }

  identify(
    userId: string | undefined,
    deviceId: string | undefined,
    properties: IdentifyProperties,
    options?: IdentifyOptions,
    extra?: MiddlewareExtra
  ) {
    const amplitudeIdentify = new AmplitudeIdentify();
    for (const [key, value] of Object.entries({ ...properties })) {
      if (value !== undefined) {
        amplitudeIdentify.set(key, value);
      }
    }
    const promise = this.isInitializedAndEnabled()
      ? this.amplitude.logEvent(
          { ...options, ...amplitudeIdentify.identifyUser(userId || null, deviceId) }, extra,
        )
      : getDefaultPromiseResponse();
    return { promise };
  }

  track(userId: string | undefined, event: Event, options?: EventOptions, extra?: MiddlewareExtra) {
    const promise = this.isInitializedAndEnabled()
      ? this.amplitude.logEvent({ ...options, ...event,  user_id: userId }, extra)
      : getDefaultPromiseResponse();
    return { promise };
  }

  flush() {
    const promise = this.isInitializedAndEnabled()
      ? this.amplitude.flush()
      : getDefaultPromiseResponse();
    return { promise };
  }

  /**
   * EventMaxIntForTest
   *
   * [View in Tracking Plan](https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/0.0.0/EventMaxIntForTest)
   *
   * Event to test schema validation
   * 
   * Owner: Test codegen
   *
   * @param userId The user's ID.
   * @param properties The event's properties (e.g. intMax10)
   * @param options Amplitude event options.
   * @param extra Extra untyped parameters for use in middleware.
   */
  eventMaxIntForTest(
    userId: string | undefined,
    properties: EventMaxIntForTestProperties,
    options?: EventOptions,
    extra?: MiddlewareExtra,
  ) {
    return this.track(userId, new EventMaxIntForTest(properties), options, extra);
  }

  /**
   * Event No Properties
   *
   * [View in Tracking Plan](https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/0.0.0/Event%20No%20Properties)
   *
   * Event w no properties description
   * 
   * Owner: Test codegen
   *
   * @param userId The user's ID.
   * @param options Amplitude event options.
   * @param extra Extra untyped parameters for use in middleware.
   */
  eventNoProperties(
    userId: string | undefined,
    options?: EventOptions,
    extra?: MiddlewareExtra,
  ) {
    return this.track(userId, new EventNoProperties(), options, extra);
  }

  /**
   * Event Object Types
   *
   * [View in Tracking Plan](https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/0.0.0/Event%20Object%20Types)
   *
   * Event with Object and Object Array
   * 
   * Owner: Test codegen
   *
   * @param userId The user's ID.
   * @param properties The event's properties (e.g. requiredObject)
   * @param options Amplitude event options.
   * @param extra Extra untyped parameters for use in middleware.
   */
  eventObjectTypes(
    userId: string | undefined,
    properties: EventObjectTypesProperties,
    options?: EventOptions,
    extra?: MiddlewareExtra,
  ) {
    return this.track(userId, new EventObjectTypes(properties), options, extra);
  }

  /**
   * Event With All Properties
   *
   * [View in Tracking Plan](https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/0.0.0/Event%20With%20All%20Properties)
   *
   * Event w all properties description
   * 
   * Owner: Test codegen
   *
   * @param userId The user's ID.
   * @param properties The event's properties (e.g. requiredInteger)
   * @param options Amplitude event options.
   * @param extra Extra untyped parameters for use in middleware.
   */
  eventWithAllProperties(
    userId: string | undefined,
    properties: EventWithAllPropertiesProperties,
    options?: EventOptions,
    extra?: MiddlewareExtra,
  ) {
    return this.track(userId, new EventWithAllProperties(properties), options, extra);
  }

  /**
   * Event With Array Types
   *
   * [View in Tracking Plan](https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/0.0.0/Event%20With%20Array%20Types)
   *
   * Description for event with Array Types
   * 
   * Owner: Test codegen
   *
   * @param userId The user's ID.
   * @param properties The event's properties (e.g. requiredBooleanArray)
   * @param options Amplitude event options.
   * @param extra Extra untyped parameters for use in middleware.
   */
  eventWithArrayTypes(
    userId: string | undefined,
    properties: EventWithArrayTypesProperties,
    options?: EventOptions,
    extra?: MiddlewareExtra,
  ) {
    return this.track(userId, new EventWithArrayTypes(properties), options, extra);
  }

  /**
   * Event With Const Types
   *
   * [View in Tracking Plan](https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/0.0.0/Event%20With%20Const%20Types)
   *
   * Description for event with const types
   * 
   * Owner: Test codegen
   *
   * @param userId The user's ID.
   * @param options Amplitude event options.
   * @param extra Extra untyped parameters for use in middleware.
   */
  eventWithConstTypes(
    userId: string | undefined,
    options?: EventOptions,
    extra?: MiddlewareExtra,
  ) {
    return this.track(userId, new EventWithConstTypes(), options, extra);
  }

  /**
   * event withDifferent_CasingTypes
   *
   * [View in Tracking Plan](https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/0.0.0/event%20withDifferent_CasingTypes)
   *
   * Description for case with space
   * 
   * Owner: Test codegen
   *
   * @param userId The user's ID.
   * @param properties The event's properties (e.g. EnumPascalCase)
   * @param options Amplitude event options.
   * @param extra Extra untyped parameters for use in middleware.
   */
  eventWithDifferentCasingTypes(
    userId: string | undefined,
    properties: EventWithDifferentCasingTypesProperties,
    options?: EventOptions,
    extra?: MiddlewareExtra,
  ) {
    return this.track(userId, new EventWithDifferentCasingTypes(properties), options, extra);
  }

  /**
   * Event With Enum Types
   *
   * [View in Tracking Plan](https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/0.0.0/Event%20With%20Enum%20Types)
   *
   * Description for event with enum types
   * 
   * Owner: Test codegen
   *
   * @param userId The user's ID.
   * @param properties The event's properties (e.g. required enum)
   * @param options Amplitude event options.
   * @param extra Extra untyped parameters for use in middleware.
   */
  eventWithEnumTypes(
    userId: string | undefined,
    properties: EventWithEnumTypesProperties,
    options?: EventOptions,
    extra?: MiddlewareExtra,
  ) {
    return this.track(userId, new EventWithEnumTypes(properties), options, extra);
  }

  /**
   * Event With Optional Array Types
   *
   * [View in Tracking Plan](https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/0.0.0/Event%20With%20Optional%20Array%20Types)
   *
   * Description for event with optional array types
   * 
   * Owner: Test codegen
   *
   * @param userId The user's ID.
   * @param properties The event's properties (e.g. optionalStringArray)
   * @param options Amplitude event options.
   * @param extra Extra untyped parameters for use in middleware.
   */
  eventWithOptionalArrayTypes(
    userId: string | undefined,
    properties?: EventWithOptionalArrayTypesProperties,
    options?: EventOptions,
    extra?: MiddlewareExtra,
  ) {
    return this.track(userId, new EventWithOptionalArrayTypes(properties), options, extra);
  }

  /**
   * Event With Optional Properties
   *
   * [View in Tracking Plan](https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/0.0.0/Event%20With%20Optional%20Properties)
   *
   * Event w optional properties description
   * 
   * Owner: Test codegen
   *
   * @param userId The user's ID.
   * @param properties The event's properties (e.g. optionalNumber)
   * @param options Amplitude event options.
   * @param extra Extra untyped parameters for use in middleware.
   */
  eventWithOptionalProperties(
    userId: string | undefined,
    properties?: EventWithOptionalPropertiesProperties,
    options?: EventOptions,
    extra?: MiddlewareExtra,
  ) {
    return this.track(userId, new EventWithOptionalProperties(properties), options, extra);
  }
}

export const ampli = new Ampli();
