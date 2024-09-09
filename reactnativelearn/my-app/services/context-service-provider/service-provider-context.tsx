import React, {createContext, useContext} from "react";
import {ServiceProvider} from "@/services/context-service-provider/service-provider";

const ServicesContext = createContext<ServiceProvider>(undefined!);

export interface IBuildable<T>{
    Build() : T;
    Key() : string;
}

export interface Factory<T>{
    new(n: ServiceProvider) : IBuildable<T>;
}

// @ts-ignore
export const ServicesProvider = ({ children }) => {
    // const client = useWeatherClient(new ServiceProvider());

    return (
        <ServicesContext.Provider value={new ServiceProvider()}>
            {children}
        </ServicesContext.Provider>
    );
};

export const useServicesContext = () => {
    return useContext(ServicesContext);
};
