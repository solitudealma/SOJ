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

type CompileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompileLogic {
	return &CompileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CompileLogic) Compile(in *pb.CompileReq) (*pb.CompileResp, error) {

	var resp []*judge.Result
	var cmd judge.Cmd
	var stdinEmptyContent string = ""
	var stdoutFileName string = "stdout"
	var stdoutMax int64 = 1024 * 1024 * 32
	var stderrFileName string = "stderr"
	var stderr int64 = 1024 * 1024 * 32

	languageConfig := judge.LanguageConfigMap[in.Language]

	cmd = judge.Cmd{
		Args: languageConfig.CompileCommand,
		Env:  languageConfig.CompileEnvs,
		Files: []*judge.CmdFile{{
			Content: &stdinEmptyContent,
		}, {
			Name: &stdoutFileName,
			Max:  &stdoutMax,
		}, {
			Name: &stderrFileName,
			Max:  &stderr,
		}},
		CpuLimit:    uint64(languageConfig.MaxCpuTime),
		ClockLimit:  uint64(languageConfig.MaxRealTime),
		MemoryLimit: uint64(languageConfig.MaxMemory),
		StackLimit:  134217728, // 128m
		ProcLimit:   128,
		CopyIn: map[string]judge.CmdFile{
			languageConfig.SrcName: {
				Content: &in.Code,
			},
		},
		CopyOut:       []string{"stdout", "stderr"},
		CopyOutCached: []string{languageConfig.ExeName},
	}

	client := resty.New()
	r, err := client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string][]judge.Cmd{"cmd": {cmd}}).
		SetResult(&resp). // or SetResult(AuthSuccess{}).
		Post("http://127.0.0.1:5050/run")

	if err != nil {
		l.Logger.Errorf("compile api for compile request failed, statusCode: %d, err: %v", r.StatusCode(), err)
		return &pb.CompileResp{}, err
	}

	if r.StatusCode() != 200 || len(resp) == 0 {
		var body map[string]interface{}
		json.Unmarshal(r.Body(), &body)
		errStr, _ := json.Marshal(body)
		l.Logger.Errorf("compile api for compile request failed, statusCode: %d, requestError: %v", r.StatusCode(), string(errStr))
		return &pb.CompileResp{}, errors.New(string(errStr))
	}

	if resp[0].Status != "Accepted" {
		stderr := resp[0].Files["stderr"]
		l.Logger.Errorf("compile rpc for debug request failed, status: %v, stderr: %v", resp[0].Status, stderr)
	}

	var fileError = []*pb.FileError{}
	_ = copier.Copy(&fileError, &resp[0].FileError)

	return &pb.CompileResp{
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
