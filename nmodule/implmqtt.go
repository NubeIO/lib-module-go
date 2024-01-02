package nmodule

import (
	"github.com/NubeIO/lib-module-go/nhttp"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/datatype"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/dto"
)

func (g *GRPCMarshaller) Publish(topic string, qos datatype.QOS, retain bool, payload string, opts ...*Opts) error {
	api := "/api/mqtt/publish"
	body := dto.MqttBody{
		Topic:   topic,
		Qos:     qos,
		Retain:  retain,
		Payload: payload,
	}

	_, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
	return err
}

func (g *GRPCMarshaller) PublishNonBuffer(topic string, qos datatype.QOS, retain bool, payload string, opts ...*Opts) error {
	api := "/api/mqtt/publish-non-buffer"
	body := dto.MqttBody{
		Topic:   topic,
		Qos:     qos,
		Retain:  retain,
		Payload: payload,
	}

	_, err := g.CallDBHelperWithParser(nhttp.POST, api, body, opts...)
	return err
}
