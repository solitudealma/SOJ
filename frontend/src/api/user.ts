import service from '@/utils/request'
import { ResponseInterface } from './types'
import { LoginRequest, LoginReponse, RegisterRequest, RegisterResponse, GetUserInfoResponse } from './types/user'

// @Summary 用户登录
// @Produce  application/json
// @Param data body {username:"string",password:"string"}
// @Router /user/login [post]
export const login = (data: LoginRequest): Promise<ResponseInterface<LoginReponse>> => {
    return service({
        url: '/user/login',
        method: 'post',
        data: data
    })
}

// @Summary 用户注册
// @Produce  application/json
// @Param data body {username:"string",password:"string",passwordConfirm:"string"}
// @Router /user/register [post]
export const register = (data: RegisterRequest): Promise<ResponseInterface<RegisterResponse>> => {
    return service({
        url: '/user/register',
        method: 'post',
        data: data
    })
}

// @Summary 用户注销登录
// @Produce  application/json
// @Param data body {}
// @Router /user/logout [post]
export const logout = (): Promise<ResponseInterface<null>> => {
    return service({
        url: '/user/logout',
        method: 'post',
        data: {},
    })
}

// @Summary 获取用户信息
// @Produce  application/json
// @Param data body {}
// @Router /user/info [get]
export const getUserInfo = (): Promise<ResponseInterface<GetUserInfoResponse>> => {
    return service({
        url: '/user/info',
        method: 'get',
        data: {},
    })
}