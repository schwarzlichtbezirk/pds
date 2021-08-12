package main

import (
	"context"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/schwarzlichtbezirk/pds/pb"
)

// API error codes.
// Each error code have unique source code point,
// so this error code at service reply exactly points to error place.
const (
	AECnull = iota
	AECbadbody
	AECnoreq
	AECbadjson

	// api/tool/ping
	AECtoolpingcallfail

	// api/port/set

	AECportsetnodata
	AECportsetcallfail

	// api/port/get

	AECportgetnodata
	AECportgetcallfail

	// api/port/name

	AECportnamenodata
	AECportnamecallfail

	// api/port/near

	AECportnearnodata
	AECportnearcallfail

	// api/port/circle

	AECportcircnodata
	AECportcirccallfail

	// api/port/text

	AECportcitynodata
	AECportcitycallfail
)

// HTTP error messages
var (
	ErrNoJSON = errors.New("data not given")
	ErrNoData = errors.New("data is empty")
)

func apiToolPing(w http.ResponseWriter, r *http.Request) {
	var err error
	var body, _ = io.ReadAll(r.Body)
	var ret *pb.Content

	// limit execution time of the action
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// make rpc call
	if ret, err = grpcTool.Ping(ctx, &pb.Content{Value: body}); err != nil {
		WriteError500(w, err, AECtoolpingcallfail)
		return
	}

	w.WriteHeader(http.StatusOK)
	WriteJSONHeader(w)
	w.Write(ret.Value)
}

func apiPortSet(w http.ResponseWriter, r *http.Request) {
	var err error
	var arg pb.Port
	var ret *pb.Key

	// get arguments
	if err = AjaxGetArg(r, &arg); err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}
	if arg.Name == "" || len(arg.Unlocs) == 0 {
		WriteError400(w, err, AECportsetnodata)
		return
	}

	// limit execution time of the action
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// make rpc call
	if ret, err = grpcClient.SetByKey(ctx, &arg); err != nil {
		WriteError500(w, err, AECportsetcallfail)
		return
	}

	WriteOK(w, ret)
}

func apiPortGet(w http.ResponseWriter, r *http.Request) {
	var err error
	var arg pb.Key
	var ret *pb.Port

	// get arguments
	if err = AjaxGetArg(r, &arg); err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}
	if arg.Value == "" {
		WriteError400(w, err, AECportgetnodata)
		return
	}

	// limit execution time of the action
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// make rpc call
	if ret, err = grpcClient.GetByKey(ctx, &arg); err != nil {
		WriteError500(w, err, AECportgetcallfail)
		return
	}

	WriteOK(w, ret)
}

func apiPortName(w http.ResponseWriter, r *http.Request) {
	var err error
	var arg pb.Name
	var ret *pb.Port

	// get arguments
	if err = AjaxGetArg(r, &arg); err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}
	if arg.Value == "" {
		WriteError400(w, err, AECportnamenodata)
		return
	}

	// limit execution time of the action
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// make rpc call
	if ret, err = grpcClient.GetByName(ctx, &arg); err != nil {
		WriteError500(w, err, AECportnamecallfail)
		return
	}

	WriteOK(w, ret)
}

func apiPortNear(w http.ResponseWriter, r *http.Request) {
	var err error
	var arg pb.Point
	var ret *pb.Port

	// get arguments
	if err = AjaxGetArg(r, &arg); err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}
	if arg.Latitude == 0 && arg.Longitude == 0 {
		WriteError400(w, err, AECportnearnodata)
		return
	}

	// limit execution time of the action
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// make rpc call
	if ret, err = grpcClient.FindNearest(ctx, &arg); err != nil {
		WriteError500(w, err, AECportnearcallfail)
		return
	}

	WriteOK(w, ret)
}

func apiPortCircle(w http.ResponseWriter, r *http.Request) {
	var err error
	var arg pb.Circle
	var ret *pb.Ports

	// get arguments
	if err = AjaxGetArg(r, &arg); err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}
	if (arg.Center.Latitude == 0 && arg.Center.Longitude == 0) || arg.Radius <= 0 {
		WriteError400(w, err, AECportcircnodata)
		return
	}

	// limit execution time of the action
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// make rpc call
	if ret, err = grpcClient.FindInCircle(ctx, &arg); err != nil {
		WriteError500(w, err, AECportcirccallfail)
		return
	}

	WriteOK(w, ret)
}

func apiPortText(w http.ResponseWriter, r *http.Request) {
	var err error
	var arg pb.Quest
	var ret *pb.Ports

	// get arguments
	if err = AjaxGetArg(r, &arg); err != nil {
		WriteJSON(w, http.StatusBadRequest, err)
		return
	}
	if arg.Value == "" {
		WriteError400(w, err, AECportcitynodata)
		return
	}

	// limit execution time of the action
	var ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// make rpc call
	if ret, err = grpcClient.FindText(ctx, &arg); err != nil {
		WriteError500(w, err, AECportcitycallfail)
		return
	}

	WriteOK(w, ret)
}
