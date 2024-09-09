export class WeatherClientConfig {
    private readonly _BaseUri: string;

    get BaseUri(): string {
        return this._BaseUri;
    }


    constructor(baseUri: string ) {
        this._BaseUri = baseUri;
    }
}