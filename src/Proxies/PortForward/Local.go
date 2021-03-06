package PortForward

import (
	"fmt"
	"github.com/shoriwe/FullProxy/src/ConnectionStructures"
	"github.com/shoriwe/FullProxy/src/MasterSlave"
	"github.com/shoriwe/FullProxy/src/Proxies/Basic"
	"github.com/shoriwe/FullProxy/src/Sockets"
	"net"
)

func CreateLocalPortForwardSession(clientConnection net.Conn, clientReader ConnectionStructures.SocketReader, clientWriter ConnectionStructures.SocketWriter, args...interface{}){
	targetAddress := args[0].(*string)
	targetPort := args[1].(*string)
	targetConnection := Sockets.Connect(targetAddress, targetPort)
	if targetConnection != nil{
		targetReader, targetWriter := ConnectionStructures.CreateSocketConnectionReaderWriter(targetConnection)
		Basic.Proxy(clientConnection, targetConnection, clientReader, clientWriter, targetReader, targetWriter)
	} else {
		_ = clientConnection.Close()
	}
}


func StartLocalPortForward(targetAddress *string, targetPort *string, masterAddress *string, masterPort *string){
	if !(*targetAddress == "" || *targetPort == "" || *masterAddress == "" || *masterPort == ""){
		MasterSlave.GeneralSlave(masterAddress, masterPort, CreateLocalPortForwardSession, targetAddress, targetPort)
	} else {
		fmt.Println("All flags need to be in use")
	}
}