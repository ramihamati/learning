import {useEffect, useState} from "react";
import {User} from "@/services/authentication/user";
import {AuthService} from "@/services/authentication/auth_service";


// an example of a hook calling a method. We could use hooks to also return services
export function useAuthService() {
    let [user, setUser] = useState<User>();
    console.log("called useAuthService")

    useEffect(() => {
        new AuthService().GetUser()
            .then((user) => {
                console.log("called auth service")
                setUser(user);
            })
            .catch((error) => {
                // TODO: do something with the error
            });
    }, [1])

    return user;
}

