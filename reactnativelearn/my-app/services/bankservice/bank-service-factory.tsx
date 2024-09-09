import {IBuildable} from "@/services/context-service-provider/service-provider-context";
import {ServiceProvider} from "@/services/context-service-provider/service-provider";
import {BankService} from "@/services/bankservice/bank-service";

export class BankServiceFactory implements IBuildable<BankService> {
    constructor(private serviceProvider: ServiceProvider) {
    }

    public Build(): BankService {
        return new BankService();
    }

    public Key(): string {
        return "bank-service";
    }
}