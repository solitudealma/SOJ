export interface Problem {
    problemId: string;
    title: string;
	passingRate: number;
	difficulty: number;
	isAccepted: boolean;
}

export interface ProblemInfo {
    problemId: string;
    title: string;
    description: string;
	input: string;
	output: string;
	examples: string;
	difficulty: number;
    timeLimit: number;
	memoryLimit: number;
	totalNumberOfPasses: number;
	totalAttempts: number;
	source: string;
	tags: string;
}

export interface GetProblemListInfoRequest {
    currentPage: number;
}

export interface GetProblemListInfoResponse {
    problems: Array<Problem>;
    total: number;
}

export interface GetProblemInfoRequest {
    problemId: string;
}

export interface GetProblemInfoResponse {
    problemInfo: ProblemInfo;
}