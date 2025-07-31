
export class AppConfig {
    private runtime: any

    public constructor(runtime: any) {
        this.runtime = runtime
    }

    get version(): string {
        return this.runtime.public.version
    }

    get apiPath(): string {
        return this.runtime.public.apiPath
    }
}

export function useAppConfig(): AppConfig {
    return new AppConfig(useRuntimeConfig())
}