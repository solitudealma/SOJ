export interface UserInfo {
    userId: number;
    username: string;
    avatar: string;
}

export interface LoginRequest {
    username: string;
    password: string;
}

export interface LoginReponse {
    token: string;
    userInfo: UserInfo;
}

export interface RegisterRequest {
    username: string;
    password: string;
    passwordConfirmed: string;
}

export interface RegisterResponse {

}

export interface GetUserInfoRequest {

}

export interface GetUserInfoResponse {
    userInfo: UserInfo;
}
