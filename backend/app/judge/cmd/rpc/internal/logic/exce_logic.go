package logic

import (
	"context"
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/solitudealma/SOJ/backend/app/judge/cmd/rpc/internal/svc"
	"github.com/solitudealma/SOJ/backend/app/judge/cmd/rpc/pb"
	"github.com/solitudealma/SOJ/backend/common/constant/judge"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExceLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExceLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExceLogic {
	return &ExceLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ExceLogic) Exce(in *pb.ExceReq) (*pb.ExceResp, error) {

	var resp []*judge.Result
	var stdinContent string = in.UserInput
	var stdoutFileName string = "stdout"
	var stdoutMax int64 = 1024 * 1024 * 32
	var stderrFileName string = "stderr"
	var stderr int64 = 1024 * 1024 * 16
	languageConfig := judge.LanguageConfigMap[in.Language]

	if len(stdinContent)*2 > 32*1024*1024 {
		stdoutMax = int64(len(stdinContent)) * 2
	} else {
		stdoutMax = 32 * 1024 * 1024
	}

	copyIn := make(map[string]judge.CmdFile)
	if in.FileId != "" {
		copyIn[languageConfig.ExeName] = judge.CmdFile{
			FileID: &in.FileId,
		}
	} else {
		copyIn[languageConfig.ExeName] = judge.CmdFile{
			Content: &in.Content,
		}
	}

	cmd := judge.Cmd{
		Args: languageConfig.RunCommand,
		Env:  languageConfig.RunEnvs,
		Files: []*judge.CmdFile{{
			Content: &stdinContent,
		}, {
			Name: &stdoutFileName,
			Max:  &stdoutMax,
		}, {
			Name: &stderrFileName,
			Max:  &stderr,
		}},
		CpuLimit:    uint64(languageConfig.MaxCpuTime),
		MemoryLimit: uint64(languageConfig.MaxMemory),
		StackLimit:  134217728, // 128m
		ProcLimit:   128,
		CopyIn:      copyIn,
		CopyOut:     []string{"stdout", "stderr"},
	}

	client := resty.New()
	r, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string][]judge.Cmd{"cmd": {cmd}}).
		SetResult(&resp). // or SetResult(AuthSuccess{}).
		Post("http://127.0.0.1:5050/run")

	if err != nil {
		l.Logger.Errorf("exec api for run request failed, statusCode: %d, err: %v", r.StatusCode(), err)
		return &pb.ExceResp{}, err
	}

	if r.StatusCode() != 200 || len(resp) == 0 {
		var body map[string]interface{}
		json.Unmarshal(r.Body(), &body)
		errStr, _ := json.Marshal(body)
		l.Logger.Errorf("exec api for compile request failed, statusCode: %d, requestError: %v", r.StatusCode(), string(errStr))
		return &pb.ExceResp{}, errors.New(string(errStr))
	}

	var fileError = []*pb.FileError{}
	_ = copier.Copy(&fileError, &resp[0].FileError)

	return &pb.ExceResp{
		Result: []*pb.Result{{
			Status:     resp[0].Status,
			Time:       resp[0].Time / 1000000,
			Memory:     resp[0].Memory / 1024,
			Files:      resp[0].Files,
			FileIds:    resp[0].FileIds,
			FileError:  fileError,
			ExitStatus: int64(resp[0].ExitStatus),
			Error:      resp[0].Error,
			RunTime:    resp[0].RunTime,
		}},
	}, nil
}
