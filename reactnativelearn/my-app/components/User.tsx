import {TouchableOpacity, View, Text} from "react-native";
import {useAuthService} from "@/hooks/useAuthService";
import {useServicesContext} from "@/services/context-service-provider/service-provider-context";
import {BankServiceFactory} from "@/services/bankservice/bank-service-factory";

export function User() {

    var serviceProvider = useServicesContext();
    const bankService = serviceProvider.Get(BankServiceFactory);
    const bankService1 = serviceProvider.Get(BankServiceFactory);
    const bankService2 = serviceProvider.Get(BankServiceFactory);
    const bankService3 = serviceProvider.Get(BankServiceFactory);

    return (

      <View>
          <TouchableOpacity style={{backgroundColor : 'cyan', padding : 10 }}>
              <Text>
                  First Name : {useAuthService()?.FirstName ?? "loading"}
              </Text>
              <Text>
                  Last Name : {useAuthService()?.FirstName ?? "loading"}
              </Text>
              <Text>
                  Amount : {bankService.Amount()}
              </Text>
          </TouchableOpacity>
      </View>
    );
}