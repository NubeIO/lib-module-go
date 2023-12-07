package shared

import (
	"github.com/NubeIO/lib-module-go/http"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/model"
	"github.com/NubeIO/nubeio-rubix-lib-models-go/nargs"
)

func (g *GRPCMarshaller) Publish(topic string, qos model.QOS, retain bool, payload string) error {
	api := "/api/mqtt/publish"
	body := model.MqttBody{
		Topic:   topic,
		Qos:     qos,
		Retain:  retain,
		Payload: payload,
	}

	_, err := g.CallDBHelperWithParser(http.POST, api, nargs.Args{}, body)
	return err
}

func (g *GRPCMarshaller) PublishNonBuffer(topic string, qos model.QOS, retain bool, payload string) error {
	api := "/api/mqtt/publish-non-buffer"
	body := model.MqttBody{
		Topic:   topic,
		Qos:     qos,
		Retain:  retain,
		Payload: payload,
	}

	_, err := g.CallDBHelperWithParser(http.POST, api, nargs.Args{}, body)
	return err
}
