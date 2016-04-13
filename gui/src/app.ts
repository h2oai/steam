/// <reference path="references.ts" />

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


    //
    // Knockout Extensions
    //
    ko.bindingHandlers['element'] = {
        update: (element, valueAccessor, allBindings, viewModel, bindingContext) => {
            const arg = ko.unwrap(valueAccessor())
            if (arg) {
                const $element = $(element)
                $element.empty()
                $element.append(arg)
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
    // Mockers (temporary)
    // 

    const permissionTypes: string[] = [
        'ManageGroup',
        'ViewGroup',
        'ManageRole',
        'ViewRole',
        'ManageUser',
        'ViewUser',
        'ManageCloud',
        'ViewCloud',
        'ViewJob',
        'ManageProject',
        'ViewProject',
        'ManageModel',
        'ViewModel',
        'ManageScript',
        'ViewScript',
        'RunScript',
        'ManageExperiment',
        'ViewExperiment',
        'RunExperiment'
    ]

    const useCases: string[] = [
        'Wallet Share Estimation',
        'Churn',
        'Customer Segmentation',
        'Product Mix',
        'Cross Selling',
        'Up Selling',
        'Channel Optimization',
        'Discount Targeting',
        'Reactivation Likelihood',
        'Ad Optimization',
        'Lead Prioritization',
        'Demand Forecasting',
        'Credit Risk',
        'Fraud Detection',
        'Accounts Payable Recovery',
        'Anti Money Laundering',
        'Message Optimization',
        'Volume Forecasting',
        'Resume Screening',
        'Employee Churn',
        'Training Recommendation',
        'Talent Management',
        'Claims Prioritization',
        'Medicaid Fraud',
        'Prescription Compliance',
        'Physician Attrition',
        'Survival Analysis',
        'Dosage Effectiveness',
        'Readmission Risk',
        'Credit Card Fraud',
        'Claims Prediction',
        'Demand Forecasting'
    ]

    function randomUseCase(): string {
        return _.sample<string>(useCases)
    }
    function randomUseCaseSlug(): string {
        return randomUseCase().toLowerCase().replace(/\s+/g, '-')
    }

    function proxy_startCloud(cloudName: string, cloudSize: int, on: On<string>): void {
        function go() {
            on(null, `app_id_${(Math.random() * 100) | 0}`)
        }
        setTimeout(go, 3000)
    }

    function proxy_deployModel(modelId: string, port: int, on: On<string>): void {
        function go() {
            on(null, '')
        }
        setTimeout(go, 3000)
    }

    function proxy_uploadFile(file: string, on: On<string>): void {
        function go() {
            on(null, '')
        }
        setTimeout(go, 3000)
    }

    function proxy_getClouds(on: On<Cloud[]>): void {
        const clouds = _.times(_.random(5, 20), (i: int): Cloud => {
            return new Cloud(
                `${faker.name.firstName()}'s Cloud`,
                _.random(1, 10)
            )
        })
        on(null, clouds)
    }

    function proxy_getModels(on: On<Model[]>): void {
        const clouds = _.times(_.random(5, 20), (i: int): Model => {
            const useCase = randomUseCaseSlug()
            return new Model(
                `model-${useCase}`,
                `${faker.name.firstName()}'s Cloud`,
                `${useCase}.hex`,
                `${faker.lorem.word()}`
            )
        })
        on(null, clouds)
    }

    function proxy_getServices(on: On<Service[]>): void {
        const services = _.times(_.random(5, 20), (i: int): Service => {
            return new Service(
                `${randomUseCaseSlug()}`,
                `${faker.internet.ip()}/${randomUseCaseSlug()}`
            )
        })
        on(null, services)
    }

    function proxy_getEngines(on: On<Engine[]>): void {
        const engines = _.times(_.random(5, 20), (i: int): Engine => {
            const name = `h2o-3.4.2.${i}`
            return {
                name: name,
                path: `/var/engines/${name}`,
                createdDate: String(faker.date.recent())
            }
        })
        on(null, engines)
    }


    //
    // Models
    //

    class Cloud {
        constructor(
            public id: string,
            public size: int
        ) { }
    }

    class Model {
        constructor(
            public id: string,
            public cloud: string,
            public frame: string,
            public responseColumn: string
        ) { }
    }

    class Service {
        constructor(
            public id: string,
            public endpoint: string
        ) { }
    }

    interface Engine {
        name: string
        path: string
        createdDate: string
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
            newNavButton('ion-ios-cloud-outline', 'Clouds', true, ctx.showClouds),
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

    interface CloudsPane extends Pane {
        items: Sigs<Folder>
        startCloud: Act
    }


    function doAfterRender(elements: HTMLElement[]): void {
        $(elements).click(function() {
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

    interface Dialog extends Templated {
        title: string
        cancel: Act
        dispose: Act
        template: string
    }

    interface StartCloudDialog extends Dialog {
        cloudName: Sig<string>
        cloudNameError: Sig<string>
        cloudSize: Sig<string>
        cloudSizeError: Sig<string>
        canStartCloud: Sig<boolean>
        startCloud: Act
        error: Sig<string>
    }

    interface StartCloudDialogResult {
        applicationId: string
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
        file: Sig<string>
        fileError: Sig<string>
        canAddEngine: Sig<boolean>
        addEngine: Act
        error: Sig<string>
    }
    interface AddEngineDialogResult {
        success: boolean // FIXME
    }

    interface CloudPane extends Pane {
        size: string
        stopCloud: Act
    }

    interface ModelsPane extends Pane {
        items: Sigs<Folder>
    }

    interface ModelPane extends Pane {
        cloud: string
        frame: string
        responseColumn: string
        deployModel: Act
        deleteModel: Act
    }

    interface ServicesPane extends Pane {
        items: Sigs<Folder>
    }

    interface ServicePane extends Pane {
        endpoint: string
        stopService: Act
    }

    interface AssetsPane extends Pane {
        items: Folder[]
    }

    interface EnginesPane extends Pane {
        items: Sigs<Folder>
        addEngine: Act
    }

    interface EnginePane extends Pane {
        path: string
        createdDate: string
        deleteEngine: Act
    }

    //
    // Dialogs
    //


    const cloudNamePattern = /^[a-z0-9-]{1,16}$/i
    function newStartCloudDialog(ctx: Context, go: Eff<StartCloudDialogResult>): StartCloudDialog {

        const error = sig<string>('')

        const cloudName = sig<string>('')
        const cloudNameError = lift(cloudName, (cloudName): string =>
            (cloudNamePattern.test(cloudName))
                ? ''
                : "Enter a valid cloud name"
        )

        const cloudSize = sig<string>('1')
        const cloudSizeNum = lift(cloudSize, (cloudSize): int =>
            parseInt(cloudSize, 10)
        )
        const cloudSizeError = lift(cloudSizeNum, (size): string =>
            (!isNaN(size) && size > 0)
                ? ''
                : "Invalid cloud size"
        )

        const canStartCloud = lift2(cloudNameError, cloudSizeError, (e1, e2): boolean =>
            e1 === '' && e2 === ''
        )

        const startCloud: Act = () => {
            ctx.setBusy('Creating cloud...')
            proxy_startCloud(cloudName(), cloudSizeNum(), (err, applicationId) => {
                if (err) {
                    error(err.message)
                } else {
                    go({
                        applicationId: applicationId
                    })
                }
                ctx.setFree()
            })
        }
        const cancel: Act = () => {
            go(null)
        }

        return {
            title: 'Start a new cloud',
            cloudName: cloudName,
            cloudNameError: cloudNameError,
            cloudSize: cloudSize,
            cloudSizeError: cloudSizeError,
            canStartCloud: canStartCloud,
            startCloud: startCloud,
            error: error,
            cancel: cancel,
            dispose: noop,
            template: 'start-cloud-dialog'
        }
    }

    function newDeployModelDialog(ctx: Context, modelId: string, go: Eff<DeployModelDialogResult>): DeployModelDialog {

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
            proxy_deployModel(modelId, portNum(), (err) => {
                if (err) {
                    error(err.message)
                } else {
                    go({
                        success: true
                    })
                }
                ctx.setFree()
            })
        }
        const cancel: Act = () => {
            go(null)
        }

        return {
            title: `Deploy Model ${modelId}`,
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

        const file = sig<string>('')
        const fileError = lift(file, (file): string =>
            (file && file.trim() !== '')
                ? ''
                : 'Invalid file name'
        )

        const canAddEngine = lift(file, (e): boolean =>
            e === ''
        )

        const addEngine: Act = () => {
            ctx.setBusy('Uploading asset...')
            proxy_uploadFile(file(), (err) => {
                if (err) {
                    error(err.message)
                } else {
                    go({
                        success: true
                    })
                }
                ctx.setFree()
            })
        }
        const cancel: Act = () => {
            go(null)
        }

        return {
            title: `Add Engine`,
            file: file,
            fileError: fileError,
            canAddEngine: canAddEngine,
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

    function newCloudsPane(ctx: Context, clouds: Cloud[]): CloudsPane {
        const items = sigs<Folder>(_.map(clouds, (cloud): Folder => {
            return {
                title: cloud.id,
                subhead: 'Size:',
                slug: String(cloud.size),
                execute: () => { ctx.showCloud(cloud) },
                template: 'folder'
            }
        }))
        const startCloud: Act = () => {
            const dialog = newStartCloudDialog(ctx, (result: StartCloudDialogResult) => {
                // XXX use result to update cloud list
                ctx.popDialog()
            })
            ctx.pushDialog(dialog)
        }
        return {
            title: 'Clouds',
            template: 'clouds',
            dispose: noop,
            position: newPanePosition(),
            items: items,
            startCloud: startCloud,
        }
    }

    function newCloudPane(ctx: Context, cloud: Cloud): CloudPane {
        const stopCloud: Act = () => {
            alert('--- Stop Cloud ---')
        }
        return {
            title: cloud.id,
            size: String(cloud.size),
            stopCloud: stopCloud,
            template: 'cloud',
            dispose: noop,
            position: newPanePosition(650)
        }
    }

    function newModelsPane(ctx: Context, models: Model[]): ModelsPane {
        const items = sigs<Folder>(_.map(models, (model): Folder => {
            return {
                title: model.id,
                subhead: model.frame,
                slug: model.responseColumn,
                execute: () => { ctx.showModel(model) },
                template: 'folder'
            }
        }))
        return {
            title: 'Models',
            items: items,
            template: 'models',
            dispose: noop,
            position: newPanePosition(),
        }
    }

    function newModelPane(ctx: Context, model: Model): ModelPane {
        const deployModel: Act = () => {
            const dialog = newDeployModelDialog(ctx, model.id, (result: DeployModelDialogResult) => {
                // XXX use result to update services list
                ctx.popDialog()
            })
            ctx.pushDialog(dialog)
        }
        const deleteModel: Act = () => {
            alert('--- Delete selected models ---')
        }
        return {
            title: model.id,
            cloud: model.cloud,
            frame: model.frame,
            responseColumn: model.responseColumn,
            deployModel: deployModel,
            deleteModel: deleteModel,
            template: 'model',
            dispose: noop,
            position: newPanePosition(650)
        }
    }

    function newServicesPane(ctx: Context, services: Service[]): ServicesPane {
        const items = sigs<Folder>(_.map(services, (service): Folder => {
            return {
                title: service.id,
                subhead: service.endpoint,
                slug: '',
                execute: () => { ctx.showService(service) },
                template: 'folder'
            }
        }))
        return {
            title: 'Services',
            template: 'services',
            dispose: noop,
            position: newPanePosition(),
            items: items,
        }
    }

    function newServicePane(ctx: Context, service: Service): ServicePane {
        const stopService: Act = () => {
            alert('--- Deploy Service ---')
        }
        return {
            title: service.id,
            template: 'service',
            dispose: noop,
            position: newPanePosition(650),
            endpoint: service.endpoint,
            stopService: stopService,
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

    function newEnginesPane(ctx: Context, engines: Engine[]): EnginesPane {
        const items = sigs<Folder>(_.map(engines, (engine): Folder => {
            return {
                title: engine.name,
                subhead: engine.createdDate,
                slug: '',
                execute: () => { ctx.showEngine(engine) },
                template: 'folder'
            }
        }))

        const addEngine: Act = () => {
            const dialog = newAddEngineDialog(ctx, (result: AddEngineDialogResult) => {
                // XXX use result to update engine list
                ctx.popDialog()
            })
            ctx.pushDialog(dialog)
        }

        return {
            title: 'Engines',
            items: items,
            addEngine: addEngine,
            template: 'engines',
            dispose: noop,
            position: newPanePosition()
        }
    }

    function newEnginePane(ctx: Context, engine: Engine): EnginePane {
        const deleteEngine: Act = () => {
            alert('--- Delete Engine ---')
        }
        return {
            title: engine.name,
            path: engine.path,
            createdDate: engine.createdDate,
            deleteEngine: deleteEngine,
            template: 'engine',
            dispose: noop,
            position: newPanePosition(650),
        }
    }

    export class Context {
        public setBusy = uni1<string>()
        public setFree = uni()
        public pushDialog = uni1<Dialog>()
        public popDialog = uni()
        public showPane = uni2<int, Pane>()
        public showClouds = uni()
        public showCloud = uni1<Cloud>()
        public showModels = uni()
        public showModel = uni1<Model>()
        public showServices = uni()
        public showService = uni1<Service>()
        public showAssets = uni()
        public showEngines = uni()
        public showEngine = uni1<Engine>()
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
            proxy_getClouds((err, clouds) => {
                ctx.showPane(0, newCloudsPane(ctx, clouds))
            })
        })

        ctx.showCloud.on((cloud: Cloud) => {
            ctx.showPane(1, newCloudPane(ctx, cloud))
        })

        ctx.showModels.on(() => {
            proxy_getModels((err, models) => {
                ctx.showPane(0, newModelsPane(ctx, models))
            })
        })

        ctx.showModel.on((model: Model) => {
            ctx.showPane(1, newModelPane(ctx, model))
        })

        ctx.showServices.on(() => {
            proxy_getServices((err, services) => {
                ctx.showPane(0, newServicesPane(ctx, services))
            })
        })

        ctx.showService.on((service: Service) => {
            ctx.showPane(1, newServicePane(ctx, service))
        })

        ctx.showAssets.on(() => {
            ctx.showPane(0, newAssetsPane(ctx))
        })

        ctx.showEngines.on(() => {
            proxy_getEngines((err, engines) => {
                ctx.showPane(1, newEnginesPane(ctx, engines))
            })
        })

        ctx.showEngine.on((engine: Engine) => {
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
