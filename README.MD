# Fabric Envelope K6 Extension

This extension for K6 (k6.io) helps to create Fabric envelopes objects as javascript/bytes.

It is useful to be sent through Kafka topics


## Building the extension

Install golang and xk6 extension (https://github.com/grafana/xk6/)

After everything configured run the command

```xk6 build --with xk6-fabric=. --with github.com/mostafa/xk6-kafka```

It will create a custom k6 bin with our extension and kafka extension

## Testing

To test the extension, run:

* A Kafka cluster
* Run the test.js script `./k6 run testjs`

Check the messages on Kafka


More information about extensions for K6 go to https://k6.io/docs/extensions/