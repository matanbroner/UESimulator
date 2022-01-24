package ue_actions

import (
	"github.com/free5gc/nas"
	"github.com/free5gc/nas/nasMessage"
	"github.com/free5gc/ngap"
	// Nausf_UEAU_Client "github.com/free5gc/openapi/Nausf_UEAuthentication"
	// "github.com/free5gc/openapi/models"

	"free5gc_ue/src/ue/ue_ngap"
)

func GetNGSetupRequest(gnbId []byte, bitlength uint64, name string) ([]byte, error) {
	message := ue_ngap.BuildNGSetupRequest()
	// GlobalRANNodeID
	ie := message.InitiatingMessage.Value.NGSetupRequest.ProtocolIEs.List[0]
	gnbID := ie.Value.GlobalRANNodeID.GlobalGNBID.GNBID.GNBID
	gnbID.Bytes = gnbId
	gnbID.BitLength = bitlength
	// RANNodeName
	ie = message.InitiatingMessage.Value.NGSetupRequest.ProtocolIEs.List[1]
	ie.Value.RANNodeName.Value = name

	return ngap.Encoder(message)
}

func GetInitialUEMessage(ranUeNgapID int64, nasPdu []byte, fiveGSTmsi string) ([]byte, error) {
	message := ue_ngap.BuildInitialUEMessage(ranUeNgapID, nasPdu, fiveGSTmsi)
	return ngap.Encoder(message)
}

func GetUplinkNASTransport(amfUeNgapID, ranUeNgapID int64, nasPdu []byte) ([]byte, error) {
	message := ue_ngap.BuildUplinkNasTransport(amfUeNgapID, ranUeNgapID, nasPdu)
	return ngap.Encoder(message)
}

func GetInitialContextSetupResponse(amfUeNgapID int64, ranUeNgapID int64) ([]byte, error) {
	message := ue_ngap.BuildInitialContextSetupResponseForRegistraionTest(amfUeNgapID, ranUeNgapID)

	return ngap.Encoder(message)
}

func GetInitialContextSetupResponseForServiceRequest(
	amfUeNgapID int64, ranUeNgapID int64, ipv4 string) ([]byte, error) {
	message := ue_ngap.BuildInitialContextSetupResponse(amfUeNgapID, ranUeNgapID, ipv4, nil)
	return ngap.Encoder(message)
}

func GetPDUSessionResourceSetupResponse(pduSessionId int64, amfUeNgapID int64, ranUeNgapID int64, ipv4 string) ([]byte, error) {
	message := ue_ngap.BuildPDUSessionResourceSetupResponseForRegistrationTest(pduSessionId, amfUeNgapID, ranUeNgapID, ipv4)
	return ngap.Encoder(message)
}
func EncodeNasPduWithSecurity(ue *RanUeContext, pdu []byte, securityHeaderType uint8,
	securityContextAvailable, newSecurityContext bool) ([]byte, error) {
	m := nas.NewMessage()
	err := m.PlainNasDecode(&pdu)
	if err != nil {
		return nil, err
	}
	m.SecurityHeader = nas.SecurityHeader{
		ProtocolDiscriminator: nasMessage.Epd5GSMobilityManagementMessage,
		SecurityHeaderType:    securityHeaderType,
	}
	return NASEncode(ue, m, securityContextAvailable, newSecurityContext)
}

func GetUEContextReleaseComplete(amfUeNgapID int64, ranUeNgapID int64, pduSessionIDList []int64) ([]byte, error) {
	message := ue_ngap.BuildUEContextReleaseComplete(amfUeNgapID, ranUeNgapID, pduSessionIDList)
	return ngap.Encoder(message)
}

func GetUEContextReleaseRequest(amfUeNgapID int64, ranUeNgapID int64, pduSessionIDList []int64) ([]byte, error) {
	message := ue_ngap.BuildUEContextReleaseRequest(amfUeNgapID, ranUeNgapID, pduSessionIDList)
	return ngap.Encoder(message)
}

func GetPDUSessionResourceReleaseResponse(amfUeNgapID int64, ranUeNgapID int64) ([]byte, error) {
	message := ue_ngap.BuildPDUSessionResourceReleaseResponseForReleaseTest(amfUeNgapID, ranUeNgapID)
	return ngap.Encoder(message)
}
func GetPathSwitchRequest(amfUeNgapID int64, ranUeNgapID int64) ([]byte, error) {
	message := ue_ngap.BuildPathSwitchRequest(amfUeNgapID, ranUeNgapID)
	message.InitiatingMessage.Value.PathSwitchRequest.ProtocolIEs.List =
		message.InitiatingMessage.Value.PathSwitchRequest.ProtocolIEs.List[0:5]
	return ngap.Encoder(message)
}

func GetHandoverRequired(
	amfUeNgapID int64, ranUeNgapID int64, targetGNBID []byte, targetCellID []byte) ([]byte, error) {
	message := ue_ngap.BuildHandoverRequired(amfUeNgapID, ranUeNgapID, targetGNBID, targetCellID)
	return ngap.Encoder(message)
}

func GetHandoverRequestAcknowledge(amfUeNgapID int64, ranUeNgapID int64) ([]byte, error) {
	message := ue_ngap.BuildHandoverRequestAcknowledge(amfUeNgapID, ranUeNgapID)
	return ngap.Encoder(message)
}

func GetHandoverNotify(amfUeNgapID int64, ranUeNgapID int64) ([]byte, error) {
	message := ue_ngap.BuildHandoverNotify(amfUeNgapID, ranUeNgapID)
	return ngap.Encoder(message)
}

func GetPDUSessionResourceSetupResponseForPaging(amfUeNgapID int64, ranUeNgapID int64, ipv4 string) ([]byte, error) {
	message := ue_ngap.BuildPDUSessionResourceSetupResponseForPaging(amfUeNgapID, ranUeNgapID, ipv4)
	return ngap.Encoder(message)
}
