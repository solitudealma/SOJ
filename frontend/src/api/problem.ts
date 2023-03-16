import service from '@/utils/request'
import { ResponseInterface } from './types'
import { GetProblemListInfoRequest, GetProblemListInfoResponse, GetProblemInfoRequest, GetProblemInfoResponse } from './types/problem'

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