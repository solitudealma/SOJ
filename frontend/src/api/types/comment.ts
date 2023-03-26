export interface Comment {
    [key: string]: any;
    commentId: number;
    layer: number;
    rootCommentId: number;
    fromUserId: number;
    fromAvatar: string;
    fromName: string;
    toCommentId: number;
    toUserId: number;
    toName: string;
    content: string;
	createTime: number;
    inputShow: boolean;
    reply: Array<Comment>;
}

export interface CreateCommentReq {
    solutionId: number;
    layer: number;
    rootCommentId: number;
    fromUserId: number;
    fromAvatar: string;
    fromName: string;
    toCommentId: number;
    toUserId: number;
    toName: string;
    content: string;
    createTime: number;
};

export interface GetCommentListInfoRequest {
    solutionId: number;
}

export interface GetCommentListInfoResponse {
    comments: Array<Comment>;
    total: number;
}

export interface CreateCommentRequest {
    comment: CreateCommentReq;
}

export interface CreateCommentResponse {
}