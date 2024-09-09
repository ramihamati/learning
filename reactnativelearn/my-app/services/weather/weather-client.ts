import axios, {AxiosInstance} from 'axios';
import {WeatherClientConfig} from "@/services/weather/weather-client-config";



export class CurrentWeather {
    public Time: string = "";
    public Temperature_2m: string = "";
}

export class WeatherResponse {
    public Latitude: number = 0;
    public Longitude: number = 0;
    public Elevation: number = 0;
    public Current: CurrentWeather = new CurrentWeather();
}

export class WeatherClient {
    constructor(private httpClient: AxiosInstance) {
        console.log("weather client created")
    }

    public async GetWeather(): Promise<WeatherResponse> {
        const response = await this.httpClient
            .get<any>("v1/forecast?latitude=52.52&longitude=13.41&current=temperature_2m");
        let a  = new WeatherResponse();
        a.Longitude = response.data.longitude;
        a.Latitude = response.data.latitude;
        a.Elevation = response.data.elevation;
        a.Current = new CurrentWeather();

        a.Current.Time =  response.data.current.time;
        a.Current.Temperature_2m = response.data.current.temperature_2m;
        return a;
    }
}

const httpClientBuilder = (config : WeatherClientConfig) => axios.create({
    baseURL: config.BaseUri,
    headers: {
        'Content-Type': 'application/json',
        // Add any other headers here
    },
});

const weatherClientBuilder = (config : WeatherClientConfig) => {
    return new WeatherClient(httpClientBuilder(config));
}

export default weatherClientBuilder;