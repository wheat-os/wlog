package wlog

import (
	"fmt"
	"sync"

	jsoniter "github.com/json-iterator/go"
)

type jsonFormatter struct {
	ignoreBasicFields bool
	objPoll           *sync.Pool
}

func (j *jsonFormatter) object() *jsonObj {
	return j.objPoll.Get().(*jsonObj)
}

func (j *jsonFormatter) release(obj *jsonObj) {
	// 初始化对象
	obj.File, obj.Func, obj.Level, obj.FormatTime, obj.Message = "", "", "", "", ""
	j.objPoll.Put(obj)
}

func JsonFormatter(ignoreBasicFields bool) Formatter {
	return &jsonFormatter{
		ignoreBasicFields: ignoreBasicFields,
		objPoll: &sync.Pool{New: func() interface{} {
			return new(jsonObj)
		}},
	}
}

type jsonObj struct {
	Level      string `json:"level"`
	FormatTime string `json:"time"`
	File       string `json:"file"`
	Func       string `json:"func"`
	Message    string `json:"message"`
}

func (j *jsonFormatter) Format(entry *Entry) error {

	if !j.ignoreBasicFields {
		// 获取 obj json 对象
		obj := j.object()
		defer j.release(obj)

		obj.File = fmt.Sprintf("%s:%d", entry.File, entry.Line)
		obj.Level = LevelNameMapping[entry.Level]
		obj.FormatTime = entry.Time.Format(formatTime)
		obj.Func = entry.Func

		switch entry.Format {
		case FmtEmptySeparate:
			obj.Message = fmt.Sprint(entry.Args...)
		default:
			obj.Message = fmt.Sprintf(entry.Format, entry.Args...)
		}

		return jsoniter.NewEncoder(entry.Buffer).Encode(obj)
	}

	var message string
	switch entry.Format {
	case FmtEmptySeparate:
		message = fmt.Sprint(entry.Args...)
	default:
		message = fmt.Sprintf(entry.Format, entry.Args...)
	}

	entry.Buffer.WriteString(fmt.Sprintf(`{"message": "%s"}`, message))

	return nil
}
