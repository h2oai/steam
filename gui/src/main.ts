/// <reference path="references.ts" />
/// <reference path="app.ts" />
"use strict"

module Main {
    // Export app for debugging.
    // Accessible as window.Main.app from the browser's console.
    export var app: App

    function boot() {
        app = newApp()
        ko.applyBindings(app)
    }

    (document.readyState !== 'loading')
        ? boot()
        : document.addEventListener('DOMContentLoaded', boot)
}

