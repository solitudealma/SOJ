import service from '@/utils/request'
import { ResponseInterface } from './types'
import {
    GetCommentListInfoRequest, GetCommentListInfoResponse,
    CreateCommentRequest, CreateCommentResponse,
} from './types/comment'

// @Summary 获取对应题解的评论列表信息
// @Produce  application/json
// @Param data body {}
// @Router /comment [get]
export const getCommentListInfo = (params: GetCommentListInfoRequest): Promise<ResponseInterface<GetCommentListInfoResponse>> => {
    return service({
        url: '/comment',
        method: 'get',
        params: params
    })
}

// @Summary 创建对应题解的评论信息
// @Produce  application/json
// @Param data body {}
// @Router /comment [put]
export const createComment = (data: CreateCommentRequest): Promise<ResponseInterface<CreateCommentResponse>> => {
    return service({
        url: '/comment',
        method: 'put',
        data: data,
    })
}
