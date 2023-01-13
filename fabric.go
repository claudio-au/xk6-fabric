package fabric

import (
	"log"
	"reflect"
	pb "xk6-fabric/proto"

	"time"
	telemetrypb "xk6-fabric/proto/autonomic/ext/telemetry"
	"xk6-fabric/proto/autonomic/ext/telemetry/enumerations"

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

func (f *FabricEnvelope) GetWellKnownSignalMap() map[string]string {

	wksSignals := make(map[string]string)
	wksSignals["ODOMETER"] = "ODOMETER"
	wksSignals["BATTERY_VOLTAGE"] = "BATTERY_VOLTAGE"
	wksSignals["FUEL_LEVEL"] = "FUEL_LEVEL"
	wksSignals["IGNITION_STATUS"] = "IGNITION_STATUS"
	wksSignals["OIL_LIFE_REMAINING"] = "OIL_LIFE_REMAINING"
	wksSignals["OIL_LIFE"] = "OIL_LIFE"
	wksSignals["POSITION"] = "POSITION"
	wksSignals["SEAT_BELT_STATUS"] = "SEAT_BELT_STATUS"
	wksSignals["TIRE_PRESSURE"] = "TIRE_PRESSURE"
	wksSignals["TOTAL_ENGINE_TIME"] = "TOTAL_ENGINE_TIME"
	wksSignals["XEV_BATTERY_RANGE"] = "XEV_BATTERY_RANGE"
	wksSignals["XEV_BATTERY_STATE_OF_CHARGE"] = "XEV_BATTERY_STATE_OF_CHARGE"
	wksSignals["XEV_PLUG_CHARGER_STATUS"] = "XEV_PLUG_CHARGER_STATUS"
	return wksSignals
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

	wksSignals := make(map[string]telemetrypb.WellKnownSignal)
	wksSignals["ODOMETER"] = telemetrypb.WellKnownSignal_ODOMETER
	wksSignals["BATTERY_VOLTAGE"] = telemetrypb.WellKnownSignal_BATTERY_VOLTAGE
	wksSignals["FUEL_LEVEL"] = telemetrypb.WellKnownSignal_FUEL_LEVEL
	wksSignals["OIL_LIFE_REMAINING"] = telemetrypb.WellKnownSignal_OIL_LIFE_REMAINING
	wksSignals["OIL_LIFE"] = telemetrypb.WellKnownSignal_OIL_LIFE_REMAINING
	wksSignals["POSITION"] = telemetrypb.WellKnownSignal_POSITION
	wksSignals["TIRE_PRESSURE"] = telemetrypb.WellKnownSignal_TIRE_PRESSURE
	wksSignals["TOTAL_ENGINE_TIME"] = telemetrypb.WellKnownSignal_TOTAL_ENGINE_TIME
	wksSignals["XEV_BATTERY_RANGE"] = telemetrypb.WellKnownSignal_XEV_BATTERY_RANGE
	wksSignals["XEV_BATTERY_STATE_OF_CHARGE"] = telemetrypb.WellKnownSignal_XEV_BATTERY_STATE_OF_CHARGE

	wksSignals["SEAT_BELT_STATUS"] = telemetrypb.WellKnownSignal_SEAT_BELT_STATUS
	wksSignals["IGNITION_STATUS"] = telemetrypb.WellKnownSignal_IGNITION_STATUS
	wksSignals["XEV_PLUG_CHARGER_STATUS"] = telemetrypb.WellKnownSignal_XEV_PLUG_CHARGER_STATUS

	wks, exist := wksSignals[signal]
	if exist {
		return wks
	} else {
		return 0
	}

}

func getEnumeration(en string) *telemetrypb.Enumeration {
	enumValue := make(map[string]*telemetrypb.Enumeration)

	seatBelt := &telemetrypb.Enumeration_SeatbeltStatus{
		SeatbeltStatus: enumerations.SeatbeltStatus_UNBUCKLED,
	}
	enumValue["SEAT_BELT_STATUS_UNBUCKLED"] = &telemetrypb.Enumeration{
		Value: seatBelt,
	}

	seatBelt = &telemetrypb.Enumeration_SeatbeltStatus{
		SeatbeltStatus: enumerations.SeatbeltStatus_BUCKLED,
	}
	enumValue["SEAT_BELT_STATUS_BUCKLED"] = &telemetrypb.Enumeration{
		Value: seatBelt,
	}

	ignition := &telemetrypb.Enumeration_IgnitionStatus{
		IgnitionStatus: enumerations.IgnitionStatus_ON,
	}
	enumValue["IGNITION_ON"] = &telemetrypb.Enumeration{
		Value: ignition,
	}
	ignition = &telemetrypb.Enumeration_IgnitionStatus{
		IgnitionStatus: enumerations.IgnitionStatus_OFF,
	}
	enumValue["IGNITION_OFF"] = &telemetrypb.Enumeration{
		Value: ignition,
	}

	xevPlugCharger := &telemetrypb.Enumeration_XevPlugChargerStatus{
		XevPlugChargerStatus: enumerations.XevPlugChargerStatus_CONNECTED,
	}
	enumValue["XEV_PLUG_CHARGER_CONNECTED"] = &telemetrypb.Enumeration{
		Value: xevPlugCharger,
	}

	xevPlugCharger = &telemetrypb.Enumeration_XevPlugChargerStatus{
		XevPlugChargerStatus: enumerations.XevPlugChargerStatus_DISCONNECTED,
	}
	enumValue["XEV_PLUG_CHARGER_DISCONNECTED"] = &telemetrypb.Enumeration{
		Value: xevPlugCharger,
	}

	e, exist := enumValue[en]
	if exist {
		return e
	} else {
		return &telemetrypb.Enumeration{}
	}
}

func getPositionValue(v map[string]interface{}) *telemetrypb.Metric_PositionValue {

	value := &telemetrypb.Metric_PositionValue{}
	listLocation := make([]*telemetrypb.Location, 0)
	points := v["Location"].([]interface{})

	for _, p := range points {
		pAux := p.(map[string]interface{})
		point := &telemetrypb.Point{
			Latitude:  toFloat64(pAux["Latitude"]),
			Longitude: toFloat64(pAux["Longitude"]),
		}
		location := &telemetrypb.Location{
			Location: &telemetrypb.Location_Point{
				Point: point,
			},
		}
		listLocation = append(listLocation, location)
	}

	value.PositionValue = &telemetrypb.Position{
		Location: listLocation,
	}
	return value
}

func toFloat64(value interface{}) float64 {

	if value != nil && reflect.TypeOf(value).String() == "int64" {
		return float64(value.(int64))
	}
	return -1
}

func getDoubleValue(v interface{}) *telemetrypb.Metric_DoubleValue {
	double_value := &telemetrypb.Metric_DoubleValue{}
	if v != nil && reflect.TypeOf(v).String() == "int64" {
		v = float64(v.(int64))
	}
	double_value.DoubleValue = v.(float64)
	return double_value
}

func getEnumValue(v interface{}) *telemetrypb.Metric_EnumValue {
	value := &telemetrypb.Metric_EnumValue{}
	value.EnumValue = getEnumeration(v.(string))

	return value
}

func setValue(signal map[string]interface{}) *telemetrypb.Metric {
	metric := &telemetrypb.Metric{}
	v, exist := signal["DoubleValue"]
	if exist {
		metric.Value = getDoubleValue(v)
		return metric
	}
	v, exist = signal["EnumValue"]
	if exist {
		metric.Value = getEnumValue(v)
		return metric
	}
	v, exist = signal["PositionValue"]
	if exist {
		metric.Value = getPositionValue(v.(map[string]interface{}))
		return metric
	}
	return metric
}

func createMetrics(metrics []interface{}) []*pb.DataElementProto {
	elements := make([]*pb.DataElementProto, 0)

	for _, element := range metrics {
		signal := element.(map[string]interface{})

		metric := setValue(signal)

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
