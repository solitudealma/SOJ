import service from '@/utils/request'
import { ResponseInterface } from './types'
import {
    GetSolutionListInfoRequest, GetSolutionListInfoResponse, GetSolutionInfoRequest, GetSolutionInfoResponse,
    CreateSolutionInfoRequest, CreateSolutionInfoResponse, GetSavedSolutionInfoRequest, GetSavedSolutionInfoResponse,
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

// @Summary 获取保存的题解信息
// @Produce  application/json
// @Param data body {}
// @Router /solution [get]
export const getSavedSolutionInfo = (path: GetSavedSolutionInfoRequest): Promise<ResponseInterface<GetSavedSolutionInfoResponse>> => {
    return service({
        url: '/solution/saved/' + path.authorId,
        method: 'get',
    })
}

// @Summary 创建题解信息
// @Produce  application/json
// @Param data body {}
// @Router /solution [put]
export const createSolutionInfo = (data: CreateSolutionInfoRequest): Promise<ResponseInterface<CreateSolutionInfoResponse>> => {
    return service({
        url: '/solution',
        method: 'put',
        data: data,
    })
}