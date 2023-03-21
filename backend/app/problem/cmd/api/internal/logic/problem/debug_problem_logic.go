package problem

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/solitudealma/SOJ/backend/app/judge/cmd/rpc/pb"
	"github.com/solitudealma/SOJ/backend/app/problem/cmd/api/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/problem/cmd/api/internal/types"
	"github.com/solitudealma/SOJ/backend/app/problem/model"
	"github.com/solitudealma/SOJ/backend/common/constant/judge"
	"github.com/solitudealma/SOJ/backend/common/xerr"
	"github.com/zeromicro/go-zero/core/logx"
)

type DebugProblemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDebugProblemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DebugProblemLogic {
	return &DebugProblemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DebugProblemLogic) DebugProblem(req *types.DebugProblemReq) (resp *types.DebugProblemResp, err error) {

	var fileId string = ""

	var compileResp = &pb.CompileResp{}
	var execResp = &pb.ExceResp{}

	languageConfig := judge.LanguageConfigMap[req.Language]

	problem, err := l.svcCtx.ProblemModel.FindOneByProblemId(l.ctx, req.ProblemId)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "GetProblemInfo failed, db err: %v", err)
	}

	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("该题目信息不存在"), "can't find the userInfo, err: %v", err)
	}

	if languageConfig.SrcName == "" || (!strings.HasSuffix(languageConfig.SrcName, ".c") &&
		!strings.HasSuffix(languageConfig.SrcName, ".cpp")) {
		problem.TimeLimit *= 2
		problem.MemoryLimit *= 2
	}

	if languageConfig.CompileCommand != nil {
		compileResp, err = l.svcCtx.JudgeRpc.Compile(l.ctx, &pb.CompileReq{
			Code:     req.Code,
			Language: req.Language,
		})

		if err != nil {
			l.Logger.Errorf("compile rpc for debug request failed, err: %v", err)
			return &types.DebugProblemResp{}, err
		}

		if compileResp.Result[0].Status != "Accepted" {
			stderr := compileResp.Result[0].Files["stderr"]

			l.Logger.Errorf("compile rpc for debug request failed, status: %v, stderr: %v", compileResp.Result[0].Status, stderr)
			return &types.DebugProblemResp{
				Result: types.Result{
					ExpectedOutput: "22",
					Memory:         int64(compileResp.Result[0].Memory),
					Status:         compileResp.Result[0].Status,
					Stderr:         stderr,
					Time:           int64(compileResp.Result[0].Time),
					UserInput:      req.UserInput,
					UserOutput:     compileResp.Result[0].Files["stdout"],
				},
			}, err
		}

		fileId = compileResp.Result[0].FileIds[languageConfig.ExeName]
	}

	execResp, err = l.svcCtx.JudgeRpc.Exce(l.ctx, &pb.ExceReq{
		Content:   req.Code,
		FileId:    fileId,
		Language:  req.Language,
		UserInput: req.UserInput,
	})

	if err != nil {
		l.Logger.Errorf("exec rpc for debug request failed, err: %v", err)
		return &types.DebugProblemResp{}, err
	}

	return &types.DebugProblemResp{
		Result: types.Result{
			ExpectedOutput: "22",
			Memory:         int64(execResp.Result[0].Memory / 1024),
			Time:           int64(execResp.Result[0].Time / 1000000),
			Status:         execResp.Result[0].Status,
			Stderr:         execResp.Result[0].Files["stderr"],
			UserInput:      req.UserInput,
			UserOutput:     execResp.Result[0].Files["stdout"],
		},
	}, nil
}
