export class User {
    private readonly _firstName: string;
    private readonly _lastName: string;

    public get FirstName() {
        return this._firstName;
    }

    public get LastName() {
        return this._lastName;
    }

    constructor(firstName: string, lastName: string) {
        this._firstName = firstName;
        this._lastName = lastName;
    }
}