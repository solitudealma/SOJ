import service from '@/utils/request'
import { ResponseInterface } from './types'
import {
    GetProblemListInfoRequest, GetProblemListInfoResponse, GetProblemInfoRequest, GetProblemInfoResponse,
    DebugProblemRequest, DebugProblemResponse,
} from './types/problem'

// @Summary 获取题目列表信息
// @Produce  application/json
// @Param data body {}
// @Router /problem [get]
export const getProblemListInfo = (params: GetProblemListInfoRequest): Promise<ResponseInterface<GetProblemListInfoResponse>> => {
    return service({
        url: '/problem',
        method: 'get',
        params: params
    })
}

// @Summary 获取题目信息
// @Produce  application/json
// @Param data body {}
// @Router /problem [get]
export const getProblemInfo = (path: GetProblemInfoRequest): Promise<ResponseInterface<GetProblemInfoResponse>> => {
    return service({
        url: '/problem/' + path.problemId,
        method: 'get',
    })
}

// @Summary 调试题目
// @Produce  application/json
// @Param data body {}
// @Router /problem/debug [post]
export const debugProblem = (data: DebugProblemRequest): Promise<ResponseInterface<DebugProblemResponse>> => {
    return service({
        url: '/problem/debug',
        method: 'post',
        data: data,
    })
}