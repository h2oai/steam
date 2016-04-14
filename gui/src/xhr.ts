/// <reference path="references.ts" />
"use strict"

module Proxy {

    var _rpcId: number = 0
    function nextId(): number {
        return ++_rpcId
    }

    interface RpcRequest {
        method: string
        params: any[]
        id: number
    }

    interface RpcResponse {
        result: any
        error: string
        id: number
    }

    interface HttpStatusCode {
        status: any
        code: any
    }

    function invoke(method: string, param: any, headers: any, go: (error: Error, data: any) => void) {
        const req: RpcRequest = {
            method: `usr.${method}`,
            params: [param],
            id: nextId()
        }

        const settings: JQueryAjaxSettings = {
            url: "/urpc",
            type: "POST",
            data: JSON.stringify(req),
            contentType: "application/json; charset=utf-8",
            dataType: "json"
        }

        if (headers) {
            settings.headers = headers
        }

        const p = $.ajax(settings)

        p.done((data, status, xhr) => {
            if (data.error) {
                return go(new Error(data.error), null)
            }
            go(null, data.result)
        })

        p.fail((xhr, status, error) => {
            // XXX handle 401

            const res = xhr.responseJSON
            if (res && res.error) {
                return go(new Error(res.error), null)
            }

            // special-case net::ERR_CONNECTION_REFUSED
            if (status == 'error' && xhr.status == 0) {
                return go(new Error("Could not connect to Steam. The server is currently unresponsive."), null)
            }

            go(new Error(`HTTP connection failure: status=${status}, code=${xhr.status}, error=${error ? error : '?'}`), null)
        })
    }

    export function Call(method: string, param: any, go: (error: Error, data: any) => void) {
        invoke(method, param, null, go)
    }

    export function authenticate(username: string, password: string, go: (error: Error, data: any) => void) {
        const headers = {
            "Authorization": `Custom ${username} ${password}`
        }
        invoke("Ping", { status: true }, headers, go)
    }
}

