/// <reference path="references.ts" />
/// <reference path="xhr.ts" />
/// <reference path="proxy.ts" />
"use strict"

module Main {


    //
    // Prelude
    //
    const Epsilon = 1e-6

    export type int = number
    export type float = number
    export type byte = number
    export interface Dict<T> { [key: string]: T }

    export type Predicate = (...conditions: any[]) => boolean

    export type Func<T1, T> = (arg: T1) => T
    export type Func2<T1, T2, T> = (v1: T1, v2: T2) => T
    export type Func3<T1, T2, T3, T> = (v1: T1, v2: T2, v3: T3) => T
    export type Func4<T1, T2, T3, T4, T> = (v1: T1, v2: T2, v3: T3, v4: T4) => T

    export type Eff<T> = (t: T) => void
    export type Eff2<T1, T2> = (v1: T1, v2: T2) => void
    export type Eff3<T1, T2, T3> = (v1: T1, v2: T2, v3: T3) => void
    export type Eff4<T1, T2, T3, T4> = (v1: T1, v2: T2, v3: T3, v4: T4) => void

    export type Act = () => void
    export type Get<T> = () => T
    export type Go = (err?: Error) => void
    export type On<T> = (err: Error, t: T) => void
    export type Relay = Eff<Act>

    export type Sig<T> = KnockoutObservable<T>
    export type Sigs<T> = KnockoutObservableArray<T>
    export type Arrow = KnockoutSubscription

    interface Noise {
        seed: (seed: float) => void
        simplex2: (x: float, y: float) => float
    }

    declare var noise: Noise

    //
    // Unicast Delegates
    //

    interface Uni {
        (): void
        on: (f: Act) => void
        off: (f: Act) => void
        dispose: () => void
    }

    interface Uni1<T1> {
        (v1: T1): void
        on: (f: Eff<T1>) => void
        off: (f: Eff<T1>) => void
        dispose: () => void
    }

    interface Uni2<T1, T2> {
        (v1: T1, v2: T2): void
        on: (f: Eff2<T1, T2>) => void
        off: (f: Eff2<T1, T2>) => void
        dispose: () => void
    }

    interface Uni3<T1, T2, T3> {
        (v1: T1, v2: T2, v3: T3): void
        on: (f: Eff3<T1, T2, T3>) => void
        off: (f: Eff3<T1, T2, T3>) => void
        dispose: () => void
    }

    interface Uni4<T1, T2, T3, T4> {
        (v1: T1, v2: T2, v3: T3, v4: T4): void
        on: (f: Eff4<T1, T2, T3, T4>) => void
        off: (f: Eff4<T1, T2, T3, T4>) => void
        dispose: () => void
    }

    export function uni(): Uni {
        let link: Act = null
        let trigger: any = (): void => {
            if (link) {
                link()
            }
        }
        trigger.on = (f: Act): void => {
            if (link) {
                throw new Error('Delegate is unicast.')
            }
            link = f
        }
        trigger.off = (f: Act): void => {
            if (f !== link) {
                throw new Error('Invalid subscription.')
            }
            link = null
        }
        trigger.dispose = () => {
            link = null
        }
        return trigger
    }

    export function uni1<T1>(): Uni1<T1> {
        let link: Eff<T1> = null
        let trigger: any = (v1: T1): void => {
            if (link) {
                link(v1)
            }
        }
        trigger.on = (f: Eff<T1>): void => {
            if (link) {
                throw new Error('Delegate is unicast.')
            }
            link = f
        }
        trigger.off = (f: Eff<T1>): void => {
            if (f !== link) {
                throw new Error('Invalid subscription.')
            }
            link = null
        }
        trigger.dispose = () => {
            link = null
        }
        return trigger
    }

    export function uni2<T1, T2>(): Uni2<T1, T2> {
        let link: Eff2<T1, T2> = null
        let trigger: any = (v1: T1, v2: T2): void => {
            if (link) {
                link(v1, v2)
            }
        }
        trigger.on = (f: Eff2<T1, T2>): void => {
            if (link) {
                throw new Error('Delegate is unicast.')
            }
            link = f
        }
        trigger.off = (f: Eff2<T1, T2>): void => {
            if (f !== link) {
                throw new Error('Invalid subscription.')
            }
            link = null
        }
        trigger.dispose = () => {
            link = null
        }
        return trigger
    }

    export function uni3<T1, T2, T3>(): Uni3<T1, T2, T3> {
        let link: Eff3<T1, T2, T3> = null
        let trigger: any = (v1: T1, v2: T2, v3: T3): void => {
            if (link) {
                link(v1, v2, v3)
            }
        }
        trigger.on = (f: Eff3<T1, T2, T3>): void => {
            if (link) {
                throw new Error('Delegate is unicast.')
            }
            link = f
        }
        trigger.off = (f: Eff3<T1, T2, T3>): void => {
            if (f !== link) {
                throw new Error('Invalid subscription.')
            }
            link = null
        }
        trigger.dispose = () => {
            link = null
        }
        return trigger
    }

    export function uni4<T1, T2, T3, T4>(): Uni4<T1, T2, T3, T4> {
        let link: Eff4<T1, T2, T3, T4> = null
        let trigger: any = (v1: T1, v2: T2, v3: T3, v4: T4): void => {
            if (link) {
                link(v1, v2, v3, v4)
            }
        }
        trigger.on = (f: Eff4<T1, T2, T3, T4>): void => {
            if (link) {
                throw new Error('Delegate is unicast.')
            }
            link = f
        }
        trigger.off = (f: Eff4<T1, T2, T3, T4>): void => {
            if (f !== link) {
                throw new Error('Invalid subscription.')
            }
            link = null
        }
        trigger.dispose = () => {
            link = null
        }
        return trigger
    }

    //
    // Multicast Delegates
    //

    interface Multi1<T1> {
        (v1: T1): void
        on: (f: Eff<T1>) => void
        off: (f: Eff<T1>) => void
        dispose: () => void
    }

    interface Multi2<T1, T2> {
        (v1: T1, v2: T2): void
        on: (f: Eff2<T1, T2>) => void
        off: (f: Eff2<T1, T2>) => void
        dispose: () => void
    }

    interface Multi3<T1, T2, T3> {
        (v1: T1, v2: T2, v3: T3): void
        on: (f: Eff3<T1, T2, T3>) => void
        off: (f: Eff3<T1, T2, T3>) => void
        dispose: () => void
    }

    interface Multi4<T1, T2, T3, T4> {
        (v1: T1, v2: T2, v3: T3, v4: T4): void
        on: (f: Eff4<T1, T2, T3, T4>) => void
        off: (f: Eff4<T1, T2, T3, T4>) => void
        dispose: () => void
    }

    export function multi3<T1, T2, T3>(): Multi3<T1, T2, T3> {
        const links: Eff3<T1, T2, T3>[] = []
        let trigger: any = (v1: T1, v2: T2, v3: T3): void => {
            for (const f of links) {
                f(v1, v2, v3)
            }
        }
        trigger.on = (f: Eff3<T1, T2, T3>): void => {
            links.push(f)
        }
        trigger.off = (f: Eff3<T1, T2, T3>): void => {
            const i = _.indexOf(links, f)
            if (i < 0) {
                return
            }
            links.splice(i, 1)
        }
        trigger.dispose = () => {
            links.length = 0
        }
        return trigger
    }

    export function sig<T>(value: T, equalityComparer?: (a: T, b: T) => boolean): Sig<T> {
        let o = ko.observable<T>(value)
        if (equalityComparer) {
            o.equalityComparer = equalityComparer
        }
        return o
    }

    export function sigs<T>(value: T[]): Sigs<T> {
        return ko.observableArray<T>(value)
    }

    export function react<T1>(s1: Sig<T1>, a: (v1: T1) => void): Arrow {
        return s1.subscribe(a)
    }

    export function reacts<T1>(s1: Sigs<T1>, a: (v1: T1[]) => void): Arrow {
        return s1.subscribe(a)
    }

    export function react2<T1, T2>(s1: Sig<T1>, s2: Sig<T2>, f: Eff2<T1, T2>): Arrow[] {
        return [
            react(s1, (v1: T1) => { f(v1, s2()) }),
            react(s2, (v2: T2) => { f(s1(), v2) })
        ]
    }

    export function react3<T1, T2, T3>(s1: Sig<T1>, s2: Sig<T2>, s3: Sig<T3>, f: Eff3<T1, T2, T3>): Arrow[] {
        return [
            react(s1, (v1: T1) => { f(v1, s2(), s3()) }),
            react(s2, (v2: T2) => { f(s1(), v2, s3()) }),
            react(s3, (v3: T3) => { f(s1(), s2(), v3) })
        ]
    }

    export function react4<T1, T2, T3, T4>(s1: Sig<T1>, s2: Sig<T2>, s3: Sig<T3>, s4: Sig<T4>, f: Eff4<T1, T2, T3, T4>): Arrow[] {
        return [
            react(s1, (v1: T1) => { f(v1, s2(), s3(), s4()) }),
            react(s2, (v2: T2) => { f(s1(), v2, s3(), s4()) }),
            react(s3, (v3: T3) => { f(s1(), s2(), v3, s4()) }),
            react(s4, (v4: T4) => { f(s1(), s2(), s3(), v4) })
        ]
    }

    export function act<T1>(s1: Sig<T1>, f: Eff<T1>): Arrow {
        f(s1())
        return react<T1>(s1, f)
    }

    export function act2<T1, T2>(s1: Sig<T1>, s2: Sig<T2>, f: Eff2<T1, T2>): Arrow[] {
        f(s1(), s2())
        return react2<T1, T2>(s1, s2, f)
    }

    export function act3<T1, T2, T3>(s1: Sig<T1>, s2: Sig<T2>, s3: Sig<T3>, f: Eff3<T1, T2, T3>): Arrow[] {
        f(s1(), s2(), s3())
        return react3<T1, T2, T3>(s1, s2, s3, f)
    }

    export function lift<T1, T>(s1: Sig<T1>, f: Func<T1, T>): Sig<T> {
        let t = sig<T>(f(s1()))
        react<T1>(s1, (v1: T1) => { t(f(v1)) })
        return t
    }

    export function lifts<T1, T>(s1: Sigs<T1>, f: Func<T1[], T>): Sig<T> {
        let t = sig<T>(f(s1()))
        reacts<T1>(s1, (v1: T1[]) => { t(f(v1)) })
        return t
    }

    export function lift2<T1, T2, T>(s1: Sig<T1>, s2: Sig<T2>, f: Func2<T1, T2, T>): Sig<T> {
        let t = sig<T>(f(s1(), s2()))
        react2<T1, T2>(s1, s2, (v1: T1, v2: T2) => { t(f(v1, v2)) })
        return t
    }

    export function lift3<T1, T2, T3, T>(s1: Sig<T1>, s2: Sig<T2>, s3: Sig<T3>, f: Func3<T1, T2, T3, T>): Sig<T> {
        let t = sig<T>(f(s1(), s2(), s3()))
        react3<T1, T2, T3>(s1, s2, s3, (v1: T1, v2: T2, v3: T3) => { t(f(v1, v2, v3)) })
        return t
    }

    export function lift4<T1, T2, T3, T4, T>(s1: Sig<T1>, s2: Sig<T2>, s3: Sig<T3>, s4: Sig<T4>, f: Func4<T1, T2, T3, T4, T>): Sig<T> {
        let t = sig<T>(f(s1(), s2(), s3(), s4()))
        react4<T1, T2, T3, T4>(s1, s2, s3, s4, (v1: T1, v2: T2, v3: T3, v4: T4) => { t(f(v1, v2, v3, v4)) })
        return t
    }

    function noop() { }

    function isNonEmpty<T>(a: T[]): boolean {
        return a.length !== 0
    }

    function isEmpty<T>(a: T[]): boolean {
        return a.length === 0
    }

    function isNull(a: any): boolean {
        return a === null
    }
    function isUndefined(a: any): boolean {
        return a === void 0
    }

    function timestampToAge(t: number): string {
        return moment.unix(t).fromNow()
    }

    function formatTimestamp(t: number): string {
        return moment.unix(t).format("MMM D YYYY h:mm:ss a")
    }

    //
    // Knockout Extensions
    //

    // ko.bindingHandlers['element'] = {
    //     update: (element, valueAccessor, allBindings, viewModel, bindingContext) => {
    //         const arg = ko.unwrap(valueAccessor())
    //         if (arg) {
    //             const $element = $(element)
    //             $element.empty()
    //             $element.append(arg)
    //         }
    //         return
    //     }
    // }
    ko.bindingHandlers['element'] = {
        init: (element, valueAccessor, allBindings, viewModel, bindingContext) => {
            valueAccessor()(element)
        }
    }

    ko.bindingHandlers['file'] = {
        init: (element, valueAccessor, allBindings, viewModel, bindingContext) => {
            const file = valueAccessor()
            if (file) {
                const $file = $(element)
                $file.change(() => {
                    file((<HTMLInputElement>$file[0]).files[0])
                })
            }
            return
        }
    }

    ko.bindingHandlers['autoscroll'] = {
        init: (element, valueAccessor, allBindings, viewModel, bindingContext) => {
            // Bit of a hack. Attaches a method to the viewModel that scrolls the pane into view

            const $el = $(element)
            const $viewport = $el.closest('.workspace')
            viewModel.ensureVisible = () => {
                const p = $viewport.scrollLeft()
                const l = viewModel.left()
                const w = viewModel.width()
                const vw = $viewport.width()

                if (l + w > p + vw) {
                    $viewport.animate({ scrollLeft: l + w - vw }, 'fast')
                } else if (l < p) {
                    $viewport.animate({ scrollLeft: l }, 'fast')
                }
            }
            return
        }
    }

    //
    // Models
    //

    class Model {
        constructor(
            public id: string,
            public cloud: string,
            public algo: string,
            public frame: string,
            public responseColumn: string,
            public createdAt: string
        ) { }
    }

    class Service {
        constructor(
            public id: string,
            public endpoint: string,
            public createdAt: string
        ) { }
    }

    interface Engine {
        name: string
        path: string
        createdAt: string
    }

    //
    // Components
    // 

    class NavButton {
        constructor(
            public icon: string,
            public title: string,
            public isSelected: Sig<boolean>,
            public execute: Act
        ) { }
    }

    class NavBar {
        constructor(
            public buttons: NavButton[]
        ) { }
    }

    function newNavBar(ctx: Context): NavBar {

        const newNavButton = (icon: string, title: string, isSelected: boolean, execute: Act): NavButton => {
            const button = new NavButton(
                icon,
                title,
                sig(isSelected),
                () => {
                    for (const b of buttons) {
                        b.isSelected(b === button)
                    }
                    execute()
                }
            )
            return button
        }

        const buttons: NavButton[] = [
            newNavButton('ion-ios-cloud-outline', 'Clusters', true, ctx.showClouds),
            newNavButton('ion-ios-color-filter-outline', 'Models', false, ctx.showModels),
            newNavButton('ion-ios-world-outline', 'Services', false, ctx.showServices),
            newNavButton('ion-ios-paper-outline', 'Assets', false, ctx.showAssets)
        ]

        return new NavBar(buttons)
    }

    interface Templated {
        template: string
    }

    interface Disposable {
        dispose: Act
    }

    function templateOf(t: Templated): string {
        return `tmpl-${t.template}`
    }

    function px(value: int): string {
        return `${Math.round(value)}px`
    }

    interface PaneOpts {
        width?: int
    }

    interface PanePosition {
        left: Sig<int>
        width: Sig<int>
        leftPx: Sig<string>
        widthPx: Sig<string>
        ensureVisible: Act
    }

    function newPanePosition(width: float = 200): PanePosition {
        const l = sig<int>(0)
        const w = sig<int>(width)
        return {
            left: l,
            width: w,
            leftPx: lift(l, px),
            widthPx: lift(w, px),
            ensureVisible: noop
        }
    }

    function doAfterRender(elements: HTMLElement[]): void {
        $(elements).click(function () {
            const $this = $(this)
            $this.parent().children().removeClass('folder--selected')
            $this.addClass('folder--selected')
        })
    }

    interface Pane extends Templated {
        title: string
        dispose: Act
        position: PanePosition
    }

    interface Folder extends Templated {
        title: string
        subhead: string
        slug: string
        execute: Act
        template: string
    }

    // FIXME: Get rid of the "I" postfix -- what is this for?
    interface FolderI extends Folder {
        isActive: Sig<boolean>
        subheadI: string
        slugI: Sig<string>
    }

    interface Dialog extends Templated {
        title: string
        cancel: Act
        dispose: Act
        template: string
    }

    interface RegisterCloudDialog extends Dialog {
        address: Sig<string>
        addressError: Sig<string>
        canRegisterCloud: Sig<boolean>
        registerCloud: Act
        error: Sig<string>
    }

    interface RegisterCloudDialogResult {
        clusterId: int
    }

    interface StartCloudDialog extends Dialog {
        engines: Sigs<Proxy.Engine>
        engine: Sig<Proxy.Engine>
        engineError: Sig<string>
        cloudId: Sig<string>
        cloudIdError: Sig<string>
        cloudSize: Sig<string>
        cloudSizeError: Sig<string>
        cloudMemory: Sig<string>
        cloudMemoryError: Sig<string>
        canStartCloud: Sig<boolean>
        startCloud: Act
        error: Sig<string>
    }

    interface StartCloudDialogResult {
        clusterId: int
    }

    interface BuildModelDialog extends Dialog {
        frame: Sig<string>
        frameError: Sig<string>
        responseColumn: Sig<string>
        responseColumnError: Sig<string>
        maxRunTime: Sig<string>
        maxRunTimeError: Sig<string>
        canBuildModel: Sig<boolean>
        buildModel: Act
        error: Sig<string>
    }

    interface BuildModelDialogResult {
        success: boolean
    }

    interface DeployModelDialog extends Dialog {
        port: Sig<string>
        portError: Sig<string>
        canDeployModel: Sig<boolean>
        deployModel: Act
        error: Sig<string>
    }

    interface DeployModelDialogResult {
        success: boolean // FIXME
    }

    interface AddEngineDialog extends Dialog {
        form: Sig<HTMLFormElement>
        file: Sig<File>
        addEngine: Act
        error: Sig<string>
    }
    interface AddEngineDialogResult {
        success: boolean // FIXME
    }

    interface CloudsPane extends Pane {
        error: Sig<string>
        items: Sigs<FolderI>
        hasItems: Sig<boolean>
        registerCloud: Act
        startCloud: Act
    }

    interface CloudPane extends Pane {
        error: Sig<string>
        items: FolderI[]
    }

    interface CloudDetailsPane extends Pane {
        // engineName: string
        // size: string
        // applicationId: string
        // memory: string
        // username: string
        address: string
        state: Sig<string>
        createdAt: string
        cloudDetails: Sig<CloudDetail>
        yarnCluster: Sig<Proxy.YarnCluster>
        error: Sig<string>
        canStop: Sig<boolean>
        canStopCloud: boolean
        stopCloud: Act
        unregisterCloud: Act
    }

    interface CloudDetail {
        engineVersion: string
        totalMemory: string
        totalCores: string
        allowedCores: string
    }

    interface JobsPane extends Pane {
        error: Sig<string>
        items: Sigs<Folder>
        hasItems: Sig<boolean>
    }

    interface JobPane extends Pane {
        description: string
        progress: string
        createdAt: string
        finishedAt: Sig<string>
        error: Sig<string>
    }

    interface ModelsPane extends Pane {
        error: Sig<string>
        items: Sigs<Folder>
        hasItems: Sig<boolean>
    }

    interface CloudModelsPane extends ModelsPane {
        canBuildModel: boolean
        buildModel: Act
    }

    interface ModelBasePane extends Pane {
        // cloud: string
        algo: string
        frame: string
        responseColumn: string
        maxRunTime: string
        // javaModelPath: string
        createdAt: string
    }

    interface ModelPane extends ModelBasePane {
        deployModel: Act
        deleteModel: Act
    }

    interface CloudModelPane extends ModelBasePane {
        getModel: Act
    }

    interface ServicesPane extends Pane {
        error: Sig<string>
        items: Sigs<FolderI>
        hasItems: Sig<boolean>
    }

    interface ServicePane extends Pane {
        address: string
        state: string
        port: string
        url: string
        pid: string
        createdAt: string
        canStop: Sig<boolean>
        stopService: Act
    }

    interface AssetsPane extends Pane {
        items: Folder[]
    }

    interface EnginesPane extends Pane {
        error: Sig<string>
        items: Sigs<Folder>
        hasItems: Sig<boolean>
        addEngine: Act
    }

    interface EnginePane extends Pane {
        path: string
        createdAt: string
        deleteEngine: Act
    }

    //
    // Dialogs
    //

    function newRegisterCloudDialog(ctx: Context, go: Eff<RegisterCloudDialogResult>): RegisterCloudDialog {
        const
            error = sig<string>(''),
            address = sig<string>(void 0),
            addressError = lift(address, (address): string => address ? '' : 'Enter a valid IP:Port'),
            canRegisterCloud = lift(addressError, (e1): boolean => e1 === ''),
            registerCloud: Act = () => {
                if (!canRegisterCloud()) {
                    return
                }
                ctx.setBusy('Connecting to cluster...')
                ctx.remote.registerCluster(address(), (err, clusterId) => {
                    if (err) {
                        error(err.message)
                    } else {
                        go({ clusterId: clusterId })
                    }
                    ctx.setFree()
                })
            },
            cancel: Act = () => {
                go(null)
            }

        return {
            title: 'Connect to cluster',
            address: address,
            addressError: addressError,
            canRegisterCloud: canRegisterCloud,
            registerCloud: registerCloud,
            error: error,
            cancel: cancel,
            dispose: noop,
            template: 'register-cloud-dialog'
        }
    }

    const cloudIdPattern = /^[a-z0-9-]{1,16}$/i
    const cloudMemoryPattern = /^[0-9]+[kmg]$/i
    function newStartCloudDialog(ctx: Context, go: Eff<StartCloudDialogResult>): StartCloudDialog {

        const error = sig<string>('')

        const engines = sigs<Proxy.Engine>([])
        const engine = sig<Proxy.Engine>(void 0)
        const engineError = lift(engine, (engine): string =>
            engine
                ? ''
                : "Select a H2O version"
        )

        const cloudId = sig<string>('')
        const cloudIdError = lift(cloudId, (cloudId): string =>
            (cloudIdPattern.test(cloudId))
                ? ''
                : "Enter a valid cluster name"
        )

        const cloudSize = sig<string>('1')
        const cloudSizeNum = lift(cloudSize, (cloudSize): int =>
            parseInt(cloudSize, 10)
        )
        const cloudSizeError = lift(cloudSizeNum, (size): string =>
            (!isNaN(size) && size > 0)
                ? ''
                : "Invalid cluster size"
        )

        const cloudMemory = sig<string>('')
        const cloudMemoryError = lift(cloudMemory, (cloudMemory): string =>
            (cloudMemoryPattern.test(cloudMemory))
                ? ''
                : "Enter a valid Java memory specifier (e.g. 1024m, 2g, etc.)"

        )

        const canStartCloud = lift4(engineError, cloudIdError, cloudSizeError, cloudMemoryError, (e1, e2, e3, e4): boolean =>
            e1 === '' && e2 === '' && e3 === '' && e4 === ''
        )

        const startCloud: Act = () => {
            if (!canStartCloud()) {
                return
            }
            ctx.setBusy('Creating cluster...')
            ctx.remote.startYarnCluster(cloudId(), engine().id, cloudSizeNum(), cloudMemory(), ctx.principal.username, (err, clusterId) => {
                if (err) {
                    error(err.message)
                } else {
                    go({ clusterId: clusterId })
                }
                ctx.setFree()
            })
        }
        const cancel: Act = () => {
            go(null)
        }

        ctx.remote.getEngines((err, items) => {
            if (err) {
                return
            }
            engines(items)
        })

        return {
            title: 'Start a new cluster',
            engines: engines,
            engine: engine,
            engineError: engineError,
            cloudId: cloudId,
            cloudIdError: cloudIdError,
            cloudSize: cloudSize,
            cloudSizeError: cloudSizeError,
            cloudMemory: cloudMemory,
            cloudMemoryError: cloudMemoryError,
            canStartCloud: canStartCloud,
            startCloud: startCloud,
            error: error,
            cancel: cancel,
            dispose: noop,
            template: 'start-cloud-dialog'
        }
    }

    function newBuildModelDialog(ctx: Context, cloudId: int, go: Eff<BuildModelDialogResult>): BuildModelDialog {

        const error = sig<string>('')

        const frame = sig<string>('')
        const frameError = lift(frame, (f): string =>
            (f && f.trim().length > 0)
                ? ''
                : 'Enter a valid dataset path'
        )

        const responseColumn = sig<string>('')
        const responseColumnError = lift(responseColumn, (c): string =>
            (c && c.trim().length > 0)
                ? ''
                : 'Enter a valid column name'
        )

        const maxRunTime = sig<string>('1000')
        const maxRunTimeNum = lift(maxRunTime, (t) => parseInt(t, 10))
        const maxRunTimeError = lift(maxRunTimeNum, (t): string =>
            (!isNaN(t) && t > 0)
                ? ''
                : 'Invalid run time'
        )

        const canBuildModel = lift3(frameError, responseColumnError, maxRunTimeError, (e1, e2, e3): boolean =>
            e1 === '' && e2 === '' && e3 === ''
        )

        function buildModel(): void {
            ctx.setBusy('Building model...')
            ctx.remote.buildModel(cloudId, frame(), responseColumn(), maxRunTimeNum(), (err) => {
                if (err) {
                    error(err.message)
                } else {
                    go({ success: true })
                }
                ctx.setFree()
            })
        }

        const cancel: Act = () => {
            go(null)
        }

        return {
            title: 'Build a Model',
            frame: frame,
            frameError: frameError,
            responseColumn: responseColumn,
            responseColumnError: responseColumnError,
            maxRunTime: maxRunTime,
            maxRunTimeError: maxRunTimeError,
            canBuildModel: canBuildModel,
            buildModel: buildModel,
            error: error,
            cancel: cancel,
            dispose: noop,
            template: 'build-model-dialog'
        }
    }

    function newDeployModelDialog(ctx: Context, model: Proxy.Model, go: Eff<DeployModelDialogResult>): DeployModelDialog {

        const error = sig<string>('')

        const port = sig<string>('8000')
        const portNum = lift(port, (port): int =>
            parseInt(port, 10)
        )
        const portError = lift(portNum, (size): string =>
            (!isNaN(size) && size > 0)
                ? ''
                : 'Invalid port number'
        )

        const canDeployModel = lift(portError, (e): boolean =>
            e === ''
        )

        const deployModel: Act = () => {
            ctx.setBusy('Deploying model...')
            ctx.remote.startScoringService(model.id, portNum(), (err) => {
                if (err) {
                    error(err.message)
                } else {
                    go({ success: true })
                }
                ctx.setFree()
            })
        }
        const cancel: Act = () => {
            go(null)
        }

        return {
            title: `Deploy Model ${model.name}`,
            port: port,
            portError: portError,
            canDeployModel: canDeployModel,
            deployModel: deployModel,
            error: error,
            cancel: cancel,
            dispose: noop,
            template: 'deploy-model-dialog'
        }
    }

    function newAddEngineDialog(ctx: Context, go: Eff<AddEngineDialogResult>): AddEngineDialog {

        const error = sig<string>('')

        const form = sig<HTMLFormElement>(null)
        const file = sig<File>(null)

        const addEngine: Act = () => {
            const f = file()
            if (!(f && f.name)) {
                return
            }

            ctx.setBusy('Uploading asset...')
            const formData = new FormData(form())
            ctx.remote.upload(formData, (err, data) => {
                ctx.setFree()

                if (err) {
                    error(err.message)
                    return
                }

                go({ success: true })
            })

        }
        const cancel: Act = () => {
            go(null)
        }

        return {
            title: `Add Engine`,
            form: form,
            file: file,
            addEngine: addEngine,
            error: error,
            cancel: cancel,
            dispose: noop,
            template: 'add-engine-dialog'
        }
    }

    //
    // Panes
    //

    function newCloudsPane(ctx: Context): CloudsPane {
        const error = sig<string>('')
        const items = sigs<FolderI>([])
        const hasItems = lifts(items, (items) => items.length > 0)

        const registerCloud: Act = () => {
            const dialog = newRegisterCloudDialog(ctx, (result: RegisterCloudDialogResult) => {
                ctx.popDialog()
                if (result) {
                    ctx.showClouds()
                }
            })
            ctx.pushDialog(dialog)
        }

        const startCloud: Act = () => {
            const dialog = newStartCloudDialog(ctx, (result: StartCloudDialogResult) => {
                ctx.popDialog()
                if (result) {
                    ctx.showClouds()
                }
            })
            ctx.pushDialog(dialog)
        }

        ctx.remote.getClusters(0, 10000, (err, clouds) => {
            if (err) {
                error(err.message)
                return
            }
            items(_.map(clouds, (cloud): FolderI => {
                const slugI = sig<string>('NA')
                const isActive = sig<boolean>(cloud.state !== 'stopped')

                if (cloud.state !== 'stopped') {
                    // slugI(timestampToAge(cloud.activity))
                }
                return {
                    title: cloud.name,
                    subhead: 'State:',
                    slug: cloud.state,
                    execute: () => { ctx.showCloud(cloud) },
                    template: 'folderI',
                    isActive: isActive,
                    subheadI: 'Last Active:',
                    slugI: slugI
                }
            }))
        })
        return {
            title: 'Clusters',
            error: error,
            hasItems: hasItems,
            items: items,
            registerCloud: registerCloud,
            startCloud: startCloud,
            template: 'clouds',
            dispose: noop,
            position: newPanePosition()
        }
    }

    function newCloudPane(ctx: Context, cloud: Proxy.Cluster): CloudPane {
        const error = sig<string>('')
        const slugI = sig<string>('')
        const items: FolderI[] = [
            {
                title: 'Cluster Details',
                subhead: 'State:',
                slug: cloud.state,
                execute: () => { ctx.showCloudDetails(cloud) },
                template: 'folder',
                isActive: sig<boolean>(false),
                subheadI: '',
                slugI: slugI
            }
        ]

        if (cloud.state !== 'stopped') {
            const modelSlug = sig<string>('')
            const jobSlug = sig<string>('')
            const isActive = sig<boolean>(true)

            ctx.remote.getClusterModels(cloud.id, (err, models) => {
                if (err) {
                    alert(err.message)
                    return
                }
                modelSlug(String(models.length))
            })
            ctx.remote.getJobs(cloud.id, (err, jobs) => {
                if (err) {
                    alert(err.message)
                    return
                }
                jobSlug(String(jobs.length))
            })
            items.push({
                title: 'Jobs',
                subhead: 'Jobs in this cluster',
                slug: '',
                execute: () => { ctx.showCloudJobs(cloud) },
                template: 'folderI',
                isActive: isActive,
                subheadI: 'Number of Jobs:',
                slugI: jobSlug
            })

            items.push({
                title: 'Models',
                subhead: 'Models in this cluster',
                slug: '',
                template: 'folderI',
                execute: () => { ctx.showCloudModels(cloud) },
                isActive: isActive,
                subheadI: 'Number of models:',
                slugI: modelSlug
            })
        }

        return {
            title: cloud.name,
            error: error,
            items: items,
            template: 'cloud',
            dispose: noop,
            position: newPanePosition()
        }
    }

    function newCloudDetailsPane(ctx: Context, cloud: Proxy.Cluster): CloudDetailsPane {
        const error = sig<string>('')
        const state = sig<string>(cloud.state)
        const cloudDetails = sig<CloudDetail>(null)
        const yarnClusterDetails = sig<Proxy.YarnCluster>(null)
        const canStopCloud = !isExternalCluster(cloud)
        const canStop = sig<boolean>(cloud.state !== 'stopped')
        function stopCloud(): void {
            ctx.setBusy('Stopping cluster...')
            ctx.remote.stopYarnCluster(cloud.id, (err) => {
                ctx.setFree()
                if (err) {
                    error(err.message)
                    return
                }
                ctx.showClouds()
            })
        }

        function unregisterCloud(): void {
            ctx.setBusy('Disconnecting from cluster...')
            ctx.remote.unregisterCluster(cloud.id, (err) => {
                ctx.setFree()
                if (err) {
                    error(err.message)
                    return
                }
                ctx.showClouds()
            })
        }

        if (cloud.state != 'stopped') {
            ctx.remote.getClusterStatus(cloud.id, (err, h2oCloud) => {
                if (err) {
                    state('unknown')
                    error(err.message)
                    return
                }
                const cloudDetail = {
                    engineVersion: h2oCloud.version,
                    totalMemory: h2oCloud.max_memory,
                    totalCores: String(h2oCloud.total_cpu_count),
                    allowedCores: String(h2oCloud.total_allowed_cpu_count)
                }
                cloudDetails(cloudDetail)
                state(h2oCloud.status)
            })
        }

        if (!isExternalCluster(cloud)) {
            ctx.remote.getYarnCluster(cloud.id, (err, yarnCluster) => {
                if (err) {
                    error(err.message)
                    return
                }
                yarnClusterDetails(yarnCluster)
            })
        }
        return {
            title: 'Cluster Details',
            // engineName: cloud.engine_name,
            // size: String(cloud.size),
            // memory: cloud.memory,
            // applicationId: cloud.application_id,
            // username: cloud.username,
            address: `http://${cloud.address}/`,
            state: state,
            createdAt: timestampToAge(cloud.created_at),
            canStop: canStop,
            canStopCloud: canStopCloud,
            stopCloud: stopCloud,
            unregisterCloud: unregisterCloud,
            cloudDetails: cloudDetails,
            yarnCluster: yarnClusterDetails,
            template: 'cloudInfo',
            error: error,
            dispose: noop,
            position: newPanePosition(650)
        }
    }

    function newCloudJobsPane(ctx: Context, cloud: Proxy.Cluster): JobsPane {
        const error = sig<string>('')
        const items = sigs<Folder>([])
        const hasItems = lifts(items, (items) => items.length > 0)

        ctx.remote.getJobs(cloud.id, (err, jobs) => {
            if (err) {
                error(err.message)
                return
            }
            items(_.map(jobs, (job): Folder => {
                return {
                    title: job.name,
                    subhead: "Status",
                    slug: job.progress,
                    execute: () => { ctx.showCloudJob(job) },
                    template: 'folder'
                }
            }))
        })

        return {
            title: 'Cluster Jobs',
            error: error,
            items: items,
            hasItems: hasItems,
            template: 'cloudJobs',
            dispose: noop,
            position: newPanePosition(),
        }
    }

    function isExternalCluster(cloud: Proxy.Cluster): boolean {
        return cloud.detail_id === 0
    }

    function newCloudModelsPane(ctx: Context, cloud: Proxy.Cluster): CloudModelsPane {
        const error = sig<string>('')
        const items = sigs<Folder>([])
        const hasItems = lifts(items, (items) => items.length > 0)
        const canBuildModel = !isExternalCluster(cloud)
        function buildModel(): void {
            const dialog = newBuildModelDialog(ctx, cloud.id, (result: BuildModelDialogResult) => {
                ctx.popDialog()
                if (result) {
                    ctx.showModels()
                }
            })
            ctx.pushDialog(dialog)
        }
        ctx.remote.getClusterModels(cloud.id, (err, models) => {
            if (err) {
                error(err.message)
                return
            }
            items(_.map(models, (model): Folder => {
                return {
                    title: model.name,
                    subhead: model.dataset_name,
                    slug: model.response_column_name,
                    execute: () => { ctx.showCloudModel(cloud.id, model) },
                    template: 'folder'
                }
            }))
        })

        return {
            title: 'Cluster Models',
            error: error,
            items: items,
            hasItems: hasItems,
            template: 'cloudModels',
            canBuildModel: canBuildModel,
            buildModel: buildModel,
            dispose: noop,
            position: newPanePosition(),
        }
    }

    function newCloudJobPane(ctx: Context, job: Proxy.Job): JobPane {
        const error = sig<string>('')
        const finishedAt = sig<string>('')
        if (job.progress == "DONE") {
            finishedAt(timestampToAge(job.completed_at / 1000))
        }
        return {
            title: job.name,
            description: job.description,
            progress: job.progress,
            createdAt: timestampToAge(job.started_at / 1000),
            finishedAt: finishedAt,
            error: error,
            template: 'cloudJob',
            dispose: noop,
            position: newPanePosition(650),
        }
    }

    function newCloudModelPane(ctx: Context, clusterId: int, model: Proxy.Model): CloudModelPane {
        const getModel: Act = () => {
            ctx.setBusy('Getting model from h2o...')
            ctx.remote.importModelFromCluster(clusterId, model.name, (err, model) => {
                ctx.setFree()
                if (err) {
                    alert(err.message)
                    return
                }

                ctx.showModels()
            })
        }

        return {
            title: model.name,
            // cloud: model.cloud_name,
            algo: model.algorithm,
            frame: model.dataset_name,
            responseColumn: model.response_column_name,
            maxRunTime: String(model.max_runtime),
            // javaModelPath: model.java_model_path,
            createdAt: timestampToAge(model.created_at / 1000),
            getModel: getModel,
            template: 'cloudModel',
            dispose: noop,
            position: newPanePosition(650)
        }
    }

    function newModelsPane(ctx: Context): ModelsPane {
        const error = sig<string>('')
        const items = sigs<Folder>([])
        const hasItems = lifts(items, (items) => items.length > 0)
        ctx.remote.getModels(0, 10000, (err, models) => {
            if (err) {
                error(err.message)
                return
            }
            items(_.map(models, (model): Folder => {
                return {
                    title: model.name,
                    subhead: model.dataset_name,
                    slug: model.response_column_name,
                    execute: () => { ctx.showModel(model) },
                    template: 'folder'
                }
            }))
        })

        return {
            title: 'Models',
            error: error,
            items: items,
            hasItems: hasItems,
            template: 'models',
            dispose: noop,
            position: newPanePosition(),
        }
    }

    function newModelPane(ctx: Context, model: Proxy.Model): ModelPane {
        const deployModel: Act = () => {
            const dialog = newDeployModelDialog(ctx, model, (result: DeployModelDialogResult) => {
                ctx.popDialog()
                if (result) {
                    ctx.showServices()
                }
            })
            ctx.pushDialog(dialog)
        }
        const deleteModel: Act = () => {
            ctx.setBusy('Deleting model...')
            ctx.remote.deleteModel(model.id, (err) => {
                ctx.setFree()
                if (err) {
                    alert(err.message) // FIXME
                    return
                }
                ctx.showModels()
            })
        }
        return {
            title: model.name,
            //cloud: model.cloud_name,
            algo: model.algorithm,
            frame: model.dataset_name,
            responseColumn: model.response_column_name,
            maxRunTime: String(model.max_runtime),
            // javaModelPath: model.java_model_path,
            createdAt: timestampToAge(model.created_at),
            deployModel: deployModel,
            deleteModel: deleteModel,
            template: 'model',
            dispose: noop,
            position: newPanePosition(650)
        }
    }

    function newServicesPane(ctx: Context): ServicesPane {
        const error = sig<string>('')
        const items = sigs<FolderI>([])
        const hasItems = lifts(items, (items) => items.length > 0)
        ctx.remote.getScoringServices(0, 10000, (err, services) => {
            if (err) {
                error(err.message)
                return
            }
            items(_.map(services, (service): FolderI => {
                const slugI = sig<string>('')
                const isActive = sig<boolean>(service.state !== 'stopped')

                if (service.state !== 'stopped') {
                    // slugI(timestampToAge(service.activity))
                }
                return {
                    title: String(service.model_id), // FIXME
                    subhead: 'State:',
                    slug: service.state,
                    execute: () => { ctx.showService(service) },
                    template: 'folderI',
                    isActive: isActive,
                    subheadI: 'Last Activity:',
                    slugI: slugI
                }
            }))
        })
        return {
            title: 'Services',
            error: error,
            hasItems: hasItems,
            items: items,
            template: 'services',
            dispose: noop,
            position: newPanePosition(),
        }
    }

    function newServicePane(ctx: Context, service: Proxy.ScoringService): ServicePane {
        const canStop = sig<boolean>(service.state !== 'stopped')
        const stopService: Act = () => {
            ctx.setBusy('Stopping service...')
            ctx.remote.stopScoringService(service.id, (err) => {
                ctx.setFree()
                if (err) {
                    alert(err.message)
                    return
                }
                ctx.showServices()
            })
        }
        return {
            title: String(service.model_id), // FIXME
            state: service.state,
            address: service.address,
            port: String(service.port),
            url: `http://${service.address}:${service.port}/`,
            pid: String(service.process_id),
            createdAt: timestampToAge(service.created_at),
            canStop: canStop,
            stopService: stopService,
            template: 'service',
            dispose: noop,
            position: newPanePosition(650),
        }
    }

    function newAssetsPane(ctx: Context): AssetsPane {
        const items: Folder[] = [
            {
                title: 'Engines',
                subhead: 'View deployed engines',
                slug: '',
                execute: () => { ctx.showEngines() },
                template: 'folder'
            }
        ]

        return {
            title: 'Assets',
            template: 'assets',
            dispose: noop,
            position: newPanePosition(),
            items: items
        }
    }

    function newEnginesPane(ctx: Context): EnginesPane {
        const error = sig<string>('')
        const items = sigs<Folder>([])
        const hasItems = lifts(items, (items) => items.length > 0)
        const addEngine: Act = () => {
            const dialog = newAddEngineDialog(ctx, (result: AddEngineDialogResult) => {
                ctx.popDialog()

                if (result) {
                    if (result.success) {
                        ctx.showEngines()
                    }
                }
            })
            ctx.pushDialog(dialog)
        }

        ctx.remote.getEngines((err, engines) => {
            if (err) {
                error(err.message)
                return
            }
            items(_.map(engines, (engine): Folder => {
                return {
                    title: engine.name,
                    subhead: timestampToAge(engine.created_at),
                    slug: '',
                    execute: () => { ctx.showEngine(engine) },
                    template: 'folder'
                }
            }))
        })

        return {
            title: 'Engines',
            error: error,
            items: items,
            hasItems: hasItems,
            addEngine: addEngine,
            template: 'engines',
            dispose: noop,
            position: newPanePosition()
        }
    }

    function newEnginePane(ctx: Context, engine: Proxy.Engine): EnginePane {
        const deleteEngine: Act = () => {
            ctx.setBusy('Deleting engine...')
            ctx.remote.deleteEngine(engine.id, (err) => {
                ctx.setFree()
                if (err) {
                    alert(err.message) // FIXME
                    return
                }
                ctx.showEngines()
            })
        }
        return {
            title: engine.name,
            path: engine.location,
            createdAt: timestampToAge(engine.created_at),
            deleteEngine: deleteEngine,
            template: 'engine',
            dispose: noop,
            position: newPanePosition(650),
        }
    }

    interface Principal {
        username: string
    }

    export class Context {
        public remote = Proxy
        public principal = { username: 'unknown' }
        public setBusy = uni1<string>()
        public setFree = uni()
        public pushDialog = uni1<Dialog>()
        public popDialog = uni()
        public showPane = uni2<int, Pane>()
        public showClouds = uni()
        public showCloud = uni1<Proxy.Cluster>()
        public showCloudDetails = uni1<Proxy.Cluster>()
        public showCloudJobs = uni1<Proxy.Cluster>()
        public showCloudModels = uni1<Proxy.Cluster>()
        public showCloudJob = uni1<Proxy.Job>()
        public showCloudModel = uni2<int, Proxy.Model>()
        public showModels = uni()
        public showModel = uni1<Proxy.Model>()
        public showServices = uni()
        public showService = uni1<Proxy.ScoringService>()
        public showAssets = uni()
        public showEngines = uni()
        public showEngine = uni1<Proxy.Engine>()
    }

    class Breadcrumb {
        constructor(
            public title: string,
            public execute: Act
        ) { }
    }

    export interface App {
        context: Context
        navBar: NavBar
        breadcrumbs: Sigs<Breadcrumb>
        panes: Sigs<Pane>
        span: Sig<string>
        hasDialogs: Sig<boolean>
        dialogs: Sigs<Dialog>
        busyMessage: Sig<string>
        templateOf: (t: Templated) => string
        afterRender: Eff<HTMLElement[]>
    }

    export function newApp(): App {
        const ctx = new Context()
        const navBar = newNavBar(ctx)
        const breadcrumbs = sigs<Breadcrumb>([])
        const panes = sigs<Pane>([])
        const span = sig<int>(0)
        const spanPx = lift(span, px)

        const dialogs = sigs<Dialog>([])
        const hasDialogs = lifts(dialogs, isNonEmpty)
        const busyMessage = sig<string>(void 0)

        ctx.pushDialog.on((dialog: Dialog) => {
            dialogs.push(dialog)
        })

        ctx.popDialog.on(() => {
            dialogs.pop()
        })

        ctx.setBusy.on((message: string) => {
            busyMessage(message)
        })
        ctx.setFree.on(() => {
            busyMessage(void 0)
        })

        ctx.showPane.on((index: int, pane: Pane) => {
            const disposables = panes.splice(index, panes().length - index, pane)
            for (const disposable of disposables) {
                disposable.dispose()
            }
            let left = 0
            for (const p of panes()) {
                const pos = p.position
                pos.left(left)
                left += pos.width()
            }

            // Set span to max total width so that browsing panes leftward
            //  does not result in a jerky scroll to the right.
            if (span() < left) {
                span(left)
            }
            pane.position.ensureVisible()

            breadcrumbs(_.map(panes(), (pane: Pane): Breadcrumb => {
                return new Breadcrumb(
                    pane.title,
                    () => {
                        pane.position.ensureVisible()
                    }
                )
            }))
        })

        ctx.showClouds.on(() => {
            ctx.showPane(0, newCloudsPane(ctx))
        })

        ctx.showCloud.on((cloud: Proxy.Cluster) => {
            ctx.showPane(1, newCloudPane(ctx, cloud))
        })

        ctx.showCloudDetails.on((cloud: Proxy.Cluster) => {
            ctx.showPane(2, newCloudDetailsPane(ctx, cloud))
        })

        ctx.showCloudModels.on((cloud: Proxy.Cluster) => {
            ctx.showPane(2, newCloudModelsPane(ctx, cloud))
        })

        ctx.showCloudJobs.on((cloud: Proxy.Cluster) => {
            ctx.showPane(2, newCloudJobsPane(ctx, cloud))
        })

        ctx.showCloudModel.on((clusterId: int, model: Proxy.Model) => {
            ctx.showPane(3, newCloudModelPane(ctx, clusterId, model))
        })

        ctx.showCloudJob.on((job: Proxy.Job) => {
            ctx.showPane(3, newCloudJobPane(ctx, job))
        })

        ctx.showModels.on(() => {
            ctx.showPane(0, newModelsPane(ctx))
        })

        ctx.showModel.on((model: Proxy.Model) => {
            ctx.showPane(1, newModelPane(ctx, model))
        })

        ctx.showServices.on(() => {
            ctx.showPane(0, newServicesPane(ctx))
        })

        ctx.showService.on((service: Proxy.ScoringService) => {
            ctx.showPane(1, newServicePane(ctx, service))
        })

        ctx.showAssets.on(() => {
            ctx.showPane(0, newAssetsPane(ctx))
        })

        ctx.showEngines.on(() => {
            ctx.showPane(1, newEnginesPane(ctx))
        })

        ctx.showEngine.on((engine: Proxy.Engine) => {
            ctx.showPane(2, newEnginePane(ctx, engine))
        })

        ctx.showClouds()

        return {
            context: ctx,
            navBar: navBar,
            breadcrumbs: breadcrumbs,
            panes: panes,
            span: spanPx,
            hasDialogs: hasDialogs,
            dialogs: dialogs,
            busyMessage: busyMessage,
            templateOf: templateOf,
            afterRender: doAfterRender
        }
    }
}
