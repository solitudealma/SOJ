export interface Solution {
    problemSource: string;
    solutionId: number;
    problemId: string;
    title: string;
	createTime: number;
	authorId: number;
    authorName: string;
    authorAvatar: string;
    problemDifficulty: number;
	read: number;
}

export interface SolutionInfo {
    problemId: string;
    title: string;
    problemSource: string;
    content: string;
	problemLink: string;
	problemDifficulty: number;
	updateTime: number;
    authorId: number;
	authorName: string;
    authorAvatar: string;
	read: number;
}

export interface GetSolutionListInfoRequest {
    currentPage: number;
}

export interface GetSolutionListInfoResponse {
    solutions: Array<Solution>;
    total: number;
}

export interface GetSolutionInfoRequest {
    solutionId: string;
}

export interface GetSolutionInfoResponse {
    solutionInfo: SolutionInfo;
}

export interface GetSavedSolutionInfoRequest {
    authorId: number;
}

export interface GetSavedSolutionInfoResponse {
    problemId: string;
    title: string;
    problemSource: string;
    content: string;
	problemLink: string;
	problemDifficulty: number;
}

export interface CreateSolutionInfoRequest {
    problemId: string;
    title: string;
    problemSource: string;
    content: string;
	problemLink: string;
	problemDifficulty: number;
    authorId: number;
	authorName: string;
    authorAvatar: string;
    type: number;
}

export interface CreateSolutionInfoResponse {
    solutionId: number;
}