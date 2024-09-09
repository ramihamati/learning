import {useMemo} from 'react';
import {WeatherClientConfig} from "@/services/weather/weather-client-config";
import weatherClientBuilder from "@/services/weather/weather-client";


export const useWeatherClient = (config: WeatherClientConfig) => {
    const weatherClient = useMemo(() => {
        return weatherClientBuilder(config);
    }, []);

    return weatherClient;
};