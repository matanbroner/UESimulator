package ue_handler

import (
	"github.com/matanbroner/UESimulator/src/ue/ue_context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/matanbroner/UESimulator/src/ue/logger"
	ue_procedures "github.com/matanbroner/UESimulator/src/ue/ue_actions"
	"github.com/matanbroner/UESimulator/src/ue/ue_handler/ue_message"
)

var handlerLog *logrus.Entry

func init() {
	// init pool
	handlerLog = logger.HandlerLog
}

func Handle() {
	for {
		select {
		case msg, ok := <-ue_message.UeChannel:
			if ok {
				switch msg.Event {
				case ue_message.EventRegistrationProcedure:
					handlerLog.Infof("Registration Procedure Triggered")
					ueContext := ue_context.UeSelf()
					ue_procedures.HandleRegistrationProcedure(ueContext)
				case ue_message.EventDeregistrationProcedure:
					handlerLog.Infof("Start Deregistration Procedure")
					ueContext := ue_context.UeSelf()
					ue_procedures.HandleDeregistrationProcedure(ueContext)
				case ue_message.EventPDUSessionEstablishment:
					handlerLog.Infof("Start Deregistration Procedure")
					ueContext := ue_context.UeSelf()
					ue_procedures.SetupPDUSession(ueContext)
				}
			}
		case <-time.After(1 * time.Second):
		}
	}
}
