import { Ampli, ApiKey } from './ampli';

describe('Ampli Browser JS SDK tests', () => {
  /** @typedef {Ampli}*/
  let ampli;
  let userId = 'test-browser-js-ampli-user-id';
  let consoleLogMock;
  let consoleErrorMock;

  beforeEach(() => {
    consoleLogMock = jest.spyOn(console, 'log').mockImplementation();
    consoleErrorMock = jest.spyOn(console, 'error').mockImplementation();

    ampli = new Ampli();
    // Set API keys for tests
    ApiKey.production = 'test-api-key-prod';
    ApiKey.development = 'test-api-key-dev';
  });

  test('should load() without any arguments if there are ApiKeys for each environment', () => {
    expect(() => ampli.load()).not.toThrow();

    expect(consoleLogMock).toHaveBeenCalledTimes(0);
    expect(consoleErrorMock).toHaveBeenCalledTimes(0);
  });

  test('should log warning if load() without any arguments without ApiKeys for each environment', () => {
    ApiKey.production = '';
    ApiKey.development = '';
    ampli.load();

    expect(consoleLogMock).toHaveBeenCalledTimes(0);
    expect(consoleErrorMock).toHaveBeenCalledTimes(1);
    expect(consoleErrorMock.mock.calls).toEqual([
      [`ERROR: ampli.load() requires 'environment', 'client.apiKey', or 'client.instance'`],
    ]);
  });

  test('should identify()', () => {
    const mockAmp = {
      setUserId: jest.fn(),
      identify: jest.fn(),
    };
    ampli.load({ client: { instance: mockAmp } });

    ampli.identify(userId, {
      optionalArray: ["A", "ray"],
      requiredNumber: 42,
    });

    const setUserIdCalls = mockAmp.setUserId.mock.calls;
    expect(setUserIdCalls.length).toBe(1);
    expect(setUserIdCalls[0]).toEqual([userId]);

    const identifyCalls = mockAmp.identify.mock.calls;
    expect(identifyCalls.length).toBe(1);
    const [ identify ] = identifyCalls[0];
    expect(identify.properties).toEqual(['optionalArray', 'requiredNumber']);
    expect(identify.userPropertiesOperations).toEqual({
      '$set': {
        'optionalArray': ['A', 'ray'],
        'requiredNumber': 42
      }
    });

    expect(consoleLogMock).toHaveBeenCalledTimes(0);
    expect(consoleErrorMock).toHaveBeenCalledTimes(0);
  });

  test('should setGroup()', () => {
    const mockAmp = { setGroup: jest.fn() };
    ampli.load({ client: { instance: mockAmp } });

    ampli.setGroup('Group name', 'Group Value');

    const setGroupCalls = mockAmp.setGroup.mock.calls;
    expect(setGroupCalls.length).toBe(1);
    expect(setGroupCalls[0]).toEqual(["Group name", "Group Value"]);
    expect(consoleLogMock).toHaveBeenCalledTimes(0);
    expect(consoleErrorMock).toHaveBeenCalledTimes(0);
  });

  test('should groupIdentify()', () => {
    const mockAmp = { groupIdentify: jest.fn() };
    ampli.load({ client: { instance: mockAmp } });

    ampli.groupIdentify('Group name', 'Group Value', { requiredBoolean: true, optionalString: 'some-string' });

    const groupIdentifyCalls = mockAmp.groupIdentify.mock.calls;
    expect(groupIdentifyCalls.length).toBe(1);
    expect(groupIdentifyCalls[0]).toEqual([
      "Group name",
      "Group Value",
      {
        "properties": [
          "requiredBoolean",
          "optionalString"
        ],
        "userPropertiesOperations": {
          "$set": {
            "optionalString": "some-string",
            "requiredBoolean": true
          }
        }
      },
      undefined
    ]);

    expect(consoleLogMock).toHaveBeenCalledTimes(0);
    expect(consoleErrorMock).toHaveBeenCalledTimes(0);
  });

  test('should track an event with no properties', () => {
    const mockAmp = {
      setUserId: jest.fn(),
      setDeviceId: jest.fn(),
      logEvent: jest.fn(),
    };
    ampli.load({ client: { instance: mockAmp } });

    ampli.eventNoProperties({ user_id: userId, device_id: 'device-1' });

    const setUserIdCalls = mockAmp.setUserId.mock.calls;
    expect(setUserIdCalls.length).toBe(1);
    expect(setUserIdCalls[0]).toEqual([userId]);

    const setDeviceIdCalls = mockAmp.setDeviceId.mock.calls;
    expect(setDeviceIdCalls.length).toBe(1);
    expect(setDeviceIdCalls[0]).toEqual(['device-1']);

    const logEventCalls = mockAmp.logEvent.mock.calls;
    expect(logEventCalls.length).toBe(1);
    expect(logEventCalls[0]).toEqual(['Event No Properties', undefined, undefined, undefined]);

    expect(consoleLogMock).toHaveBeenCalledTimes(0);
    expect(consoleErrorMock).toHaveBeenCalledTimes(0);
  });

  test('should track an event with properties of all types', () => {
    const mockAmp = { logEvent: jest.fn() };
    ampli.load({ client: { instance: mockAmp } });

    ampli.eventWithAllProperties({
      requiredBoolean: false,
      requiredInteger: 42,
      requiredEnum: "Enum1",
      requiredNumber: 42.0,
      requiredString: "Required string",
      requiredArray: ["Required","string"],
    });

    const logEventCalls = mockAmp.logEvent.mock.calls;
    expect(logEventCalls.length).toBe(1);
    expect(logEventCalls[0]).toEqual(['Event With All Properties', {
      requiredBoolean: false,
      requiredConst: "some-const-value",
      requiredInteger: 42,
      requiredEnum: "Enum1",
      requiredNumber: 42.0,
      requiredString: "Required string",
      requiredArray: ["Required","string"],
    }, undefined, undefined]);

    expect(consoleLogMock).toHaveBeenCalledTimes(0);
    expect(consoleErrorMock).toHaveBeenCalledTimes(0);
  });

  test('should identify and track', () => {
    const mockAmp = {
      setUserId: jest.fn(),
      identify: jest.fn(),
      logEvent: jest.fn(),
    };
    ampli.load({ client: { instance: mockAmp } });

    ampli.eventWithOptionalProperties({
      optionalString: "xyz",
    }, {
      user_id: userId,
      user_properties: {"prop-1": 123, "prop-2": "abc"},
    });

    const setUserIdCalls = mockAmp.setUserId.mock.calls;
    expect(setUserIdCalls.length).toBe(1);
    expect(setUserIdCalls[0]).toEqual([userId]);

    const identifyCalls = mockAmp.identify.mock.calls;
    expect(identifyCalls.length).toBe(1);
    const [ identify ] = identifyCalls[0];
    expect(identify.properties).toEqual(['prop-1', 'prop-2']);
    expect(identify.userPropertiesOperations).toEqual({
      '$set': {
        'prop-1': 123,
        'prop-2': 'abc'
      }
    });

    const logEventCalls = mockAmp.logEvent.mock.calls;
    expect(logEventCalls.length).toBe(1);
    expect(logEventCalls[0]).toEqual(['Event With Optional Properties', { 'optionalString': 'xyz' }, undefined, undefined]);

    expect(consoleLogMock).toHaveBeenCalledTimes(0);
    expect(consoleErrorMock).toHaveBeenCalledTimes(0);
  });
});
