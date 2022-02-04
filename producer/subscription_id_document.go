// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package producer

import (
	"net/http"

	"github.com/free5gc/http_wrapper"
	"github.com/free5gc/nssf/logger"
	"github.com/free5gc/openapi/models"
)

// HandleNSSAIAvailabilityUnsubscribe - Deletes an already existing NSSAI availability notification subscription
func HandleNSSAIAvailabilityUnsubscribe(request *http_wrapper.Request) *http_wrapper.Response {
	logger.Nssaiavailability.Infof("Handle NSSAIAvailabilityUnsubscribe")

	subscriptionID := request.Params["subscriptionId"]

	problemDetails := NSSAIAvailabilityUnsubscribeProcedure(subscriptionID)

	if problemDetails == nil {
		return http_wrapper.NewResponse(http.StatusNoContent, nil, nil)
	} else if problemDetails != nil {
		return http_wrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	}
	problemDetails = &models.ProblemDetails{
		Status: http.StatusForbidden,
		Cause:  "UNSPECIFIED",
	}
	return http_wrapper.NewResponse(http.StatusForbidden, nil, problemDetails)
}
