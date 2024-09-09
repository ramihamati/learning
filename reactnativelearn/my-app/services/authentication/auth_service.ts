import {User} from "@/services/authentication/user";

export class AuthService{

    public GetUser() : Promise<User>{
        return Promise.resolve(new User(
            "Rami",
            "Hamati"
        ));
    }
}