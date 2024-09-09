import {Factory} from "@/services/context-service-provider/service-provider-context";

export class ServiceProvider {
    // storing singletons here. I tried with useMemo and useCallback and it's still lbeing called
    // multiple times. Must read more about them,
    private services: { [key: string]: any } = {};

    constructor() {
        console.log("services created")
    }

    public Get<T>(factory: Factory<T>): T {

        const factoryInstance = new factory(this);
        const key = factoryInstance.Key();

        if (!this.services[key]) {
            this.services[key] = factoryInstance.Build()
        }

        return this.services[factoryInstance.Key()];
    }
}