export interface User {
    id: number;
    username: string;
    password_hash: string;
    created_at: Date;
    last_login: Date;
};

export interface Login {
    username: string;
    password: string;
}