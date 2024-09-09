import {useHttpClientContext} from "@/services/contexts/http-client-provider";
import {useCallback, useEffect, useState} from "react";
import {WeatherResponse} from "@/services/weather/weather-client";
import {View, Text, Button} from 'react-native';

export function Weather() {

    const httpClient = useHttpClientContext();
    const httpClient1 = useHttpClientContext();
    const httpClient2 = useHttpClientContext(); // the weather client is created only once and passed multiple times. this is a global state
    let state: WeatherResponse | null = null;
    const [data, setData] = useState<WeatherResponse | null>(state)
    const [secondData, setSecondData] = useState<WeatherResponse | null>(state)

    useEffect(() => {
        try {
            const response =  httpClient.GetWeather();
            response
                .then(r  => {
                    setData(r); console.log(r);
                })
                .catch(err=> console.log(err));
        } catch (error) {
            console.error('Error fetching data:', error);
        }
    }, [httpClient]); // httpClient is a reactive value. If not provided the component will re-render on each render cycle

    const getWeather = useCallback(async () => { // this should be used as a post attached to a button. It caches the function for performance
        console.log("pressed again")
        const response = httpClient.GetWeather();
        try {
            const r = await response;
            setSecondData(r);
            console.log(r);
            return r;
        } catch (err) {
            console.log(err);
            return null;
        }
    }, [httpClient]); // httpClient is a reactive value. If not provided the component will re-render on each render cycle

    return (
        <View>
            <Text>{data?.Current?.Temperature_2m ?? "none"}</Text>
            <View>
                <Button onPress={getWeather} title="Get Weather"></Button>
                <Text>{secondData?.Current?.Temperature_2m ?? "none"}</Text>
            </View>
        </View>

    );
}