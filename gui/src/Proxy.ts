// ----------------------------------
// --- Generated with go:generate ---
// ---        DO NOT EDIT         ---
// ----------------------------------

module Proxy {

    // --- Contract ---

    export interface Service {
        ping: (status: boolean, go: (error: Error, status: boolean) => void) => void
    }

    // --- Messages ---

    interface PingIn {
        status: boolean
    }

    interface PingOut {
        status: boolean
    }

    // --- Client Stub ---

    export function ping(status: boolean, go: (error: Error, status: boolean) => void): void {
        var req: PingIn = {
            status: status
        }
        Proxy.Call("Ping", req, function(error, data) {
            return error ? go(error, null) : go(null, (<PingOut>data).status)
        })

    }
}

