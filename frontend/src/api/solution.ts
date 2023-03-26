import service from '@/utils/request'
import { ResponseInterface } from './types'
import {
    GetSolutionListInfoRequest, GetSolutionListInfoResponse, GetSolutionInfoRequest, GetSolutionInfoResponse,
} from './types/solution'

// @Summary 获取题解列表信息
// @Produce  application/json
// @Param data body {}
// @Router /solution [get]
export const getSolutionListInfo = (params: GetSolutionListInfoRequest): Promise<ResponseInterface<GetSolutionListInfoResponse>> => {
    return service({
        url: '/solution',
        method: 'get',
        params: params
    })
}

// @Summary 获取题解信息
// @Produce  application/json
// @Param data body {}
// @Router /solution [get]
export const getSolutionInfo = (path: GetSolutionInfoRequest): Promise<ResponseInterface<GetSolutionInfoResponse>> => {
    return service({
        url: '/solution/' + path.solutionId,
        method: 'get',
    })
}