// Package handler provides HTTP handler utilities for processing requests and responses.
package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

// GET handler get template
func GET[RESP any](logic func(*http.Request) (RESP, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// 仅允许 GET 请求
		// 在 HTTP/1.1 规范中，GET 请求不应该带有 Body，很多 HTTP 客户端和服务器（包括 Go 的 net/http）会忽略 GET 的 Body。
		if r.Method != http.MethodGet {
			ResponseWriterMethodError(w, http.StatusMethodNotAllowed)
			return
		}

		err := json.NewEncoder(w).Encode(NewResponse(logic(r)))
		if err != nil {
			// 遇到这样不可能发生的意外, 哪怕通知前端, 也是未定义行为, 增加了日志提醒级别
			slog.ErrorContext(r.Context(),
				"GET json.NewEncoder(w).Encode(response) panic error",
				slog.Any("error", err),
				slog.String("path", r.URL.Path),
			)
			return
		}
	}
}

func HTTP[REQ, RESP any](logic func(context.Context, REQ) (RESP, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		do(logic, w, r)
	}
}

func do[REQ, RESP any](logic func(context.Context, REQ) (RESP, error), w http.ResponseWriter, r *http.Request) {
	c := r.Context()

	var request REQ
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		// 如果恶意尝试过多, 会关闭
		slog.InfoContext(c, "do json.NewDecoder(r.Body).Decode(&request) error",
			slog.Any("error", err),
			slog.String("path", r.URL.Path),
			slog.String("method", r.Method),
		)
		ResponseWriterMethodError(w, http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(NewResponse(logic(c, request)))
	if err != nil {
		// 遇到这样不可能发生的意外, 哪怕通知前端, 也是未定义行为, 增加了日志提醒级别
		slog.ErrorContext(c, "do json.NewEncoder(w).Encode(response) panic error",
			slog.Any("error", err),
			slog.String("path", r.URL.Path),
			slog.String("method", r.Method),
		)
		return
	}
}

func POST[REQ, RESP any](logic func(context.Context, REQ) (RESP, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			ResponseWriterMethodError(w, http.StatusMethodNotAllowed)
			return
		}

		do(logic, w, r)
	}
}

func PUT[REQ, RESP any](logic func(context.Context, REQ) (RESP, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			ResponseWriterMethodError(w, http.StatusMethodNotAllowed)
			return
		}

		do(logic, w, r)
	}
}

func DELETE[REQ, RESP any](logic func(context.Context, REQ) (RESP, error)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			ResponseWriterMethodError(w, http.StatusMethodNotAllowed)
			return
		}

		do(logic, w, r)
	}
}
