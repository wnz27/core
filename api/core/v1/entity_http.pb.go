// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http 0.1.0

package v1

import (
	context "context"
	json "encoding/json"
	go_restful "github.com/emicklei/go-restful"
	http "net/http"
)

import transportHTTP "github.com/tkeel-io/kit/transport/http"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the tkeel package it is being compiled against.
// import package.context.http.go_restful.json.

const _ = transportHTTP.ImportAndUsed

type EntityHTTPServer interface {
	AppendMapper(context.Context, *AppendMapperRequest) (*EntityResponse, error)
	CreateEntity(context.Context, *CreateEntityRequest) (*EntityResponse, error)
	DeleteEntity(context.Context, *DeleteEntityRequest) (*DeleteEntityResponse, error)
	GetEntity(context.Context, *GetEntityRequest) (*EntityResponse, error)
	ListEntity(context.Context, *ListEntityRequest) (*ListEntityResponse, error)
	UpdateEntity(context.Context, *UpdateEntityRequest) (*EntityResponse, error)
}

type EntityHTTPHandler struct {
	srv EntityHTTPServer
}

func newEntityHTTPHandler(s EntityHTTPServer) *EntityHTTPHandler {
	return &EntityHTTPHandler{srv: s}
}

func (h *EntityHTTPHandler) AppendMapper(req *go_restful.Request, resp *go_restful.Response) {
	in := AppendMapperRequest{}
	if err := transportHTTP.GetBody(req, &in.Mapper); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.AppendMapper(req.Request.Context(), &in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *EntityHTTPHandler) CreateEntity(req *go_restful.Request, resp *go_restful.Response) {
	in := CreateEntityRequest{}
	if err := transportHTTP.GetBody(req, &in.Properties); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.CreateEntity(req.Request.Context(), &in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *EntityHTTPHandler) DeleteEntity(req *go_restful.Request, resp *go_restful.Response) {
	in := DeleteEntityRequest{}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.DeleteEntity(req.Request.Context(), &in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *EntityHTTPHandler) GetEntity(req *go_restful.Request, resp *go_restful.Response) {
	in := GetEntityRequest{}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.GetEntity(req.Request.Context(), &in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *EntityHTTPHandler) ListEntity(req *go_restful.Request, resp *go_restful.Response) {
	in := ListEntityRequest{}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.ListEntity(req.Request.Context(), &in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *EntityHTTPHandler) UpdateEntity(req *go_restful.Request, resp *go_restful.Response) {
	in := UpdateEntityRequest{}
	if err := transportHTTP.GetBody(req, &in.Properties); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetQuery(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	if err := transportHTTP.GetPathValue(req, &in); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	out, err := h.srv.UpdateEntity(req.Request.Context(), &in)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(out)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
	_, err = resp.Write(result)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}
}

func RegisterEntityHTTPServer(container *go_restful.Container, srv EntityHTTPServer) {
	var ws *go_restful.WebService
	for _, v := range container.RegisteredWebServices() {
		if v.RootPath() == "/v1" {
			ws = v
			break
		}
	}
	if ws == nil {
		ws = new(go_restful.WebService)
		ws.ApiVersion("/v1")
		ws.Path("/v1").Produces(go_restful.MIME_JSON)
		container.Add(ws)
	}

	handler := newEntityHTTPHandler(srv)
	ws.Route(ws.POST("/plugins/{plugin}/entities").
		To(handler.CreateEntity))
	ws.Route(ws.PUT("/plugins/{plugin}/entities/{id}").
		To(handler.UpdateEntity))
	ws.Route(ws.DELETE("/plugins/{plugin}/entities/{id}").
		To(handler.DeleteEntity))
	ws.Route(ws.GET("/plugins/{plugin}/entities/{id}").
		To(handler.GetEntity))
	ws.Route(ws.GET("/plugins/{plugin}/entities").
		To(handler.ListEntity))
	ws.Route(ws.POST("/plugins/{plugin}/entities/{id}/mappers").
		To(handler.AppendMapper))
}
