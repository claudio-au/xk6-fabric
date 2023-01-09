package fabric

import (
	"fmt"
	"log"
	"reflect"
	pb "xk6-fabric/proto"

	"time"
	telemetrypb "xk6-fabric/proto/autonomic/ext/telemetry"

	"go.k6.io/k6/js/modules"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func init() {
	modules.Register("k6/x/fabric", new(FabricEnvelope))
}

type FabricEnvelope struct{}

func (c *FabricEnvelope) IsGreater(a, b int) bool {
	if a > b {
		return true
	} else {
		return false
	}
}

func (f *FabricEnvelope) CreateFabricEnvelope(data map[string]interface{}) *pb.FabricEnvelopeProto {
	fabric, error := buildFabricEnvelope(data)
	fmt.Println("Creating Fabric envelope")
	if error != nil {
		log.Fatal(error)
	}
	return fabric
}

func (f *FabricEnvelope) ToByte(data *pb.FabricEnvelopeProto) []byte {
	response, err := proto.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	return response
}

func convertToTimestamp(timestampValue int64) *timestamppb.Timestamp {
	t := time.UnixMilli(timestampValue)
	return timestamppb.New(t)
}

func isNotNull(data interface{}) bool {
	return data != nil
}

func buildFabricEnvelope(data map[string]interface{}) (*pb.FabricEnvelopeProto, error) {

	fabric := &pb.FabricEnvelopeProto{}
	fabric.AssetId = data["AssetId"].(string)
	fabric.Source = data["Source"].(string)

	if isNotNull(data["EnvelopeTime"]) {
		fabric.EnvelopeTime = convertToTimestamp(data["EnvelopeTime"].(int64))
	}
	if isNotNull(data["EnvelopeCreationTime"]) {
		fabric.EnvelopeCreationTime = convertToTimestamp(data["EnvelopeCreationTime"].(int64))
	}
	if isNotNull(data["DataMinTime"]) {
		fabric.DataMinTime = convertToTimestamp(data["DataMinTime"].(int64))
	}
	if isNotNull(data["DataMaxTime"]) {
		fabric.DataMaxTime = convertToTimestamp(data["DataMaxTime"].(int64))
	}

	if data["Tags"] != nil {
		tags := map[string]string{}

		for k, v := range data["Tags"].(map[string]interface{}) {
			tags[k] = v.(string)
		}
		fabric.Tags = tags
	}
	fabric.Data = createMetrics(data["Data"].([]interface{}))
	return fabric, nil
}
func getWksSignal(signal string) telemetrypb.WellKnownSignal {
	switch signal {
	case "ODOMETER":
		return telemetrypb.WellKnownSignal_ODOMETER
	default:
		return 0
	}

}

func createMetrics(metrics []interface{}) []*pb.DataElementProto {
	elements := make([]*pb.DataElementProto, 0)

	for _, element := range metrics {
		signal := element.(map[string]interface{})
		metric := &telemetrypb.Metric{}

		double_value := &telemetrypb.Metric_DoubleValue{}
		v := signal["DoubleValue"]
		if v != nil && reflect.TypeOf(v).String() == "int64" {
			v = float64(v.(int64))
		}
		double_value.DoubleValue = v.(float64)
		metric.Value = double_value

		if signal["MetricKind"].(int64) == 1 {
			metric.MetricKind = telemetrypb.Metric_GAUGE
		}
		s := &telemetrypb.Signal{}
		s.Signal = &telemetrypb.Signal_WksSignal{
			WksSignal: getWksSignal(signal["Signal"].(string)),
		}

		metric.Signal = s

		any := &anypb.Any{}
		any.TypeUrl = "type.googleapis.com/autonomic.ext.telemetry.Metric"

		buffer, e := proto.Marshal(metric)
		if e != nil {
			log.Fatal(e)
		}
		any.Value = buffer

		d := &pb.DataElementProto{}
		d.Element = any

		elements = append(elements, d)
	}

	return elements
}
