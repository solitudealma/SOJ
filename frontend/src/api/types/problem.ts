export interface ProblemInfo {
    problemId: string;
    title: string;
	passingRate: number;
	difficulty: number;
	isAccepted: boolean;
}

export interface GetProblemListInfoRequest {
    currentPage: number;
}

export interface GetProblemListInfoResponse {
    problems: Array<ProblemInfo>;
    total: number;
}