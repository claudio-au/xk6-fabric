import fabric from "k6/x/fabric"
import {
    Writer,
    Reader,
    Connection,
    SchemaRegistry,
    SCHEMA_TYPE_BYTES,
    SCHEMA_TYPE_PROTOBUF
} from "k6/x/kafka"; // import kafka extension

export let options = { maxRedirects: 4 };

const brokers = ["localhost:9092"];
const topic = "xk6_kafka_byte_array_topic";

const writer = new Writer({
    brokers: brokers,
    topic: topic,
    autoCreateTopic: true,
});
const reader = new Reader({
    brokers: brokers,
    topic: topic,
});
const connection = new Connection({
    address: brokers[0],
});
const schemaRegistry = new SchemaRegistry();

if (__VU == 0) {
    connection.createTopic({ topic: topic });
}

const getTimeStamp = () => {
    return new Date().getTime()
}
const createOdometer = () => {
    let tags = [{
        "Name": "METRIC_UNITS",
        "Value": "km",
    }]
    return {
        "DoubleValue": 10.1,
        "MetricKind": 1,
        "Signal": "ODOMETER",
        "StartTime": getTimeStamp(),
        "Tags": tags,
    }
}

export default () => {

    let signals = [createOdometer(), createOdometer()]
    let data = {
        "AssetId": "aui:asset:vehicle:1234",
        "Source": "chat-telemetry-test",
        "EnvelopeTime": getTimeStamp(),
        "EnvelopeCreationTime": getTimeStamp(),
        "DataMinTime": getTimeStamp(),
        "DataMaxTime": getTimeStamp(),
        "Tags": {
            "Name": "TAG1",
            "Value": "Tag1"
        },
        "Data": signals,
        "DeviceAssetAui": "deviceAssetAui",
        "VechicleAssetAui": "vehicleAssetAui",
    }
    const message = fabric.createFabricEnvelope(data);
    const buffer = fabric.toByte(message)

    let messages = [
        {
            value: schemaRegistry.serialize({
                data: Array.from(buffer, (x) => x),
                schemaType: SCHEMA_TYPE_BYTES,
            }),
        }
    ];

    writer.produce({
        messages: messages,
    });
}

export function teardown(data) {
    //  if (__VU == 0) {
    //    // Delete the topic
    //    connection.deleteTopic(topic);
    //  }
    writer.close();
    reader.close();
    connection.close();
}