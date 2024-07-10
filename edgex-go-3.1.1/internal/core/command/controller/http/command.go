//
// Copyright (C) 2021-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package http

import (
	"fmt"
	"math"
	"net/http"

	"github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/container"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	responseDTO "github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/responses"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"

	"github.com/edgexfoundry/edgex-go/internal/core/command/application"
	commandContainer "github.com/edgexfoundry/edgex-go/internal/core/command/container"
	"github.com/edgexfoundry/edgex-go/internal/pkg"
	"github.com/edgexfoundry/edgex-go/internal/pkg/utils"

	"github.com/labstack/echo/v4"
)

type CommandController struct {
	dic *di.Container
}

// NewCommandController creates and initializes an CommandController
func NewCommandController(dic *di.Container) *CommandController {
	return &CommandController{
		dic: dic,
	}
}

func (cc *CommandController) AllCommands(c echo.Context) error {
	lc := container.LoggingClientFrom(cc.dic.Get)
	r := c.Request()
	w := c.Response()
	ctx := r.Context()
	config := commandContainer.ConfigurationFrom(cc.dic.Get)

	// parse URL query string for offset, limit
	offset, limit, _, err := utils.ParseGetAllObjectsRequestQueryString(c, 0, math.MaxInt32, -1, config.Service.MaxResultCount)
	if err != nil {
		return utils.WriteErrorResponse(w, ctx, lc, err, "")

	}
	commands, totalCount, err := application.AllCommands(offset, limit, cc.dic)
	if err != nil {
		return utils.WriteErrorResponse(w, ctx, lc, err, "")
	}

	response := responseDTO.NewMultiDeviceCoreCommandsResponse("", "", http.StatusOK, totalCount, commands)
	utils.WriteHttpHeader(w, ctx, http.StatusOK)
	// encode and send out the response
	return pkg.EncodeAndWriteResponse(response, w, lc)
}

func (cc *CommandController) CommandsByDeviceName(c echo.Context) error {
	lc := container.LoggingClientFrom(cc.dic.Get)
	r := c.Request()
	w := c.Response()
	ctx := r.Context()

	// URL parameters
	name := c.Param(common.Name)
	deviceCoreCommand, err := application.CommandsByDeviceName(name, cc.dic)
	if err != nil {
		return utils.WriteErrorResponse(w, ctx, lc, err, "")
	}

	response := responseDTO.NewDeviceCoreCommandResponse("", "", http.StatusOK, deviceCoreCommand)
	utils.WriteHttpHeader(w, ctx, http.StatusOK)
	// encode and send out the response
	return pkg.EncodeAndWriteResponse(response, w, lc)
}

func validateGetCommandParameters(r *http.Request) (err errors.EdgeX) {
	dsReturnEvent := utils.ParseQueryStringToString(r, common.ReturnEvent, common.ValueTrue)
	dsPushEvent := utils.ParseQueryStringToString(r, common.PushEvent, common.ValueFalse)
	if dsReturnEvent != common.ValueTrue && dsReturnEvent != common.ValueFalse {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid query parameter, %s has to be %s or %s", dsReturnEvent, common.ValueTrue, common.ValueFalse), nil)
	}
	if dsPushEvent != common.ValueTrue && dsPushEvent != common.ValueFalse {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid query parameter, %s has to be %s or %s", dsPushEvent, common.ValueTrue, common.ValueFalse), nil)
	}
	return nil
}

func (cc *CommandController) IssueGetCommandByName(c echo.Context) error {
	lc := container.LoggingClientFrom(cc.dic.Get)
	r := c.Request()
	w := c.Response()
	ctx := r.Context()

	// URL parameters
	deviceName := c.Param(common.Name)
	commandName := c.Param(common.Command)

	// Query params
	queryParams := r.URL.RawQuery
	err := validateGetCommandParameters(r)
	if err != nil {
		return utils.WriteErrorResponse(w, ctx, lc, err, "")
	}

	response, err := application.IssueGetCommandByName(deviceName, commandName, queryParams, cc.dic)
	if err != nil {
		return utils.WriteErrorResponse(w, ctx, lc, err, "")
	}
	// encode and send out the response
	if response != nil {
		utils.WriteHttpHeader(w, ctx, response.StatusCode)
		return pkg.EncodeAndWriteResponse(response, w, lc)
	}
	// If dsReturnEvent is no, there will be no content returned in the http response
	utils.WriteHttpHeader(w, ctx, http.StatusOK)
	return nil
}

func (cc *CommandController) IssueSetCommandByName(c echo.Context) error {
	lc := container.LoggingClientFrom(cc.dic.Get)
	r := c.Request()
	w := c.Response()
	ctx := r.Context()

	// URL parameters
	deviceName := c.Param(common.Name)
	commandName := c.Param(common.Command)
	// Query params
	queryParams := r.URL.RawQuery

	// Request body
	settings, err := utils.ParseBodyToMap(r)
	if err != nil {
		return utils.WriteErrorResponse(w, ctx, lc, err, "")
	}
	response, err := application.IssueSetCommandByName(deviceName, commandName, queryParams, settings, cc.dic)
	if err != nil {
		return utils.WriteErrorResponse(w, ctx, lc, err, "")
	}

	utils.WriteHttpHeader(w, ctx, response.StatusCode)
	// encode and send out the response
	return pkg.EncodeAndWriteResponse(response, w, lc)
}
