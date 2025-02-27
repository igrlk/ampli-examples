//
// Ampli - A strong typed wrapper for your Analytics
//
// This file is generated by Amplitude.
// To update run 'ampli pull jre-java-ampli'
//
// Required dependencies: com.amplitude:java-sdk:1.8.0, org.json:json:20201115
// Tracking Plan Version: 0
// Build: 1.0.0
// Runtime: jre:java-ampli
//
// [View Tracking Plan](https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest)
//
// [Full Setup Instructions](https://data.amplitude.com/test-codegen/Test%20Codegen/implementation/jre-java-ampli)
//
package com.amplitude.ampli;

import java.lang.reflect.Method;
import java.util.HashMap;
import java.util.Map;
import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import com.amplitude.Amplitude;
import com.amplitude.MiddlewareExtra;
import com.amplitude.Plan;

public class Ampli {
    private static volatile Ampli singleton = null;

    public static Ampli getInstance() {
        if (singleton == null) {
            createSingleton();
        }
        return singleton;
    }

    private synchronized static void createSingleton() {
        if (singleton == null) {
            singleton = new Ampli();
        }
    }

    public enum Environment {
        DEVELOPMENT,
        PRODUCTION
    }

    public static final Map<Environment, String> API_KEY = new HashMap<>();

    static {
        API_KEY.put(Environment.DEVELOPMENT, "");
        API_KEY.put(Environment.PRODUCTION, "");
    }

    private boolean disabled = false;

    public boolean isLoaded() {
        return this.client != null;
    }

    private Amplitude client;

    public Amplitude getClient() {
        this.isInitializedAndEnabled();
        return this.client;
    }

    private final Plan observePlan = new Plan()
        .setBranch("main")
        .setSource("jre-java-ampli")
        .setVersion("0")
        .setVersionId("79154a50-f057-4db5-9755-775e4e9f05e6");

    public void load() {
        this.load(null);
    }

    public void load(LoadOptions options) {
        Boolean disabled = options != null ? options.getDisabled() : null;
        this.disabled = disabled != null ? disabled : false;

        if (this.isLoaded()) {
            System.err.println("Warning: Ampli is already initialized. Ampli.getInstance().load() should be called once at application start up.");
            return;
        }

        Environment env = options != null ? options.getEnvironment() : null;
        if (env == null) {
            env = Environment.DEVELOPMENT;
        }

        LoadClientOptions clientOptions = options != null ? options.getClient() : null;

        String apiKey = Ampli.API_KEY.get(env);
        Amplitude client = null;
        if (clientOptions != null) {
            String optionsApiKey = clientOptions.getApiKey();
            if (optionsApiKey != null) {
                apiKey = optionsApiKey;
            }

            client = clientOptions.getInstance();
        }

        if (client != null) {
            this.client = client;
        } else if (apiKey != null && !apiKey.equals("")) {
            this.client = Amplitude.getInstance();
            this.client.init(apiKey);
        } else {
            System.err.println("Ampli.getInstance().load() requires 'environment', 'client.apiKey', or 'client.instance'");
            return;
        }

        Plan plan = observePlan;
        if (clientOptions != null) {
            Plan optionsPlan = clientOptions.getPlan();
            if (optionsPlan != null) {
                plan = optionsPlan;
            }
        }
        this.client.setPlan(plan);

        // set IngestionMetadata with backwards compatibility, min Java SDK version 1.10.1.
        try {
            Class<?> clazz = Class.forName("com.amplitude.IngestionMetadata");
            Method setSourceNameMethod = clazz.getMethod("setSourceName", String.class);
            Method setSourceVersionMethod = clazz.getMethod("setSourceVersion", String.class);
            Object ingestionMetadata = clazz.newInstance();
            setSourceNameMethod.invoke(ingestionMetadata, "jre-java-ampli");
            setSourceVersionMethod.invoke(ingestionMetadata, "1.0.0");
            Method setIngestionMetadata = Amplitude.class.getMethod("setIngestionMetadata", clazz);
            setIngestionMetadata.invoke(this.client, ingestionMetadata);
        } catch (ClassNotFoundException | NoSuchMethodException | SecurityException e) {
            System.out.println("com.amplitude.IngestionMetadata is available starting from Java SDK 1.10.1 version");
        } catch (Exception e) {
            System.err.println("Unexpected error when setting IngestionMetadata");
        }
    }

    public void track(String userId, Event event) {
        this.track(userId, event, null, null);
    }

    public void track(String userId, Event event, EventOptions options) {
        this.track(userId, event, options, null);
    }

    public void track(String userId, Event event, MiddlewareExtra extra) {
        this.track(userId, event, null, extra);
    }

    public void track(String userId, Event event, EventOptions options, MiddlewareExtra extra) {
        if (!this.isInitializedAndEnabled()) {
            return;
        }
        com.amplitude.Event amplitudeEvent = this.createAmplitudeEvent(event.eventType, options, userId);
        amplitudeEvent.eventProperties = this.getEventPropertiesJson(event);

        this.client.logEvent(amplitudeEvent, extra);
    }

    public void identify(String userId, Identify event) {
        this.identify(userId, event, null, null);
    }

    public void identify(String userId, Identify event, EventOptions options) {
        this.identify(userId, event, options, null);
    }

    public void identify(String userId, Identify event, MiddlewareExtra extra) {
        this.identify(userId, event, null, extra);
    }

    public void identify(String userId, Identify event, EventOptions options, MiddlewareExtra extra) {
        if (!this.isInitializedAndEnabled()) {
            return;
        }
        com.amplitude.Event amplitudeEvent = this.createAmplitudeEvent(event.eventType, options, userId);
        amplitudeEvent.userProperties = this.getEventPropertiesJson(event);

        this.client.logEvent(amplitudeEvent, extra);
    }

    public void setGroup(String userId, String name, String value) {
        this.setGroup(userId, name, value, null, null);
    }

    public void setGroup(String userId, String name, String value, EventOptions options) {
        this.setGroup(userId, name, value, options, null);
    }

    public void setGroup(String userId, String name, String value, MiddlewareExtra extra) {
        this.setGroup(userId, name, value, null, extra);
    }

    public void setGroup(String userId, String name, String value, EventOptions options, MiddlewareExtra extra) {
        this.setGroupProperties(userId, name, value, options, extra);
    }

    public void setGroup(String userId, String name, String[] value) {
        this.setGroup(userId, name, value, null, null);
    }

    public void setGroup(String userId, String name, String[] value, EventOptions options) {
        this.setGroup(userId, name, value, options, null);
    }

    public void setGroup(String userId, String name, String[] value, MiddlewareExtra extra) {
        this.setGroup(userId, name, value, null, extra);
    }

    public void setGroup(String userId, String name, String[] value, EventOptions options, MiddlewareExtra extra) {
        JSONArray jsonValue;
        try {
            jsonValue = new JSONArray(value);
        } catch (JSONException e) {
            System.err.printf("Error converting value to JSONArray: %s%n", e.getMessage());
            return;
        }
        this.setGroupProperties(userId, name, jsonValue, options, extra);
    }

    private void setGroupProperties(String userId, String name, Object value, EventOptions options, MiddlewareExtra extra) {
        if (!this.isInitializedAndEnabled()) {
            return;
        }

        JSONObject groupProperties = new JSONObject();
        try {
            groupProperties.put(name, value);
        } catch (JSONException e) {
            System.err.printf("Error converting value to JSONArray: %s%n", e.getMessage());
        }

        com.amplitude.Event amplitudeEvent = this.createAmplitudeEvent("$identify", options, userId);
        amplitudeEvent.groupProperties = groupProperties;

        this.client.logEvent(amplitudeEvent, extra);
    }

    public void flush() {
        if (!this.isInitializedAndEnabled()) {
            return;
        }
        this.client.flushEvents();
    }

    /**
     * EventMaxIntForTest
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/EventMaxIntForTest">View in Tracking Plan</a>
     * <p>
     * Event to test schema validation
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     */
    public void eventMaxIntForTest(String userId, EventMaxIntForTest event) {
        this.eventMaxIntForTest(userId, event, null, null);
    }

    /**
     * EventMaxIntForTest
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/EventMaxIntForTest">View in Tracking Plan</a>
     * <p>
     * Event to test schema validation
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     */
    public void eventMaxIntForTest(String userId, EventMaxIntForTest event, EventOptions options) {
        this.eventMaxIntForTest(userId, event, options, null);
    }

    /**
     * EventMaxIntForTest
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/EventMaxIntForTest">View in Tracking Plan</a>
     * <p>
     * Event to test schema validation
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventMaxIntForTest(String userId, EventMaxIntForTest event, MiddlewareExtra extra) {
        this.eventMaxIntForTest(userId, event, null, extra);
    }

    /**
     * EventMaxIntForTest
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/EventMaxIntForTest">View in Tracking Plan</a>
     * <p>
     * Event to test schema validation
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventMaxIntForTest(String userId, EventMaxIntForTest event, EventOptions options, MiddlewareExtra extra) {
        this.track(userId, event, options, extra);
    }

    /**
     * Event No Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20No%20Properties">View in Tracking Plan</a>
     * <p>
     * Event w no properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     */
    public void eventNoProperties(String userId) {
        this.eventNoProperties(userId, null, null);
    }

    /**
     * Event No Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20No%20Properties">View in Tracking Plan</a>
     * <p>
     * Event w no properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param options The event's options
     */
    public void eventNoProperties(String userId, EventOptions options) {
        this.eventNoProperties(userId, options, null);
    }

    /**
     * Event No Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20No%20Properties">View in Tracking Plan</a>
     * <p>
     * Event w no properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventNoProperties(String userId, MiddlewareExtra extra) {
        this.eventNoProperties(userId, null, extra);
    }

    /**
     * Event No Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20No%20Properties">View in Tracking Plan</a>
     * <p>
     * Event w no properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param options The event's options
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventNoProperties(String userId, EventOptions options, MiddlewareExtra extra) {
        this.track(userId, EventNoProperties.builder().build(), options, extra);
    }

    /**
     * Event Object Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20Object%20Types">View in Tracking Plan</a>
     * <p>
     * Event with Object and Object Array
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     */
    public void eventObjectTypes(String userId, EventObjectTypes event) {
        this.eventObjectTypes(userId, event, null, null);
    }

    /**
     * Event Object Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20Object%20Types">View in Tracking Plan</a>
     * <p>
     * Event with Object and Object Array
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     */
    public void eventObjectTypes(String userId, EventObjectTypes event, EventOptions options) {
        this.eventObjectTypes(userId, event, options, null);
    }

    /**
     * Event Object Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20Object%20Types">View in Tracking Plan</a>
     * <p>
     * Event with Object and Object Array
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventObjectTypes(String userId, EventObjectTypes event, MiddlewareExtra extra) {
        this.eventObjectTypes(userId, event, null, extra);
    }

    /**
     * Event Object Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20Object%20Types">View in Tracking Plan</a>
     * <p>
     * Event with Object and Object Array
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventObjectTypes(String userId, EventObjectTypes event, EventOptions options, MiddlewareExtra extra) {
        this.track(userId, event, options, extra);
    }

    /**
     * Event With All Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20All%20Properties">View in Tracking Plan</a>
     * <p>
     * Event w all properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     */
    public void eventWithAllProperties(String userId, EventWithAllProperties event) {
        this.eventWithAllProperties(userId, event, null, null);
    }

    /**
     * Event With All Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20All%20Properties">View in Tracking Plan</a>
     * <p>
     * Event w all properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     */
    public void eventWithAllProperties(String userId, EventWithAllProperties event, EventOptions options) {
        this.eventWithAllProperties(userId, event, options, null);
    }

    /**
     * Event With All Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20All%20Properties">View in Tracking Plan</a>
     * <p>
     * Event w all properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithAllProperties(String userId, EventWithAllProperties event, MiddlewareExtra extra) {
        this.eventWithAllProperties(userId, event, null, extra);
    }

    /**
     * Event With All Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20All%20Properties">View in Tracking Plan</a>
     * <p>
     * Event w all properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithAllProperties(String userId, EventWithAllProperties event, EventOptions options, MiddlewareExtra extra) {
        this.track(userId, event, options, extra);
    }

    /**
     * Event With Array Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Array%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with Array Types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     */
    public void eventWithArrayTypes(String userId, EventWithArrayTypes event) {
        this.eventWithArrayTypes(userId, event, null, null);
    }

    /**
     * Event With Array Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Array%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with Array Types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     */
    public void eventWithArrayTypes(String userId, EventWithArrayTypes event, EventOptions options) {
        this.eventWithArrayTypes(userId, event, options, null);
    }

    /**
     * Event With Array Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Array%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with Array Types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithArrayTypes(String userId, EventWithArrayTypes event, MiddlewareExtra extra) {
        this.eventWithArrayTypes(userId, event, null, extra);
    }

    /**
     * Event With Array Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Array%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with Array Types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithArrayTypes(String userId, EventWithArrayTypes event, EventOptions options, MiddlewareExtra extra) {
        this.track(userId, event, options, extra);
    }

    /**
     * Event With Const Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Const%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with const types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     */
    public void eventWithConstTypes(String userId) {
        this.eventWithConstTypes(userId, null, null);
    }

    /**
     * Event With Const Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Const%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with const types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param options The event's options
     */
    public void eventWithConstTypes(String userId, EventOptions options) {
        this.eventWithConstTypes(userId, options, null);
    }

    /**
     * Event With Const Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Const%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with const types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithConstTypes(String userId, MiddlewareExtra extra) {
        this.eventWithConstTypes(userId, null, extra);
    }

    /**
     * Event With Const Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Const%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with const types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param options The event's options
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithConstTypes(String userId, EventOptions options, MiddlewareExtra extra) {
        this.track(userId, EventWithConstTypes.builder().build(), options, extra);
    }

    /**
     * event withDifferent_CasingTypes
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/event%20withDifferent_CasingTypes">View in Tracking Plan</a>
     * <p>
     * Description for case with space
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     */
    public void eventWithDifferentCasingTypes(String userId, EventWithDifferentCasingTypes event) {
        this.eventWithDifferentCasingTypes(userId, event, null, null);
    }

    /**
     * event withDifferent_CasingTypes
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/event%20withDifferent_CasingTypes">View in Tracking Plan</a>
     * <p>
     * Description for case with space
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     */
    public void eventWithDifferentCasingTypes(String userId, EventWithDifferentCasingTypes event, EventOptions options) {
        this.eventWithDifferentCasingTypes(userId, event, options, null);
    }

    /**
     * event withDifferent_CasingTypes
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/event%20withDifferent_CasingTypes">View in Tracking Plan</a>
     * <p>
     * Description for case with space
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithDifferentCasingTypes(String userId, EventWithDifferentCasingTypes event, MiddlewareExtra extra) {
        this.eventWithDifferentCasingTypes(userId, event, null, extra);
    }

    /**
     * event withDifferent_CasingTypes
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/event%20withDifferent_CasingTypes">View in Tracking Plan</a>
     * <p>
     * Description for case with space
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithDifferentCasingTypes(String userId, EventWithDifferentCasingTypes event, EventOptions options, MiddlewareExtra extra) {
        this.track(userId, event, options, extra);
    }

    /**
     * Event With Enum Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Enum%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with enum types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     */
    public void eventWithEnumTypes(String userId, EventWithEnumTypes event) {
        this.eventWithEnumTypes(userId, event, null, null);
    }

    /**
     * Event With Enum Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Enum%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with enum types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     */
    public void eventWithEnumTypes(String userId, EventWithEnumTypes event, EventOptions options) {
        this.eventWithEnumTypes(userId, event, options, null);
    }

    /**
     * Event With Enum Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Enum%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with enum types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithEnumTypes(String userId, EventWithEnumTypes event, MiddlewareExtra extra) {
        this.eventWithEnumTypes(userId, event, null, extra);
    }

    /**
     * Event With Enum Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Enum%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with enum types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithEnumTypes(String userId, EventWithEnumTypes event, EventOptions options, MiddlewareExtra extra) {
        this.track(userId, event, options, extra);
    }

    /**
     * Event With Optional Array Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Optional%20Array%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with optional array types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     */
    public void eventWithOptionalArrayTypes(String userId, EventWithOptionalArrayTypes event) {
        this.eventWithOptionalArrayTypes(userId, event, null, null);
    }

    /**
     * Event With Optional Array Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Optional%20Array%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with optional array types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     */
    public void eventWithOptionalArrayTypes(String userId, EventWithOptionalArrayTypes event, EventOptions options) {
        this.eventWithOptionalArrayTypes(userId, event, options, null);
    }

    /**
     * Event With Optional Array Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Optional%20Array%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with optional array types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithOptionalArrayTypes(String userId, EventWithOptionalArrayTypes event, MiddlewareExtra extra) {
        this.eventWithOptionalArrayTypes(userId, event, null, extra);
    }

    /**
     * Event With Optional Array Types
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Optional%20Array%20Types">View in Tracking Plan</a>
     * <p>
     * Description for event with optional array types
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithOptionalArrayTypes(String userId, EventWithOptionalArrayTypes event, EventOptions options, MiddlewareExtra extra) {
        this.track(userId, event, options, extra);
    }

    /**
     * Event With Optional Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Optional%20Properties">View in Tracking Plan</a>
     * <p>
     * Event w optional properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     */
    public void eventWithOptionalProperties(String userId, EventWithOptionalProperties event) {
        this.eventWithOptionalProperties(userId, event, null, null);
    }

    /**
     * Event With Optional Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Optional%20Properties">View in Tracking Plan</a>
     * <p>
     * Event w optional properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     */
    public void eventWithOptionalProperties(String userId, EventWithOptionalProperties event, EventOptions options) {
        this.eventWithOptionalProperties(userId, event, options, null);
    }

    /**
     * Event With Optional Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Optional%20Properties">View in Tracking Plan</a>
     * <p>
     * Event w optional properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithOptionalProperties(String userId, EventWithOptionalProperties event, MiddlewareExtra extra) {
        this.eventWithOptionalProperties(userId, event, null, extra);
    }

    /**
     * Event With Optional Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Optional%20Properties">View in Tracking Plan</a>
     * <p>
     * Event w optional properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithOptionalProperties(String userId, EventWithOptionalProperties event, EventOptions options, MiddlewareExtra extra) {
        this.track(userId, event, options, extra);
    }

    /**
     * Event With Template Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Template%20Properties">View in Tracking Plan</a>
     * <p>
     * Event with template properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     */
    public void eventWithTemplateProperties(String userId, EventWithTemplateProperties event) {
        this.eventWithTemplateProperties(userId, event, null, null);
    }

    /**
     * Event With Template Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Template%20Properties">View in Tracking Plan</a>
     * <p>
     * Event with template properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     */
    public void eventWithTemplateProperties(String userId, EventWithTemplateProperties event, EventOptions options) {
        this.eventWithTemplateProperties(userId, event, options, null);
    }

    /**
     * Event With Template Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Template%20Properties">View in Tracking Plan</a>
     * <p>
     * Event with template properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithTemplateProperties(String userId, EventWithTemplateProperties event, MiddlewareExtra extra) {
        this.eventWithTemplateProperties(userId, event, null, extra);
    }

    /**
     * Event With Template Properties
     * <p>
     * <a href="https://data.amplitude.com/test-codegen/Test%20Codegen/events/main/latest/Event%20With%20Template%20Properties">View in Tracking Plan</a>
     * <p>
     * Event with template properties description
     * <p>
     * Owner: Test codegen
     *
     * @param userId The user's ID
     * @param event The event
     * @param options The event's options
     * @param extra Extra untyped parameters for use in middleware
     */
    public void eventWithTemplateProperties(String userId, EventWithTemplateProperties event, EventOptions options, MiddlewareExtra extra) {
        this.track(userId, event, options, extra);
    }

    private com.amplitude.Event createAmplitudeEvent(String eventType, EventOptions options, String userId) {
        return new com.amplitude.Event(
            eventType,
            userId != null ? userId : (options != null ? options.userId : null),
            options != null ? options.deviceId : null
        );
    }

    private boolean isInitializedAndEnabled() {
        if (!this.isLoaded()) {
            System.err.println("Ampli is not yet initialized. Have you called `Ampli.getInstance().load()` on app start?");
            return false;
        }
        return !this.disabled;
    }

    private JSONObject getEventPropertiesJson(Event event) {
        if (event == null || event.eventProperties == null) {
            return null;
        }

        JSONObject json = new JSONObject();

        for (Map.Entry<String, Object> entry : event.eventProperties.entrySet()) {
            String key = entry.getKey();
            Object value = entry.getValue();
            try {
                if (value != null) {
                    json.put(key, value.getClass().isArray() ? new JSONArray(value) : value);
                } else {
                    json.put(key, JSONObject.NULL);
                }
            } catch (JSONException e) {
                System.err.printf("Error converting properties to JSONObject: %s%n", e.getMessage());
            }
        }

        return json;
    }
}
