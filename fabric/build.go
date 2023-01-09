package fabric

import (
	pb "fabric/proto"
	"log"
	"time"
	telemetrypb "xk6-fabric/proto/autonomic/ext/telemetry"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func BuildFabricEnvelope(data map[string]interface{}) (*pb.FabricEnvelopeProto, error) {

	fabric := &pb.FabricEnvelopeProto{}
	fabric.AssetId = data["AssetId"].(string)
	fabric.Source = data["Source"].(string)
	fabric.EnvelopeTime = timestamppb.New(data["EnvelopeTime"].(time.Time))
	fabric.EnvelopeCreationTime = timestamppb.New(data["EnvelopeCreationTime"].(time.Time))
	fabric.DataMinTime = timestamppb.New(data["DataMinTime"].(time.Time))
	fabric.DataMaxTime = timestamppb.New(data["DataMaxTime"].(time.Time))
	fabric.Tags = data["Tags"].(map[string]string)
	fabric.Data = createMetrics(data["Data"].([]map[string]interface{}))
	return fabric, nil
}

func createMetrics(metrics []map[string]interface{}) []*pb.DataElementProto {
	elements := make([]*pb.DataElementProto, 0)

	for _, signal := range metrics {

		metric := &telemetrypb.Metric{}

		double_value := &telemetrypb.Metric_DoubleValue{}
		double_value.DoubleValue = signal["DoubleValue"].(float64)
		metric.Value = double_value

		if signal["MetricKind"].(int) == 1 {
			metric.MetricKind = telemetrypb.Metric_GAUGE
		}
		s := &telemetrypb.Signal{}
		s.Signal = &telemetrypb.Signal_StringSignal{
			StringSignal: signal["Signal"].(string),
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
