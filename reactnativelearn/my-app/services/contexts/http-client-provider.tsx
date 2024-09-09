import React, {createContext, useContext} from 'react';
import {useWeatherClient} from "@/services/weather/weather-client-hook";
import {WeatherClientConfig} from "@/services/weather/weather-client-config";
import {WeatherClient} from "@/services/weather/weather-client";

const HttpClientContext = createContext<WeatherClient>(undefined!);

const config = new WeatherClientConfig("https://api.open-meteo.com")

// @ts-ignore
export const HttpClientProvider = ({ children }) => {
    const client = useWeatherClient(config);

    return (
        <HttpClientContext.Provider value={client}>
            {children}
        </HttpClientContext.Provider>
    );
};

export const useHttpClientContext = () => {
    return useContext(HttpClientContext);
};
