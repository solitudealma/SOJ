import service from '@/utils/request'
import { ResponseInterface } from './types'
import { GetProblemListInfoRequest, GetProblemListInfoResponse } from './types/problem'

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