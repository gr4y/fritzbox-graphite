package lib

import (
	"github.com/huin/goupnp"
	"github.com/huin/goupnp/dcps/internetgateway2"
	"github.com/huin/goupnp/soap"
	"net/url"
)

// Fritz!Box Routers have some extra methods, so I have wrapped WANCommonInterfaceConfig1 into
// an own struct and reimplemented parts of internetgateway2.
type WANCommonInterfaceConfig1 struct {
	internetgateway2.WANCommonInterfaceConfig1
}

// NewWANCommonInterfaceConfig1Clients discovers instances of the service on the network,
// and returns clients to any that are found. errors will contain an error for
// any devices that replied but which could not be queried, and err will be set
// if the discovery process failed outright.
//
// This is a typical entry calling point into this package.
func NewWANCommonInterfaceConfig1Clients() (clients []*WANCommonInterfaceConfig1, errors []error, err error) {
	var genericClients []goupnp.ServiceClient
	if genericClients, errors, err = goupnp.NewServiceClients(internetgateway2.URN_WANCommonInterfaceConfig_1); err != nil {
		return
	}
	clients = newWANCommonInterfaceConfig1ClientsFromGenericClients(genericClients)
	return
}

// NewWANCommonInterfaceConfig1ClientsByURL discovers instances of the service at the given
// URL, and returns clients to any that are found. An error is returned if
// there was an error probing the service.
//
// This is a typical entry calling point into this package when reusing an
// previously discovered service URL.
func NewWANCommonInterfaceConfig1ClientsByURL(loc *url.URL) ([]*WANCommonInterfaceConfig1, error) {
	genericClients, err := goupnp.NewServiceClientsByURL(loc, internetgateway2.URN_WANCommonInterfaceConfig_1)
	if err != nil {
		return nil, err
	}
	return newWANCommonInterfaceConfig1ClientsFromGenericClients(genericClients), nil
}

// NewWANCommonInterfaceConfig1ClientsFromRootDevice discovers instances of the service in
// a given root device, and returns clients to any that are found. An error is
// returned if there was not at least one instance of the service within the
// device. The location parameter is simply assigned to the Location attribute
// of the wrapped ServiceClient(s).
//
// This is a typical entry calling point into this package when reusing an
// previously discovered root device.
func NewWANCommonInterfaceConfig1ClientsFromRootDevice(rootDevice *goupnp.RootDevice, loc *url.URL) ([]*WANCommonInterfaceConfig1, error) {
	genericClients, err := goupnp.NewServiceClientsFromRootDevice(rootDevice, loc, internetgateway2.URN_WANCommonInterfaceConfig_1)
	if err != nil {
		return nil, err
	}
	return newWANCommonInterfaceConfig1ClientsFromGenericClients(genericClients), nil
}

func newWANCommonInterfaceConfig1ClientsFromGenericClients(genericClients []goupnp.ServiceClient) []*WANCommonInterfaceConfig1 {
	clients := make([]*WANCommonInterfaceConfig1, len(genericClients))
	for i := range genericClients {
		clients[i] = &WANCommonInterfaceConfig1{internetgateway2.WANCommonInterfaceConfig1{genericClients[i]}}
	}
	return clients
}

func (client *WANCommonInterfaceConfig1) GetAddonInfos() (NewByteSendRate uint32, NewByteReceiveRate uint32,
	NewPacketSendRate uint32, NewPacketReceiveRate uint32, NewTotalBytesSent uint32, NewTotalBytesReceived uint32,
	NewAutoDisconnectTime uint32, NewIdleDisconnectTime uint32, NewDNSServer1 string, NewDNSServer2 string,
	NewVoipDNSServer1 string, NewVoipDNSServer2 string, NewUpnpControlEnabled bool, NewRoutedBridgedModeBoth bool, err error) {
	request := interface{}(nil)
	response := &struct {
		NewByteSendRate          string
		NewByteReceiveRate       string
		NewPacketSendRate        string
		NewPacketReceiveRate     string
		NewTotalBytesSent        string
		NewTotalBytesReceived    string
		NewAutoDisconnectTime    string
		NewIdleDisconnectTime    string
		NewDNSServer1            string
		NewDNSServer2            string
		NewVoipDNSServer1        string
		NewVoipDNSServer2        string
		NewUpnpControlEnabled    string
		NewRoutedBridgedModeBoth string
	}{}

	// Perform the SOAP call.
	if err = client.SOAPClient.PerformAction(internetgateway2.URN_WANCommonInterfaceConfig_1, "GetAddonInfos", request, response); err != nil {
		return
	}

	// BEGIN Unmarshal arguments from response.
	if NewAutoDisconnectTime, err = soap.UnmarshalUi4(response.NewAutoDisconnectTime); err != nil {
		return
	}
	if NewIdleDisconnectTime, err = soap.UnmarshalUi4(response.NewIdleDisconnectTime); err != nil {
		return
	}
	if NewByteSendRate, err = soap.UnmarshalUi4(response.NewByteSendRate); err != nil {
		return
	}
	if NewByteReceiveRate, err = soap.UnmarshalUi4(response.NewByteReceiveRate); err != nil {
		return
	}
	if NewPacketSendRate, err = soap.UnmarshalUi4(response.NewPacketSendRate); err != nil {
		return
	}
	if NewPacketReceiveRate, err = soap.UnmarshalUi4(response.NewPacketReceiveRate); err != nil {
		return
	}
	if NewTotalBytesSent, err = soap.UnmarshalUi4(response.NewTotalBytesSent); err != nil {
		return
	}
	if NewTotalBytesReceived, err = soap.UnmarshalUi4(response.NewTotalBytesReceived); err != nil {
		return
	}
	if NewDNSServer1, err = soap.UnmarshalString(response.NewDNSServer1); err != nil {
		return
	}
	if NewDNSServer2, err = soap.UnmarshalString(response.NewDNSServer2); err != nil {
		return
	}
	if NewVoipDNSServer1, err = soap.UnmarshalString(response.NewVoipDNSServer1); err != nil {
		return
	}
	if NewVoipDNSServer2, err = soap.UnmarshalString(response.NewVoipDNSServer2); err != nil {
		return
	}
	if NewUpnpControlEnabled, err = soap.UnmarshalBoolean(response.NewUpnpControlEnabled); err != nil {
		return
	}
	if NewRoutedBridgedModeBoth, err = soap.UnmarshalBoolean(response.NewRoutedBridgedModeBoth); err != nil {
		return
	}

	// END Unmarshal arguments from response.
	return
}
