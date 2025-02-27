import { Ampli, ApiKey } from './ampli';
import * as amplitude from '@amplitude/analytics-browser';

describe('Ampli Browser TS SDK V2 tests', () => {
  let userId = 'test-browser-ts-ampli-v2-user-id';
  let consoleLogMock: jest.SpyInstance;
  let consoleErrorMock: jest.SpyInstance;

  beforeEach(() => {
    consoleLogMock = jest.spyOn(console, 'log').mockImplementation();
    consoleErrorMock = jest.spyOn(console, 'error').mockImplementation();
  });

  afterEach(() => {
    consoleLogMock.mockRestore();
    consoleErrorMock.mockRestore();
  });

  test('should log warning if load() without any arguments without ApiKeys for each environment', async () => {
    const ampli = new Ampli();
    await ampli.load().promise;

    expect(consoleLogMock).toHaveBeenCalledTimes(0);
    expect(consoleErrorMock).toHaveBeenCalledTimes(1);
    expect(consoleErrorMock.mock.calls).toEqual([
      [`ERROR: ampli.load() requires 'environment', 'client.apiKey', or 'client.instance'`],
    ]);
  });

  test('should identify()', async () => {
    const send = jest.fn().mockReturnValue({
      status: 'success',
      statusCode: 200,
      body: {
        eventsIngested: 1,
        payloadSizeBytes: 1,
        serverUploadTime: 1,
      },
    });
    const ampli = new Ampli();
    await ampli.load({
      client: {
        apiKey: 'test-api-key',
        configuration: {
          transportProvider: {
            send,
          },
        },
      },
    })?.promise;
    const response = (await ampli.identify(userId, {
      optionalArray: ['A', 'ray'],
      requiredNumber: 42,
    })?.promise) as amplitude.Types.Result;

    expect(response.event.event_type).toBe('$identify');
    expect(response.event.user_id).toBe(userId);
    expect(response.event.user_properties).toEqual({
      $set: {
        optionalArray: ['A', 'ray'],
        requiredNumber: 42,
      },
    });
    expect(consoleLogMock).toHaveBeenCalledTimes(0);
    expect(consoleErrorMock).toHaveBeenCalledTimes(0);
  });

  test('should setGroup()', async () => {
    const send = jest.fn().mockReturnValue({
      status: 'success',
      statusCode: 200,
      body: {
        eventsIngested: 1,
        payloadSizeBytes: 1,
        serverUploadTime: 1,
      },
    });
    const ampli = new Ampli();
    await ampli.load({
      client: {
        apiKey: 'test-api-key',
        configuration: {
          transportProvider: {
            send,
          },
        },
      },
    })?.promise;

    const response = (await ampli.setGroup('Group name', 'Group Value', { user_id: userId })?.promise) as amplitude.Types.Result;

    expect(response.event.event_type).toBe('$identify');
    expect(response.event.user_id).toBe(userId);
    expect(response.event.groups).toEqual({
      'Group name': 'Group Value',
    });
    expect(consoleLogMock).toHaveBeenCalledTimes(0);
    expect(consoleErrorMock).toHaveBeenCalledTimes(0);
  });

  test('should groupIdentify()', async () => {
    const send = jest.fn().mockReturnValue({
      status: 'success',
      statusCode: 200,
      body: {
        eventsIngested: 1,
        payloadSizeBytes: 1,
        serverUploadTime: 1,
      },
    });
    const ampli = new Ampli();
    await ampli.load({
      client: {
        apiKey: 'test-api-key',
        configuration: {
          transportProvider: {
            send,
          },
        },
      },
    })?.promise;

    const response = (await ampli.groupIdentify('Group name', 'Group Value', {
      requiredBoolean: true,
      optionalString: 'some-string',
    }, { user_id: userId })?.promise) as amplitude.Types.Result;

    expect(response.event.event_type).toBe('$groupidentify');
    expect(response.event.user_id).toBe(userId);
    expect(response.event.groups).toEqual({
      'Group name': 'Group Value',
    });
    expect(response.event.group_properties).toEqual({
      '$set': {
        requiredBoolean: true,
        optionalString: 'some-string',
      },
    });
    expect(consoleLogMock).toHaveBeenCalledTimes(0);
    expect(consoleErrorMock).toHaveBeenCalledTimes(0);
  });

  test('should track an event with no properties', async () => {
    const send = jest.fn().mockReturnValue({
      status: 'success',
      statusCode: 200,
      body: {
        eventsIngested: 1,
        payloadSizeBytes: 1,
        serverUploadTime: 1,
      },
    });
    const ampli = new Ampli();
    await ampli.load({
      client: {
        apiKey: 'test-api-key',
        configuration: {
          transportProvider: {
            send,
          },
        },
      },
    } as any)?.promise;

    const response: amplitude.Types.Result = (await ampli.eventNoProperties()?.promise) as amplitude.Types.Result;

    expect(response.event.event_type).toBe('Event No Properties');
    expect(response.event.event_properties).toBe(undefined);

    expect(consoleLogMock).toHaveBeenCalledTimes(0);
    expect(consoleErrorMock).toHaveBeenCalledTimes(0);
  });

  test('should track an event with properties of all types', async () => {
    const send = jest.fn().mockReturnValue({
      status: 'success',
      statusCode: 200,
      body: {
        eventsIngested: 1,
        payloadSizeBytes: 1,
        serverUploadTime: 1,
      },
    });
    const ampli = new Ampli();
    await ampli.load({
      client: {
        apiKey: 'test-api-key',
        configuration: {
          transportProvider: {
            send,
          },
        },
      },
    } as any)?.promise;

    const response = await ampli.eventWithAllProperties({
      requiredBoolean: false,
      requiredInteger: 42,
      requiredEnum: 'Enum1',
      requiredNumber: 42.0,
      requiredString: 'Required string',
      requiredArray: ['Required', 'string'],
    })?.promise;

    expect((response as amplitude.Types.Result).event.event_type).toBe('Event With All Properties');
    expect((response as amplitude.Types.Result).event.event_properties).toEqual({
      requiredBoolean: false,
      requiredConst: 'some-const-value',
      requiredInteger: 42,
      requiredEnum: 'Enum1',
      requiredNumber: 42.0,
      requiredString: 'Required string',
      requiredArray: ['Required', 'string'],
    });

    expect(consoleLogMock).toHaveBeenCalledTimes(0);
    expect(consoleErrorMock).toHaveBeenCalledTimes(0);
  });
});
